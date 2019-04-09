package models

import (
	"comma-rpc-services/dao"
	"time"
	sq "github.com/Masterminds/squirrel"
)

type CourseScheduleTrainer struct {
	ID          int64
	ScheduleId  int64
	TrainerId   int64
	TrainerName string
	DataStatus  int8
	CreateTime  int64
	UpdateTime  int64
}

func AddCourseScheduleTrainer(cst *CourseScheduleTrainer)  error {
	r, err := dao.Db.Insert("t_course_schedule_trainer").
		Columns("schedule_id", "trainer_id", "trainer_name", "data_status", "create_time", "update_time").
		Values(cst.ScheduleId, cst.TrainerId, cst.TrainerName, cst.DataStatus, cst.CreateTime, cst.UpdateTime).
		Exec()
	if err != nil {
		return err
	}

	cst.ID, err = r.LastInsertId()

	return err
}

func PutCourseSchedule2Trainer(scheduleId, trainerId int64, trainerName string) error {
	r, err := dao.Db.Update("t_course_schedule_trainer").
		Set("trainer_id", trainerId).
		Set("trainer_name", trainerName).
		Set("update_time", time.Now().Unix()).
		Where(sq.Eq{"schedule_id": scheduleId}).
		Exec()
	if err != nil {
		return err
	}

	_, err = r.RowsAffected()

	return err
}