// This file is based on code written by The Go Authors.
// See original source: https://github.com/golang/go/blob/go1.22.7/src/sync/map_test.go
//
// --- Original License Notice ---
//
// Copyright (c) 2009 The Go Authors. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
// * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
// * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
// * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
package main

import (
	"math/rand"
	"reflect"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"testing/quick"

	"utils/testutil"
)

type mapOp string

const (
	opLoad             = mapOp("Load")
	opStore            = mapOp("Store")
	opLoadOrStore      = mapOp("LoadOrStore")
	opLoadAndDelete    = mapOp("LoadAndDelete")
	opDelete           = mapOp("Delete")
	opSwap             = mapOp("Swap")
	opCompareAndSwap   = mapOp("CompareAndSwap")
	opCompareAndDelete = mapOp("CompareAndDelete")
)

var mapOps = [...]mapOp{
	opLoad,
	opStore,
	opLoadOrStore,
	opLoadAndDelete,
	opDelete,
	opSwap,
	opCompareAndSwap,
	opCompareAndDelete,
}

// mapCall is a quick.Generator for calls on mapInterface.
type mapCall struct {
	op   mapOp
	k, v any
}

func (c mapCall) apply(m mapInterface) (any, bool) {
	switch c.op {
	case opLoad:
		return m.Load(c.k)
	case opStore:
		m.Store(c.k, c.v)
		return nil, false
	case opLoadOrStore:
		return m.LoadOrStore(c.k, c.v)
	case opLoadAndDelete:
		return m.LoadAndDelete(c.k)
	case opDelete:
		m.Delete(c.k)
		return nil, false
	case opSwap:
		return m.Swap(c.k, c.v)
	case opCompareAndSwap:
		if m.CompareAndSwap(c.k, c.v, rand.Int()) {
			m.Delete(c.k)
			return c.v, true
		}
		return nil, false
	case opCompareAndDelete:
		if m.CompareAndDelete(c.k, c.v) {
			if _, ok := m.Load(c.k); !ok {
				return nil, true
			}
		}
		return nil, false
	default:
		panic("invalid mapOp")
	}
}

type mapResult struct {
	value any
	ok    bool
}

func randValue(r *rand.Rand) any {
	b := make([]byte, r.Intn(4))
	for i := range b {
		b[i] = 'a' + byte(rand.Intn(26))
	}
	return string(b)
}

func (mapCall) Generate(r *rand.Rand, size int) reflect.Value {
	c := mapCall{op: mapOps[rand.Intn(len(mapOps))], k: randValue(r)}
	switch c.op {
	case opStore, opLoadOrStore:
		c.v = randValue(r)
	}
	return reflect.ValueOf(c)
}

func applyCalls(m mapInterface, calls []mapCall) (results []mapResult, final map[any]any) {
	for _, c := range calls {
		v, ok := c.apply(m)
		results = append(results, mapResult{v, ok})
	}

	final = make(map[any]any)
	m.Range(func(k, v any) bool {
		final[k] = v
		return true
	})

	return results, final
}

func applyMap(calls []mapCall) ([]mapResult, map[any]any) {
	return applyCalls(new(sync.Map), calls)
}

func applyRWMutexMap(calls []mapCall) ([]mapResult, map[any]any) {
	return applyCalls(new(RWMutexMap), calls)
}

func applyDeepCopyMap(calls []mapCall) ([]mapResult, map[any]any) {
	return applyCalls(new(DeepCopyMap), calls)
}

func TestMapMatchesRWMutex(t *testutil.TestRunner) {
	if err := quick.CheckEqual(applyMap, applyRWMutexMap, nil); err != nil {
		t.Error(err)
	}
}

func TestMapMatchesDeepCopy(t *testutil.TestRunner) {
	if err := quick.CheckEqual(applyMap, applyDeepCopyMap, nil); err != nil {
		t.Error(err)
	}
}

