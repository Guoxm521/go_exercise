package db

func NewAccount() *Account {
	_self := new(Account)
	return _self
}

type Account struct {
	BaseHeader  `xorm:"extends"`
	AccountId   string `xorm:"varchar(40) default '' unique(account_id) comment('账号id')" json:"account_id"` //账号Id
	Account     string `xorm:"varchar(30) default '' not null comment('登录账号')" json:"account"`              //登录账号
	Password    string `xorm:"varchar(100) default '' not null comment('登陆密码')" json:"password"`            //登陆密码
	Avatar      string `xorm:"varchar(255) not null default '' comment('图片')" json:"avatar"`                //图片
	Name        string `xorm:"varchar(50) default '' not null comment('账号名称')" json:"name"`                 //账号名称
	PersonnelId int64  `xorm:"int(11) not null default 0 index comment('员工id')" json:"personnel_id"`        //员工Id
	IsSuper     int64  `xorm:"int(3) not null default 1 comment('是否超管')" json:"is_super"`                   //是否超管
	IsEnable    int64  `xorm:"int(3) not null default 1 comment('状态')" json:"is_enable"`                    //状态
	VerifySign  int64  `xorm:"int(3) not null default 0 comment('验签')" json:"verify_sign"`                  //验签
	ModeFooter  `xorm:"extends"`
}

func (that *Account) GetByAccountID(accountID interface{}) (bool, error) {
	return engine.Where("account_id = ?", accountID).Get(that)
}

func (that *Account) AddAccount(data interface{}) (int64, error) {
	return engine.Insert(data)
}

func (that *Account) GetByAccount(account interface{}) (bool, error) {
	return engine.Where("account= ?", account).Get(that)
}
