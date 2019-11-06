package main

import (
	"encoding/xml"
	"time"
)

type youTubeImporter struct {
	importerBase
	channelID string
}

func (t *youTubeImporter) getURL() string {
	return "https://www.youtube.com/feeds/videos.xml?channel_id=" + t.channelID
}

func (t *youTubeImporter) Parse(rawContent []byte) (*importFeed, error) {
	var yt youtubeFeed
	if err := xml.Unmarshal(rawContent, &yt); err != nil {
		return nil, err
	}

	feedToImport := &importFeed{
		Title: yt.Title,
		Link:  "https://www.youtube.com/channel/" + yt.ChannelID,
		Image: importFeedImage{
			Link:  "https://www.youtube.com/channel/" + yt.ChannelID,
			Title: yt.Title,
			URL:   "",
		},
		Description:   "",
		Language:      "",
		Copyright:     "@ " + yt.Title,
		LastBuildDate: t.getDate(""),
		Category:      "",
		PubDate:       t.getDate(yt.PublishedDate),
		Keywords:      "",
		Items:         []importFeedItem{},
	}

	for _, entry := range yt.Entries {
		feedToImport.Items = append(feedToImport.Items, importFeedItem{
			Title:       entry.Title,
			Guid:        entry.VideoID,
			Link:        "https://www.youtube.com/watch?v=" + entry.VideoID,
			PubDate:     t.getDate(entry.PublishedDate),
			Description: "",
			Enclosure: importFeedItemEnclosure{
				URL:  entry.VideoID,
				Type: "video/x-youtube",
			},
			Duration: "0:00",
			Image: importFeedItemImage{
				Href: entry.Media.Thumbnail.URL,
			},
		})
	}

	return feedToImport, nil
}

func (t *youTubeImporter) getDate(rawDate string) time.Time {
	ti, err := time.Parse("2006-01-02T03:04:05+00:00", rawDate)
	if err != nil {
		return time.Now()
	}

	return ti
}

func newYouTubeImporter(channelID string) *youTubeImporter {
	ytImporter := &youTubeImporter{
		channelID: channelID,
	}
	ytImporter.downloader = &downloader{
		url: ytImporter.getURL(),
	}

	return ytImporter
}

type youtubeFeed struct {
	XMLName       xml.Name           `xml:"feed"`
	ChannelID     string             `xml:"channelId"`
	Title         string             `xml:"title"`
	PublishedDate string             `xml:"published"`
	Author        youtubeFeedAuthor  `xml:"author"`
	Entries       []youtubeFeedEntry `xml:"entry"`
}

type youtubeFeedAuthor struct {
	XMLName xml.Name `xml:"author"`
	Name    string   `xml:"name"`
	Uri     string   `xml:"uri"`
}

type youtubeFeedEntry struct {
	XMLName       xml.Name              `xml:"entry"`
	VideoID       string                `xml:"videoId"`
	Title         string                `xml:"title"`
	PublishedDate string                `xml:"published"`
	UpdatedDate   string                `xml:"updated"`
	Media         youtubeFeedEntryMedia `xml:"group"`
}

type youtubeFeedEntryMedia struct {
	XMLName     xml.Name                       `xml:"group"`
	Title       string                         `xml:"title"`
	Description string                         `xml:"description"`
	Thumbnail   youtubeFeedEntryMediaThumbnail `xml:"thumbnail"`
}

type youtubeFeedEntryMediaThumbnail struct {
	XMLName xml.Name `xml:"thumbnail"`
	URL     string   `xml:"url,attr"`
}
