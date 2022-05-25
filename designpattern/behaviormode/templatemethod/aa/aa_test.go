package aa

import "testing"

func TestName(t *testing.T) {

	var downloader Downloader = NewHTTPDownloader()
	downloader.Download("")
}
