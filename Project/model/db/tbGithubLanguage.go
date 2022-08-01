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
