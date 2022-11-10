package main

import (
	"context"
	"main/share"
	"net/http"
)

func main() {
	ctx := context.Background()
	ref := share.GetRef(ctx)

	_ = http.ListenAndServe(":9000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := r.URL.Path
		if p == "/api/articles" {
			share.ListArticles(ref, w, r)
		}
	}))
}
