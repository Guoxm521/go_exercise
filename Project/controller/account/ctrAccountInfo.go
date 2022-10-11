package account

import (
	"example.com/m/v2/common"
	"example.com/m/v2/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AccountInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_logic := model.NewLogic().NewAccount()
		_account, _ := ctx.Get("account")
		fmt.Println("====================", _account)
		_data, _exp := _logic.GetInfo(ToString(_account))
		if _exp != nil {
			common.Response(ctx, _exp)
			return
		}
		common.Response(ctx, _data)
		return

	}
}

type Stringer interface {
	String() string
}

func ToString(any interface{}) string {
	if v, ok := any.(Stringer); ok {
		return v.String()
	}
	switch v := any.(type) {
	case int:
		return strconv.Itoa(v)
	case string:
		return v
	}
	return ""
}
