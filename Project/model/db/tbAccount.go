package db

type Account struct {
	BaseHeader  `xorm:"extends"`
	CompanyId   int64  `xorm:"int(11) index default 0 not null comment('公司id')" json:"company_id"`             //公司id
	AccountId   int64  `xorm:"int(11) not null unique(account_id) comment('账号id')" json:"account_id"`          //账号Id
	Account     string `xorm:"varchar(30) default '' not null unique(account) comment('登录账号')" json:"account"` //登录账号
	Password    string `xorm:"varchar(100) default '' not null comment('登陆密码')" json:"password"`               //登陆密码
	Name        string `xorm:"varchar(50) default '' not null comment('账号名称')" json:"name"`                    //账号名称
	PersonnelId int64  `xorm:"int(11) not null default 0 index comment('员工id')" json:"personnel_id"`           //员工Id
	IsSuper     int64  `xorm:"int(3) not null default 1 comment('是否超管')" json:"is_super"`                      //是否超管
	IsEnable    int64  `xorm:"int(3) not null default 1 comment('状态')" json:"is_enable"`                       //状态
	VerifySign  int64  `xorm:"int(3) not null default 0 comment('验签')" json:"verify_sign"`                     //验签
	ModeFooter  `xorm:"extends"`
}

