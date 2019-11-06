package main

type iTunesImporter struct {
	importerBase
	url string
}

func (t *iTunesImporter) getURL() string {
	return t.url
}

func (t *iTunesImporter) Parse(rawContent []byte) (*importFeed, error) {
	return &importFeed{}, nil
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
