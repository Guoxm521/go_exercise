package model

import (
	"example.com/m/v2/model/db"
)

func (that *Logic) NewMessage() *Message {
	return &Message{}
}

type Message struct {
	Logic
	mode *db.Message
}

func (that *Message) init() *Message {
	if that.mode == nil {
		_self := new(db.Message)
		that.mode = _self
	}
	return that
}

func (that *Message) SetTableField(name string, value interface{}) {
	that.init()
	that.parseTableField(that.mode, name, value)
}

func (that *Message) List(page int, size int) (interface{}, error) {
	that.init()
	_mode := that.mode
	_accountMode := db.NewAccount()
	_search := make(map[string]interface{}, 0)
	_search["group"] = that.mode.Group
	_search["uid"] = that.mode.UId
	_data, _err := _mode.List(page, size, _search)
	_uids := []string{}
	for _, v := range _data {
		_b := IsContain(_uids, v.UId)
		if !_b {
			_uids = append(_uids, v.UId)
		}
	}
	_accountlist, _ := _accountMode.GetAccountList(map[string]interface{}{
		"uids": _uids,
	})
	_accountMap := make(map[string]*db.Account, 0)
	for _, account := range _accountlist {
		if _accountMap[account.AccountId] == nil {
			_accountMap[account.AccountId] = account
		}
	}

	_dataList := []interface{}{}
	for _, datum := range _data {
		_dataMap := make(map[string]interface{}, 0)
		_dataMap["account"] = _accountMap[datum.UId].Account
		_dataMap["account_id"] = _accountMap[datum.UId].AccountId
		_dataMap["avatar"] = _accountMap[datum.UId].Avatar
		_dataMap["name"] = _accountMap[datum.UId].Name
		_dataMap["u_id"] = datum.UId
		_dataMap["c_time"] = datum.CTime
		_dataMap["content"] = datum.Content
		_dataMap["to_uid"] = datum.ToUid
		_dataMap["group"] = datum.Group
		_dataList = append(_dataList, _dataMap)
	}
	_count, _err := _mode.Count(_search)
	_mp := make(map[string]interface{}, 0)
	_mp["data"] = _dataList
	_mp["count"] = _count
	_mp["page"] = page
	_mp["size"] = size
	return _mp, _err
}

func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}
