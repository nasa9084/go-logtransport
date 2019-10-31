package logtransport_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/nasa9084/go-logtransport"
)

func Example() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, world!"))
	}
	srv := httptest.NewServer(http.HandlerFunc(handler))
	defer srv.Close()

	c := &http.Client{
		Transport: &logtransport.Transport{
			// you can use other logger like sirupsen/logrus
			Logger: log.New(os.Stdout, "", log.LstdFlags),
		},
	}
	c.Get(srv.URL)
}
