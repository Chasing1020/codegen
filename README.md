# CodeGen

CodeGen是一个基于Go实现的后端代码生成器

允许使用yaml定义数据结构自动生成RESTful接口

支持MySQL与Redis。

1.快速开始

本项目需要go语言环境
```shell
git clone https://github.com/Chasing1020/codegen.git
go mod tidy
go build
```
即安装成功

2. 使用方式

修改项目的`conf/config.yaml`文件，参考示例，修改为期望的数据类型，`go run main.go`运行程序
当出现以下提示信息后

(base) \$ go run main.go

2022/04/07 22:32:09 code generation finished

2022/04/07 22:32:12 go mod tidy finished

2022/04/07 22:32:12 use command cd dist && go run main.go to start

在dist目录下即为生成的项目，运行项目后，会自动生成结构体对应数据库的表

3. 文档生成

可以使用`go get -u github.com/swaggo/swag/cmd/swag`下载Swagger，切换至生成的项目根目录，将`./dist/router/router.go`下的import取消注释，输入命令
```shell
swag init
go mod tidy
```
运行项目后访问 http://localhost:8080/swagger/index.html 即可


4. 待改进

支持消息队列，JWT，参数校验规则等。

5. 项目依赖

```
github.com/gin-gonic/gin
github.com/go-redis/redis/v8
github.com/go-sql-driver/mysql
gopkg.in/yaml.v2
gorm.io/driver/mysql
gorm.io/gorm
```