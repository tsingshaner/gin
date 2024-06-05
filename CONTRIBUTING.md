# 贡献指南

请先阅读 `README.md` 完成开发工具的安装。

## 目录结构

### 根目录

| 文件夹名      | 内容                                                                    |
| ------------- | ----------------------------------------------------------------------- |
| `cmd`         | 存放应用程序的入口                                                      |
| `pkg`         | 存放可以被其他项目使用的库代码。                                        |
| `internal`    | 业务代码，按照不同的功能模块进行划分，如 `user`、`class`、`course` 等。 |
| `api`         | 存放 `API` 定义和协议文件，如 `proto` 文件和 `swagger` 文件。           |
| `logs`        | 存放输出的日志文件。                                                    |
| `config`      | 存放配置文件和处理配置的代码。                                          |
| `deployments` | 存放与部署相关的文件，如 `Dockerfile` 和 `docker-compose.yml`。         |
| `.github`     | 存放 `GitHub Actions` 相关的配置文件                                    |
| `target`      | 存放编译后的二进制文件。                                                |

### 业务模块

模块文件夹命名参考 `DDD` 的模块划分，不同模块之间尽量只依赖 `export` 文件夹中的接口。
不同模块间尽量不要使用 `gorm`自带的一对多，多对多 `model`，使用 `id` 来关联，以保持模块的独立性。

| 模块文件夹名       | 内容                                                                                |
| ------------------ | ----------------------------------------------------------------------------------- |
| `constant`         | 存放常量，例如响应代码。                                                            |
| `domain`           | 存放业务规则, 在这里声明 `repository` 的接口声明。                                  |
| `entity`           | 存放实体类，尽量将只操作当前实体的方法在这里实现。                                  |
| `export`           | 用于给其它模块调用的接口实现。                                                      |
| `handler`          | 处理 `HTTP` 请求接口实现，将请求参数转为 `entity`，调用领域层的代码来执行业务逻辑。 |
| `infra`            | 存放基础设施代码，例如数据库操作。                                                  |
| `infra/model`      | 存放数据库模型。                                                                    |
| `infra/repository` | 存放 `repository` 接口的实现，仓库接口定义了与数据库交互的方法。                    |
| `interface`        | 存放接口定义，可能包括用于不同交互方式（如 `REST`，`gRPC` 等）的接口。              |
| `interface/rest`   | 存放 `HTTP` 请求接口定义, 声明 `handler` 的接口。                                   |
| `interface/export` | 提供给外部模块的接口。                                                              |
| `xxx.go`           | 模块入口文件，在这里声明路由和创建实例                                              |

## API 文档注释

为了尽量避免 `API` 文档和代码不一致的情况, 和重复进行请求/响应结构体的声明, 使用 `swag` 生成 `swagger` 文档再自动导入 `APIFox`。
基于 `swag` 生成 `swagger` 文档，需要在 `handler` 中的方法上添加注释，例如(详细说明请看官方仓库)：

### 文档注释

```go
// Register 为正常格式化代码，需要在请求注释前添加 // xxxx \n//
//
// @Summary 获取用户信息 对应 APIFox 中的接口名称
// @Description.markdown XXXCustomCode // 会读取 `XXXCustomCode.md` 文件作为描述信息， 对应 APIFox 中的接口说明
// @Tags    user // 对应 APIFox 中的接口标签与文件夹名
// @Accept  json
// @Produce json
// @Param   id     path      int   true  "用户ID"
// @Success 200    {object}  User
// @Failure 400    {object}  Error
// @Failure 500    {object}  Error
// @Router /v1/user/{id} [get]
func Handler(c *gin.Context) {
 // ...
}
```

### 自定义业务代码描述信息

由于请求自定义描述信息前端通常不直接使用，所以对于成功的请求默认只传自定义业务代码 `code`, 不传多余的描述信息(`msg`, `message`等)
各接口的自定义业务代码的定义在 `constant` 中，例如：

```go
package constant

import // ...

const domainCode = 100 // 模块代码

// 成功请求的自定义业务代码 resp.CodeSuccess(20000) + domainCode(100) + iota = 20100
const (
	SUCCESS_REGISTER = resp.CodeSuccess + domainCode + iota
	SUCCESS_LOGIN
)

// 请求错误的自定义业务代码 resp.CodeBadRequest(40000) + domainCode(100) + iota = 40100
const (
	USER_ALREADY_EXISTS = resp.CodeBadRequest + domainCode + iota
	USER_NOT_EXISTS
	USER_PASSWORD_ERROR
)

// 服务端错误的自定义业务代码 resp.CodeServerError(50000) + domainCode(100) + iota = 50100
const (
	DB_ERROR = resp.CodeServerError + domainCode + iota
	PASSWORD_HASH_ERROR
	TOKEN_GEN_ERROR
)

// 错误信息 会在响应体中返回 { code: xxx, err: map[xxx] }
var ErrorMessage = map[int]string{
	USER_ALREADY_EXISTS: "user already exists",
	USER_NOT_EXISTS:     "user not exists",
	USER_PASSWORD_ERROR: "user password error",

	DB_ERROR:            "server database error",
	PASSWORD_HASH_ERROR: "password hash error",
	TOKEN_GEN_ERROR:     "token gen error",
}

// 请求错误 自定义业务代码对应的 http 状态码
var StatusMap = map[int]int{
	USER_ALREADY_EXISTS: http.StatusBadRequest,
	USER_NOT_EXISTS:     http.StatusBadRequest,
	USER_PASSWORD_ERROR: http.StatusBadRequest,

	DB_ERROR:            http.StatusInternalServerError,
	PASSWORD_HASH_ERROR: http.StatusInternalServerError,
	TOKEN_GEN_ERROR:     http.StatusInternalServerError,
}

/// 自定义代码的描述信息, 会输出到 `swagger` 文档中
/// 该部分应只在文档生成时需要, 故不定义为常量, 在编译时会被优化掉
/// 每个模块的自定义代码都需要一个类似的方法
/// 返回的结构体会自动生成markdown文档, 参考上面的 `@Description.markdown` 注释
func GetCodeMap() gen.CodeMap {
	return gen.CodeMap{
    // 会在swag输出目录生成 `Register.md` 文件
		"Register": {
			SUCCESS_REGISTER: "注册成功",

			USER_ALREADY_EXISTS: "用户已存在",

			DB_ERROR:            "服务端数据库链接错误",
			PASSWORD_HASH_ERROR: "服务端密码加密错误",
			TOKEN_GEN_ERROR:     "服务端token生成错误",
		},
    // 会在swag输出目录生成 `Login.md` 文件
		"Login": {
			SUCCESS_LOGIN: "登录成功",

			USER_NOT_EXISTS:     "用户不存在",
			USER_PASSWORD_ERROR: "用户密码错误",

			DB_ERROR:            "服务端数据库链接错误",
			PASSWORD_HASH_ERROR: "服务端密码加密错误",
			TOKEN_GEN_ERROR:     "服务端token生成错误",
		},
	}
}
```

