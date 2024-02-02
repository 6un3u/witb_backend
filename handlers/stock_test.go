package handlers_test

import (
	"testing"

	"github.com/6un3u/witb_backend/handlers"
	"github.com/6un3u/witb_backend/utils"
	"github.com/joho/godotenv"
)

func TestMakeStockResult(t *testing.T) {
	err := godotenv.Load("../.env")
	utils.HandleErr(err)

	stockResult := handlers.MakeStockResult("S000000780090")

	if len(stockResult) == 0 {
		t.Error("There isn't any result")
	}
}
