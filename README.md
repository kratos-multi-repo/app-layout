# Kratos Project Template For Multi-repo

[kratos](https://go-kratos.dev/) 项目管理是大仓/单仓模式（monorepo），这也是业界比较流行的最佳实践。但是，在某些特殊情况下需要对敏感项目单独管理，
于是就做了一个小仓/多仓模式的尝试，将 `protocol-buffers`，`sdk` 公开，但是对项目实现细节隐藏。为敏感保密项目的仓库管理
提供一种方式。

## Multi-repo Structure

多仓模式的项目组成

| 项目 | 样例 | 可见性 | 说明 |
| ---- | ---- | ---- | ---- |
| 协议 | [protocol-buffers](https://github.com/kratos-multi-repo/protocol-buffers) | 公开 |      |
| SDK | [helloworld-sdk](https://github.com/kratos-multi-repo/helloworld-sdk) | 公开 | |
| 保密项目 | 该模板仓库 | 私有 | |

一般开发流程：

通过 `protocol-buffers` 约定协议，并生成 `sdk (helloworld-sdk)` ，在隐私项目中引入 `skd` 即可。

## Run

```shell
go mod download
kratos run

```

Testing

```shell
curl 'http://127.0.0.1:8000/helloworld/v1/ping'
# {"message":"pong"}

curl 'http://127.0.0.1:8000/helloworld/v1/visitors/kratos'
# {"message":"Hello kratos"}

```

## Link

[kratos docs](https://go-kratos.dev/)

[kratos repo](https://github.com/go-kratos/kratos)

[kratos monorepo layout](https://github.com/go-kratos/kratos-layout)
