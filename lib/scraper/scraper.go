package scraper

// Interface Main interface for scraper
type Interface interface {
	Get(URL string) (string, error)
	Parse(html string) (Article, error)
}

// Scraper Implementation of ScraperInterface
type Scraper struct {
	UserAgent   string
	Subjects    []Subject
	HasSubjects bool
}

// NewScraper Creates a new scraper with default subjects
func NewScraper(userAgent string, subjects []Subject) Scraper {
	return Scraper{
		UserAgent:   userAgent,
		Subjects:    subjects,
		HasSubjects: true,
	}
}

// NewEmptyScraper Creates new scraper with no default subjects
func NewEmptyScraper(userAgent string) Scraper {
	return Scraper{
		UserAgent:   userAgent,
		HasSubjects: false,
	}
}

// Get Retrives the raw html from a page
func (scraper Scraper) Get(URL string) (string, error) {
	return URL, nil
}

// Parse Parses raw html and returns a Page struct if successfull
func (scraper Scraper) Parse(html string) (Article, error) {
	return Article{}, nil
}
