package pkg

import (
	"github.com/go-yaml/yaml"
	"github.com/siriusol/plantscale_db/model"
	"io/ioutil"
	"os"
)

func GetDBConfig(path ...string) model.MySQLConf {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	confPath := dir + "/conf/db.yml"
	if len(path) > 0 {
		confPath = path[0]
	}
	confFile, err := os.Open(confPath)
	if err != nil {
		panic(err)
	}
	fileBytes, err := ioutil.ReadAll(confFile)
	if err != nil {
		panic(err)
	}
	var conf model.Conf
	err = yaml.Unmarshal(fileBytes, &conf)
	if err != nil {
		panic(err)
	}
	return conf.MySQL
}
