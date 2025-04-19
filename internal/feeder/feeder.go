package feeder

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

func (c *Client) FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "gator")

	resp, err := c.handler.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	feed, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var rssFeed RSSFeed
	xml.Unmarshal(feed, &rssFeed)
	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)
	for _, item := range rssFeed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
	}

	return &rssFeed, nil
}
