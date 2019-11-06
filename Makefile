run-yt:
	./bin/openwire-xml-importer -output /tmp/openwire.xml -yt-channel-id UCRuqcoHGbRDOvgKwJ8B87cw
install:
	go install openwire.tv/openwire-xml-importer
test:
	go test -v openwire.tv/openwire-xml-importer
