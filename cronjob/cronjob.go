package cronjob

import (
	"context"
	"github.com/jasonlvhit/gocron"
	"github.com/siriusol/plantscale_db/handler/open"
)

func Init() {
	ctx := context.Background()
	err := gocron.Every(1).Hour().Do(open.NewHeartBeat, ctx, "cronjob")
	if err != nil {
		panic(err)
	}
	<-gocron.Start()
}
