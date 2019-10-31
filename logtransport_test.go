package logtransport_test

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/nasa9084/go-logtransport"
)

const date = "Thu, 31 Oct 2019 06:00:00 GMT"

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Date", date)
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("foobar")); err != nil {
		panic(err)
	}
}

func TestRoundTrip(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(testHandler))
	defer srv.Close()

	var buf bytes.Buffer
	c := &http.Client{
		Transport: &logtransport.Transport{
			Transport: http.DefaultTransport,
			Logger:    log.New(&buf, "", 0),
		},
	}

	req, err := http.NewRequest(
		http.MethodGet,
		srv.URL,
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := c.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: %d != %d", resp.StatusCode, http.StatusOK)
		return
	}

	got := buf.String()
	want := fmt.Sprintf(`GET / HTTP/1.1
Host: %s
User-Agent: Go-http-client/1.1
Accept-Encoding: gzip

HTTP/1.1 200 OK
Content-Length: 6
Content-Type: text/plain; charset=utf-8
Date: %s

foobar
`, strings.TrimPrefix(srv.URL, "http://"), date)
	if got != want {
		t.Errorf("unexpected log:\ngot:\n%s\nwant:\n%s", got, want)
		return
	}
}
