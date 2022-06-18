package main

import (
	"github.com/siriusol/plantscale_db/cronjob"
	"github.com/siriusol/plantscale_db/db"
)

func Init() {
	db.Init()
	cronjob.Init()
}
