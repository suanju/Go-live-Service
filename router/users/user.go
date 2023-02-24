package users

import (
	"Go-Live/controllers/users"
	"Go-Live/middlewares"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (s *LoginRouter) InitRouter(Router *gin.RouterGroup) {
	router := Router.Group("user").Use(middlewares.VerificationToken())
	{
		usersControllers := new(users.UserControllers)
		router.POST("/getUserInfo", usersControllers.GetUserInfo)
		router.POST("/setUserInfo", usersControllers.SetUserInfo)
		router.POST("/determineNameExists", usersControllers.DetermineNameExists)
		router.POST("/updateAvatar", usersControllers.UpdateAvatar)
		router.POST("/upload", usersControllers.Upload)
		router.POST("/getLiveData", usersControllers.GetLiveData)
		router.POST("/saveLiveData", usersControllers.SaveLiveData)
		router.POST("/sendEmailVerificationCodeByChangePassword", usersControllers.SendEmailVerificationCodeByChangePassword)
		router.POST("/changePassword", usersControllers.ChangePassword)
		router.POST("/attention", usersControllers.Attention)
		router.POST("/createFavorites", usersControllers.CreateFavorites)
		router.POST("/getFavoritesList", usersControllers.GetFavoritesList)
		router.POST("/deleteFavorites", usersControllers.DeleteFavorites)
		router.POST("/favoriteVideo", usersControllers.FavoriteVideo)
		router.POST("/getFavoritesListByFavoriteVideo", usersControllers.GetFavoritesListByFavoriteVideo)
	}
}
