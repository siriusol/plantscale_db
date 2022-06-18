package pkg

import "testing"

func TestGetDBConfig(t *testing.T) {
	config := GetDBConfig("../conf/db.yml")
	t.Logf("%+v", config)
}
