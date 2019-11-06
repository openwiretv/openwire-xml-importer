package main

import (
	"io/ioutil"
	"time"

	"github.com/beevik/etree"
)

type openwireWriter struct {
	feed     *importFeed
	filepath string
}

func (t *openwireWriter) ToBytes() ([]byte, error) {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	rss := doc.CreateElement("rss")
	rss.CreateAttr("version", "2.0")
	rss.CreateAttr("xmlns:openwire", "http://www.openwire.tv/1.0/feed")

	channel := rss.CreateElement("channel")
	channel.CreateElement("title").CreateText(t.feed.Title)
	channel.CreateElement("link").CreateText(t.feed.Link)
	img := channel.CreateElement("image")
	img.CreateElement("url").CreateText(t.feed.Image.URL)
	img.CreateElement("title").CreateText(t.feed.Image.Title)
	img.CreateElement("link").CreateText(t.feed.Image.Link)
	channel.CreateElement("description").CreateText(t.feed.Description)
	channel.CreateElement("language").CreateText(t.feed.Language)
	channel.CreateElement("copyright").CreateText(t.feed.Copyright)
	channel.CreateElement("lastBuildDate").CreateText(t.feed.LastBuildDate.Format(time.RFC1123Z))
	channel.CreateElement("category").CreateText(t.feed.Category)
	channel.CreateElement("pubDate").CreateText(t.feed.PubDate.Format(time.RFC1123Z))
	channel.CreateElement("openwire:keywords").CreateText(t.feed.Keywords)
	rssSuggestions := channel.CreateElement("openwire:suggestions")

	for _, feedChannel := range t.feed.Suggestions.Channels {
		rssChannel := rssSuggestions.CreateElement("openwire:channel")
		rssChannel.CreateElement("openwire:title").CreateText(feedChannel.Title)
		rssChannel.CreateElement("openwire:url").CreateText(feedChannel.URL)
	}

	for _, item := range t.feed.Items {
		rssItem := channel.CreateElement("item")
		rssItem.CreateElement("guid").CreateText(item.Guid)
		rssItem.CreateElement("title").CreateText(item.Title)
		rssItem.CreateElement("link").CreateText(item.Link)
		rssItem.CreateElement("pubDate").CreateText(item.PubDate.Format(time.RFC1123Z))
		rssItem.CreateElement("description").CreateText(item.Description)
		rssItem.CreateElement("openwire:duration").CreateText(item.Duration)
		rssItem.CreateElement("openwire:image").CreateAttr("href", item.Image.Href)
		enc := rssItem.CreateElement("enclosure")
		enc.CreateAttr("url", item.Enclosure.URL)
		enc.CreateAttr("type", item.Enclosure.Type)
		enc.CreateAttr("length", "")
	}

	// doc.Indent(2)
	// doc.WriteTo(os.Stdout)

	return doc.WriteToBytes()
}

func (t *openwireWriter) Write() error {
	content, err := t.ToBytes()
	if err != nil {
		return err
	}

	return ioutil.WriteFile(t.filepath, content, 0644)
}
