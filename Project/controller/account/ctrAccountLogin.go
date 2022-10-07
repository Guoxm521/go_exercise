package account

import (
	"example.com/m/v2/common"
	"example.com/m/v2/model"
	"github.com/gin-gonic/gin"
)

func AccountLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		that := AccountControllers{}
		if _err := ctx.ShouldBind(&that); _err != nil {
			common.Response(ctx, _err)
			return
		}
		_logic := model.NewLogic().NewAccount()
		_logic.SetTableField("account", that.Account)
		_logic.SetTableField("password", that.Password)
		_data, _exp := _logic.Login()
		if _exp != nil {
			common.Response(ctx, _exp)
			return
		}
		common.Response(ctx, _data)
		return

	}
}
