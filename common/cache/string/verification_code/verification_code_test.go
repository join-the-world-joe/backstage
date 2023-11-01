package verification_code

import (
	"backstage/common/conf"
	"backstage/utils/random_number"
	"github.com/BurntSushi/toml"
	"testing"
)

var cluster_conf = `
	[Redis.test]
	Name = "Clustered Redis Server"
	Servers = ["192.168.130.128:7001", "192.168.130.128:7002", "192.168.130.128:7003"]
	Password = "123456"
	[Redis.test.Param]
	db = 0
`

func TestFormat(t *testing.T) {
	countryCode := "86"
	phoneNumber := "110"
	behavior := "Register"
	t.Log(GetKey(behavior, countryCode, phoneNumber))
}

func TestCreate(t *testing.T) {
	countryCode := "86"
	phoneNumber := "110"
	behavior := "Register"
	code, err := random_number.Generate(0, 9, 4)
	if err != nil {
		t.Fatal(err)
	}
	cf := &conf.CacheConf{}
	if err := toml.Unmarshal([]byte(cluster_conf), &cf); err != nil {
		t.Fatal(err)
	}

	if err := Create(behavior, countryCode, phoneNumber, code); err != nil {
		t.Fatal(err)
	}
}
