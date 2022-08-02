package github

import (
	"example.com/m/v2/Project/common"
	"example.com/m/v2/Project/model"
	"github.com/gin-gonic/gin"
)

type GithubTrendingListControllers struct {
	CompanySyncId int64   `form:"company_sync_id" desc:"company_sync_id"`
	CompanyName   string  `form:"company_name" desc:"公司名称"`
	ProvinceId    int64   `form:"province_id" desc:"省"` //州/省
	CityId        int64   `form:"city_id" desc:"市"`     //州/省
	CountyId      int64   `form:"county_id" desc:"区县"`  //区县
	Latitude      float64 `form:"latitude" desc:"经度"`   //经度
	Longitude     float64 `form:"longitude" desc:"经度"`  //经度
	Address       string  `form:"address" desc:"地址"`    //地址
	Name          string  `form:"name" desc:"联系人名称"`
	Phone         string  `form:"phone" desc:"联系方式"`
	Remark        string  `form:"remark" desc:"备注"`
	IsEnable      int64   `form:"is_enable" desc:"状态:1启用2停用"`
}

func GithubTrendingList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		that := GithubTrendingListControllers{}
		if _err := ctx.ShouldBind(&that); _err != nil {
			common.Response(ctx, _err)
			return
		}
		_logic := model.NewLogic().NewGithubTrending()
		_data, _exp := _logic.List()
		//响应
		if _exp != nil {
			common.Response(ctx, _exp)
			return
		}
		common.Response(ctx, _data)
		return

	}
}
