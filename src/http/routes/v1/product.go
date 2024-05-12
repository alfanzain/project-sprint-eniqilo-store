package v1

import (
	"github.com/alfanzain/eniqilo-store/src/http/handlers"
	"github.com/alfanzain/eniqilo-store/src/http/middlewares"
	"github.com/alfanzain/eniqilo-store/src/repositories"
	"github.com/alfanzain/eniqilo-store/src/services"
)

func (i *V1Routes) MountProduct() {
	g := i.Echo.Group("/product")

	productHandler := handlers.NewProductHandler(services.NewProductService(
		repositories.NewProductRepository(),
	))

	g.POST("", productHandler.Store, middlewares.Authorized())
}
