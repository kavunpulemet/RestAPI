package handler

import (
	"RESTAPIService2/pkg/service"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	services *service.Service
}

// Функция-обертка для применения h.userIdentity в качестве middleware
func (h *Handler) userIdentityMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.userIdentity(w, r)
		next.ServeHTTP(w, r)
	})
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() http.Handler {
	router := mux.NewRouter()

	auth := router.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/sign-up", h.signUp).Methods(http.MethodPost)
	auth.HandleFunc("/sign-in", h.signIn).Methods(http.MethodPost)

	api := router.PathPrefix("/api").Subrouter()
	api.Use(h.userIdentityMW)
	lists := api.PathPrefix("/lists").Subrouter()
	lists.HandleFunc("/", h.createList).Methods(http.MethodPost)
	lists.HandleFunc("/", h.getAllLists).Methods(http.MethodGet)
	lists.HandleFunc("/{id}", h.getListById).Methods(http.MethodGet)
	lists.HandleFunc("/{id}", h.updateList).Methods(http.MethodPut)
	lists.HandleFunc("/{id}", h.deleteList).Methods(http.MethodDelete)

	items := lists.PathPrefix("/{id}/items").Subrouter()
	items.HandleFunc("/", h.createItem).Methods(http.MethodPost)
	items.HandleFunc("/", h.getAllItems).Methods(http.MethodGet)
	items.HandleFunc("/{item_id}", h.getItemById).Methods(http.MethodGet)
	items.HandleFunc("/{item_id}", h.updateItem).Methods(http.MethodPut)
	items.HandleFunc("/{item_id}", h.deleteItem).Methods(http.MethodDelete)

	return router
}
