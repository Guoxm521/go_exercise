package db

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"xorm.io/xorm"
)

var (
	initStatus bool
	engine     *xorm.Engine
)

//SetORM ...
func SetORM(xormEngine *xorm.Engine) {
	engine = xormEngine
	Init()
}

func Init() {
	if initStatus {
		return
	}
	_slice := []interface{}{
		new(Account),
		new(GithubTrending),
		new(GithubLanguage),
		new(GithubSince),
		new(Message),
	}
	if err := engine.Sync2(_slice...); err != nil {
		panic("db init fail!" + err.Error())
	}
	//同步字段类型、注释，默认关闭
	_syncComment := false
	if _syncComment {
		_regComment := regexp.MustCompile(`comment\('([^']+)'\)`)
		_regType := regexp.MustCompile(`[\s]*((int|smallint|varchar)\([\d]+\))\s+`)
		_notNull := regexp.MustCompile(`not[\s]+null`)
		_default := regexp.MustCompile(`default[\s]+([^\s]+)\s+`)
		_session := engine.NewSession()
		defer _session.Close()
		for _, _v := range _slice {
			_t := reflect.TypeOf(_v).Elem()
			fieldNum := _t.NumField()
			k := 0
			for i := 0; i < fieldNum; i++ {
				if k > 1000 {
					break
				}
				k = k + 1
				_f := _t.Field(i)
				_xorm := _f.Tag.Get("xorm")
				if _xorm == "extends" {

				}
				_matchesComment := _regComment.FindAllStringSubmatch(_xorm, -1)
				_matchesType := _regType.FindAllStringSubmatch(_xorm, -1)
				if len(_matchesComment) > 0 && len(_matchesType) > 0 {
					_sqlSlice := []string{
						"alter table",
						"`" + _t.Name() + "`",
						"modify column",
						"`" + _f.Name + "`",
						_matchesType[0][1],
					}
					//
					_matchesNotNull := _notNull.FindAllStringSubmatch(_xorm, -1)
					if len(_matchesNotNull) > 0 {
						_sqlSlice = append(_sqlSlice, "not null")
					}
					_matchesDefault := _default.FindAllStringSubmatch(_xorm, -1)
					if len(_matchesDefault) > 0 {
						_sqlSlice = append(_sqlSlice, []string{
							"default",
							_matchesDefault[0][1],
						}...)
					}
					//
					_sqlSlice = append(_sqlSlice, []string{
						"comment",
						"'" + _matchesComment[0][1] + "'",
					}...)
					_sql := strings.Join(_sqlSlice, " ")
					_, _err := _session.Exec(_sql)
					fmt.Println(_sql, _err)
				}
			}
		}
	}
	initStatus = true
}
