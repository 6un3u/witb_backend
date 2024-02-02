package utils_test

import (
	"fmt"
	"testing"

	"github.com/6un3u/witb_backend/utils"
)

func TestValidateStockId(t *testing.T) {
	bookID := "S000000780090"
	result := utils.ValidateStockId(bookID)
	fmt.Printf("%s, %v\n", bookID, result)

	bookID = "S000000800"
	result = utils.ValidateStockId(bookID)
	fmt.Printf("%s, %v\n", bookID, result)

	bookID = "String"
	result = utils.ValidateStockId(bookID)
	fmt.Printf("%s, %v\n", bookID, result)

	bookID = "S000000'80090"
	result = utils.ValidateStockId(bookID)
	fmt.Printf("%s, %v\n", bookID, result)
}
