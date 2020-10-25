# inventory-toko-ijah


Instalation: 
1. Make sure Go already [installed](https://golang.org/doc/install) in your machine and setup your [Go Workspace](https://golang.org/doc/code.html#Workspaces)
2. Install [Govendor](https://github.com/kardianos/govendor) `go get github.com/kardianos/govendor`
3. Go to inside $GOPATH/src/. Run `git clone https://github.com/vinbyte/inventory-toko-ijah.git`
4. Go to inside $GOPATH/src/inventory-toko-ijah
5. Run `govendor sync` to pull all package needed and wait until finish
6. Install [gcc](https://www.guru99.com/c-gcc-install.html) for compile the sqlite driver
7. Install Sqlite driver :
    - For Linux/Mac :
        - `go get github.com/mattn/go-sqlite3`
        - `go install github.com/mattn/go-sqlite3`
    - For Windows, follow this tutorial : https://medium.com/@yaravind/go-sqlite-on-windows-f91ef2dacfe
8. Adjust your setting in `config/config.json` file
9. (optional) Install [Fresh](https://github.com/gravityblast/fresh) for auto restart server.
10. You're ready to go. Run `fresh` or `go run main.go`
11. Now your application is running in your specific port in config

*Dont forget to change mode to `production` in config/config.json

Depedencies :
- [Gin](https://github.com/gin-gonic/gin)
- [Viper](https://github.com/spf13/viper)
- [CORS](https://github.com/gin-contrib/cors)
- [GORM](https://github.com/jinzhu/gorm)
- [go-sqlite3](https://github.com/mattn/go-sqlite3)
