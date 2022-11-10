package share

import (
	"context"
	"encoding/json"
	"net/http"

	"firebase.google.com/go/v4/db"
)

func ListArticles(ref *db.Ref, w http.ResponseWriter, r *http.Request) {
	var v map[string]interface{}
	err := ref.Child("articles").OrderByChild("updateAt").Get(context.Background(), &v)
	if err != nil {
		writeError(w, err)
		return
	}
	js, err := json.Marshal(transformArticles(v))
	if err != nil {
		writeError(w, err)
		return
	}
	w.Write(js)
}
func transformArticles(v map[string]interface{}) []map[string]interface{} {
	var j []map[string]interface{}
	for n, m := range v {
		k := make(map[string]interface{})
		k["id"] = n
		k["updateAt"] = m.(map[string]interface{})["updateAt"]
		j = append(j, k)
	}
	return j
}
