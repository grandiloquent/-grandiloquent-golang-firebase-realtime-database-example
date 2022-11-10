## 命令

通过代理安装依赖项

```
$env:GO111MODULE="auto";$env:https_proxy="http://127.0.0.1:10809";go get -v firebase.google.com/go
```

## 引用

- https://github.com/firebase/firebase-admin-go