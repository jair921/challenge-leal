package main

import (
	"github.com/jair921/challenge-leal/internal/application"
	"os"
)

func main() {
	application.Run(os.Getenv("APP_ENV"))
}
