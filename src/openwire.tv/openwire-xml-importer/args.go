package main

import (
	"flag"
)

type args struct {
	output           string
	youtubeChannelID string
	iTunesPodcastURL string
}

func (t *args) isValid() bool {
	if t.output == "" || !isWritable(t.output) {
		return false
	}

	return t.isYouTubeImport() || t.isITunesImport()
}
func (t *args) isYouTubeImport() bool {
	return t.youtubeChannelID != ""
}
func (t *args) isITunesImport() bool {
	return t.iTunesPodcastURL != ""
}
func (t *args) getNewImporter() importer {
	if t.isYouTubeImport() {
		return newYouTubeImporter(t.youtubeChannelID)
	} else if t.isITunesImport() {
		return newITunesImporter(t.iTunesPodcastURL)
	}

	panic("unable to get importer")
}

func getCmdArguments() *args {
	output := flag.String("output", "", "Path to output file")
	ytChannelID := flag.String("yt-channel-id", "", "YouTube Channel ID")
	itPodcastURL := flag.String("it-podcast-url", "", "iTunes Podcast Url")

	flag.Parse()

	return &args{
		output:           *output,
		youtubeChannelID: *ytChannelID,
		iTunesPodcastURL: *itPodcastURL,
	}
}
