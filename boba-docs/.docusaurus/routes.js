import React from 'react';
import ComponentCreator from '@docusaurus/ComponentCreator';

export default [
  {
    path: '/__docusaurus/debug',
    component: ComponentCreator('/__docusaurus/debug', '5ff'),
    exact: true
  },
  {
    path: '/__docusaurus/debug/config',
    component: ComponentCreator('/__docusaurus/debug/config', '5ba'),
    exact: true
  },
  {
    path: '/__docusaurus/debug/content',
    component: ComponentCreator('/__docusaurus/debug/content', 'a2b'),
    exact: true
  },
  {
    path: '/__docusaurus/debug/globalData',
    component: ComponentCreator('/__docusaurus/debug/globalData', 'c3c'),
    exact: true
  },
  {
    path: '/__docusaurus/debug/metadata',
    component: ComponentCreator('/__docusaurus/debug/metadata', '156'),
    exact: true
  },
  {
    path: '/__docusaurus/debug/registry',
    component: ComponentCreator('/__docusaurus/debug/registry', '88c'),
    exact: true
  },
  {
    path: '/__docusaurus/debug/routes',
    component: ComponentCreator('/__docusaurus/debug/routes', '000'),
    exact: true
  },
  {
    path: '/blog',
    component: ComponentCreator('/blog', 'b2f'),
    exact: true
  },
  {
    path: '/blog/archive',
    component: ComponentCreator('/blog/archive', '182'),
    exact: true
  },
  {
    path: '/blog/authors',
    component: ComponentCreator('/blog/authors', '0b7'),
    exact: true
  },
  {
    path: '/blog/authors/all-sebastien-lorber-articles',
    component: ComponentCreator('/blog/authors/all-sebastien-lorber-articles', '4a1'),
    exact: true
  },
  {
    path: '/blog/authors/yangshun',
    component: ComponentCreator('/blog/authors/yangshun', 'a68'),
    exact: true
  },
  {
    path: '/blog/first-blog-post',
    component: ComponentCreator('/blog/first-blog-post', '89a'),
    exact: true
  },
  {
    path: '/blog/long-blog-post',
    component: ComponentCreator('/blog/long-blog-post', '9ad'),
    exact: true
  },
  {
    path: '/blog/mdx-blog-post',
    component: ComponentCreator('/blog/mdx-blog-post', 'e9f'),
    exact: true
  },
  {
    path: '/blog/tags',
    component: ComponentCreator('/blog/tags', '287'),
    exact: true
  },
  {
    path: '/blog/tags/docusaurus',
    component: ComponentCreator('/blog/tags/docusaurus', '704'),
    exact: true
  },
  {
    path: '/blog/tags/facebook',
    component: ComponentCreator('/blog/tags/facebook', '858'),
    exact: true
  },
  {
    path: '/blog/tags/hello',
    component: ComponentCreator('/blog/tags/hello', '299'),
    exact: true
  },
  {
    path: '/blog/tags/hola',
    component: ComponentCreator('/blog/tags/hola', '00d'),
    exact: true
  },
  {
    path: '/blog/welcome',
    component: ComponentCreator('/blog/welcome', 'd2b'),
    exact: true
  },
  {
    path: '/markdown-page',
    component: ComponentCreator('/markdown-page', '3d7'),
    exact: true
  },
  {
    path: '/',
    component: ComponentCreator('/', 'a80'),
    routes: [
      {
        path: '/',
        component: ComponentCreator('/', '194'),
        routes: [
          {
            path: '/',
            component: ComponentCreator('/', '373'),
            routes: [
              {
                path: '/account-abstraction/biconomy',
                component: ComponentCreator('/account-abstraction/biconomy', '4b0'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/account-abstraction/bundler-api',
                component: ComponentCreator('/account-abstraction/bundler-api', '81b'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/account-abstraction/bundler-config',
                component: ComponentCreator('/account-abstraction/bundler-config', 'be2'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/account-abstraction/bundler-sdk',
                component: ComponentCreator('/account-abstraction/bundler-sdk', '870'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/account-abstraction/index',
                component: ComponentCreator('/account-abstraction/index', '7df'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/account-abstraction/paymasters',
                component: ComponentCreator('/account-abstraction/paymasters', 'ffc'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/addresses/',
                component: ComponentCreator('/addresses/', '287'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/addresses/aa-contract-addresses',
                component: ComponentCreator('/addresses/aa-contract-addresses', 'ef0'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/addresses/contract/network-bnb',
                component: ComponentCreator('/addresses/contract/network-bnb', '38d'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/addresses/contract/network-eth',
                component: ComponentCreator('/addresses/contract/network-eth', '4f1'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/addresses/custom-fee-token',
                component: ComponentCreator('/addresses/custom-fee-token', '7e5'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/addresses/index',
                component: ComponentCreator('/addresses/index', '621'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/addresses/token-addresses',
                component: ComponentCreator('/addresses/token-addresses', 'e32'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/basics/add-new-token',
                component: ComponentCreator('/basics/add-new-token', 'b0d'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/basics/basic-ops',
                component: ComponentCreator('/basics/basic-ops', 'a87'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/basics/bridging/boba-token-bridge',
                component: ComponentCreator('/basics/bridging/boba-token-bridge', '72c'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/basics/bridging/index',
                component: ComponentCreator('/basics/bridging/index', 'c9b'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/basics/bridging/light-bridge',
                component: ComponentCreator('/basics/bridging/light-bridge', '018'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/basics/bridging/standard-bridge',
                component: ComponentCreator('/basics/bridging/standard-bridge', '6f8'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/basics/deploy-contract',
                component: ComponentCreator('/basics/deploy-contract', '285'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/basics/json-rpc-api',
                component: ComponentCreator('/basics/json-rpc-api', 'c00'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/basics/verify-sc',
                component: ComponentCreator('/basics/verify-sc', '9a8'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/faq',
                component: ComponentCreator('/faq', '885'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/faucets',
                component: ComponentCreator('/faucets', '815'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/hc/addsub/',
                component: ComponentCreator('/hc/addsub/', 'c91'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/hc/addsub/write-contract',
                component: ComponentCreator('/hc/addsub/write-contract', '3e7'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/hc/addsub/write-handler',
                component: ComponentCreator('/hc/addsub/write-handler', 'd0b'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/hc/deploy-contract',
                component: ComponentCreator('/hc/deploy-contract', '11b'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/hc/index',
                component: ComponentCreator('/hc/index', '4c8'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/hc/price-fetch',
                component: ComponentCreator('/hc/price-fetch', '672'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/hc/server-setup',
                component: ComponentCreator('/hc/server-setup', '300'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/node-operators/index',
                component: ComponentCreator('/node-operators/index', 'c7a'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/node-operators/network-upgrades',
                component: ComponentCreator('/node-operators/network-upgrades', 'a8c'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/node-operators/run-node-docker',
                component: ComponentCreator('/node-operators/run-node-docker', 'f82'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/node-operators/run-node-source',
                component: ComponentCreator('/node-operators/run-node-source', '425'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/node-operators/snapshot-downloads',
                component: ComponentCreator('/node-operators/snapshot-downloads', '6e7'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/node-operators/software-release',
                component: ComponentCreator('/node-operators/software-release', '1a0'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/subgraph',
                component: ComponentCreator('/subgraph', 'bb7'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/',
                component: ComponentCreator('/', 'c92'),
                exact: true,
                sidebar: "tutorialSidebar"
              }
            ]
          }
        ]
      }
    ]
  },
  {
    path: '*',
    component: ComponentCreator('*'),
  },
];
