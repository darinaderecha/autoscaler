package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

var DollarNow float64

type createDollarRequest struct {
}

type DollarResponse struct {
	Rate string `json:"rate"`
}

func fetchDollarRate() (string, error) {
	resp, err := http.Get("https://minfin.com.ua/ua/currency/nbu/")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New("failed to fetch dollar rate")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	rate := ""
	doc.Find("td.sc-1x32wa2-10.ccSsXh").Each(func(i int, s *goquery.Selection) {
		if s.Text() == "Долар США" {
			rate = s.Next().Find("div[type='nbu'].sc-1x32wa2-9.fevpFL").Contents().Not("div").Text()
		}
	})

	if rate == "" {
		return "", errors.New("failed to find dollar rate in the page")
	}
	fmt.Println()
	fmt.Println()
	fmt.Println(rate)
	fmt.Println()
	return strings.TrimSpace(rate), nil
}

func (s *Server) getDollar(ctx *gin.Context) {
	var req createDollarRequest
	var err error
	if err = ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	dollar, err := fetchDollarRate()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, DollarResponse{Rate: dollar})
}
