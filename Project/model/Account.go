package model

import (
	"errors"
	"example.com/m/v2/middleware"
	"example.com/m/v2/model/db"
	uuid "github.com/satori/go.uuid"
)

func (that *Logic) NewAccount() *Account {
	return &Account{}
}

type Account struct {
	Logic
	mode *db.Account
}

func (that *Account) init() *Account {
	if that.mode == nil {
		_self := new(db.Account)
		that.mode = _self
	}
	return that
}

func (that *Account) SetTableField(name string, value interface{}) {
	that.init()
	that.parseTableField(that.mode, name, value)
}

func (that *Account) Add() (interface{}, error) {
	that.init()
	_mode := that.mode
	_b, _err := _mode.GetByAccount(that.mode.Account)
	if _b {
		return nil, errors.New("账号已存在")
	}
	if _err != nil {
		return nil, _err
	}
	that.mode.AccountId = uuid.NewV4().String()
	_, _err = _mode.AddAccount(that.mode)
	if _err != nil {
		return "", _err
	}
	_mode.GetByAccountID(that.mode.AccountId)
	return _mode, nil
}

func (that *Account) Login() (interface{}, error) {
	that.init()
	_mode := that.mode
	_newMode := db.NewAccount()
	_b, _err := _newMode.GetByAccount(_mode.Account)
	if !_b {
		//账号不存在
		that.mode.AccountId = uuid.NewV4().String()
		_, _err = _mode.AddAccount(that.mode)
	} else {
		if _newMode.Password != _mode.Password {
			return nil, errors.New("密码错误")
		}
	}
	if _err != nil {
		return nil, _err
	}
	_mode.GetByAccountID(that.mode.AccountId)
	user := middleware.UserInfo{
		Account:   _mode.Account,
		Password:  _mode.Password,
		AccountId: _mode.AccountId,
	}
	tokenString, _err := middleware.GenerateToken(user)
	_mp := make(map[string]interface{}, 0)
	_mp["account"] = _mode.Account
	_mp["token"] = tokenString
	return _mp, nil
}

func (that *Account) GetInfo(account string) (interface{}, error) {
	that.init()
	_newMode := db.NewAccount()
	_b, _err := _newMode.GetByAccount(account)
	if !_b {
		return nil, errors.New("账号不存在")
	}
	if _err != nil {
		return nil, _err
	}
	return _newMode, nil
}
