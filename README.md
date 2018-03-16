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

## 三、数据库ClearDb mysql
### 1. 代码准备

### 2. Heroku创建数据库
```
heroku addons:create cleardb:ignite

heroku config | grep CLEARDB_DATABASE_URL

heroku config:set DB_USERNAME=b8564e6d9298c9
heroku config:set DB_PASSWORD=password
heroku config:set DB_HOST=us-cdbr-iron-east-05.cleardb.net
heroku config:set DB_DATABASE=heroku_2a001fea285944b

```
程序会将上述环境变量，拼出来DATABASE_URL结构如下：
DATABASE_URL='user:pass@tcp(us-cdbr-iron-east-05.cleardb.net:3306)/your_heroku_database'

### 3. 链接数据库，创建初始数据
用上述DATABASE_UR中的`user:pass@host/databse`连接到数据库，手动创建Table `Task`
```sql

CREATE TABLE `task` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(30) NOT NULL,
  `done` TINYINT(1) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `task` VALUES (1, 'GO', 0);
```

4. 获取初始数据
`https://zy-go-todo.herokuapp.com/task/1`

## 四、问题
Beego & Angular JS ajax Redirect
* 问题描述
点击【退出】之后，页面并未进行跳转

* 具体细节
** 点击【退出】，路由跳转`/logout`
** 随后会重定向到`/login`
** task.js里面打印的data可以看出已经是`login.html`的页面代码

* 原因
ajax无法获得控制页面跳转的参数，可以参见(beego的局限性和解决登录跳转问题)[http://blog.csdn.net/github_37320188/article/details/79107380]

* 暂定措施
改成`<form>`提交，可以参见(beego redirect跳转问题)[https://github.com/astaxie/beego/issues/2565]

## 参考
* [Building Web Apps with Go](https://www.gitbook.com/book/codegangsta/building-web-apps-with-go/details)
* [Getting Started with Go on Heroku](https://github.com/heroku/go-getting-started)
* [The Vendor Tool for Go](https://github.com/kardianos/govendor)
* [Heroku cleardb connection default addr for network unknown
](https://stackoverflow.com/questions/43562091/heroku-cleardb-connection-default-addr-for-network-unknown)