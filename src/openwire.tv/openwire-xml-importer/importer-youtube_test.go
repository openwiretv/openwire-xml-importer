package main

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestImporterYoutube(t *testing.T) {
	Convey("Test dates", t, func() {
		importer := newYouTubeImporter("")
		ti := importer.getDate("2019-11-04T05:59:59+00:00")

		So(ti.Format(time.RFC3339), ShouldEqual, "2019-11-04T05:59:59Z")
	})
	Convey("Test unmarshaling XML feed", t, func() {
		importer := newYouTubeImporter("")
		feed, err := importer.Parse([]byte(testYoutubeFeed))

		So(err, ShouldEqual, nil)
		So(feed.Title, ShouldEqual, "Test channel name")
		So(feed.Link, ShouldEqual, "https://www.youtube.com/channel/UCRuqcoHGbRDOvgKwJ8B87c")
		So(feed.Description, ShouldEqual, "")
		So(feed.Language, ShouldEqual, "")
		So(feed.Copyright, ShouldEqual, "@ Test channel name")
		So(feed.Category, ShouldEqual, "")
		So(feed.Keywords, ShouldEqual, "")
		So(len(feed.Items), ShouldEqual, 1)
		So(feed.Items[0].Title, ShouldEqual, "Test entry title")
		So(feed.Items[0].Enclosure.URL, ShouldEqual, "3_dlQju4J")
		So(feed.Items[0].Enclosure.Type, ShouldEqual, "video/x-youtube")
	})
}

var testYoutubeFeed = `
<?xml version="1.0" encoding="UTF-8"?>
<feed xmlns:yt="http://www.youtube.com/xml/schemas/2015" xmlns:media="http://search.yahoo.com/mrss/" xmlns="http://www.w3.org/2005/Atom">
 <link rel="self" href="http://www.youtube.com/feeds/videos.xml?channel_id=UCRuqcoHGbRDOvgKwJ8B87c"/>
 <id>yt:channel:UCRuqcoHGbRDOvgKwJ8B87c</id>
 <yt:channelId>UCRuqcoHGbRDOvgKwJ8B87c</yt:channelId>
 <title>Test channel name</title>
 <link rel="alternate" href="https://www.youtube.com/channel/UCRuqcoHGbRDOvgKwJ8B87c"/>
 <author>
  <name>Test channel name</name>
  <uri>https://www.youtube.com/channel/UCRuqcoHGbRDOvgKwJ8B87c</uri>
 </author>
 <published>2013-11-12T23:44:50+00:00</published>
 <entry>
  <id>yt:video:3_dlQju4J</id>
  <yt:videoId>3_dlQju4J</yt:videoId>
  <yt:channelId>UCRuqcoHGbRDOvgKwJ8B87c</yt:channelId>
  <title>Test entry title</title>
  <link rel="alternate" href="https://www.youtube.com/watch?v=3_dlQju4J"/>
  <author>
   <name>Test author</name>
   <uri>https://www.youtube.com/channel/UCRuqcoHGbRDOvgKwJ8B87c</uri>
  </author>
  <published>2019-11-04T05:00:00+00:00</published>
  <updated>2019-11-04T06:20:18+00:00</updated>
  <media:group>
   <media:title>Test title</media:title>
   <media:content url="https://www.youtube.com/v/3_dlQju4J?version=3" type="application/x-shockwave-flash" width="640" height="390"/>
   <media:thumbnail url="https://i4.ytimg.com/vi/3_dlQju4J/hqdefault.jpg" width="480" height="360"/>
   <media:description>test</media:description>
   <media:community>
    <media:starRating count="48" average="4.75" min="1" max="5"/>
    <media:statistics views="1036"/>
   </media:community>
  </media:group>
 </entry>
</feed>
`
