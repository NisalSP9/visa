package routes

import (
	"github.com/gorilla/mux"
	"../api"
)

func VisaRoutes() *mux.Router  {
	var router = mux.NewRouter()
	router = mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/visa/get", api.GetAllData)
	router.HandleFunc("/api/visa/post", api.CreateData)
	router.HandleFunc("/api/visa/delete", api.DeleteData)
	router.HandleFunc("/api/visa/update", api.UpdateData)
	return router
}
