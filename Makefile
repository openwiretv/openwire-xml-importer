run-yt:
	./bin/openwire-xml-importer -output /tmp/openwire.xml -yt-channel-id UCRuqcoHGbRDOvgKwJ8B87cw
run-it:
	./bin/openwire-xml-importer -output /tmp/openwire.xml -it-podcast-url https://feeds.captivate.fm/ecommerce-masterplan/
install:
	go get github.com/beevik/etree
	go install openwire.tv/openwire-xml-importer
test:
	go test -v openwire.tv/openwire-xml-importer
