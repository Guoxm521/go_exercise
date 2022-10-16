package message

import (
	"example.com/m/v2/common"
	"example.com/m/v2/model"
	"github.com/gin-gonic/gin"
)

type MessageListControllers struct {
	Group string `form:"group" desc:"分组群聊"`
	Uid   string `form:"uid" desc:"uid"`
	Page  int    `form:"page,default=1"  binding:"required" desc:"页码"`
	Size  int    `form:"size,default=100" binding:"required" desc:"每页数量"`
}

func MessageList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		that := MessageListControllers{}
		if _err := ctx.ShouldBind(&that); _err != nil {
			common.Response(ctx, _err)
			return
		}
		_logic := model.NewLogic().NewMessage()
		_logic.SetTableField("group", that.Group)
		_logic.SetTableField("uid", that.Uid)
		_data, _exp := _logic.List(that.Page, that.Size)
		if _exp != nil {
			common.Response(ctx, _exp)
			return
		}
		common.Response(ctx, _data)
		return
	}
}
