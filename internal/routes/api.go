package routes

import (
	"server/internal/db"
	"server/internal/handlers"

	"github.com/go-chi/chi/v5"
)

type RouterService struct {
	handlerService *handlers.HandlerService
	Router         *chi.Mux
}

func NewRouterService(dbService *db.DBService) *RouterService {
	router := chi.NewRouter()
	handlerService := handlers.NewHandlerService(dbService)
	return &RouterService{handlerService, router}
}

func (thisService *RouterService) APIRouter() *chi.Mux {
	thisService.Router.Route("/user", thisService.UserRouter())
	thisService.Router.Route("/auth", thisService.AuthRouter())
	return thisService.Router
}

func (thisService *RouterService) UserRouter() func(chi.Router) {
	return func(router chi.Router) {
		router.Route("/{username}", func(subrouter chi.Router) {
			subrouter.Get("/", thisService.handlerService.GetUser())
		})
	}
}

func (thisService *RouterService) AuthRouter() func(chi.Router) {
	return func(router chi.Router) {
		router.Post("/signup", thisService.handlerService.SignUpUser())
		router.Post("/login", thisService.handlerService.LoginUser())
	}
}
