package models

import (
	"comma-rpc-services/thrift/schedule"
	"comma-rpc-services/dao"
	sq "github.com/Masterminds/squirrel"
	"fmt"
)

const CourseTeamAndPersonal = 3


func GetTrainersByGym(gymId int64, trainerType int8, trainerName string, phone string) (r []*schedule.Trainer, err error) {
	sql := fmt.Sprintf("SELECT t.id, t.trainer_name, t.phone FROM t_trainer t LEFT JOIN t_trainer_place rp ON t.id = rp.trainer_id WHERE rp.relation_status = %d AND rp.gym_id = %d AND (t.trainer_type = %d OR t.trainer_type = %d)", VALID, gymId, trainerType, CourseTeamAndPersonal)

	sql += " AND (t.trainer_name LIKE '%"+ trainerName +"%' OR t.phone LIKE '%"+ phone +"%')"

	rows, err := dao.GetSqlDb().Query(sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		t := schedule.Trainer{}
		err := rows.Scan(&t.ID, &t.Name, &t.Phone)
		if err != nil {
			return nil, err
		}
		r = append(r, &t)
	}

	return r, err
}

func GetTrainersByScheduleTime(gymId, startTime, endTime int64) ([]*schedule.Trainer, error) {
	r1, err := Get1TrainersByScheduleTime(gymId, startTime)
	if err != nil {
		return nil, err
	}

	r2, err := Get2TrainersByScheduleTime(gymId, startTime)
	if err != nil {
		return nil, err
	}

	r3, err := Get3TrainersByUnavailableTime(startTime, endTime)
	if err != nil {
		return nil, err
	}

	r1 = append(r1, r2...)
	r1 = append(r1, r3...)
	return r1, err
}

func Get1TrainersByScheduleTime(gymId, startTime int64) (r []*schedule.Trainer, err error) {
	rows, err := dao.Db.Select("cs.trainer_id,t.trainer_name, t.phone").
		From("t_course_schedule cs").
		Join("t_trainer t on t.id=cs.trainer_id").
		Where(sq.And{sq.Eq{"cs.schedule_status": VALID}, sq.Eq{"cs.gym_id": gymId}, sq.Eq{"cs.start_time": startTime}}).
		Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		t := schedule.Trainer{}
		err := rows.Scan(&t.ID, &t.Name, &t.Phone)
		if err != nil {
			return nil, err
		}
		r = append(r, &t)
	}

	return r, nil
}

func Get2TrainersByScheduleTime(gymId, startTime int64) (r []*schedule.Trainer, err error) {
	rows, err := dao.Db.Select("rt.trainer_id,rt.trainer_name,t.phone").
		From("t_course_schedule_trainer rt").
		Join("t_course_schedule cs ON cs.id=rt.schedule_id").
		Join("t_trainer t ON t.id=rt.trainer_id").
		Where(sq.And{sq.Eq{"cs.schedule_status": VALID}, sq.Eq{"rt.data_status": VALID}, sq.Eq{"cs.gym_id": gymId}, sq.Eq{"cs.start_time": startTime}}).
		Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		t := schedule.Trainer{}
		err := rows.Scan(&t.ID, &t.Name, &t.Phone)
		if err != nil {
			return nil, err
		}
		r = append(r, &t)
	}

	return r, nil
}

func Get3TrainersByUnavailableTime(startTime, endTime int64) (r []*schedule.Trainer, err error) {
	rows, err := dao.Db.Select("rt.trainer_id,rt.trainer_name,t.phone").
		From("t_unavailable_date rt").
		Join("t_trainer t ON t.id=rt.trainer_id").
		Where(sq.And{sq.Eq{"rt.data_status": VALID}, sq.Or{sq.LtOrEq{"rt.start_time" : startTime}, sq.GtOrEq{"rt.end_time" : endTime}}}).
		Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		t := schedule.Trainer{}
		err := rows.Scan(&t.ID, &t.Name, &t.Phone)
		if err != nil {
			return nil, err
		}
		r = append(r, &t)
	}

	return r, nil
}


