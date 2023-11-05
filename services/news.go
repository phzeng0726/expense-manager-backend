package services

import (
	"expense-manager-backend/constants"
	"expense-manager-backend/models"
	"expense-manager-backend/utils"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func extractEnNews(doc *goquery.Document) models.News {
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

func extractTwNews(doc *goquery.Document) models.News {
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

func FetchAndExtractNews(language string, idStr string) (models.News, error) {
	var news models.News
	var url string
	if language == "en_US" {
		url = fmt.Sprintf("%s/news/%s", constants.EnDomain, idStr)
	} else {
		url = fmt.Sprintf("%s/news/detail.php?%s", constants.ZhDomain, idStr)
	}

	doc, err := utils.FetchHTMLContent(url)
	if err != nil {
		return news, err
	}

	if language == "en_US" {
		news = extractEnNews(doc)
		news.Id = idStr
	} else {
		news = extractTwNews(doc)
		news.Id = idStr
	}

	return news, nil
}
