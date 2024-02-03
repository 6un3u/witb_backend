package main

import (
	"os"

	"github.com/6un3u/witb_backend/router"
)

type indexData struct {
	PageTitle string
}

func main() {
	e := router.Router()
	print(os.Getenv("ENV"))

	if os.Getenv("ENV") == "DEV" {
		e.Debug = true
	} else {
		e.Debug = false
	}

	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
