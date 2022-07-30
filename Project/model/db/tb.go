package db

//BaseHeader ..
type BaseHeader struct {
	Id       int64     `xorm:"INT(11) pk autoincr" json:"id"` //表自增id
}

//ModeFooter ...
type ModeFooter struct {
	CTime int64 `xorm:"int(11) not null created comment('添加时间')" json:"c_time"` //创建时间
	UTime int64 `xorm:"int(11) not null updated comment('更新时间')" json:"u_time"` //最后更新时间
}
