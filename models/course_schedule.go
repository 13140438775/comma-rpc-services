package models

import (
	"comma-rpc-services/dao"
	sq "github.com/Masterminds/squirrel"
	"time"
	"comma-rpc-services/thrift/schedule"
)

const (
	CourseStatusWait = 1 // 待确认
	CourseStatusNO = 2   // 待开课
	PersonalCourse   = 5 // 私教课
	SelfManager      = 2 // 自主课
	BackProcess		 = 1 // 后台排课

	// CourseType
	CourseTypeLittleTeam = 4

	// course_status
	CancelCourse = 7
)

func AddCourseSchedule(cs *schedule.CourseSchedule) error {
	r, err := dao.Db.Insert("t_course_schedule").
		Columns("course_id", "course_name", "course_intensity", "calorie", "gym_id", "province_id", "province_name", "city_id", "city_name", "district_id", "district_name", "trainer_id", "trainer_name", "course_date", "price", "start_time", "end_time", "course_type", "people_num", "schedule_type", "remark", "course_status", "create_time", "update_time").
		Values(cs.CourseId, cs.CourseName, cs.CourseIntensity, cs.Calorie, cs.GymId, cs.ProvinceId, cs.ProvinceName, cs.CityId, cs.CityName, cs.DistrictId, cs.DistrictName, cs.TrainerId, cs.TrainerName, cs.CourseDate, cs.Price, cs.StartTime, cs.EndTime, cs.CourseType, cs.PeopleNum, cs.ScheduleType, cs.Remark, cs.CourseStatus, cs.CreateTime, cs.UpdateTime).
		Exec()
	if err != nil {
		return err
	}

	cs.ID, err = r.LastInsertId()
	return err
}

func GetCourseSchedule(id int64) (*schedule.CourseSchedule, error) {
	cs := schedule.CourseSchedule{}

	err := dao.Db.Select("id,start_time,schedule_status,course_status").
		From("t_course_schedule").
		Where(sq.Eq{"id": id}).
		QueryRow().
		Scan(&cs.ID, &cs.StartTime, &cs.ScheduleStatus, &cs.CourseStatus)

	return &cs, err
}

func CancelCourseSchedule(id int64) error {
	r, err := dao.Db.Update("t_course_schedule").
		Set("course_status", CancelCourse).
		Where(sq.Eq{"id": id}).
		Exec()
	if err != nil {
		return err
	}

	_, err = r.RowsAffected()
	return err
}

func PutCourseSchedule1Trainer(id, trainerId int64, trainerName string) error {
	r, err := dao.Db.Update("t_course_schedule").
		Set("trainer_id", trainerId).
		Set("trainer_name", trainerName).
		Set("update_time", time.Now().Unix()).
		Where(sq.Eq{"id": id}).
		Exec()
	if err != nil {
		return err
	}

	_, err = r.RowsAffected()
	return err
}


func PutCourseSchedule(cs *schedule.CourseSchedule) error {
	r, err := dao.Db.Update("t_course_schedule").
		Set("course_id", cs.CourseId).
		Set("course_name", cs.CourseName).
		Set("course_intensity", cs.CourseIntensity).
		Set("gym_id", cs.GymId).
		Set("province_id", cs.ProvinceId).
		Set("province_name", cs.ProvinceName).
		Set("city_id", cs.CityId).
		Set("city_name", cs.CityName).
		Set("district_id", cs.DistrictId).
		Set("district_name", cs.DistrictName).
		Set("trainer_id", cs.TrainerId).
		Set("trainer_name", cs.TrainerName).
		Set("start_time", cs.StartTime).
		Set("end_time", cs.EndTime).
		Set("people_num", cs.PeopleNum).
		Set("update_time", cs.UpdateTime).
		Where(sq.Eq{"id": cs.ID}).
		Exec()

	if err != nil {
		return err
	}

	_, err = r.RowsAffected()
	return err
}
