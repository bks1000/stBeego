#本项目采用bee生成的基于beego框架的项目

#beego入门文档
https://my.oschina.net/astaxie/blog/124040  


#golang mysql 
安装golang mysql 驱动
C:\Users\lenovo>go get github.com/go-sql-driver/mysql
安装beego orm 
C:\Users\lenovo>go get github.com/astaxie/beego/orm

https://blog.csdn.net/z1134145881/article/details/52180550    orm
https://blog.csdn.net/lengyuezuixue/article/details/79148762


#golang jwt (json web token)
go get github.com/codegangsta/negroni 
Idiomatic HTTP Middleware for Golang 
http的一个中间件

go get github.com/dgrijalva/jwt-go 
Golang implementation of JSON Web Tokens (JWT)

go get github.com/dgrijalva/jwt-go/request
https://blog.csdn.net/wangshubo1989/article/details/74529333


go高级话题
http://www.360doc.com/content/17/1127/09/7534118_707483730.shtml
https://www.cnblogs.com/yinzhengjie/p/7079626.html

golang的命名规范及大小写的访问权限
1.golang的命名需要使用驼峰命名法，且不能出现下划线
2.golang中根据首字母的大小写来确定可以访问的权限。无论是方法名、常量、变量名还是结构体的名称，如果首字母大写，则可以被其他的包访问；如果首字母小写，则只能在本包中使用。
可以简单的理解成，首字母大写是公有的，首字母小写是私有的
3、结构体中属性名的大写
如果属性名小写则在数据解析（如json解析,或将结构体作为请求或访问参数）时无法解析


从零开始搭建一个简单的基于webpack的vue开发环境
https://segmentfault.com/a/1190000012789253

