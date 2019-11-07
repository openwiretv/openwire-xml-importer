package main

import (
	"encoding/xml"
	"time"
)

type importFeed struct {
	XMLName       xml.Name              `xml:"channel"`
	Title         string                `xml:"title"`
	Link          string                `xml:"link"`
	Image         importFeedImage       `xml:"image"`
	Description   string                `xml:"description"`
	Language      string                `xml:"language"`
	Copyright     string                `xml:"copyright"`
	LastBuildDate time.Time             `xml:"lastBuildDate"`
	Category      string                `xml:"category"`
	PubDate       time.Time             `xml:"pubDate"`
	Keywords      string                `xml:"keywords"` // namespace
	Items         []importFeedItem      `xml:"item"`
	Suggestions   importFeedSuggestions `xml:"suggestions"` // namespace
}

type importFeedImage struct {
	XMLName xml.Name `xml:"image"`
	URL     string   `xml:"url"`
	Title   string   `xml:"title"`
	Link    string   `xml:"link"`
}
type importFeedSuggestions struct {
	Channels []importFeedSuggestionChannel `xml:"channel"`
}
type importFeedSuggestionChannel struct {
	XMLName xml.Name `xml:"channel"` // namespace
	Title   string   `xml:"title"`   // namespace
	URL     string   `xml:"url"`     // namespace
}
type importFeedItem struct {
	XMLName     xml.Name                `xml:"item"`
	Title       string                  `xml:"title"`
	GUID        string                  `xml:"guid"`
	Link        string                  `xml:"link"`
	PubDate     time.Time               `xml:"pubDate"`
	Description string                  `xml:"description"`
	Enclosure   importFeedItemEnclosure `xml:"enclosure"`
	Duration    string                  `xml:"duration"` // namespace
	Image       importFeedItemImage     `xml:"image"`    // namespace
}
type importFeedItemEnclosure struct {
	XMLName xml.Name `xml:"enclosure"`
	URL     string   `xml:"url,attr"`
	Type    string   `xml:"type,attr"`
}
type importFeedItemImage struct {
	XMLName xml.Name `xml:"image"` // namespace
	Href    string   `xml:"href,attr"`
}
