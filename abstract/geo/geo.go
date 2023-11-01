package geo

type Geo interface {
	Name() string
	Country(string) (string, error)
	CountryCode(string) (string, error)
	City(string) (string, error)
	CityCode(string) (string, error)
	Province(string) (string, error)
	GetCoordinates(ip string) (float64, float64, error)
	//Address(string) string
}
