package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/vklymniuk/geo-locator/internal/locator"
	"log"
	"net/http"
	"os"
)

// GEO_DATA_PATH - /var/lib/GeoLite2-Country.mmdb
var ipLocator = locator.New(os.Getenv("GEO_DATA_PATH"))

func LookupCountry(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ip := ps.ByName("ip")
	location := ipLocator.LocateByIp(ip)
	js, err := json.Marshal(location)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	router := httprouter.New()
	router.GET("/locate/:ip", LookupCountry)

	log.Fatal(http.ListenAndServe(":8080", router))
}
