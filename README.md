# Simple API Server Architecture

不含邏輯處理的HttpServer專案結構

## _Tech_

- [Gin](https://github.com/gin-gonic/gin) - Web framework written in Go
- [Logrus](https://github.com/sirupsen/logrus) - Structured logger for Go
- [Viper](https://github.com/spf13/viper) - Complete configuration solution for Go
- [GoMail](https://github.com/go-gomail/gomail) - Simple and efficient package to send emails

## _Docker_

The Docker will expose port 8080, Change this within the
docker-compose if necessary. 

```sh
cd my-arch
docker-compose up --build -d
```