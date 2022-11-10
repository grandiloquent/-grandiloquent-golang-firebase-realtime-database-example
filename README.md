1. 访问 [Firebase console](https://console.firebase.google.com/) 添加一个项目
2. 构建 > Realtime Database 创建一个实时数据库
3. 修改规则

    ```
    {
        "rules": {
            ".read": false,
            ".write": false,
            "datas": {
                "articles": {
                    ".indexOn": [
                        "updateAt"
                    ]
                }
            }
        }
    }
    ```
4. 项目概览设置图标 > 项目设置 > 服务帐号
5. 执行 `go get firebase.google.com/go/v4`
 
## 命令

通过代理安装依赖项

```
$env:GO111MODULE="auto";$env:https_proxy="http://127.0.0.1:10809";go get -v firebase.google.com/go
```

## 引用

- https://github.com/firebase/firebase-admin-go