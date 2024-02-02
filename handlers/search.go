package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/6un3u/witb_backend/utils"
	"github.com/labstack/echo/v4"
)

type Search struct {
	SearchQuery string `json:"s"`
}

type SearchResponse struct {
	Data struct {
		ReturnCode      int              `json:"returnCode"`
		TotalSize       int              `json:"totalSize"`
		RealSize        int              `json:"realSize"`
		ErrorMessage    string           `json:"errorMessage"`
		ResultDocuments []ResultDocument `json:"resultDocuments"`
	} `json:"data"`
}

type ResultDocument struct {
	ISBNCode    string `json:"CMDTCODE"`
	BookName    string `json:"CMDT_NAME"`
	BookId      string `json:"SALE_CMDTID"`
	Country     string `json:"SALE_CMDT_DVSN_CODE_SUB"`
	Score       string `json:"SCORE"`
	HTMLlist    string `json:"RELATE_HTML_LIST"`
	DVSNCode    string `json:"SALE_CMDT_GRP_DVSN_CODE"`
	DqId        string `json:"DQ_ID"`
	SearchType  string `json:"SALE_CMDT_DVSN_CODE"`
	TotHTMLList string `json:"TOT_RELATE_HTML_LIST"`
}

type Book struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Price     string `json:"price"`
	Img       string `json:"cover"`
	Id        string `json:"id"`
}

// @Summery Search Book
// @Accept json
// @Produce json
// @Param s body Search true "Book Name"
// @Success 200 {object} []Book
// @Router /search [post]
func SearchHandler(c echo.Context) error {
	var search Search
	err := c.Bind(&search)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid Request")
	}

	bookList := makeSearchResult(search.SearchQuery)
	return c.JSON(http.StatusOK, bookList)
}

func makeSearchResult(s string) []Book {
	search_url := os.Getenv("SEARCH_URL") + url.QueryEscape(s)
	log.Println(search_url)
	body := utils.GetHTTP(search_url)

	strBody := string(body)
	strBody = strings.TrimPrefix(strBody, "autocompleteShop(")
	strBody = strings.TrimSuffix(strBody, ");")

	var searchResponse SearchResponse
	err := json.Unmarshal([]byte(strBody), &searchResponse)
	utils.HandleErr(err)

	var books []Book
	for _, resultDoc := range searchResponse.Data.ResultDocuments {
		if strings.HasPrefix(resultDoc.BookId, "S") { // except ebook that starts with "E"
			book := extractBookInfo(resultDoc)
			books = append(books, book)
		}
	}

	return books
}

func extractBookInfo(doc ResultDocument) Book {
	var book Book
	book.Id = doc.BookId

	splitInfo := strings.Split(doc.TotHTMLList, "$@")
	book.Title = splitInfo[2]
	book.Author = splitInfo[3]
	book.Publisher = splitInfo[4]
	book.Price = splitInfo[7]
	book.Img = splitInfo[14]

	return book
}