func TestConcurrentRange(t *testutil.TestRunner) {
	const mapSize = 1 << 10

	m := new(sync.Map)
	for n := int64(1); n <= mapSize; n++ {
		m.Store(n, int64(n))
	}

	done := make(chan struct{})
	var wg sync.WaitGroup
	defer func() {
		close(done)
		wg.Wait()
	}()
	for g := int64(runtime.GOMAXPROCS(0)); g > 0; g-- {
		r := rand.New(rand.NewSource(g))
		wg.Add(1)
		go func(g int64) {
			defer wg.Done()
			for i := int64(0); ; i++ {
				select {
				case <-done:
					return
				default:
				}
				for n := int64(1); n < mapSize; n++ {
					if r.Int63n(mapSize) == 0 {
						m.Store(n, n*i*g)
					} else {
						m.Load(n)
					}
				}
			}
		}(g)
	}

	//iters := 1 << 10
	//if testing.Short() {
	//	iters = 16
	//}
	iters := 16
	for n := iters; n > 0; n-- {
		seen := make(map[int64]bool, mapSize)

		m.Range(func(ki, vi any) bool {
			k, v := ki.(int64), vi.(int64)
			if v%k != 0 {
				t.Fatalf("while Storing multiples of %v, Range saw value %v", k, v)
			}
			if seen[k] {
				t.Fatalf("Range visited key %v twice", k)
			}
			seen[k] = true
			return true
		})

		if len(seen) != mapSize {
			t.Fatalf("Range visited %v elements of %v-element Map", len(seen), mapSize)
		}
	}
}

func TestIssue40999(t *testutil.TestRunner) {
	var m sync.Map

	// Since the miss-counting in missLocked (via Delete)
	// compares the miss count with len(m.dirty),
	// add an initial entry to bias len(m.dirty) above the miss count.
	m.Store(nil, struct{}{})

	var finalized uint32

	// Set finalizers that count for collected keys. A non-zero count
	// indicates that keys have not been leaked.
	for atomic.LoadUint32(&finalized) == 0 {
		p := new(int)
		runtime.SetFinalizer(p, func(*int) {
			atomic.AddUint32(&finalized, 1)
		})
		m.Store(p, struct{}{})
		m.Delete(p)
		runtime.GC()
	}
}

func TestMapRangeNestedCall(t *testutil.TestRunner) { // Issue 46399
	var m sync.Map
	for i, v := range [3]string{"hello", "world", "Go"} {
		m.Store(i, v)
	}
	m.Range(func(key, value any) bool {
		m.Range(func(key, value any) bool {
			// We should be able to load the key offered in the Range callback,
			// because there are no concurrent Delete involved in this tested map.
			if v, ok := m.Load(key); !ok || !reflect.DeepEqual(v, value) {
				t.Fatalf("Nested Range loads unexpected value, got %+v want %+v", v, value)
			}

			// We didn't keep 42 and a value into the map before, if somehow we loaded
			// a value from such a key, meaning there must be an internal bug regarding
			// nested range in the Map.
			if _, loaded := m.LoadOrStore(42, "dummy"); loaded {
				t.Fatalf("Nested Range loads unexpected value, want store a new value")
			}

			// Try to Store then LoadAndDelete the corresponding value with the key
			// 42 to the Map. In this case, the key 42 and associated value should be
			// removed from the Map. Therefore any future range won't observe key 42
			// as we checked in above.
			val := "sync.Map"
			m.Store(42, val)
			if v, loaded := m.LoadAndDelete(42); !loaded || !reflect.DeepEqual(v, val) {
				t.Fatalf("Nested Range loads unexpected value, got %v, want %v", v, val)
			}
			return true
		})

		// Remove key from Map on-the-fly.
		m.Delete(key)
		return true
	})

	// After a Range of Delete, all keys should be removed and any
	// further Range won't invoke the callback. Hence length remains 0.
	length := 0
	m.Range(func(key, value any) bool {
		length++
		return true
	})

	if length != 0 {
		t.Fatalf("Unexpected sync.Map size, got %v want %v", length, 0)
	}
}

func TestCompareAndSwap_NonExistingKey(t *testutil.TestRunner) {
	m := &sync.Map{}
	if m.CompareAndSwap(m, nil, 42) {
		// See https://go.dev/issue/51972#issuecomment-1126408637.
		t.Fatalf("CompareAndSwap on a non-existing key succeeded")
	}
}

func TestMapRangeNoAllocations(t *testutil.TestRunner) { // Issue 62404
	var m sync.Map
	allocs := testing.AllocsPerRun(10, func() {
		m.Range(func(key, value any) bool {
			return true
		})
	})
	if allocs > 0 {
		t.Errorf("AllocsPerRun of m.Range = %v; want 0", allocs)
	}
}
