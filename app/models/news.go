package models

type News struct {
	Id       string `json:"id"`
	Reporter string `json:"reporter"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Date     string `json:"date"`
}
