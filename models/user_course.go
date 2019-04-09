package models

import (
	"comma-rpc-services/dao"
	sq "github.com/Masterminds/squirrel"
	"comma-rpc-services/thrift/schedule"
	"fmt"
	"strings"
)

func GetUserCourse(id int64) (*schedule.UserCourse, error) {
	uc := schedule.UserCourse{}

	err := dao.Db.Select("id, remain_times, course_type").
		From("t_user_course").
		Where(sq.Eq{"id": id, "course_status": VALID}).
		QueryRow().
		Scan(&uc.ID, &uc.RemainTimes, &uc.CourseType)

	return &uc, err
}

func GetUserCourseByType(gymId, trainerId, courseId, userId int64, courseType int8) (*schedule.UserCourse, error) {
	uc := schedule.UserCourse{}

	err := dao.Db.Select("id, remain_times, course_type").
		From("t_user_course").
		Where(sq.Eq{"course_status": VALID, "gym_id": gymId, "trainer_id": trainerId, "course_id": courseId, "user_id": userId, "course_type": courseType}).QueryRow().Scan(&uc.ID, &uc.RemainTimes, &uc.CourseType)

	return &uc, err
}

func PutUserCourseRemainTimes(id int64, remainTimes int32) error {
	rs, err := dao.Db.Update("t_user_course").
		Set("remain_times", remainTimes).
		Where(sq.Eq{"id": id}).
		Exec()
	if err != nil {
		return err
	}

	_, err = rs.RowsAffected()

	return err
}

func BatchPutUserCourseRemainTimes(rs []*schedule.Reservation) error {
	sql := "UPDATE t_user_course SET remain_times = `remain_times`+1 WHERE "
	var args []string
	for _, r := range rs {
		args = append(args, fmt.Sprintf("id = %d", r.UserCourseId))
	}
	sql += strings.Join(args, " OR ")

	_, err := dao.GetSqlDb().Exec(sql)

	return err
}
