logtransport
===

[![GoDoc](https://godoc.org/github.com/nasa9084/go-logtransport?status.svg)](https://godoc.org/github.com/nasa9084/go-logtransport)

logtransport is a thin wrapper of http.RoundTripper interface, which logs request and response using given logger.

## SYNOPSIS

``` go
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

```
