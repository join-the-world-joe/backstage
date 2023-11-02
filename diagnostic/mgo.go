package diagnostic

import (
	"backstage/common/conf"
	"backstage/global/config"
	"github.com/BurntSushi/toml"
)

var mgo_replica_conf = `
[Mongo.test]
	Servers = ["119.23.224.221:27021", "119.23.224.221:27022", "119.23.224.221:27023"]
	User = "root"
	Password = "123456"
`

var mgo_server = `
[Mongo.test]
	URI = 'mongodb://root:123456@119.23.224.221:27001/?directConnection=true'
`

func SetupMongoDB() {
	cf := &conf.MongoConf{}
	if err := toml.Unmarshal([]byte(mgo_server), &cf); err != nil {
		panic(err)
	}
	config.SetMongoConf(cf)
}
