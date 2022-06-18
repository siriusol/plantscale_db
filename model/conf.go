package model

type Conf struct {
	MySQL MySQLConf `json:"mysql"`
}

type MySQLConf struct {
	Host     string
	Username string
	Password string
	DBName   string
}
