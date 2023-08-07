package services

import (
	"expense-manager-backend/models"
	"expense-manager-backend/utils"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func extractNewsList(doc *goquery.Document) []models.News {
	var newsList []models.News

	newsItems := doc.Find("div.index-content-left.col0 div.cardshap.redius_bg div.news_list")
	newsItems.Each(func(_ int, newsItem *goquery.Selection) {
		title, _ := newsItem.Find("h3 a").Attr("title")
		link, _ := newsItem.Find("h3 a").Attr("href")
		date := strings.Trim(newsItem.Find(".fb_search_btn small").Text(), "()")
		id, _ := strconv.Atoi(strings.Split(link, "?")[1])

		news := models.News{
			Title: title,
			Date:  date,
			Id:    id,
		}
		newsList = append(newsList, news)
	})

	return newsList
}

func FetchAndExtractNewsList(url string) ([]models.News, error) {
	doc, err := utils.FetchHTMLContent(url)
	if err != nil {
		return nil, err
	}

	newsList := extractNewsList(doc)
	return newsList, nil
}
