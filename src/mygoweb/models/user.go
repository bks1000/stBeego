package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id      string `orm:"pk;column(id)"`
	Name    string
	Age     int
	Address string
	Phone   string `orm:"size(15)"`
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL) //注册驱动
	/*
			第一个参数是数据库的别名，用来切换数据库使用。
		    第二个是driverName，在RegisterDriver时注册的
			第三数据库连接字符串: test:123456@/test?charset=utf8 <=> 用户名:密码@数据库地址+名称?字符集 (test数据库提前建好)
			第四个参数相当于:
		    orm.SetMaxIdleConns("default", 30)//设置数据库的最大空闲连接。
			第五个参数相当于：
		    orm.SetMaxOpenConns("default", 30)//设置数据库的最大数据库连接。
	*/
	orm.RegisterDataBase("default", "mysql",
		"root:AT*-=uik123@tcp(localhost:3306)/test?charset=utf8", 30, 30) //注册默认数据库
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User)) //注册 model
	//orm.RunSyncdb("default", false, true)
}

func UserQuery(paras ...interface{}) ([]User, error) {
	//o orm.Ormer
	o := orm.NewOrm()
	sql := "select * from user where name=?"
	// query
	var users []User
	num, err := o.Raw(sql, paras).QueryRows(&users)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	return users, err
}

/*
	注意，如果要实现事务，o对象最好从外面传入，可以控制o.Begin()/o.Rollback()/o.Commit()

	// insert
 	o.Begin()
    id, err := models.UserAdd(o,user)
    fmt.Printf("ID: %d, ERR: %v\n", id, err)
    if err != nil {
        o.Rollback()
        fmt.Println("插入user表出错,事务回滚")
    } else {
        o.Commit()
        fmt.Println("插入user表成功,事务提交")
    }
*/
func UserAdd(user User) (int64, error) {
	// insert
	o := orm.NewOrm()
	return o.Insert(&user)
}

func UserCut(user User) (int64, error) {
	// delete
	o := orm.NewOrm()
	return o.Delete(&user, "name")
}

func UserUpdate(user User) (int64, error) {
	// update
	o := orm.NewOrm()
	return o.Update(&user, "name")
}
