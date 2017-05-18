package main


import (
	"github.com/alknopfler/tdd_api_mongodb/driver"
	"github.com/alknopfler/tdd_api_mongodb/database"
	"github.com/alknopfler/tdd_api_mongodb/api"
	"net/http"
)

func main() {
	//the database and driver data, will be passed by environment
	var drv = mongodb.MongoDB{}
	db := database.ConnectionDB{"localhost","service"}

	http.ListenAndServe(":8080", api.HandlerController(&drv,db))
}
