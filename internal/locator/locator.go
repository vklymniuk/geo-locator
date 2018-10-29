package locator

import (
	"github.com/oschwald/geoip2-golang"

	"net"

	"log"
)

type IpLocator struct {
	reader *geoip2.Reader
}

type Tier struct {
	Name string `json:"name"`
}

type Location struct {
	CountryNames map[string]string `json:"country_names"`
	IsoCode      string            `json:"iso_code"`
	Tier         *Tier             `json:"tier"`
}

type FetchLocation interface {
	LocateByIp(ip string) *Location
}

func (locator *IpLocator) LocateByIp(ip string) *Location {
	country, countryErr := locator.reader.Country(net.ParseIP(ip))

	if countryErr != nil {
		return &Location{
			Tier: &Tier{"tier-2"},
		}
	}

	return &Location{
		CountryNames: country.Country.Names,
		IsoCode:      country.Country.IsoCode,
		Tier:         &Tier{"some-tier-name"},
	}
}

func New(filePath string) *IpLocator {
	reader, err := geoip2.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	return &IpLocator{reader: reader}
}
