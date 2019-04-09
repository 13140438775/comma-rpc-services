package models

import (
	"comma-rpc-services/dao"
	sq "github.com/Masterminds/squirrel"
	"comma-rpc-services/thrift/schedule"
)

func GetCoursePersonal(id int64) (*schedule.CoursePersonal, error) {
	cp := schedule.CoursePersonal{}

	err := dao.Db.Select("id, course_name, intensity, calorie_per_hour").
		From("t_course_personal").
		Where(sq.Eq{"course_status": VALID, "id": id}).
		QueryRow().
		Scan(&cp.ID, &cp.CourseName, &cp.Intensity, &cp.CaloriePerHour)

	return &cp, err
}
