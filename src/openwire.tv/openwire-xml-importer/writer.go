package main

type writer interface {
	Write() error
	ToBytes() ([]byte, error)
}

func newWriter(feed *importFeed, filepath string) writer {
	return &openwireWriter{
		feed:     feed,
		filepath: filepath,
	}
}
