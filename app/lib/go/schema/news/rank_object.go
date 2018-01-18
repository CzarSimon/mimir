package news

// RankObject contains info to scrape and rank an article.
type RankObject struct {
	Urls     []string  `json:"urls"`
	Subjects []Subject `json:"subjects"`
	Author   Author    `json:"author"`
	Language string    `json:"language"`
}
