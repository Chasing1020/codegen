# CodeGen

CodeGen是一个基于Go实现的后端代码生成器，并支持生成文档与调试接口。

只需要`更改yaml定义的数据结构`，即可自动生成对应的所有RESTful接口与文档。

## 1. 使用方式
本项目需要 Go 语言环境。
1. 修改当前项目下的 `conf/config.yaml` 文件（当前项目以普通的选课系统作为作为示例，包括：课程表，选课表，教师表，学院表，学生表等），改为期望的数据表。并修改好项目名称，MySQL和Redis等必要配置信息。

2. 修改好配置文件后在当前项目根目录下执行 `go run main.go` ，至出现 `Codegen Success!` 提示

3. 运行成功后，在 `dist` 目录下即为生成的项目。

4. `cd dist` 进入生成好的项目目录，使用 `make fmt` 格式化并下载相关的项目依赖

5. 使用 `make run` 即可启动项目

## 2. 文档生成

运行项目后访问 `http://localhost:8080/swagger/index.html` 即可，可见如下页面：
<img src="https://github.com/Chasing1020/codegen/assets/swagger_index.png">

在当前页面下可以直接执行请求，此外，也可以通过PostMan导入。链接为`http://localhost:8080/swagger/doc.json`，示意图如下：

<img src="https://github.com/Chasing1020/codegen/assets/postman_import.png">

导入成功后，即可查看所有接口与测试数据：

<img src="https://github.com/Chasing1020/codegen/assets/postman_demo.png">

## 3. 待改进

支持消息队列，JWT，参数校验规则等。

## 4. 关于
这个项目的灵感来自各种数据库相关的大作业，基于`避免重复造轮子`的原则，简化了新建项目的成本，并增加了对数据库与缓存的支持。

## 5. 许可
本项目基于 2.0 版本的 APACHE 许可证，链接：http://www.apache.org/licenses/LICENSE-2.0