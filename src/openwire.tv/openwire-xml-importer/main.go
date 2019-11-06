package main

import (
	"fmt"
	"os"
)

func main() {
	cmdArguments := getCmdArguments()

	if !cmdArguments.isValid() {
		returnError("Command is not valid")
	}

	feedImporter := cmdArguments.getNewImporter()
	feedrawContent, err := feedImporter.Download()
	if err != nil {
		returnError("Unable to download: " + err.Error())
	}
	feed, err := feedImporter.Parse(feedrawContent)
	if err != nil {
		returnError("Unable to import: " + err.Error())
	}
	writer := newWriter(feed, cmdArguments.output)
	if err := writer.Write(); err != nil {
		returnError("Unable to save file: " + err.Error())
	}
}

func returnError(msg string) {
	fmt.Println(msg)

	os.Exit(1)
}
