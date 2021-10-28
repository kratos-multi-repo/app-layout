# Kratos Project Template For Multi-repo

[kratos](https://go-kratos.dev/) 项目管理是大仓/单仓模式（monorepo），这也是业界比较流行的最佳实践。但是，在某些特殊情况下需要对敏感项目单独管理，
于是就做了一个小仓/多仓模式的尝试，将 `protocol-buffers`，`sdk` 公开，但是对项目实现细节隐藏。旨在为敏感保密项目的仓库管理
提供一种实现。

__Multi-repo Structure__

多仓模式的项目组成

| 项目 | 样例 | 可见性 | 说明 |
| ---- | ---- | ---- | ---- |
| 协议 | [protocol-buffers](https://github.com/kratos-multi-repo/protocol-buffers) | 公开 |      |
| SDK | [helloworld-sdk](https://github.com/kratos-multi-repo/helloworld-sdk) | 公开 | |
| 保密项目 | 该模板仓库 | 私有 | |

一般开发流程：

通过 `protocol-buffers` 约定协议，并生成 `sdk (helloworld-sdk)` ，在隐私项目中引入 `sdk` 即可。

## Usage

```shell
kratos new helloworld -r https://github.com/kratos-multi-repo/app-layout.git

```

### Run

```shell
go mod download
kratos run

```

### Testing

```shell
curl 'http://127.0.0.1:8000/helloworld/v1/ping'
# {"message":"pong"}

curl 'http://127.0.0.1:8000/helloworld/v1/visitors/kratos'
# {"message":"Hello kratos"}

```

## Project Structure

```
.
├── cmd  // 整个项目启动的入口文件
│    ├── server
│    │    ├── main.go  // server 启动入口
│    │    ├── wire.go  // 我们使用 wire 来维护依赖注入
│    │    └── wire_gen.go
│    ├── job
│    └── cron
├── configs  // 这里通常维护一些本地调试用的样例配置文件
│    ├── config.template.yaml
│    └── config.yaml
├── internal // 该服务所有不对外暴露的代码，通常的业务逻辑都在这下面，使用internal避免错误引用
│    ├── biz // 业务逻辑的组装层，类似 DDD 的 domain 层，data 类似 DDD 的 repo，repo 接口在这里定义，使用依赖倒置的原则。
│    │    ├── greeter // 模块名
│    │    │    ├── entity.go    // 领域值对象
│    │    │    ├── interface.go // 领域接口，有 db 接口 IRepo，用例接口 IUseCase 等
│    │    │    └── use_case.go  // 领域用例，具体业务逻辑实现。注入 service 层与之交互，通过 IRepo 实现 IUseCase
│    │    ├── main.go // wire provider
│    │    └── README.md
│    ├── conf  // 内部使用的 config 的结构定义，使用 proto 格式生成
│    │    ├── conf.pb.go
│    │    └── conf.proto
│    ├── data  // 业务数据访问，包含 cache、db 等封装，实现了 biz 的 repo 接口。
│    │    ├── model
│    │    │    └── greeter.go // db 模型
│    │    ├── client.go  // db 链接端口
│    │    ├── greeter.go // 数据 dao，实现 biz 层的 IRepo 接口
│    │    ├── main.go    // wire provider
│    │    └── README.md
│    ├── server  // http 和 grpc 实例的创建和配置
│    │    ├── grpc.go
│    │    ├── http.go
│    │    └── main.go // wire provider
│    └── service  // 实现了 sdk 中 api 定义的服务层，类似 DDD 的 application 层，处理 DTO 到 biz 领域实体的转换(DTO -> DO)，
│         │       // 同时协同各类 biz 交互，但是不应处理复杂逻辑
│         ├── README.md
│         ├── greeter.go
│         └── main.go // wire provider
├── Dockerfile
├── go.mod
├── go.sum
├── LICENSE
├── Makefile
└── README.md

```

## Link

[kratos docs](https://go-kratos.dev/)

[kratos repo](https://github.com/go-kratos/kratos)

[kratos monorepo layout](https://github.com/go-kratos/kratos-layout)
