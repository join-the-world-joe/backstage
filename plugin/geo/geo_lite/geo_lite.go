package geo_lite

import (
	"github.com/oschwald/geoip2-golang"
	"go-micro-framework/abstract/geo"
	"net"
)

const (
	Name = "Geo Lite"
)

type _geo struct {
	file   string
	reader *geoip2.Reader
	opts   *Options
}

func NewGeo(opts ...Option) (geo.Geo, error) {
	var err error
	reader := new(geoip2.Reader)
	file := ""
	options := Options{}

	for _, o := range opts {
		o(&options)
	}

	if options.file != "" {
		reader, err = geoip2.Open(options.file)
		if err != nil {
			return nil, err
		}
		file = options.file
	}

	return &_geo{
		opts:   &options,
		file:   file,
		reader: reader,
	}, nil
}

func (p *_geo) Name() string {
	return Name
}

func (p *_geo) Country(ip string) (string, error) {
	record, err := p.reader.City(net.ParseIP(ip))
	if err != nil {
		return "", err
	}
	if v, ok := record.Country.Names["en"]; ok {
		return v, nil
	}
	return "", nil
}

func (p *_geo) CountryCode(ip string) (string, error) {
	record, err := p.reader.City(net.ParseIP(ip))
	if err != nil {
		return "", err
	}
	if v, ok := record.Country.Names["en"]; ok {
		return v, nil
	}
	return "", nil
}

func (p *_geo) City(ip string) (string, error) {
	record, err := p.reader.City(net.ParseIP(ip))
	if err != nil {
		return "", err
	}
	if v, ok := record.City.Names["en"]; ok {
		return v, nil
	}
	return "", nil
}

func (p *_geo) CityCode(ip string) (string, error) {
	record, err := p.reader.City(net.ParseIP(ip))
	if err != nil {
		return "", err
	}
	if v, ok := record.City.Names["en"]; ok {
		return v, nil
	}
	return "", nil
}

func (p *_geo) Province(ip string) (string, error) {
	record, err := p.reader.City(net.ParseIP(ip))
	if err != nil {
		return "", err
	}
	if len(record.Subdivisions) > 0 {
		subdivision := record.Subdivisions[0]
		if v, ok := subdivision.Names["en"]; ok {
			return v, nil
		}
	}

	if len(record.Subdivisions) > 0 {
		return record.Subdivisions[0].Names["en"], nil
	}
	return "", nil
}

func (p *_geo) GetCoordinates(ip string) (float64, float64, error) {
	record, err := p.reader.City(net.ParseIP(ip))
	if err != nil {
		return 0, 0, err
	}

	return record.Location.Latitude, record.Location.Longitude, nil
}
