package main

import (
	"fmt"
	"github.com/kataras/iris"
    "github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"encoding/json"
)

type User struct {
	Name string `json:"name"`
	Description string `json:"Description"`
}

func main() {
		app := iris.New();
		app.Use(logger.New());
		app.Use(recover.New());

		db, err := sql.Open("mysql", "root:abc123@/dbname?charset=utf8");
		if (err != nil) {
			fmt.Println("err: ", err);
			panic(err);
		}
		defer db.Close();

		app.Get("/", func(ctx iris.Context) {
			m := make(map[string]string);
			statement := "SELECT * from user;";
			rows, err := db.Query(statement);
			if (err != nil) {
				fmt.Println("err: ", err);
				panic(err);
			}
			for rows.Next() {
				var Name string;
				var Description string;
				err = rows.Scan(&Name, &Description);
				m[Name] = Description;
			}
			b, err := json.Marshal(m);
			fmt.Println("b = ", string(b));
			ctx.Writef(string(b));
		});

		app.Post("/", func(ctx iris.Context) {
			var user []User;
			ctx.ReadJSON(&user);
			res, err := db.Prepare("insert into user (Name, Description) values (?,?);");
			if (err != nil) {
				fmt.Println("err: ", err);
				panic(err);
			}
			res.Exec(user[0].Name, user[0].Description);
			defer res.Close();
			ctx.Writef("Method: %s and path: %s", ctx.Method(), ctx.Path());
		})

		app.Put("/{Descr: string}", func(ctx iris.Context) {
			Descr := ctx.Params().Get("Descr");
			var user []User;
			fmt.Println("Descr: ", Descr);
			err := ctx.ReadJSON(&user);
			if (err != nil) {
				ctx.JSON(iris.Map{"response: ": err.Error()});
			} else {
				res, err := db.Prepare("UPDATE user SET Description=? WHERE Name=?");
				if (err != nil) {
					fmt.Println("err: ", err);
					panic(err);
				}
				res.Exec(Descr, user[0].Name);
				defer res.Close();
				ctx.Writef("Method: %s and path: %s", ctx.Method(), ctx.Path());
			}
		})

		app.Delete("/", func(ctx iris.Context) {
			res, err := db.Query("delete from user;");
			if (err != nil) {
				fmt.Println("err: ", err);
				panic(err);
			}
			fmt.Println("res:",res);
			ctx.Writef("Method: %s and path: %s", ctx.Method(), ctx.Path());
		})

		fmt.Println("Server running on http://localhost:8099");
		app.Run(iris.Addr(":8099"));
}