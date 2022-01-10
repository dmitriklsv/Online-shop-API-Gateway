package main

import (
	"log"
	"net/http"

	"github.com/arturzhamaliyev/Online-shop-API-Gateway/internal/config"
	"github.com/arturzhamaliyev/Online-shop-API-Gateway/internal/controller/rest"
)

//	@title			online-shop API
//	@version		0.0.1
//	@description	API Gateway for online-shop services

//	@contact.name	arturzhamaliyev
//	@contact.email	artur.zhamaliev@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/
func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := rest.Run(c)

	err = http.ListenAndServe(c.Port, h)
	if err != nil {
		log.Fatalln(err)
	}
}
