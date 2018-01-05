package qc

import "encoding/json"

// QueueItem Struct to
type QueueItem struct {
	URL      string         `json:"url"`
	Depth    int16          `json:"depth"`
	Subjects []news.Subject `json:"subjects"`
}

// GetKey Creates key for a QueueItem
func (item QueueItem) GetKey() string {
	return news.CreateURLHash(item.URL)
}

// Serialize Creates a json marshalled string of bytes for a QueueItem
func (item QueueItem) Serialize() (string, error) {
	bytes, err := json.Marshal(&item)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// UnserializeQueueItem Creates a QueueItem from a serialized string
func UnserializeQueueItem(byteStr string) (QueueItem, error) {
	var item QueueItem
	bytes := []bytes(byteStr)
	err = json.Unmarshall(bytes, &item)
	if err != nil {
		return item, err
	}
	return item, nil
}
