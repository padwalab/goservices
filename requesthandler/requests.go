package requesthandler

import (
	"net/http"
)

// ExtractRequestDetails extacts endpoint specific request details
func ExtractRequestDetails(r *http.Request) string {
	values := r.URL.Query()
	name := values.Get("name")
	return name
}
