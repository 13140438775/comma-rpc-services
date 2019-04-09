package models

import (
	"comma-rpc-services/dao"
	sq "github.com/Masterminds/squirrel"
	"comma-rpc-services/thrift/schedule"
)

func GetGym(id int64) (*schedule.Gym, error) {
	cs := schedule.Gym{}

	err := dao.Db.Select("id, province_id, province_name, city_id, city_name, district_id, district_name").
		From("t_gym").
		Where(sq.Eq{"gym_status": VALID, "id": id}).
		QueryRow().
		Scan(&cs.ID, &cs.ProvinceId, &cs.ProvinceName, &cs.CityId, &cs.CityName, &cs.DistrictId, &cs.DistrictName)

	return &cs, err
}
