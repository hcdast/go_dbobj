package mysqlOrm

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"os"
	"strconv"
	"time"
)

var (
	DataSourceName string
	DriverName	string
	MaxOpenConns int
	MaxIdleConns int
	Location  	 string
	Err			error
	cstZone = time.FixedZone("CST", 8*3600)
	MysqlEngin *xorm.Engine
	NewEngine = Engine{MaxOpenConns: 1000, MaxIdleConns: 1000, Location: "Asia/Shanghai"}
)

type Engine struct {
	Engine *xorm.Engine
	MaxOpenConns int 	// 最大打开连接数
	MaxIdleConns int 	// 连接池的空闲数大小
	Location	 string // 时区
	State        bool   // 链接状态
}

type ShortEngine struct {
	Host 		string  // ip
	Port 	   	string  // 端口
	User 	   	string  // 用户账号
	Pwd 		string  // 密码
	Charset 	string  // 编码
	DriverName 	string  // mysql
	DbName		string  // 数据库名字
}

func (e *Engine)createEngine() (engine *xorm.Engine, err error) {
	engine, err = xorm.NewEngine(DriverName, DataSourceName)
	fmt.Println(engine, err)
	if err != nil {
		return nil, err
	}

	pingState := make(chan bool)

	defer close(pingState)
	go func() {
		if err := engine.Ping(); err != nil {
			fmt.Println("connection db error --> ", err.Error())
		}
		pingState <- true
	}()

	t := time.AfterFunc(5 * time.Second, func() {
		pingState <- false
	})

	select {
	case state := <-pingState:
		if state == false {
			return nil, errors.New("connection db error")
		} else {
			t.Stop()
			goto END
		}
	}

	END:
	engine.ShowSQL(true)
	//engine.SetMaxOpenConns(1000)
	engine.SetMaxOpenConns(e.MaxOpenConns)
	engine.SetMaxIdleConns(e.MaxIdleConns)

	// 设置时区
	engine.DatabaseTZ = cstZone // 必须
	engine.TZLocation = cstZone // 必须
	e.State = true

	if err != nil {
		fmt.Println("set orm engine location err --> ", err.Error())
	}

	return engine, nil
}

func InitConnect()(engine *xorm.Engine, err error){
	time.Local = cstZone
	DataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&loc=%s",
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"), os.Getenv("DB_NAME"), os.Getenv("DB_CHARSET"), "Asia%2FShanghai")
	DriverName = os.Getenv("Driver_Name")

	MaxOpenConnsStr := os.Getenv("MaxOpenConns")
	if MaxOpenConnsStr != "" {
		if MaxOpenConns,Err = strconv.Atoi(MaxOpenConnsStr); Err == nil {
			NewEngine.MaxOpenConns = MaxOpenConns
		}
	}

	MaxIdleConnsStr := os.Getenv("MaxIdleConns")
	if MaxIdleConnsStr != "" {
		if MaxIdleConns,Err = strconv.Atoi(MaxIdleConnsStr); Err == nil {
			NewEngine.MaxOpenConns = MaxIdleConns
		}
	}

	Location = os.Getenv("Location")
	if Location != "" {
		NewEngine.Location = Location
	}

	MysqlEngin, Err = NewEngine.createEngine()
	if Err != nil {
		return nil, errors.New("mysql connection db error")
	}
	return MysqlEngin, nil

}