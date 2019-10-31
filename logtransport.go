// Package logtransport implements a http.RoundTripper which logs request before the request is sent and response after the request.
package logtransport

import (
	"bufio"
	"bytes"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

var defaultLogger = log.New(os.Stdout, "", log.LstdFlags)

// Logger is an interface of a logger, which is used for logging
// outgoing HTTP request and incoming HTTP response.
type Logger interface {
	Print(args ...interface{})
}

// Transport is a wrapper of http.RoundTripper.
type Transport struct {
	Transport http.RoundTripper
	Logger    Logger
}

// RoundTrip implements http.RoundTripper interface.
func (t *Transport) RoundTrip(r *http.Request) (*http.Response, error) {
	if err := t.logRequest(r); err != nil {
		return nil, err
	}
	transport := t.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}
	resp, err := transport.RoundTrip(r)
	if err != nil {
		return nil, err
	}
	if err := t.logResponse(resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *Transport) Print(args ...interface{}) {
	logger := t.Logger
	if logger == nil {
		logger = defaultLogger
	}
	logger.Print(args...)
}

func (t *Transport) logRequest(r *http.Request) error {
	b, err := httputil.DumpRequestOut(r, true)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(bytes.NewReader(b))
	for scanner.Scan() {
		t.Print(scanner.Text())
	}
	return scanner.Err()
}

func (t *Transport) logResponse(resp *http.Response) error {
	b, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(bytes.NewReader(b))
	for scanner.Scan() {
		t.Print(scanner.Text())
	}
	return scanner.Err()
}
