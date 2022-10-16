package db

func NewMessage() *Message {
	_self := new(Message)
	return _self
}

type Message struct {
	BaseHeader `xorm:"extends"`
	UId        string `xorm:"varchar(40) default '' comment('账号id')" json:"uid"`
	ToUid      string `xorm:"varchar(40) default ''  comment('私聊账号id')" json:"to_uid"`
	Content    string `xorm:"text  comment('聊天记录')" json:"content"`
	Group      string `xorm:"varchar(10) default '' comment('房间id')" json:"group"`
	ModeFooter `xorm:"extends"`
}

func (that *Message) Add(data interface{}) (int64, error) {
	return engine.Insert(data)
}

func (that *Message) List(page, size int, search interface{}) ([]*Message, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 100
	}
	_start := (page - 1) * size
	_dataList := make([]*Message, 0)
	_group := search.(map[string]interface{})["group"].(string)
	_uid := search.(map[string]interface{})["uid"].(string)
	_session := engine.NewSession()
	defer _session.Close()
	if _group != "" {
		_session.Where("`group` = ?", _group)
	}
	if _uid != "" {
		_session.Where("`uid` = ?", _uid)
	}
	if _err := _session.Limit(size, _start).Asc("c_time").Find(&_dataList); _err != nil {
		return nil, _err
	}
	return _dataList, nil
}

func (that *Message) Count(search interface{}) (int64, error) {
	_session := engine.NewSession()
	defer _session.Close()
	_group := search.(map[string]interface{})["group"].(string)
	_uid := search.(map[string]interface{})["uid"].(string)
	if _group != "" {
		_session.Where("`group` = ?", _group)
	}
	if _uid != "" {
		_session.Where("`uid` = ?", _uid)
	}
	return _session.Count(that)
}
