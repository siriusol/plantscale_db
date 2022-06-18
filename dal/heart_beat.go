package dal

import (
	"context"
	"github.com/siriusol/plantscale_db/db"
	"github.com/siriusol/plantscale_db/model"
)

const Table_HeartBeat = "heart_beat"

func GetHeartBeat(ctx context.Context, offset, limit int) (beats []model.HeartBeat, err error) {
	err = db.Client().WithContext(ctx).Debug().Table(Table_HeartBeat).Order("id desc").Offset(offset).Limit(limit).Find(&beats).Error
	return
}

func CreateHeartBeat(ctx context.Context, beat model.HeartBeat) (err error) {
	err = db.Client().WithContext(ctx).Debug().Table(Table_HeartBeat).Save(&beat).Error
	return
}
