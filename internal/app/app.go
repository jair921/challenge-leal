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

	//Conexión mysql
	db, err := InitDB(cfg)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	defer db.Close()

	// Configurar las dependencias
	deps := app.SetupDependencies(db)

	// Configurar el router
	r := routes.NewRouter(deps)

	// Arrancar el servidor
	port := cfg.Server.Port
	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
