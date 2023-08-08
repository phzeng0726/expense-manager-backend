package services

import (
	"expense-manager-backend/constants"
	"expense-manager-backend/models"
	"expense-manager-backend/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func extractNews(doc *goquery.Document) models.News {
	head := strings.TrimSpace(doc.Find("div.col-md-8.col-12 p").Text())
	parts := strings.Split(head, "報導")

	news := models.News{
		Reporter: strings.TrimSpace(parts[0]),
		Date:     strings.TrimSpace(parts[1]),
		Title:    strings.TrimSpace(doc.Find("div.pt-3.mx-3.detail_title h1").Text()),
		Content:  strings.TrimSpace(doc.Find("div.pb-3.mx-3.detail_content").Text()),
	}

	return news
}

func FetchAndExtractNews(idStr string) (models.News, error) {
	url := fmt.Sprintf("%s/news/detail.php?%s", constants.Domain, idStr)

	doc, err := utils.FetchHTMLContent(url)
	if err != nil {
		return models.News{}, err
	}

	news := extractNews(doc)
	id, _ := strconv.Atoi(idStr)
	news.Id = id

	return news, nil
}
