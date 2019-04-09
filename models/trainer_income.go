package models

import (
	"comma-rpc-services/dao"
	"comma-rpc-services/thrift/schedule"
)

func AddTrainerIncome(ti *schedule.TrainerIncome) error {
	rti, err := dao.Db.Insert("t_trainer_income").
		Columns("gym_id", "schedule_id", "trainer_id", "trainer_name", "income", "create_time", "update_time").
		Values(ti.GymId, ti.ScheduleId, ti.TrainerId, ti.TrainerName, ti.Income, ti.CreateTime, ti.UpdateTime).
		Exec()
	if err != nil {
		return err
	}

	ti.ID, err = rti.LastInsertId()

	return err
}
