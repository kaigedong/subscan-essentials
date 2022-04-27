# Subscan Essentials

![License: GPL](https://img.shields.io/badge/license-GPL-blue.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/kaigedong/subscan)](https://goreportcard.com/report/github.com/kaigedong/subscan)
![subscan](https://github.com/kaigedong/subscan/workflows/subscan/badge.svg)

Subscan Essentials 是一个高精度的区块链浏览器脚手架项目，它具有开发人员友好的界面和自定义模块解析功能，支持基于 substrate 的区块链网络。
它由 Subscan 团队开发，并在 subscan.io 中使用。开发人员可以自由使用代码库来扩展功能并为其受众开发独特的用户体验。

## Contents

- [Subscan Essentials](#subscan-essentials)
  - [Contents](#contents)
    - [Feature](#feature)
  - [QuickStart](#quickstart)
    - [Requirement](#requirement)
    - [Install](#install)
    - [UI](#ui)
    - [Config](#config)
      - [Init config file](#init-config-file)
      - [Set](#set)
    - [Usage](#usage)
    - [Docker](#docker)
    - [Test](#test)
  - [LICENSE](#license)
  - [Resource](#resource)

### Feature

1. 支持 substrate 网络[自定义](/custom_type.md)type 注册
2. 支持索引 block, Extrinsic, Event, log
3. 可自定义[插件](/plugins)索引更多的数据
4. [Gen](https://github.com/kaigedong/subscan-plugin/tree/master/tool) 工具可自动生成插件模版
5. 内置默认的 HTTP API [DOC](/docs/index.md)

## QuickStart

### Requirement

- Linux / Mac OSX
- Git
- Golang 1.12.4+
- Redis 3.0.4+
- MySQL 5.6+
- Node 8.9.0+

### Install

```bash
./build.sh build

//UI
cd ui && yarn && yarn dev
```

### UI

UI 部分使用 [nuxt.js](https://nuxtjs.org/) 和 [amis](https://github.com/baidu/amis)

Demo: [blocks](/ui/plugins/blocks.js), 更多配置请参考 [amis docs](https://baidu.gitee.io/amis/docs/index)

[在线示例](https://crab.demo.subscan.io/)

在开发环境中请替换 nuxt.config.js 中的 proxy target

```js
proxy: {
   "/api": {
      target: "https://your_server_name.com",
      secure: false,
      changeOrigin: true,
      pathRewrite: {
         "^/api": "/api"
      }
   },
}
```

在生产环境中请替换 nuxt.config.js 中的 browserBaseURL

```js
axios: {
   proxy: process.env.NODE_ENV !== 'production',
    browserBaseURL: process.env.NODE_ENV !== 'production' ? "" : "https://your_server_name.com"
},
```

![ui_demo](./ui_demo.png)

### Config

#### Init config file

```bash
cp configs/redis.toml.example configs/redis.toml && cp configs/mysql.toml.example configs/mysql.toml && cp configs/http.toml.example configs/http.toml
```

#### Set

1. Redis configs/redis.toml

> addr： redis host and port (default: 127.0.0.1:6379)

2. Mysql configs/mysql.toml

> host: mysql host (default: 127.0.0.1)
> user: mysql user (default: root)
> pass: mysql user passwd (default: "")
> db: mysql db name (default: "subscan")

3. Http configs/http.toml

> addr: local http server port (default: 0.0.0.0:4399)

### Usage

- Start DB

**Make sure you have started redis and mysql**

- Substrate Daemon

```bash
cd cmd
./subscan start substrate
```

- Api Server

```bash
cd cmd
./subscan
```

- Help

```
NAME:
   SubScan - SubScan Backend Service, use -h get help

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   1.0

DESCRIPTION:
   SubScan Backend Service, substrate blockchain explorer

COMMANDS:
     start    Start one worker, E.g substrate
     stop     Stop one worker, E.g substrate
     install  Create database and create default conf file
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --conf value   (default: "../configs")
   --help, -h     show help
   --version, -v  print the version


```

### Docker

创建本地 docker 网络

```
docker network create app_net
```

启动 mysql 和 redis 的 container

```bash
docker-compose -f docker-compose.db.yml up  -d
```

启动 subscan 服务

```bash
docker-compose build
docker-compose up -d
```

### Test

```bash
go test ./...

//UI
cd ui && yarn && yarn test
```

## LICENSE

GPL-3.0

## Resource

- [ITERING] https://github.com/itering
- [Darwinia] https://github.com/darwinia-network/darwinia
