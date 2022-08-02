package db

func NewGithubLanguage() *GithubTrending {
	_self := new(GithubTrending)
	return _self
}

type GithubLanguage struct {
	BaseHeader `xorm:"extends"`
	Name       string `xorm:"varchar(50) default '' comment('标题')" json:"name"`
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
		size = 50
	}
	_start := (page - 1) * size
	_dataList := make([]*GithubLanguage, 0)
	engine.Limit(size, _start)
	if _err := engine.Desc("id").Find(&_dataList); _err != nil {
		return nil, _err
	}
	return _dataList, nil
}
