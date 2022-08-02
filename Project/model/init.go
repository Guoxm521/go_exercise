package model

import "reflect"

type Logic struct{}

func NewLogic() *Logic {
	_logic := &Logic{}
	return _logic
}

//设置表模型属性
func (that *Logic) parseTableField(tableMode interface{}, name string, value interface{}) {
	v := reflect.ValueOf(tableMode)
	if !v.IsNil() {
		t := reflect.TypeOf(tableMode)
		v := reflect.ValueOf(tableMode)
		fieldNum := t.Elem().NumField()
		for i := 0; i < fieldNum; i++ {
			_name := t.Elem().Field(i).Name
			if _fieldName := t.Elem().Field(i).Tag.Get("json"); _fieldName != "" {
				if _fieldName == name {
					_t := v.Elem().Field(i).Kind().String()
					switch _t {
					case "string":
						if _b, _ok := value.(string); _ok {
							v.Elem().FieldByName(_name).SetString(_b)
						}
					case "bool":
						if _b, _ok := value.(bool); _ok {
							v.Elem().FieldByName(_name).SetBool(_b)
						}
					case "int", "int8", "int16", "int32", "int64", "bigint":
						//v.Elem().FieldByName(_name).SetInt(utils.ToInt64(value))
					default:
						//待补充
					}
				}
			}
		}
	}
}
