package db

func NewGithubTrending() *GithubTrending {
	_self := new(GithubTrending)
	return _self
}

type GithubTrending struct {
	BaseHeader   `xorm:"extends"`
	Author       string `xorm:"varchar(100) default '' comment('作者')" json:"author"`
	Repo         string `xorm:"varchar(100) default '' comment('项目仓库')" json:"repo"`
	Url          string `xorm:"varchar(255) default '' comment('项目链接')" json:"url"`
	Desc         string `xorm:"varchar(255) default '' comment('简介')" json:"desc"`
	Starts       string `xorm:"varchar(100) default '' comment('目前start数')" json:"starts"`
	Forks        string `xorm:"varchar(100) default '' comment('目前start数')" json:"forks"`
	Language     string `xorm:"varchar(100) default '' comment('语言')" json:"language"`
	AddedStars   string `xorm:"varchar(100) default '' comment('今天或者这周或者这个月的starts数')" json:"added_stars"`
	Avatars      string `xorm:"text default '' comment('项目贡献者的头像地址集合')" json:"avatars"`
	LanguageType int    `xorm:"int(2) index default 0  comment('语言类型')" json:"language_type"`
	SinceType    int    `xorm:"int(2) index default 0 comment('搜索时间')" json:"since_type"`
	ModeFooter   `xorm:"extends"`
}

func (that *GithubTrending) Add(data interface{}) (int64, error) {
	return engine.Insert(data)
}

func (that *GithubTrending) List(page, size int) ([]*GithubTrending, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 50
	}
	_start := (page - 1) * size
	_dataList := make([]*GithubTrending, 0)
	engine.Limit(size, _start)
	if _err := engine.Desc("id").Find(&_dataList); _err != nil {
		return nil, _err
	}
	return _dataList, nil
}
