package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestImporterItunes(t *testing.T) {
	Convey("Test unmarshaling XML feed", t, func() {
		importer := newITunesImporter("")
		feed, err := importer.Parse([]byte(testItunesFeed))

		So(err, ShouldEqual, nil)
		So(feed.Title, ShouldEqual, "Test")
		So(feed.Link, ShouldEqual, "https://openwire.tv/podcast")
		So(feed.Description, ShouldEqual, "zxc")
		So(feed.Language, ShouldEqual, "en")
		So(feed.Copyright, ShouldEqual, "Copyright 2019 AAA")
		So(feed.Category, ShouldEqual, "Business")
		So(feed.Keywords, ShouldEqual, "")
		So(feed.Image.URL, ShouldEqual, "https://cdn.captivate.fm/artwork/original.png")
		So(feed.Image.Title, ShouldEqual, "Test image")
		So(len(feed.Items), ShouldEqual, 1)
		So(feed.Items[0].Title, ShouldEqual, "BBB")
		So(feed.Items[0].Enclosure.URL, ShouldEqual, "https://openwire.tv/test.mp3")
		So(feed.Items[0].Enclosure.Type, ShouldEqual, "audio/mpeg")
	})
}

var testItunesFeed = `
<?xml version="1.0" encoding="UTF-8"?>
<rss xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns:atom="http://www.w3.org/2005/Atom" version="2.0" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd" xmlns:sy="http://purl.org/rss/1.0/modules/syndication/">
<channel>
<atom:link href="https://openwire.tv/test-itunes-feed/" rel="self" type="application/rss+xml"/>
<title>Test</title>
<lastBuildDate>Thu, 07 November 2019 00:51:56 GMT</lastBuildDate>
<generator>openwire.tv</generator>
<language>en</language>
<copyright>Copyright 2019 AAA</copyright>
<managingEditor>contact@openwire.tv</managingEditor>
<itunes:summary>zxc</itunes:summary>
<image>
	<url>https://cdn.captivate.fm/artwork/original.png</url>
	<title>Test image</title>
	<link>https://openwire.tv/podcast</link>
</image>
<itunes:image href="https://cdn.captivate.fm/artwork/original.png"/>
<itunes:owner>
	<itunes:name>John Doe</itunes:name>
	<itunes:email>contact@openwire.tv</itunes:email>
</itunes:owner>
<itunes:author>John Doe</itunes:author>
<description>zxc</description>
<link>https://openwire.tv/podcast</link>
<atom:link href="https://pubsubhubbub.appspot.com" rel="hub"/>
<itunes:subtitle>VVVVV</itunes:subtitle>
<itunes:explicit>no</itunes:explicit>
<itunes:type>episodic</itunes:type>
<itunes:category text="Business"/>
<itunes:category text="Business">
	<itunes:category text="Marketing"/>
</itunes:category>
<itunes:category text="Business">
	<itunes:category text="Entrepreneurship"/>
</itunes:category>
<itunes:new-feed-url>https://openwire.tv/test-itunes-feed/</itunes:new-feed-url>
<item>
	<title>BBB</title>
	<itunes:title>BBB</itunes:title>
	<description>NNNN</description>
	<content:encoded>NNNN</content:encoded>
	<link>https://ecommercemasterplan.com/podcast</link>
	<guid isPermaLink="false">3f868894-c136-494a-b3bc-6d7547ce6005</guid>
	<itunes:image href="https://cdn.captivate.fm/artwork/original.png"/>
	<dc:creator>John Doe</dc:creator>
	<pubDate>Mon, 04 November 2019 01:00:00 +0000</pubDate>
	<enclosure url="https://openwire.tv/test.mp3" length="28421302" type="audio/mpeg"/>
	<itunes:duration>26:43</itunes:duration>
	<itunes:explicit>no</itunes:explicit>
	<itunes:episodeType>full</itunes:episodeType>
	<itunes:episode>240</itunes:episode>
	<itunes:summary>NNNN</itunes:summary>
	<itunes:author>John Doe</itunes:author>
</item>
</channel>
</rss>
`
