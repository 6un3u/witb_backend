package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/6un3u/witb_backend/utils"
	"github.com/labstack/echo/v4"
)

type SearchBook struct {
	ID string `json:"id" validate:"bookIdValidation"`
}

type stores struct {
	StrAreaGrpCode string
	List           []storeInfo
}

type storeInfo struct {
	Barcode             string
	SaleCmdtId          string
	SaleCmdtGrpDvsnCode string
	SaleCmdtDvsnCode    string
	StrRdpCode          string
	StrName             string
	StrAreaGrpCode      string
	StrAdrs             string
	StrTlnm             string
	RealInvnQntt        int
	DlvrRqrmDyCont      int
	PlorRqrmDyCont      int
}

type stockResponse struct {
	Data          []stores
	StatusCode    int
	ResultCode    string
	ResultMessage string
	DetailMessage interface{}
}

type StockByStore struct {
	Store   string `json:"store"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Stock   int    `json:"stock"`
}

// @Summery Stock of Book
// @Description Get the book stock with bookID
// @Accept json
// @Produce json
// @Param id body SearchBook true "Book ID"
// @Success 200 {object} [][]StockByStore
// @Router /stock [post]
func StockHandler(c echo.Context) error {
	var searchBook SearchBook
	err := c.Bind(&searchBook)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid Request")
	}

	err = c.Validate(searchBook)
	if err != nil {
		log.Printf("Validator Error: %s\n", err)
	}

	bookStocks := makeStockResult(searchBook.ID)
	return c.JSON(http.StatusOK, bookStocks)
}

func makeStockResult(id string) [][]StockByStore {
	stock_url := strings.Replace(os.Getenv("STOCK_URL"), "[book_id]", id, -1)
	log.Println(stock_url)
	body := utils.GetHTTP(stock_url)

	var stockResponse stockResponse
	err := json.Unmarshal(body, &stockResponse)
	utils.HandleErr(err)

	var bookStocks [][]StockByStore

	for _, strGrp := range stockResponse.Data {
		var stock []StockByStore
		for _, store := range strGrp.List {
			stock = append(stock, extractStock(store))
		}
		bookStocks = append(bookStocks, stock)
	}

	return bookStocks
}

func extractStock(strInfo storeInfo) StockByStore {
	var stockInfo StockByStore
	stockInfo.Store = strInfo.StrName
	stockInfo.Address = strInfo.StrAdrs
	stockInfo.Phone = strInfo.StrTlnm
	stockInfo.Stock = strInfo.RealInvnQntt
	return stockInfo
}
