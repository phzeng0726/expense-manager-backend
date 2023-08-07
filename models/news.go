package models

var Domain = "https://www.cardu.com.tw"

type News struct {
	Id       int    `json:"id"`
	Reporter string `json:"reporter"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Date     string `json:"date"`
}
