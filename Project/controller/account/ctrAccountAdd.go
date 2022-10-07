package account

import (
	"example.com/m/v2/common"
	"example.com/m/v2/model"
	"github.com/gin-gonic/gin"
)

type AccountControllers struct {
	Account  string `form:"account"`
	Password string `form:"password"`
	Avatar   string `form:"avatar"`
}

func AccountAdd() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		that := AccountControllers{}
		if _err := ctx.ShouldBind(&that); _err != nil {
			common.Response(ctx, _err)
			return
		}
		_logic := model.NewLogic().NewAccount()
		_logic.SetTableField("account", that.Account)
		_logic.SetTableField("password", that.Password)
		_logic.SetTableField("avatar", that.Avatar)
		_data, _exp := _logic.Add()
		if _exp != nil {
			common.Response(ctx, _exp)
			return
		}
		common.Response(ctx, _data)
		return

	}
}
