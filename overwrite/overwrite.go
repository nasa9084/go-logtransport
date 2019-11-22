package overwrite

import (
	"log"
	"net/http"
	"os"

	"github.com/nasa9084/go-logtransport"
)

func init() {
	http.DefaultTransport = &logtransport.Transport{
		Logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}
