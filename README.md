[![open-api](https://github.com/Jon-a-than/gin-template/actions/workflows/open-api.yml/badge.svg)](https://qingshaner.com/gin-template/swagger.json)

## 命令说明

使用 `task` 作为命令管理工具, 配置文件已开启 `go modules` 以及 `goproxy` 的环境变量

若直接使用 `go` 命令行，先检查是否开启 `go modules` 环境变量

安装 `task`：

```bash
go install github.com/go-task/task/v3/cmd/task@latest
```

安装 `swag` 用于生成 `swagger` 文档：

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### 安装依赖包

```bash
task install
```

### rsa 生成

使用 `openssl` 生成 `JWT` `RSA` 密钥对，需先安装 `openssl`

```bash
task rsa
```

若仅用于开发环境，且不想安装 `openssl`，可使用在线生成工具生成密钥对，将私钥保存至根目录 `private_key.pem`，将公钥保存至根目录 `public_key.pem`。路径可在 `config/app.yaml` 中修改

### 数据库表结构更新

```bash
task migrate
```

## 开发环境运行

```bash
task dev # go run + 文件监听
task run # go run
```

## docker compose 运行

```bash
task docker:up
```

## 编译

```bash
task build
```

## API 文档

使用 `swag` 生成文档信息，使用 `APIFox` 作为文档展示 / 调试工具。

### 生成文档

```bash
task docs
```
