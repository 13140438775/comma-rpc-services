package models

import (
	"comma-rpc-services/dao"
	sq "github.com/Masterminds/squirrel"
	"comma-rpc-services/thrift/schedule"
)

func GetCourseTeam(id int64) (*schedule.CourseTeam, error) {
	ct := schedule.CourseTeam{}

	err := dao.Db.Select("id, course_name, calorie_per_hour,intensity").
		From("t_course_team").
		Where(sq.Eq{"course_status": VALID, "id": id}).
		QueryRow().
		Scan(&ct.ID, &ct.CourseName, &ct.CaloriePerHour, &ct.Intensity)

	return &ct, err
}
