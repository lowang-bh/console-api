# console-api

LAIN 3.0 Console API Server.

## 根据 [swagger.json](swagger.json) 生成/更新代码

```
go get -u github.com/go-swagger/go-swagger/cmd/swagger  # 安装 go-swagger
swagger generate server -f swagger.json -t gen  # 根据 swagger.json 生成/更新代码
```

> - 每次更改 [swagger.json](swagger.json) 后：
>     - [gen/](gen/) 下的文件除了 [gen/restapi/configure_console.go](gen/restapi/configure_console.go)
>       外每次均应该由 swagger 重新生成；
>     - [gen/restapi/configure_console.go](gen/restapi/configure_console.go) 里配置了实际的 handler，
>       需要酌情手动修改。

## 编译

```
docker build . -t lain/console-api:0.0.8
# go get -u github.com/golang/dep/cmd/dep  # 安装 dep
# dep ensure  # 安装依赖
```

## 运行

```
kubectl -f deployment.yaml
```

> 详情请参考 [go-swagger](https://github.com/go-swagger/go-swagger)。
