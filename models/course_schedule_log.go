package models

import (
	"comma-rpc-services/dao"
	"comma-rpc-services/thrift/schedule"
)

const (
	BACKSTAGE = 1 // 后台
	USER      = 2 // 用户
	TRAINER   = 3 // 教练
)

func AddCourseScheduleLog(csl *schedule.CourseScheduleLog) error {
	r, err := dao.Db.Insert("t_course_schedule_log").
		Columns("course_id", "schedule_id", "operate_id", "operate_type", "create_time", "update_time").
		Values(csl.CourseId, csl.ScheduleId, csl.OperateId, csl.OperateType, csl.CreateTime, csl.UpdateTime).
		Exec()
	if err != nil {
		return err
	}

	csl.ID, err = r.LastInsertId()

	return err
}