在`internal/gen_code_docs.go`中调用

```go
package app

import (
	user "github.com/tsingshaner/gin-starter/internal/user/constant"
	"github.com/tsingshaner/gin-starter/pkg/gen"
)

func GetCodeMaps() []gen.CodeMap {
	return []gen.CodeMap{
		user.GetCodeMap(),
	}
}

```

### 文档导入

正式文档使用 `APIFox`。正式文档在本地代码推送后从 `github pages` 定时同步，**不要**将生成的文档上传到 `git`, 文档从 `CI` 自动生成。

开发环境在开启本地服务器后访问`http://127.0.0.1:4936/open-api/doc.json`应会看到 `openapi` 格式的文档，将其导入到本地的 `APIFox` 项目中即可，在本地开发环境为避免污染正式文档，建议自行建一个 `APIFox` 项目用于调试。

**更改接口名称导入后不会删除原有接口，需要手动删除。**

`APIFox` -> 项目设置 -> 数据管理/导入数据 -> 定时导入 -> 新建 -> 手动触发 + `openapi(swagger)` + 文档地址

### 自定义结构体说明

`swag` 可以将自定义结构体转换为文档，无需去 `APIFox` 中手动添加。

```go
// swag 可以读取字段标签 example 作为示例, binding 标签作为验证规则, 注释作为说明
type RegisterReqBody struct {
	UserID   string `json:"userID" binding:"required" example:"20240126"`                // 学号或者工号
	Username string `json:"username" binding:"required" example:"杏鸣"`                    // 姓名
	Password string `json:"password" binding:"required,max=64,min=8" example:"12345678"` // 密码
}

var RegisterValidator = middleware.Validator(&middleware.ValidatorOptions{
	Body: &RegisterReqBody{},
})

// swag当前不支持泛型，需要手写，不可使用 resp.FailedRespBody[string] 来声明
type RegisterFailedRespBody struct {
	resp.BaseRespBody
	Err string `json:"err"  binding:"required" example:"错误信息"`
}

// Register
//
//	@Summary				注册
//	@Description.markdown	Register
//	@Tags					user
//	@Accept					json
//	@Produce				json
//	@Param					Body	body		RegisterReqBody	true
//	@Success				201		{object}	resp.BaseRespBody
//	@Failure				400		{object}	RegisterFailedRespBody
//	@Router					/v1/user/ [post]
func (h *Handler) Register(c *gin.Context) {}
```

## 参数校验

使用 `middleware.Validator` 中间件进行参数校验。 在每个handler中声明一个验证器，例如：

```go
type RegisterReqBody struct {
	UserID   string `json:"userID" binding:"required" example:"20240126"`                // 学号或者工号
	Username string `json:"username" binding:"required" example:"杏鸣"`                    // 姓名
	Password string `json:"password" binding:"required,max=64,min=8" example:"12345678"` // 密码
}

// 创建验证器
var RegisterValidator = middleware.Validator(&middleware.ValidatorOptions{
	Body: &RegisterReqBody{},
})

// ...

func (h *Handler) Register(c *gin.Context) {
  // 通过验证器验中间件才会到当前handler，故使用 c.MustGet 获取请求参数，默认key为"body"，可在ValidatorOptions中修改BodyKey
	user := c.MustGet(middleware.KeyBody).(*RegisterReqBody)
  // ...
}
```
不要忘记在路由中注册验证器。
```go
func (u *UserRoutes) Register(r *gin.RouterGroup) {
	user := r.Group("/user/")

	user.DELETE(":id", u.DeleteUser)
	user.GET("", u.GetUserList)
	user.GET(":id", u.GetUserProfile)
	user.PATCH(":id", u.UpdateUser)
  // 这里 Register 方法名冲突了，所以使用 u.UserHandler.Register
	user.POST("", handler.RegisterValidator, u.UserHandler.Register)
	user.POST("login", handler.LoginValidator, u.Login)
	user.PUT(":id", u.UpdateUser)
}

```
