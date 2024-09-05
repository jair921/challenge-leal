package app

import (
	"fmt"
	"github.com/jair921/challenge-leal/infrastructure/http/routes"
	"github.com/jair921/challenge-leal/internal/dependencies"
	"log"
	"net/http"
)

func Run(env string) {
	// Cargar la configuración
	cfg, err := LoadConfig(env)
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// Configurar las dependencias
	deps := &dependencies.Dependencies{}

	// Configurar el router
	r := routes.NewRouter(deps)

	// Arrancar el servidor
	port := cfg.Server.Port
	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
