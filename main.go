package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"log"
)

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gradleapk?charset=utf8")
	if err != nil {
		log.Fatal("connect mysql fail ! [%s]", err)
	} else {
		fmt.Println("connect to mysql success")
	}

	app := iris.Default()

	// 获取所有的客户度
	app.Get("/clientids", func(ctx context.Context) {

		sql := "select * from clientid "

		println(sql)
		rows, err := db.Query(sql)
		if err != nil {
			fmt.Printf("select fail [%s]", err)
		}
		var _id int
		var _clientid string
		clientids := make([]string, 0)
		for rows.Next() {
			rows.Columns()
			err := rows.Scan(
				&_id,
				&_clientid,
			)
			if err != nil {
				fmt.Printf("get user info error [%s]", err)
			}
			clientids = append(clientids, _clientid)
		}
		ctx.JSON(clientids)
	})

	// 获取客户端的所有渠道

	app.Get("/channels", func(ctx context.Context) {
		clientid := ctx.URLParam("clientid")
		if "" == clientid {
			ctx.JSON("err clientid null")
			return
		}

		sql := "select _id,channel from client_channel where clientid='" + clientid + "'"

		println(sql)

		rows, err := db.Query(sql)
		if err != nil {
			fmt.Printf("select fail [%s]", err)
		}
		var _id int
		var channel string

		channels := make([]string, 0)

		for rows.Next() {
			rows.Columns()
			err := rows.Scan(
				&_id,
				&channel,
			)
			if err != nil {
				fmt.Printf("get user info error [%s]", err)
			}
			channels = append(channels, channel)
		}
		ctx.JSON(channels)
	})

	// 获取打包信息
	app.Get("/apkinfo", func(ctx iris.Context) {
		clientid := ctx.URLParam("clientid")
		if "" == clientid {
			ctx.JSON("err clientid null")
			return
		}

		channel := ctx.URLParam("channel")
		if "" == channel {
			ctx.JSON("err channel null")
			return
		}

		println(clientid)
		println(channel)

		sql := "select * from info where clientid= '" + clientid + "' and channel = '" + channel + "'"

		println(sql)
		rows, err := db.Query(sql)
		if err != nil {
			fmt.Printf("select fail [%s]", err)
		}

		var _id int
		var _clientid string
		var _channel string
		var applicationid string
		var versionname string
		var maxsdkversion string
		var minsdkversion string
		var targetsdkversion string
		var versioncode int
		var maxsdkversion_enable bool
		for rows.Next() {

			rows.Columns()
			err := rows.Scan(
				&_id,
				&_clientid,
				&_channel,
				&applicationid,
				&versionname,
				&maxsdkversion,
				&minsdkversion,
				&targetsdkversion,
				&versioncode,
				&maxsdkversion_enable,
			)
			if err != nil {
				fmt.Printf("get user info error [%s]", err)
			}
		}

		ctx.JSON(iris.Map{
			"applicationid":        applicationid,
			"versioncode":          versioncode,
			"versionname":          versionname,
			"minsdkversion":        minsdkversion,
			"targetsdkversion":     targetsdkversion,
			"maxsdkversion":        maxsdkversion,
			"maxsdkversion_enable": maxsdkversion_enable,
		})
	})

	app.Post("changeinfo", func(ctx context.Context) {

		clientid := ctx.URLParam("clientid")

		if "" == clientid {
			ctx.JSON("err clientid null")
			return
		}

		channel := ctx.URLParam("channel")

		if "" == channel {
			ctx.JSON("err channel null")
			return
		}

		applicationid := ctx.URLParam("applicationid")
		versionname := ctx.URLParam("versionname")
		versioncode := ctx.URLParam("versioncode")
		minsdkversion := ctx.URLParam("minsdkversion")
		targetsdkversion := ctx.URLParam("targetsdkversion")
		maxsdkversion := ctx.URLParam("maxsdkversion")
		maxsdkversioneenable := ctx.URLParam("maxsdkversioneenable")

		println(clientid)
		println(channel)
		println(applicationid)
		println(versionname)
		println(versioncode)
		println(minsdkversion)
		println(targetsdkversion)
		println(maxsdkversion)
		println(maxsdkversioneenable)

		if "" != applicationid {

			sql := "update info set applicationid = '" + applicationid + "' where clientid= '" + clientid + "' and channel = '" + channel + "'"
			println(sql)
			db.Query(sql)

		}
		if "" != versionname {
			sql := "update info set versionname = '" + versionname + "' where clientid= '" + clientid + "' and channel = '" + channel + "'"
			println(sql)
			db.Query(sql)
		}
		if "" != versioncode {
			sql := "update info set versioncode = '" + versioncode + "' where clientid= '" + clientid + "' and channel = '" + channel + "'"
			println(sql)
			db.Query(sql)
		}
		if "" != minsdkversion {
			sql := "update info set minsdkversion = '" + minsdkversion + "' where clientid= '" + clientid + "' and channel = '" + channel + "'"
			println(sql)
			db.Query(sql)
		}
		if "" != targetsdkversion {
			sql := "update info set targetsdkversion = '" + targetsdkversion + "' where clientid= '" + clientid + "' and channel = '" + channel + "'"
			println(sql)
			db.Query(sql)
		}
		if "" != maxsdkversion {
			sql := "update info set maxsdkversion = '" + maxsdkversion + "' where clientid= '" + clientid + "' and channel = '" + channel + "'"
			println(sql)
			db.Query(sql)
		}
		if "" != maxsdkversioneenable {
			sql := "update info set maxsdkversioneenable = '" + maxsdkversioneenable + "' where clientid= '" + clientid + "' and channel = '" + channel + "'"
			println(sql)
			db.Query(sql)
		}

		ctx.JSON(iris.Map{
			"status": "success",
		})
	})

	app.Run(iris.Addr(":80"))

}
