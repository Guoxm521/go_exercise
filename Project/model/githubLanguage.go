package model

import (
	"example.com/m/v2/Project/model/db"
)

func (that *Logic) NewGithubLanguage() *GithubLanguage {
	return &GithubLanguage{}
}

type GithubLanguage struct {
	mode *db.GithubLanguage
}

func (that *GithubLanguage) init() *GithubLanguage {
	if that.mode == nil {
		_self := new(db.GithubLanguage)
		that.mode = _self
	}
	return that
}

func (that *GithubLanguage) SetGithubLanguageName(value string) *GithubLanguage {
	that.init()
	that.mode.Name = value
	return that
}

func (that *GithubLanguage) Add() {
	that.init()
	that.SetGithubLanguageName("zhangsan")
	_mode := that.mode
	_mode.Add(that.mode)
}
