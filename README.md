# 교보문고 재고 확인

## SetUp

.env 파일에 포트, DEBUG 모드, 교보문고 URL 설정

```
PORT=":4000"
DEBUG=true

SEARCH_URL=''
STOCK_URL=''
```

## Swagger

swag 설치 후 swag init

```
$ go install github.com/swaggo/swag/cmd/swag@latest
$ swag init -g router/router.go
```

dev 모드에서만 접근 가능

```
http://localhost:4000/swagger/index.html
```
