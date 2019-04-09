package models

import (
	"comma-rpc-services/dao"
	sq "github.com/Masterminds/squirrel"
	"comma-rpc-services/thrift/schedule"
)

const (
	// reservation_status
	ReservationStatusReservation = 1
	ReservationStatusCancel      = 2
	ReservationStatusUnconfirmed = 3
)

func GetReservation(id int64) (*schedule.Reservation, error) {
	r := schedule.Reservation{}

	err := dao.Db.Select("id,schedule_id,user_id,user_name,user_course_id,reservation_status,create_time,update_time").
		From("t_reservation").
		Where(sq.Eq{"id": id}).
		QueryRow().
		Scan(&r.ID, &r.ScheduleId, &r.UserId, &r.UserName, &r.UserCourseId, &r.ReservationStatus, &r.CreateTime, &r.UpdateTime)

	return &r, err
}

func GetReservationBySchedule(scheduleId int64) (rs []*schedule.Reservation, err error) {
	rows, err := dao.Db.Select("id,schedule_id,user_course_id").
		From("t_reservation").
		Where(sq.Eq{"schedule_id": scheduleId}).
		Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		r := schedule.Reservation{}
		err := rows.Scan(&r.ID, &r.ScheduleId, &r.UserCourseId)
		if err != nil {
			return nil, err
		}
		rs = append(rs, &r)
	}

	return rs, nil
}

func AddReservation(r *schedule.Reservation) error {
	rr, err := dao.Db.Insert("t_reservation").
		Columns("schedule_id", "user_id", "user_name", "user_course_id", "reservation_status", "create_time", "update_time").
		Values(r.ScheduleId, r.UserId, r.UserName, r.UserCourseId, r.ReservationStatus, r.CreateTime, r.UpdateTime).
		Exec()
	if err != nil {
		return err
	}

	r.ID, err = rr.LastInsertId()

	return err
}

func CancelReservation(scheduleId, userId int64) error {
	r, err := dao.Db.Update("t_reservation").
		Set("reservation_status", ReservationStatusCancel).
		Where(sq.Eq{"schedule_id": scheduleId, "user_id": userId}).
		Exec()
	if err != nil {
		return err
	}

	_, err = r.RowsAffected()

	return err
}

func CancelReservationByScheduleId(scheduleId int64) error {
	r, err := dao.Db.Update("t_reservation").
		Set("reservation_status", ReservationStatusCancel).
		Where(sq.Eq{"schedule_id": scheduleId}).
		Exec()
	if err != nil {
		return err
	}

	_, err = r.RowsAffected()

	return err
}

func GetReservationByCourse(userId, startTime int64, courseType int8) (*schedule.Reservation, error) {
	r := schedule.Reservation{}

	err := dao.Db.Select("r.id, r.schedule_id, r.user_course_id, r.user_id, r.user_name, r.create_time, r.update_time").
		From("t_reservation r").
		Join("t_course_schedule cs ON r.schedule_id=cs.id").
		Where(sq.Eq{"r.user_id": userId, "cs.start_time": startTime, "cs.course_type": courseType, "cs.schedule_status": VALID}, sq.NotEq{"r.reservation_status": ReservationStatusCancel}).
		QueryRow().
		Scan(&r.ID, &r.ScheduleId, &r.UserCourseId, &r.UserId, &r.UserName, &r.CreateTime, &r.UpdateTime)

	return &r, err
}
