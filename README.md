# inventory-toko-ijah


Instalation: 
1. Make sure Go already [installed](https://golang.org/doc/install) in your machine and setup your [Go Workspace](https://golang.org/doc/code.html#Workspaces)
2. Install [Govendor](https://github.com/kardianos/govendor) `go get github.com/kardianos/govendor`
3. Go to your $GOPATH/src/. Run `git clone https://github.com/501army/inventory-toko-ijah.git`
4. Go to inside $GOPATH/src/github.com/501army/inventory-toko-ijah
5. Run `govendor sync` to pull all package needed and wait until finish
6. Install Sqlite driver
    a. For Linux/Mac just run :
        `go get github.com/mattn/go-sqlite3`
        `go install github.com/mattn/go-sqlite3`
    b. For Windows, follow this tutorial : https://medium.com/@yaravind/go-sqlite-on-windows-f91ef2dacfe
7. Adjust your setting in `config/config.json` file
8. (optional) Install [Fresh](https://github.com/gravityblast/fresh) for auto restart server.
9. You're ready to go. Run `fresh` or `go run main.go`
10. Now your application is running in your specific port in config

*Dont forget to change mode to `production` in config/config.json

Depedencies :
- [Gin](https://github.com/gin-gonic/gin)
- [Viper](https://github.com/spf13/viper)
- [CORS](https://github.com/gin-contrib/cors)
- [GORM](https://github.com/jinzhu/gorm)
- [SqliteDriver](https://github.com/mattn/go-sqlite3)
<!-- - [gjson](https://github.com/tidwall/gjson)
- [uniqueId](https://github.com/rs/xid)
- [random_string](https://github.com/chr4/pwgen)
- [govalidator](https://github.com/asaskevich/govalidator)
- [sentry](https://github.com/getsentry/sentry-go) -->