package db

func NewGithubLanguage() *GithubLanguage {
	_self := new(GithubLanguage)
	return _self
}

type GithubLanguage struct {
	BaseHeader `xorm:"extends"`
	Name       string `xorm:"varchar(50) default '' comment('标题')" json:"name"`
	Type       int    `xorm:"int(5) not null default 0 index comment('类型')" json:"type"`
	IsEnable   int    `xorm:"int(5) not null default 2 index comment('权限')" json:"is_enable"` // 1启用  2禁用
	ModeFooter `xorm:"extends"`
}

func (that *GithubLanguage) Add(data interface{}) (int64, error) {
	return engine.Insert(data)
}

func (that *GithubLanguage) List(page, size int) ([]*GithubLanguage, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 500
	}
	_start := (page - 1) * size
	_dataList := make([]*GithubLanguage, 0)
	engine.Limit(size, _start)
	if _err := engine.Desc("id").Find(&_dataList); _err != nil {
		return nil, _err
	}
	return _dataList, nil
}

func (that *GithubLanguage) SearchAll() ([]*GithubLanguage, error) {
	_dataList := make([]*GithubLanguage, 0)
	if _err := engine.Desc("id").Where("is_enable = ?", 1).Find(&_dataList); _err != nil {
		return nil, _err
	}
	return _dataList, nil
}
