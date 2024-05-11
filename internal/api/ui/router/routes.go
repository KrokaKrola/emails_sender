package router

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	handlers "github.com/krokakrola/emails_sender/internal/api/ui/handlers/health"
)

type ApiRoutes struct {
	router *chi.Mux
}

func NewApiRouter() *ApiRoutes {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(30 * time.Second))

	return &ApiRoutes{r}
}

func (a *ApiRoutes) InitRoutes() {
	port := os.Getenv("PORT")

	server := &http.Server{Addr: fmt.Sprintf("0.0.0.0%s", port), Handler: service(a.router)}

	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sig

		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown time out... forcing exit")
			}
		}()

		log.Println("Graceful server shutdown...")

		err := server.Shutdown(shutdownCtx)
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}

		serverStopCtx()
		log.Println("Server succesfully shutdowned")
	}()

	log.Printf("Starting server on port %q\n", port)
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	<-serverCtx.Done()
}

func service(router *chi.Mux) *chi.Mux {
	healthHander := handlers.NewHealthHandlers()

	router.Route("/v1", func(r chi.Router) {
		r.Get("/health", healthHander.GetHealth)
	})

	return router
}
