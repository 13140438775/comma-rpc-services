package models

import (
	"comma-rpc-services/dao"
	"comma-rpc-services/thrift/schedule"
)

const (
	// operator_type
	OperatorTypeUser    = 1
	TrainerOperatorType = 2
	AdminOperatorType   = 3

	// flow_status
	FlowStatusCancel = 7
)

func AddReservationFlow(rf *schedule.ReservationFlow) error {
	row, err := dao.Db.Insert("t_reservation_flow").
		Columns("schedule_id", "operator_id", "operator_type", "create_time", "update_time").
		Values(rf.ScheduleId, rf.OperatorId, rf.OperatorType, rf.CreateTime, rf.UpdateTime).
		Exec()
	if err != nil {
		return err
	}

	rf.ID, err = row.LastInsertId()

	return err
}
