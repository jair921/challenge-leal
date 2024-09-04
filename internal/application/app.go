package application

import (
	"fmt"
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
	deps := Setup(cfg)

	// Configurar el router
	r := SetupRouter(deps)

	// Arrancar el servidor
	port := cfg.Server.Port
	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
