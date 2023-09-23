package http

import (
	"log"
	"net/http"
	"testing"
)

func TestHttpServer(t *testing.T) {

	path := "/Users/jordonqm/Downloads/【高清影视之家发布 www.HDBTHD.com】速度与激情10[简繁英字幕].Fast.X.2023.1080p.BluRay.x264.Atmos.TrueHD7.1-CTRLHD"

	fs := http.FileServer(http.Dir(path)) // Specify the directory you want to serve

	http.Handle("/", fs)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
