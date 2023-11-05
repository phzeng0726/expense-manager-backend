package services

import (
	"expense-manager-backend/models"
	"expense-manager-backend/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func extractEnglishNews(doc *goquery.Document) models.News {
	var paragraphs string
	newsDoc := doc.Find("div.bg-white.p-4.col-span-5")

	newsDoc.Find("p").Each(func(i int, s *goquery.Selection) {
		paragraphs += s.Text()
	})

	news := models.News{
		Date:    strings.TrimSpace(newsDoc.Find("h3 span").Text()),
		Title:   strings.TrimSpace(newsDoc.Find(".news_title").Text()),
		Content: paragraphs,
	}

	return news
}

func extractTraditionalChineseNews(doc *goquery.Document) models.News {
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

func FetchAndExtractNews(language string, newsId string) (models.News, error) {
	var news models.News
	url, err := constructURL(language, &newsId)
	if err != nil {
		return news, err
	}

	doc, err := utils.FetchHTMLContent(url)
	if err != nil {
		return news, err
	}

	if language == "en_US" {
		news = extractEnglishNews(doc)
		news.Id = newsId
	} else {
		news = extractTraditionalChineseNews(doc)
		news.Id = newsId
	}

	return news, nil
}
