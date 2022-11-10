## 命令

通过代理安装依赖项

```
$env:GO111MODULE="auto";$env:https_proxy="http://127.0.0.1:10809";go get -v firebase.google.com/go
```

## 引用

- https://github.com/firebase/firebase-admin-go

```
package share

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"firebase.google.com/go/v4/db"
)

func ListArticles(ref *db.Ref, ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var v []map[string]interface{}
	nodes, err := ref.Child("articles").OrderByChild("updateAt").GetOrdered(ctx)
	for _, n := range nodes {
		var x map[string]interface{}
		fmt.Println(n.Key())
		err = ref.Child("articles/"+n.Key()).Get(ctx, &x)
		if err != nil {
			writeError(w, err)
			return
		}
		v = append(v, x)
	}
	if err != nil {
		writeError(w, err)
		return
	}
	js, err := json.Marshal(v)
	if err != nil {
		writeError(w, err)
		return
	}
	w.Write(js)
}
```"# -grandiloquent-golang-firebase-realtime-database-example" 
