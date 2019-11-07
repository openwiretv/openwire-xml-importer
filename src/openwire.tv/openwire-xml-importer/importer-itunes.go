package main

import (
	"encoding/xml"
	"time"
)

type iTunesImporter struct {
	importerBase
	url string
}

func (t *iTunesImporter) getURL() string {
	return t.url
}

func (t *iTunesImporter) Parse(rawContent []byte) (*importFeed, error) {
	var it itunesFeed
	if err := xml.Unmarshal(rawContent, &it); err != nil {
		return nil, err
	}

	feedToImport := &importFeed{
		Title: it.Channel.Title,
		Link:  it.Channel.Image.Link, // ?
		Image: importFeedImage{
			Link:  it.Channel.Image.Link,
			Title: it.Channel.Image.Title,
			URL:   it.Channel.Image.URL,
		},
		Description:   it.Channel.Description,
		Language:      it.Channel.Language,
		Copyright:     it.Channel.Copyright,
		LastBuildDate: t.getDate(it.Channel.LastBuildDate),
		Category:      "",
		PubDate:       t.getDate(""),
		Keywords:      "",
		Items:         []importFeedItem{},
	}

	if len(it.Channel.Categories) > 0 {
		feedToImport.Category = it.Channel.Categories[0].Text
	}

	for _, item := range it.Channel.Items {
		feedToImport.Items = append(feedToImport.Items, importFeedItem{
			Title:       item.Title,
			GUID:        item.GUID,
			Link:        item.Link,
			PubDate:     item.getPubDate(),
			Description: item.Description,
			Enclosure: importFeedItemEnclosure{
				URL:  item.Enclosure.URL,
				Type: item.Enclosure.Type,
			},
			Duration: item.Duration,
			Image: importFeedItemImage{
				Href: item.Image.Href,
			},
		})
	}

	return feedToImport, nil
}

func (t *iTunesImporter) getDate(rawDate string) time.Time {
	// ti, err := time.Parse("2006-01-02T03:04:05+00:00", rawDate)
	ti, err := time.Parse(time.RFC1123Z, rawDate)
	if err != nil {
		return time.Now()
	}

	return ti
}

func newITunesImporter(podcastURL string) importer {
	itImporter := &iTunesImporter{
		url: podcastURL,
	}
	itImporter.downloader = &downloader{
		url: itImporter.getURL(),
	}

	return itImporter
}

type itunesFeed struct {
	XMLName xml.Name          `xml:"rss"`
	Channel itunesFeedChannel `xml:"channel"`
}

type itunesFeedChannel struct {
	XMLName       xml.Name             `xml:"channel"`
	Title         string               `xml:"title"`
	LastBuildDate string               `xml:"lastBuildDate"`
	Language      string               `xml:"language"`
	Copyright     string               `xml:"copyright"`
	Description   string               `xml:"description"`
	Link          []string             `xml:"link"`
	Image         itunesFeedImage      `xml:"image"`
	Categories    []itunesFeedCategory `xml:"category"`
	Items         []itunesFeedItem     `xml:"item"`
}

type itunesFeedCategory struct {
	Text string `xml:"text,attr"`
}

type itunesFeedImage struct {
	XMLName xml.Name `xml:"image"`
	Title   string   `xml:"title"`
	URL     string   `xml:"url"`
	Link    string   `xml:"link"`
}

type itunesFeedItem struct {
	XMLName     xml.Name                `xml:"item"`
	Title       string                  `xml:"title"`
	Description string                  `xml:"description"`
	Link        string                  `xml:"link"`
	GUID        string                  `xml:"guid"`
	PubDate     string                  `xml:"pubDate"` // Tue, 24 Sep 2019 09:00:00 -0000; Mon, 04 November 2019 01:00:00 +0000
	Duration    string                  `xml:"duration"`
	Image       itunesFeedItemImage     `xml:"image"`
	Enclosure   itunesFeedItemEnclosure `xml:"enclosure"`
}

func (t *itunesFeedItem) getPubDate() time.Time {
	if ti, err := time.Parse(time.RFC1123Z, t.PubDate); err == nil {
		return ti
	}
	if ti, err := time.Parse("Mon, 02 January 2006 15:04:05 -0700", t.PubDate); err == nil {
		return ti
	}

	return time.Now()
}

type itunesFeedItemImage struct {
	XMLName xml.Name `xml:"image"`
	Href    string   `xml:"href,attr"`
}

type itunesFeedItemEnclosure struct {
	XMLName xml.Name `xml:"enclosure"`
	URL     string   `xml:"url,attr"`
	Type    string   `xml:"type,attr"`
}
