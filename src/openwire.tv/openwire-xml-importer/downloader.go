package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type downloader struct {
	url      string
	filepath string
}

func (t *downloader) download() error {
	resp, err := http.Get(t.url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(t.getFilepath())
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func (t *downloader) getFilepath() string {
	if t.filepath != "" {
		return t.filepath
	}
	t.filepath = "/tmp/" + randSeq(10) + ".xml"

	return t.filepath
}

func (t *downloader) getContent() ([]byte, error) {
	return ioutil.ReadFile(t.getFilepath())
}

func (t *downloader) delete() error {
	return os.Remove(t.getFilepath())
}
