package main

import (
        _ "github.com/romans1212notes/beego-mysql-api/routers"

        "github.com/astaxie/beego"
        "github.com/astaxie/beego/orm"
        _ "github.com/go-sql-driver/mysql"
        "time"
)

func init() {
        orm.RegisterDriver("mysql", orm.DRMySQL)
        // set default database
        orm.RegisterDataBase("default", "mysql", "root:<my password>@/testdb?charset=utf8", 30)
        // Set to UTC time
        orm.DefaultTimeLoc = time.UTC
}

func main() {
        if beego.BConfig.RunMode == "dev" {
                beego.BConfig.WebConfig.DirectoryIndex = true
                beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
        }
        beego.Run()
}
