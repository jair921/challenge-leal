package main

import (
	"github.com/jair921/challenge-leal/internal/app"
	"os"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	app.Run(env)
}
