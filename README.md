# golang-testing-example

## 簡述

有關 `testing` 的相關操作練習。

## 安裝 docker 環境及執行程式

* clone GitHub repository
```
$ git https://github.com/yuyuancha/golang-testing-example.git
```

* 透過 docker-compose 開啟 docker 容器
```
$ docker-compose up -d
```

* 執行 main.go
```
$ docker-compose exec app go run main.go
```

* 關閉 docker 容器
```
docker-compose down
```
