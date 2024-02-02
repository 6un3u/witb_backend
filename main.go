package main

import (
	"os"

	"github.com/6un3u/witb_backend/router"
	"github.com/6un3u/witb_backend/utils"
	"github.com/joho/godotenv"
)

type indexData struct {
	PageTitle string
}

func main() {
	err := godotenv.Load(".env")
	utils.HandleErr(err)

	e := router.Router()

	if os.Getenv("DEBUG") == "true" {
		e.Debug = true
	} else {
		e.Debug = false
	}

	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
