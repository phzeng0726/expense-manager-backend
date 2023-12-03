package services

import (
	"expense-manager-backend/models"
	"expense-manager-backend/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func extractAllEnglishNews(doc *goquery.Document) []models.News {
	var extractedNewsList []models.News
	newsItems := doc.Find("div.mt-4.mx-4.mb-12.cursor-pointer.border-b.border-solid.border-slate-200.transition.duration-150.ease-in")

	newsItems.Each(func(i int, newsItem *goquery.Selection) {
		title := strings.TrimSpace(newsItem.Find(".news_title_col").Text())
		link, _ := newsItem.Find("a").Attr("href")
		date := newsItem.Find("h4").Text()
		id := strings.Split(link, "news/")[1]

		news := models.News{
			Title: title,
			Date:  date,
			Id:    id,
		}
		extractedNewsList = append(extractedNewsList, news)
	})

	return extractedNewsList
}

func extractAllTraditionalChineseNews(doc *goquery.Document) []models.News {
	var extractedNewsList []models.News

	newsItems := doc.Find("div.index-content-left.col0 div.cardshap.redius_bg div.news_list")
	newsItems.Each(func(_ int, newsItem *goquery.Selection) {
		title, _ := newsItem.Find("h3 a").Attr("title")
		link, _ := newsItem.Find("h3 a").Attr("href")
		date := strings.Trim(newsItem.Find(".fb_search_btn small").Text(), "()")
		id := strings.Split(link, "?")[1]

		news := models.News{
			Title: title,
			Date:  date,
			Id:    id,
		}
		extractedNewsList = append(extractedNewsList, news)
	})

	return extractedNewsList
}

func FetchAndExtractNewsList(language string) ([]models.News, error) {
	var extractedNewsList []models.News

	url, err := constructURL(language, nil)
	if err != nil {
		return extractedNewsList, err
	}

	doc, err := utils.FetchHTMLContent(url)
	if err != nil {
		return extractedNewsList, err
	}

	if language == "en_US" {
		extractedNewsList = extractAllEnglishNews(doc)
	} else {
		extractedNewsList = extractAllTraditionalChineseNews(doc)
	}

	return extractedNewsList, nil
}
