package geo_lite

import "testing"

func TestExtractGeoInfo(t *testing.T) {
	file := "GeoLite2-City.mmdb"
	geo, err := NewGeo(
		WithFile(file),
	)
	if err != nil {
		t.Error(err)
		return
	}

	//ip := "192.168.11.12"
	//ip := "1.160.19.42" // Taiwan
	ip := "1.36.25.25" // Hong Kong
	//ip := "1.32.199.25" // Singapore
	//ip := "101.51.34.18" // thai
	//ip := "119.92.244.146" // philipines
	//ip := "2404:8458:f11d:c7c7:8412:801e:5179:157d"

	str, err := geo.Country(ip)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("Country: ", str)

	str, err = geo.CountryCode(ip)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("CountryCode: ", str)

	str, err = geo.Province(ip)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("Province: ", str)

	str, err = geo.City(ip)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("City: ", str)

	str, err = geo.CityCode(ip)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("CityCode: ", str)

	lat, long, err := geo.GetCoordinates(ip)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("Lat: ", lat, ", long: ", long)
}
