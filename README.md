# OpenWire XML Importer

Command line tool for parsing XML files from remote services (eg. YouTube) and generating
a proper XML file for OpenWireTV.

Supported remote services:

* YouTube feed
* iTunes feed (podcasts)

## Installation

```
cd /tmp
git clone https://github.com/openwiretv/openwire-xml-importer
export GOPATH=/tmp/openwire-xml-importer
cd openwire-xml-importer/
make install
```

### Examples

Argument name | Description
--- | ---
`output` | Path to file, where final XML will be saved eg. `/tmp/openwire.xml`
`yt-channel-id` | Channel ID from YouTube eg. `UCNT3teib54LOt98CSlpzxjQ`
`it-podcast-url` | Url to iTunes feed file

```
./bin/openwire-xml-importer -output <output-file-path> -yt-channel-id <youtube-channel-id>
./bin/openwire-xml-importer -output <output-file-path> -it-podcast-url <url-to-itunes-xml>
```

### Tests

Run `make test`

### TODO

* Argument to pass base `openwire.xml` file for copying base fields + previous videos/podcasts

### License

MIT