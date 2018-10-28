package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/vklymniuk/geo-locator/internal/locator"
	"log"
	"net/http"
	"os"
)

//var ipLocator = locator.New(os.Getenv("geo_data_path"))
var ipLocator = locator.New("/var/lib/GeoLite2-Country.mmdb")

func LookupCountry(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ip := ps.ByName("ip")
	//location := ipLocator.GetLocation(ip)
	//fmt.Fprintf(w, "hello, %s!\n", os.Getenv("GEO_DATA_PATH"))
	//fmt.Fprintf(w, "hello, %s!\n", location.GetTier())
	////fmt.Fprintf(w, "Your tier is, %s!\n", location)
}

func main() {
	router := httprouter.New()
	router.GET("/locate/:ip", LookupCountry)

	log.Fatal(http.ListenAndServe(":8080", router))
}

//
//import (
//"fmt"
//"github.com/oschwald/geoip2-golang"
//"log"
//"net"
//"os"
//)
//
//
//
//func main() {
//	geoDataPath := os.Getenv("geo_data_path")
//
//	db, err := geoip2.Open(geoDataPath)
//	//db, err := geoip2.Open("GeoIP2-City.mmdb")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()
//	// If you are using strings that may be invalid, check that ip is not nil
//	ip := net.ParseIP("81.2.69.142")
//	record, err := db.City(ip)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Printf("Portuguese (BR) city name: %v\n", record.City.Names["pt-BR"])
//}