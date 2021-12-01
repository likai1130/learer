package moconfig

import (
	"github.com/flyleft/gprofile"
	"log"
)

var AppConfig = ApplicationConfig{}

/**
MongoDB 配置
*/
type MongoConf struct {
	Hosts       []string `profile:"hosts" profileDefault:"[127.0.0.1:27017]"`
	UserName    string   `profile:"userName"`
	Password    string   `profile:"password"`
	MaxPoolSize uint64      `profile:"maxPoolSize" profileDefault:"100" `
}

type ApplicationConfig struct {
	MongoConf MongoConf `profile:"mongodb"`
}

func init() {
	config, err := gprofile.Profile(&ApplicationConfig{}, "/Users/likai/hisun/gospace/src/learner/pkg/mongo/application.yaml", true)
	if err != nil {
		log.Fatalf("Profile execute error: %s", err.Error())
	}
	AppConfig = *config.(*ApplicationConfig)
}
