// smoke test all fixes 202349
package api

import (
	"fmt"
	"os"
)

// API_TOKEN is loaded from the environment; never hardcode secrets in source.
var API_TOKEN = os.Getenv("API_TOKEN")

func Health() string {
	return fmt.Sprintf("ok")
}
