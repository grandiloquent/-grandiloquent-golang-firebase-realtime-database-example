package share

import (
	"fmt"
	"net/http"
)

func writeError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(501)
	fmt.Fprintln(w, err)
}
