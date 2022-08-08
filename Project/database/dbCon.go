package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"runtime"
	"strconv"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var OrmEngine *xorm.Engine

//serverModel get server information from config.yml
type MysqlModel struct {
	Host        string `yaml:"host"`          // 主机
	Port        string `yaml:"port"`          // 端口
	User        string `yaml:"user"`          // 用户名
	Password    string `yaml:"password"`      // 密码
	Dbname      string `yaml:"dbname"`        // 数据库名
	Prefix      string `yaml:"prefix"`        // 前缀
	MaxIdleConn int    `yaml:"max_idle_conn"` // 最大连接数
	MaxOpenConn int    `yaml:"max_open_conn"` // 设置最大打开连接数
	ShowSql     bool   `yaml:"show_sql"`      // 控制台是否打印SQL语句
}

func ConnMySQL(configInfo *MysqlModel) {
	if configInfo == nil {
		fmt.Println("mysql配置文件初始化失败")
		return
	}
	host := configInfo.Host
	port := configInfo.Port
	user := configInfo.User
	password := configInfo.Password
	dbName := configInfo.Dbname
	prefix := configInfo.Prefix
	maxIdleConn := configInfo.MaxIdleConn
	maxOpenConn := configInfo.MaxOpenConn
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", user, password, host, port, dbName)
	_orm, _err := xorm.NewEngine("mysql", mysqlInfo)
	if _err != nil {
		_, file, line, _ := runtime.Caller(1)
		println(file+":"+strconv.Itoa(line), "MySql 数据库连接失败：%s", _err.Error())
		return
	}
	OrmEngine = _orm
	//连接测试
	if err := OrmEngine.Ping(); err != nil {
		_, file, line, _ := runtime.Caller(1)
		println(file+":"+strconv.Itoa(line), "MySql 数据库连接测试：%s", err.Error())
		return
	}
	//日志打印SQL
	if configInfo.ShowSql {
		OrmEngine.ShowSQL(true)
	}
	//设置连接池的空闲数大小
	OrmEngine.SetMaxIdleConns(maxIdleConn)
	//设置最大打开连接数
	OrmEngine.SetMaxOpenConns(maxOpenConn)
	//名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射
	OrmEngine.SetTableMapper(names.SnakeMapper{})
	//orm.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	tbMapper := names.NewPrefixMapper(names.SnakeMapper{}, prefix)
	OrmEngine.SetTableMapper(tbMapper)
	//连接成功
	_, file, line, _ := runtime.Caller(0)
	println(file+":"+strconv.Itoa(line), "Mysql %v", "连接成功")
	return
}
