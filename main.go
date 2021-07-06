package dbobj

import (
	mysqlOrm "dbobj/mysql"
	"github.com/go-xorm/xorm"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"log"
)

var(
	MysqlEngin *xorm.Engine
	err			error
)

func init()  {
	err := godotenv.Load("./app.conf")
	if err != nil {
		log.Fatal(err)
	}
	MysqlEngin, err = mysqlOrm.InitConnect()
}
