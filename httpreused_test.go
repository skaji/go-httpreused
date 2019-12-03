package httpreused_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/skaji/go-httpreused"
)

func TestBasic(t *testing.T) {
	c := httpreused.Wrap(http.DefaultClient)

	// TODO: redirect?
	res, err := c.Get("http://www.yahoo.co.jp/index.html")
	if err != nil {
		t.Fatal(err)
	}
	reused := res.Header.Get("X-Connection-Reused")
	if reused == "" {
		t.Fail()
	}
	ip := res.Header.Get("X-Connection-IP")
	if ip == "" {
		t.Fail()
	}
	fmt.Println(ip, reused)

}
