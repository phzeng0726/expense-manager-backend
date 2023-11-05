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

func extractEnNewsList(doc *goquery.Document) []models.News {
	var newsList []models.News

	title := doc.Find(".news_title_col").Text()

	// newsItems.Each(func(i int, s *goquery.Selection) {
	// 	fmt.Println(s.Text())
	// })
	// newsItems.Each(func(_ int, newsItem *goquery.Selection) {

	// 	title := newsItem.Text()
	// 	// link, _ := newsItem.Find("h3 a").Attr("href")
	// 	// date := strings.Trim(newsItem.Find(".fb_search_btn small").Text(), "()")
	// 	// id, _ := strconv.Atoi(strings.Split(link, "?")[1])
	// 	fmt.Println(title)

	// })
	news := models.News{
		Title: title,
		// Date:  date,
		// Id:    id,
	}
	newsList = append(newsList, news)
	return newsList
}

func extractTwNewsList(doc *goquery.Document) []models.News {
	var newsList []models.News

	temp := doc.Find("div.md:col-span-2")

	newsItems := doc.Find("div.index-content-left.col0 div.cardshap.redius_bg div.news_list")
	newsItems.Each(func(_ int, newsItem *goquery.Selection) {
		fmt.Println(temp) // 使用 %+v 來印出詳細信息

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

func FetchAndExtractNewsList(language string) ([]models.News, error) {
	var newsList []models.News
	var url string

	if language == "en_US" {
		url = fmt.Sprintf("%s/tag/FIN/1", constants.EnDomain)
	} else {
		url = fmt.Sprintf("%s/news/list.php?nt_pk=7", constants.ZhDomain)
	}

	doc, err := utils.FetchHTMLContent(url)
	if err != nil {
		return newsList, err
	}

	if language == "en_US" {
		newsList = extractEnNewsList(doc)
	} else {
		newsList = extractTwNewsList(doc)

	}

	return newsList, nil
}
