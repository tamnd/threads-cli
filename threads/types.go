package threads

import "time"

// Profile is a Threads user.
type Profile struct {
	ID             string    `json:"id,omitempty"`
	Username       string    `json:"username"`
	Name           string    `json:"name,omitempty"`
	Biography      string    `json:"biography,omitempty"`
	ProfilePicURL  string    `json:"profile_pic_url,omitempty"`
	IsVerified     bool      `json:"is_verified"`
	FollowerCount  int64     `json:"follower_count,omitempty"`
	FollowingCount int64     `json:"following_count,omitempty"`
	URL            string    `json:"url"`
	FetchedAt      time.Time `json:"fetched_at"`
}

// Post is a single Threads post.
type Post struct {
	ID           string    `json:"id"`
	Shortcode    string    `json:"shortcode,omitempty"`
	Text         string    `json:"text,omitempty"`
	MediaType    string    `json:"media_type,omitempty"` // TEXT_POST, IMAGE, VIDEO, CAROUSEL_ALBUM
	MediaURLs    []string  `json:"media_urls,omitempty"`
	Permalink    string    `json:"permalink,omitempty"`
	Username     string    `json:"username,omitempty"`
	UserID       string    `json:"user_id,omitempty"`
	Timestamp    time.Time `json:"timestamp,omitempty"`
	LikeCount    int64     `json:"like_count"`
	ReplyCount   int64     `json:"reply_count"`
	RepostCount  int64     `json:"repost_count"`
	QuoteCount   int64     `json:"quote_count"`
	IsQuotePost  bool      `json:"is_quote_post,omitempty"`
	IsReply      bool      `json:"is_reply,omitempty"`
	QuotedPostID string    `json:"quoted_post_id,omitempty"`
	ReplyToID    string    `json:"reply_to_id,omitempty"`
	HasMedia     bool      `json:"has_media,omitempty"`
	FetchedAt    time.Time `json:"fetched_at"`
}

// Reply is a reply to a post, carrying its thread linkage.
type Reply struct {
	ID         string    `json:"id"`
	ParentID   string    `json:"parent_id,omitempty"`
	RootID     string    `json:"root_id,omitempty"`
	Shortcode  string    `json:"shortcode,omitempty"`
	Text       string    `json:"text,omitempty"`
	Username   string    `json:"username,omitempty"`
	UserID     string    `json:"user_id,omitempty"`
	Permalink  string    `json:"permalink,omitempty"`
	Timestamp  time.Time `json:"timestamp,omitempty"`
	LikeCount  int64     `json:"like_count"`
	ReplyCount int64     `json:"reply_count"`
	MediaType  string    `json:"media_type,omitempty"`
	MediaURLs  []string  `json:"media_urls,omitempty"`
	FetchedAt  time.Time `json:"fetched_at"`
}

// SearchResult is one hit from a keyword search.
type SearchResult struct {
	ID          string    `json:"id"`
	Query       string    `json:"query"`
	Text        string    `json:"text,omitempty"`
	Username    string    `json:"username,omitempty"`
	Permalink   string    `json:"permalink,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
	MediaType   string    `json:"media_type,omitempty"`
	IsReply     bool      `json:"is_reply,omitempty"`
	IsQuotePost bool      `json:"is_quote_post,omitempty"`
	SearchedAt  time.Time `json:"searched_at"`
}

// asReply converts a parsed post (a reply lives in the same thread_items shape
// as a post) into a Reply under the given root/parent.
func (p Post) asReply(parentID, rootID string) Reply {
	return Reply{
		ID:         p.ID,
		ParentID:   parentID,
		RootID:     rootID,
		Shortcode:  p.Shortcode,
		Text:       p.Text,
		Username:   p.Username,
		UserID:     p.UserID,
		Permalink:  p.Permalink,
		Timestamp:  p.Timestamp,
		LikeCount:  p.LikeCount,
		ReplyCount: p.ReplyCount,
		MediaType:  p.MediaType,
		MediaURLs:  p.MediaURLs,
		FetchedAt:  p.FetchedAt,
	}
}
