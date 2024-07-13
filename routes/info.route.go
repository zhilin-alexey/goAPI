package routes

import (
	"github.com/gin-gonic/gin"
	"goAPI/controllers"
)

type InfoRoute struct {
	peopleController controllers.PeopleController
}

func NewInfoRoute(peopleController controllers.PeopleController) InfoRoute {
	return InfoRoute{peopleController}
}

func (infoRoute *InfoRoute) Register(router *gin.Engine) {
	router.GET("/info", infoRoute.peopleController.GetByPassport)
}
