package routes

import (
	"github.com/gin-gonic/gin"
	"goAPI/controllers"
)

type PeopleRoute struct {
	peopleController controllers.PeopleController
}

func NewPeopleRoute(peopleController controllers.PeopleController) PeopleRoute {
	return PeopleRoute{peopleController}
}

func (peopleRoute *PeopleRoute) Register(router *gin.Engine) {
	router.GET("/people", peopleRoute.peopleController.GetMultiple)
	router.DELETE("/people", peopleRoute.peopleController.Delete)
	router.PATCH("/people/:id", peopleRoute.peopleController.Edit)
	router.POST("/people", peopleRoute.peopleController.Create)
}
