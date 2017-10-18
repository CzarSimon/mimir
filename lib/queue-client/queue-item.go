package qc

import (
	"github.com/CzarSimon/mimir/lib/news"
)

// QueueItem Struct to
type QueueItem struct {
	URL      string         `json:"url"`
	Depth    int16          `json:"depth"`
	Subjects []news.Subject `json:"subjects"`
}
