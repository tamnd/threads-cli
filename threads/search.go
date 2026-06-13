package threads

import (
	"context"
	"iter"
	"time"
)

// Search streams keyword search hits via the logged-out search query. It is
// honest when the current doc_id does not expose search anonymously: the stream
// simply ends with no error after yielding what it found.
func (c *Client) Search(ctx context.Context, query string, limit int) iter.Seq2[SearchResult, error] {
	return func(yield func(SearchResult, error) bool) {
		posts, err := c.graphqlSearch(ctx, query)
		if err != nil {
			yield(SearchResult{}, err)
			return
		}
		n := 0
		for _, p := range posts {
			r := SearchResult{
				ID:          p.ID,
				Query:       query,
				Text:        p.Text,
				Username:    p.Username,
				Permalink:   p.Permalink,
				Timestamp:   p.Timestamp,
				MediaType:   p.MediaType,
				IsReply:     p.IsReply,
				IsQuotePost: p.IsQuotePost,
				SearchedAt:  time.Now(),
			}
			if !yield(r, nil) {
				return
			}
			n++
			if limit > 0 && n >= limit {
				return
			}
		}
	}
}
