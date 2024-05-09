package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ApiRoutes struct {
	mux *chi.Mux
}

func NewApiRouter() *ApiRoutes {
	r := chi.NewRouter()

	return &ApiRoutes{r}
}

func (a *ApiRoutes) InitRoutes() {
	a.mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	log.Println("Server starting on port 3000")
	err := http.ListenAndServe(":3000", a.mux)

	if err != nil {
		log.Fatal("Error starting server on port :3000")
	}
}
