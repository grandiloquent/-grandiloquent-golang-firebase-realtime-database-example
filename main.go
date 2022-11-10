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
	// var data map[string]interface{}
	// if err := ref.Child("").Get(ctx, &data); err != nil {
	// 	log.Fatalln("Error reading from database:", err)
	// }
	// m := make(map[string]interface{})
	// m["updateAt"] = time.Now().Unix()
	// d, err := ref.Child("articles").Push(ctx, m)
	// fmt.Println(d.Key)
	// q := ref.Child("articles").OrderByChild("updateAt")
	// qn, err := q.GetOrdered(ctx)
	// for _, qx := range qn {
	// 	var x map[string]interface{}
	// 	ref.Child("articles/"+qx.Key()).Get(ctx, &x)
	// 	fmt.Println(x)
	// }

}
