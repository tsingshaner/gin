[![openapi](https://github.com/Jon-a-than/gin-template/actions/workflows/openapi.yml/badge.svg)](https://qingshaner.com/gin-template/swagger.json)

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

若仅用于开发环境，且不想安装 `openssl`，可使用在线生成工具生成密钥对，将私钥保存至 `config` 目录 `private_key.pem`，将公钥保存至 `config` 目录 `public_key.pem`。路径可在 `config/app.yaml` 中修改
```
/config
  app.yaml
  config.go
  private_key.pem
  public_key.pem
  // ...
```


❗`docker compose` 配置文件挂载在 `deployments/docker/config` 中, 配置文件路径如下：

```
/deployments/docker/config
  app.yaml
  private_key.pem
  public_key.pem
```

### 数据库表结构更新

```bash
task migrate
```

### 本地开发环境运行

需自行配置 `postgresql` 数据库，配置文件路径为 `config/app.yaml`

若使用 `GoLand` 可以将 `fmt`

```bash
task dev # task run + 文件监听
task run # go run + fmt + docs
```

### docker compose 运行

```bash
task docker:build    # 构建镜像
task docker:up       # 启动容器
task docker:rebuild  # 重建镜像并启动容器
```

### 本地编译

```bash
task build
```

### 部署环境编译(alpine 容器)

```bash
task build:alpine # 用于镜像构建
```

### 生成文档

使用 `swag` 生成文档信息，使用 `APIFox` 作为文档展示 / 调试工具。

```bash
task docs
```
