package db

func NewGithubSince() *GithubSince {
	_self := new(GithubSince)
	return _self
}

type GithubSince struct {
	BaseHeader `xorm:"extends"`
	Name       string `xorm:"varchar(50) default '' comment('标题')" json:"name"`
	Type       string `xorm:"varchar(50) not default '' index comment('类型')" json:"type"`
	SinceType  int    `xorm:"int(5) not null default 0 index comment('类型字段')" json:"since_type"`
	ModeFooter `xorm:"extends"`
}

func (that *GithubSince) Add(data interface{}) (int64, error) {
	return engine.Insert(data)
}

func (that *GithubSince) List(page, size int) ([]*GithubSince, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 50
	}
	_start := (page - 1) * size
	_dataList := make([]*GithubSince, 0)
	engine.Limit(size, _start)
	if _err := engine.Desc("id").Find(&_dataList); _err != nil {
		return nil, _err
	}
	return _dataList, nil
}

func (that *GithubSince) SearchAll() ([]*GithubSince, error) {
	_dataList := make([]*GithubSince, 0)
	if _err := engine.Desc("id").Find(&_dataList); _err != nil {
		return nil, _err
	}
	return _dataList, nil
}
