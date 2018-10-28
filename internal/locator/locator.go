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
	name string
}

type Location struct {
	countryNames  map[string]string
	isoCode string
	tier *Tier
}

type GetLocation interface {
	GetLocation(ip string) *Location
}

type GetTier interface {
	GetTier(location *Location) *Tier
}

func (location *Location) GetTier() (*Tier) {
	return location.tier
}

func (locator *IpLocator) GetLocation(ip string) (*Location) {
	country, countryErr := locator.reader.Country(net.ParseIP(ip))

	if countryErr != nil {
		return &Location{
			tier: &Tier{"default-tier-name"},
		}
	}

	return &Location{
		countryNames: country.Country.Names,
		isoCode: country.Country.IsoCode,
		tier: &Tier{"some-tier-name"},
	}
}

func New(filePath string) *IpLocator {
	reader, err := geoip2.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	return &IpLocator{reader:reader}
}
