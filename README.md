# console-api

LAIN 3.0 Console API Server.

## 根据 [swagger.yml](swagger.yml) 生成/更新代码

```
go get -u github.com/go-swagger/go-swagger/cmd/swagger  # 安装 go-swagger
swagger generate server  # 根据 swagger.yml 生成/更新代码
```

## 编译

```
go get -u github.com/golang/dep/cmd/dep  # 安装 dep
dep ensure  # 安装依赖
go install github.com/laincloud/console-api/cmd/console-server
```

## 运行

```
console-server --host=0.0.0.0 --port=8080
```

> 详情请参考 [go-swagger](https://github.com/go-swagger/go-swagger)。
