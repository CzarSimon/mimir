package news

// Author contains info about an article refererer.
type Author struct {
	ID            int64 `json:"id"`
	FollowerCount int64 `json:"followerCount"`
}
