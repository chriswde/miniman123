package api

import (
	"net/http"
)

func Shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
}
