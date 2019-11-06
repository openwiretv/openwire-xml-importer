package main

import (
	"strings"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWriter(t *testing.T) {
	Convey("Generating full xml", t, func() {

		timeNow := time.Now()
		timeFormated := timeNow.Format(time.RFC1123Z)

		feedToImport := &importFeed{
			Title: "OpenWireTV Demo",
			Link:  "https://www.openwire.tv",
			Image: importFeedImage{
				Link:  "https://www.openwire.tv",
				Title: "OpenWireTV Demo",
				URL:   "https://www.openwire.tv/icon.png",
			},
			Description:   "Demonstracyjny kanał pokazujący przykładowe filmiki bezpośrednio z plików mp4 oraz z serwisu Vimeo",
			Language:      "pl-pl",
			Copyright:     "OpenWireTV",
			LastBuildDate: timeNow,
			Category:      "Education",
			PubDate:       timeNow,
			Keywords:      "openwiretv,demo",
			Items: []importFeedItem{
				importFeedItem{
					Title:       "Las pokryty śniegiem",
					Guid:        "12",
					Link:        "https://openwire.tv",
					PubDate:     timeNow,
					Description: "asd",
					Enclosure: importFeedItemEnclosure{
						URL:  "https://openwire.tv/videos/Pexels_Videos_1858244.mp4",
						Type: "video/mp4",
					},
					Duration: "0:14",
					Image: importFeedItemImage{
						Href: "https://i.vimeocdn.com/video/756241775.jpg",
					},
				},
			},
			Suggestions: importFeedSuggestions{
				Channels: []importFeedSuggestionChannel{
					importFeedSuggestionChannel{
						Title: "Filmiki ekonomiczne",
						URL:   "https://www.openwire.tv/openwire-ekonomia.xml",
					},
				},
			},
		}
		feedWriter := newWriter(feedToImport, "")
		feedStr, _ := feedWriter.ToBytes()
		expectedXml := `<?xml version="1.0" encoding="UTF-8"?><rss version="2.0" xmlns:openwire="http://www.openwire.tv/1.0/feed"><channel><title>OpenWireTV Demo</title><link>https://www.openwire.tv</link><image><url>https://www.openwire.tv/icon.png</url><title>OpenWireTV Demo</title><link>https://www.openwire.tv</link></image><description>Demonstracyjny kanał pokazujący przykładowe filmiki bezpośrednio z plików mp4 oraz z serwisu Vimeo</description><language>pl-pl</language><copyright>OpenWireTV</copyright><lastBuildDate>[DATE-TO-REPLACE]</lastBuildDate><category>Education</category><pubDate>[DATE-TO-REPLACE]</pubDate><openwire:keywords>openwiretv,demo</openwire:keywords><openwire:suggestions><openwire:channel><openwire:title>Filmiki ekonomiczne</openwire:title><openwire:url>https://www.openwire.tv/openwire-ekonomia.xml</openwire:url></openwire:channel></openwire:suggestions><item><guid>12</guid><title>Las pokryty śniegiem</title><link>https://openwire.tv</link><pubDate>[DATE-TO-REPLACE]</pubDate><description>asd</description><openwire:duration>0:14</openwire:duration><openwire:image href="https://i.vimeocdn.com/video/756241775.jpg"/><enclosure url="https://openwire.tv/videos/Pexels_Videos_1858244.mp4" type="video/mp4" length=""/></item></channel></rss>`
		expectedXml = strings.Replace(expectedXml, "[DATE-TO-REPLACE]", timeFormated, 3)
		So(string(feedStr), ShouldEqual, expectedXml)
	})
}
