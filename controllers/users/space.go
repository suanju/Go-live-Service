package users

import (
	"Go-Live/controllers"
	receive "Go-Live/interaction/receive/users"
	"Go-Live/logic/users"
	"Go-Live/utils/response"
	"Go-Live/utils/validator"
	"github.com/gin-gonic/gin"
)

type SpaceControllers struct {
	controllers.BaseControllers
}

//GetSpaceIndividual 获取个人空间
func (sp SpaceControllers) GetSpaceIndividual(ctx *gin.Context) {
	userID := ctx.GetUint("currentUserID")
	GetSpaceIndividualReceive := new(receive.GetSpaceIndividualReceiveStruct)
	if err := ctx.ShouldBind(GetSpaceIndividualReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := users.GetSpaceIndividual(GetSpaceIndividualReceive, userID)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//GetReleaseInformation 获取发布信息（视频and专栏）
func (sp SpaceControllers) GetReleaseInformation(ctx *gin.Context) {
	GetReleaseInformationReceive := new(receive.GetReleaseInformationReceiveStruct)
	if err := ctx.ShouldBind(GetReleaseInformationReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := users.GetReleaseInformation(GetReleaseInformationReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}