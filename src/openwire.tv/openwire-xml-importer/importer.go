package main

type importer interface {
	Download() ([]byte, error)
	Parse(rawContent []byte) (*importFeed, error)
}

type importerBase struct {
	downloader *downloader
}

func (t *importerBase) Download() ([]byte, error) {
	if err := t.downloader.download(); err != nil {
		return nil, err
	}

	return t.downloader.getContent()
}
