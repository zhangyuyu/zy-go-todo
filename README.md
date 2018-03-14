Golang ToDo web application on Heroku.
--- 

## 一、关键字
* go
* govender
* heroku
* web
* todo

## 二、部署
### 1. 准备项目
* 准备项目源码
* 获取环境变量`PORT`  
  Heroku给我们了一个PORT环境变量，并期望我们的web应用绑定到该端口
    
    ```
    port := os.Getenv("PORT")
        if port == "" {
            port = "9090"
        }
    ```

### 2. 准备Procfile
  在文件Procfile中，配置运行的应用程序，例如：

    ```
    web: hello
    ```

### 3. govendor依赖管理
* 安装govendor
* 编写vendor.json
* 添加`GOPATH`下的相关依赖到vendor目录下，例如：
    ```
    govendor add github.com/astaxie/beego
    govendor add github.com/astaxie/beego/config
    govendor add github.com/astaxie/beego/context
    govendor add github.com/astaxie/beego/context/param
    govendor add github.com/astaxie/beego/logs
    govendor add github.com/astaxie/beego/toolbox
    govendor add github.com/astaxie/beego/utils
    govendor add github.com/astaxie/beego/session
    govendor add github.com/astaxie/beego/grace
    
    ```

### 4. 创建Heroku应用
* 登录
    ```
    heroku login
    ```
* 创建
    ```
    heroku create
    ```
    Heroku会创建一个随机的名字，可以自己在website上修改，例如修改为`zy-go-hello`
* 设置Go buildpack
    ```
    heroku git:remote -a zy-go-todo
    ```
    
### 5. 部署
    ```
    git push heroku master
    ```

### 6. 访问
* 通过命令行，访问web应用
    ```
    heroku open
    ```

* 浏览器直接输入，访问web应用
    ```
    https://zy-go-todo.herokuapp.com/
    ```

* 查看heroku应用状况
    ```
    https://dashboard.heroku.com/apps/zy-go-todo
    ```

## 参考
* [Building Web Apps with Go](https://www.gitbook.com/book/codegangsta/building-web-apps-with-go/details)
* [Getting Started with Go on Heroku](https://github.com/heroku/go-getting-started)
* [The Vendor Tool for Go](https://github.com/kardianos/govendor)