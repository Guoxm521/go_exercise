package model

import "example.com/m/v2/model/db"

func (that *Logic) NewGithubTrending() *GithubTrending {
	return &GithubTrending{}
}

type GithubTrending struct {
	Logic
	mode *db.GithubTrending
}

func (that *GithubTrending) init() *GithubTrending {
	if that.mode == nil {
		_self := new(db.GithubTrending)
		that.mode = _self
	}
	return that
}

func (that *GithubTrending) SetGithubTrending(value string) *GithubTrending {
	that.init()
	that.mode.Author = value
	return that
}
func (that *GithubTrending) SetTableField(name string, value interface{}) {
	that.init()
	that.parseTableField(that.mode, name, value)
}

func (that *GithubTrending) Add() {
	that.init()
	that.SetGithubTrending("zhangsan")
	_mode := that.mode
	_mode.Add(that.mode)
}

func (that *GithubTrending) List() (interface{}, error) {
	that.init()
	_mode := that.mode
	_data, _err := _mode.List(1, 10)
	_count, _err := _mode.Count()
	_mp := make(map[string]interface{}, 0)
	_mp["count"] = _count
	_mp["data"] = _data
	return _mp, _err
}
