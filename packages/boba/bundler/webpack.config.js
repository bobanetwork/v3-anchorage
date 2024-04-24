const path = require('path')
const { IgnorePlugin, ProvidePlugin } = require('webpack')

module.exports = {
  plugins: [
    new IgnorePlugin({ resourceRegExp: /electron/ }),
    new IgnorePlugin({ resourceRegExp: /^scrypt$/ }),
    new IgnorePlugin({ resourceRegExp: /solidity-analyzer/ }),
    new IgnorePlugin({ resourceRegExp: /fsevents/ }),
    new ProvidePlugin({
      WebSocket: 'ws',
      fetch: ['node-fetch', 'default'],
    }),
  ],
  resolve: {
    modules: [
      'node_modules',
      '../bundler_sdk/node_modules',
      '../bundler_utils/node_modules',
      '../accountabstraction/node_modules',
    ],
    alias: {
      'ethereum-cryptography/secp256k1': require.resolve(
        'ethereum-cryptography/secp256k1'
      ),
    },
  },
  target: 'node',
  entry: './dist/src/exec.js',
  mode: 'production',
  module: {
    rules: [
      {
        test: /\.tsx?$/,
        use: 'ts-loader',
        exclude: /node_modules/,
      },
    ],
  },
  output: {
    path: path.resolve(__dirname),
    filename: 'bundler.js',
  },
  stats: 'errors-only',
}
