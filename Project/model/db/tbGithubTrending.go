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
	Starts       string `xorm:"varchar(10) default '' comment('目前start数')" json:"starts"`
	Forks        string `xorm:"varchar(10) default '' comment('目前start数')" json:"forks"`
	Language     string `xorm:"varchar(10) default '' comment('语言')" json:"language"`
	AddedStars   string `xorm:"varchar(100) default '' comment('今天或者这周或者这个月的starts数')" json:"added_stars"`
	Avatars      string `xorm:"varchar(255) default '' comment('项目贡献者的头像地址集合')" json:"avatars"`
	LanguageType int    `xorm:"int(2) index default 0  comment('语言类型')" json:"language_type"`
	SinceType    int    `xorm:"int(2) index default 0 comment('搜索时间')" json:"since_type"`
	ModeFooter   `xorm:"extends"`
}

//GetByAccountID ...
func (that *GithubTrending) GetByAccountID(accountID interface{}) (bool, error) {
	return engine.Where("account_id = ?", accountID).Get(that)
}
