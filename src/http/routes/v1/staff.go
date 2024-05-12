package v1

import (
	"github.com/alfanzain/project-sprint-eniqilo-store/src/http/handlers"
	"github.com/alfanzain/project-sprint-eniqilo-store/src/repositories"
	"github.com/alfanzain/project-sprint-eniqilo-store/src/services"
)

func (i *V1Routes) MountStaff() {
	g := i.Echo.Group("/staff")

	staffHandler := handlers.NewStaffHandler(services.NewStaffService(
		repositories.NewStaffRepository(),
	))

	g.POST("/register", staffHandler.Register)
	g.POST("/login", staffHandler.Login)
}
