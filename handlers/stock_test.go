package handlers_test

import (
	"testing"

	"github.com/6un3u/witb_backend/handlers"
)

func TestMakeStockResult(t *testing.T) {
	stockResult := handlers.MakeStockResult("S000000780090")

	if len(stockResult) == 0 {
		t.Error("There isn't any result")
	}
}
