package ws

import (
	"example.com/m/v2/common"
	"example.com/m/v2/websocket"
	"github.com/gin-gonic/gin"
)

func GetWebsocketInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_data := make(map[string]interface{}, 0)
		_data["demo"] = 1231
		_mp := websocket.WebsocketManager.Info()
		common.Response(ctx, _mp)
		return

	}
}
