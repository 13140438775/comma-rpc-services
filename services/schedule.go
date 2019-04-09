package services

import (
	"context"
	"comma-rpc-services/thrift/schedule"
	"comma-rpc-services/models"
	"github.com/pkg/errors"
	"time"
	"comma-rpc-services/dao"
	"strconv"
	"fmt"
)

type ScheduleService struct {
}

func (ss *ScheduleService) GetUser(ctx context.Context, id int64) (u *schedule.User, err error) {
	u, err = models.GetUser(id)
	if err != nil {
		err = fmt.Errorf("id %d, err %s", id, err)
	}

	return
}

func (ss *ScheduleService) GetGym(ctx context.Context, id int64) (g *schedule.Gym, err error) {
	g, err = models.GetGym(id)
	if err != nil {
		err = fmt.Errorf("id %d, err %s", id, err)
	}
	return
}

func (ss *ScheduleService) GetCourseTeam(ctx context.Context, id int64) (ct *schedule.CourseTeam, err error) {
	ct, err = models.GetCourseTeam(id)
	if err != nil {
		err = fmt.Errorf("id %d, err %s", id, err)
	}

	return
}

func (ss *ScheduleService) GetUserCourse(ctx context.Context, id int64) (uc *schedule.UserCourse, err error) {
	uc, err = models.GetUserCourse(id)
	if err != nil {
		err = fmt.Errorf("id %d, err %s", id, err)
	}

	return
}

func (ss *ScheduleService) GetReservation(ctx context.Context, id int64) (r *schedule.Reservation, err error) {
	r, err = models.GetReservation(id)
	if err != nil {
		err = fmt.Errorf("id %d, err %s", id, err)
	}

	return
}

func (ss *ScheduleService) GetCoursePersonal(ctx context.Context, id int64) (cp *schedule.CoursePersonal, err error) {
	cp, err = models.GetCoursePersonal(id)
	if err != nil {
		err = fmt.Errorf("id %d, err %s", id, err)
	}

	return
}

func (ss *ScheduleService) GetReservationsBySchedule(ctx context.Context, scheduleId int64) (rs []*schedule.Reservation, err error) {
	rs, err = models.GetReservationBySchedule(scheduleId)
	if err != nil {
		err = fmt.Errorf("scheduleId %d, err %s", scheduleId, err)
	}

	return
}

func (ss *ScheduleService) GetUserCourseByType(ctx context.Context, gymId, trainerId, courseId, userId int64, courseType int8) (uc *schedule.UserCourse, err error) {
	uc, err = models.GetUserCourseByType(gymId, trainerId, courseId, userId, courseType)
	if err != nil {
		err = fmt.Errorf("gymId %d, trainerId %d, courseId %d, userId %d, courseType %d, err %s",
			gymId, trainerId, courseId, userId, courseType, err)
	}

	return
}

func (ss *ScheduleService) GetTrainersByGym(ctx context.Context, gymId int64, trainerType int8,trainerName string,phone string) (ts []*schedule.Trainer, err error) {
	ts, err = models.GetTrainersByGym(gymId, trainerType, trainerName, phone)
	if err != nil {
		err = fmt.Errorf("gymId %d, trainerType %d, err %s", gymId, trainerType, err)
	}

	return
}

func (ss *ScheduleService) GetTrainersByScheduleTime(ctx context.Context, gymId int64, startTime int64, endTime int64) (ts []*schedule.Trainer, err error) {
	ts, err = models.GetTrainersByScheduleTime(gymId, startTime, endTime)
	if err != nil {
		err = fmt.Errorf("gymId %d, startTime %d, endTime %d,err %s", gymId, startTime, endTime, err)
	}

	return
}

func (ss *ScheduleService) GetUserReservationByCourse(ctx context.Context, userId, startTime int64, courseType int8) (r *schedule.Reservation, err error) {
	r, err = models.GetReservationByCourse(userId, startTime, courseType)
	if err != nil {
		err = fmt.Errorf("userId %d, startTime %d, courseType %d,err %s", userId, startTime, courseType, err)
	}

	return
}

func (ss *ScheduleService) IsEnReservation(ctx context.Context, scheduleId int64) (bool, error) {
	cs, err := models.GetCourseSchedule(scheduleId)
	if err != nil {
		return false, fmt.Errorf("scheduleId %d, err %s", scheduleId, err)
	}

	if cs.CourseStatus != models.CourseStatusNO {
		return false, fmt.Errorf("scheduleId %d, course_schedule course_status %d", scheduleId, cs.CourseStatus)
	}

	if cs.ScheduleStatus != models.VALID {
		return false, errors.Errorf("scheduleId %d, course_schedule schedule_status %d", scheduleId, cs.ScheduleStatus)
	}

	nowTime := time.Now().Unix()
	if nowTime < (cs.StartTime-7*24*3600) || nowTime > cs.StartTime {
		return false, errors.Errorf("scheduleId %d, course_schedule start_time=%d, now_time=%d", scheduleId, cs.StartTime, nowTime)
	}

	return true, err
}

func (ss *ScheduleService) AddLittleTeamCourse(ctx context.Context, g *schedule.Gym, ct *schedule.CourseTeam, startTime int64, endTime int64, trainerId int64, trainerName string, assistantTrainerId int64, assistantTrainerName string, num int8) (r int64, err error) {
	err = dao.Transaction(func() error {
		nowTime := time.Now().Unix()
		date, err := strconv.Atoi(time.Unix(startTime, 0).Format("20060102"))

		cs := &schedule.CourseSchedule{
			CourseId:        ct.ID,
			CourseName:      ct.CourseName,
			CourseIntensity: ct.Intensity,
			Calorie:         ct.CaloriePerHour,
			GymId:           g.ID,
			ProvinceId:      g.ProvinceId,
			ProvinceName:    g.DistrictName,
			CityId:          g.CityId,
			CityName:        g.CityName,
			DistrictId:      g.DistrictId,
			DistrictName:    g.DistrictName,
			TrainerId:       trainerId,
			TrainerName:     trainerName,
			CourseDate:      int64(date),
			StartTime:       startTime,
			EndTime:         endTime,
			PeopleNum:       num,
			CourseType:      models.CourseTypeLittleTeam,
			CourseStatus:    models.CourseStatusNO,
			ScheduleType:    models.BackProcess,
			CreateTime:      nowTime,
			UpdateTime:      nowTime}
		err = models.AddCourseSchedule(cs)
		if err != nil {
			return err
		}

		// need return schedule id
		r = cs.ID

		cst := &models.CourseScheduleTrainer{
			ScheduleId:  cs.ID,
			TrainerId:   assistantTrainerId,
			TrainerName: assistantTrainerName,
			DataStatus:  models.VALID,
			CreateTime:  nowTime,
			UpdateTime:  nowTime}
		if err = models.AddCourseScheduleTrainer(cst); err != nil {
			return err
		}

		csl := &schedule.CourseScheduleLog{
			CourseId:    ct.ID,
			ScheduleId:  cs.ID,
			OperateId:   trainerId,
			OperateType: models.BACKSTAGE,
			CreateTime:  nowTime,
			UpdateTime:  nowTime}
		err = models.AddCourseScheduleLog(csl)
		if err != nil {
			return err
		}

		return nil
	})

	return
}

func (ss *ScheduleService) PutLittleTeamCourse(ctx context.Context, g *schedule.Gym, ct *schedule.CourseTeam, scheduleId int64, startTime int64, endTime int64, trainerId int64, trainerName string, scheduleTrainerId int64, assistantTrainerId int64, assistantTrainerName string, num int8) (bool, error) {
	err := dao.Transaction(func() error {
		nowTime := time.Now().Unix()

		cs := &schedule.CourseSchedule{
			ID:              scheduleId,
			CourseId:        ct.ID,
			CourseName:      ct.CourseName,
			CourseIntensity: ct.Intensity,
			Calorie:         ct.CaloriePerHour,
			GymId:           g.ID,
			ProvinceId:      g.ProvinceId,
			ProvinceName:    g.DistrictName,
			CityId:          g.CityId,
			CityName:        g.CityName,
			DistrictId:      g.DistrictId,
			DistrictName:    g.DistrictName,
			TrainerId:       trainerId,
			TrainerName:     trainerName,
			StartTime:       startTime,
			EndTime:         endTime,
			PeopleNum:       num,
			UpdateTime:      nowTime}
		err := models.PutCourseSchedule(cs)
		if err != nil {
			return err
		}

		if trainerId != 0 {
			if scheduleTrainerId != 0 {
				err = models.PutCourseSchedule2Trainer(scheduleId, trainerId, assistantTrainerName)
			} else {
				cst := &models.CourseScheduleTrainer{
					ScheduleId:  scheduleId,
					TrainerId:   assistantTrainerId,
					TrainerName: assistantTrainerName,
					DataStatus:  models.VALID,
					CreateTime:  nowTime,
					UpdateTime:  nowTime}
				err = models.AddCourseScheduleTrainer(cst)
			}
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return false, err
	}
	return true, nil
}

func (ss *ScheduleService) UserReservationTeamCourse(ctx context.Context, scheduleId int64, userId int64, userName string, userCourseId int64, remainTimes int32) (bool, error) {
	err := dao.Transaction(func() error {
		nowTime := time.Now().Unix()

		r := &schedule.Reservation{
			ScheduleId:        scheduleId,
			UserId:            userId,
			UserName:          userName,
			UserCourseId:      userCourseId,
			ReservationStatus: models.ReservationStatusReservation,
			CreateTime:        nowTime,
			UpdateTime:        nowTime}
		if err := models.AddReservation(r); err != nil {
			return err
		}

		rf := &schedule.ReservationFlow{
			ScheduleId:   scheduleId,
			OperatorId:   userId,
			OperatorType: models.OperatorTypeUser,
			CreateTime:   nowTime,
			UpdateTime:   nowTime}
		err := models.AddReservationFlow(rf)
		if err != nil {
			return err
		}

		// userCourse类型为小团体次数,才需要userCourseId,remainTimes参数
		if userCourseId > 0 {
			err = models.PutUserCourseRemainTimes(userCourseId, remainTimes)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return false, err
	}
	return true, nil
}

func (ss *ScheduleService) UserReservationPersonalCourse(ctx context.Context, g *schedule.Gym, cp *schedule.CoursePersonal, u *schedule.User, userCourseId int64, trainerId int64, trainerName string, price int32, remainTimes int32, startTime int64, endTime int64, scheduleType int8) (r int64,err error) {
	err = dao.Transaction(func() error {
		timestamp := time.Now().Unix()
		date, err := strconv.Atoi(time.Unix(startTime, 0).Format("20060102"))
		if err != nil {
			return err
		}

		cs := &schedule.CourseSchedule{
			CourseId:        cp.ID,
			CourseName:      cp.CourseName,
			CourseIntensity: cp.Intensity,
			Calorie:         cp.CaloriePerHour,
			GymId:           g.ID,
			ProvinceId:      g.ProvinceId,
			ProvinceName:    g.DistrictName,
			CityId:          g.CityId,
			CityName:        g.CityName,
			DistrictId:      g.DistrictId,
			DistrictName:    g.DistrictName,
			TrainerId:       trainerId,
			TrainerName:     trainerName,
			CourseDate:      int64(date),
			Price:           price,
			StartTime:       startTime,
			EndTime:         endTime,
			CourseType:      models.PersonalCourse,
			CourseStatus:    models.CourseStatusNO,
			ScheduleType:    scheduleType,
			CreateTime:      timestamp,
			UpdateTime:      timestamp}
		err = models.AddCourseSchedule(cs)
		if err != nil {
			return err
		}

		// need return schedule id
		r = cs.ID

		csl := &schedule.CourseScheduleLog{
			CourseId:    cp.ID,
			ScheduleId:  cs.ID,
			OperateId:   u.ID,
			OperateType: models.USER,
			CreateTime:  timestamp,
			UpdateTime:  timestamp}
		err = models.AddCourseScheduleLog(csl)
		if err != nil {
			return err
		}

		r := &schedule.Reservation{
			ScheduleId:        cs.ID,
			UserId:            u.ID,
			UserName:          u.UserName,
			UserCourseId:      userCourseId,
			ReservationStatus: models.ReservationStatusReservation,
			CreateTime:        timestamp,
			UpdateTime:        timestamp}
		err = models.AddReservation(r)
		if err != nil {
			return err
		}

		rf := &schedule.ReservationFlow{
			ScheduleId:   cs.ID,
			OperatorId:   u.ID,
			OperatorType: models.OperatorTypeUser,
			CreateTime:   timestamp,
			UpdateTime:   timestamp}
		err = models.AddReservationFlow(rf)
		if err != nil {
			return err
		}

		// 教练预计收入
		trainerIncome := price * 30 / 100
		ti := &schedule.TrainerIncome{
			GymId:       g.ID,
			ScheduleId:  cs.ID,
			TrainerId:   trainerId,
			TrainerName: trainerName,
			Income:      trainerIncome,
			CreateTime:  timestamp,
			UpdateTime:  timestamp}
		err = models.AddTrainerIncome(ti)
		if err != nil {
			return err
		}

		return models.PutUserCourseRemainTimes(userCourseId, remainTimes)
	})

	return
}

func (ss *ScheduleService) TrainerReservationPersonalCourse(ctx context.Context, g *schedule.Gym, cp *schedule.CoursePersonal, u *schedule.User, userCourseId int64, trainerId int64, trainerName string, price int32, remainTimes int32, startTime int64, endTime int64, remark string, actionIds []int64) (scheduleId int64, err error) {
	err = dao.Transaction(func() error {
		timestamp := time.Now().Unix()
		date, err := strconv.Atoi(time.Unix(startTime, 0).Format("20060102"))
		if err != nil {
			return err
		}

		cs := &schedule.CourseSchedule{
			CourseId:        cp.ID,
			CourseName:      cp.CourseName,
			CourseIntensity: cp.Intensity,
			Calorie:         cp.CaloriePerHour,
			GymId:           g.ID,
			ProvinceId:      g.ProvinceId,
			ProvinceName:    g.DistrictName,
			CityId:          g.CityId,
			CityName:        g.CityName,
			DistrictId:      g.DistrictId,
			DistrictName:    g.DistrictName,
			TrainerId:       trainerId,
			TrainerName:     trainerName,
			CourseDate:      int64(date),
			Price:           price,
			StartTime:       startTime,
			EndTime:         endTime,
			CourseType:      models.PersonalCourse,
			Remark:          remark,
			CourseStatus:    models.CourseStatusWait,
			ScheduleType:    models.SelfManager,
			CreateTime:      timestamp,
			UpdateTime:      timestamp}
		err = models.AddCourseSchedule(cs)
		if err != nil {
			return err
		}

		scheduleId = cs.ID

		csl := &schedule.CourseScheduleLog{
			CourseId:    cp.ID,
			ScheduleId:  cs.ID,
			OperateId:   trainerId,
			OperateType: models.TRAINER,
			CreateTime:  timestamp,
			UpdateTime:  timestamp}
		err = models.AddCourseScheduleLog(csl)
		if err != nil {
			return err
		}

		r := &schedule.Reservation{
			ScheduleId:        cs.ID,
			UserId:            u.ID,
			UserName:          u.UserName,
			UserCourseId:      userCourseId,
			ReservationStatus: models.ReservationStatusUnconfirmed,
			CreateTime:        timestamp,
			UpdateTime:        timestamp}
		err = models.AddReservation(r)
		if err != nil {
			return err
		}

		rf := &schedule.ReservationFlow{
			ScheduleId:   cs.ID,
			OperatorId:   trainerId,
			OperatorType: models.TrainerOperatorType,
			CreateTime:   timestamp,
			UpdateTime:   timestamp}
		err = models.AddReservationFlow(rf)
		if err != nil {
			return err
		}

		// 添加训练计划
		if len(actionIds) != 0 {
			ib := dao.Db.Insert("t_user_report").
				Columns("schedule_id, user_id, user_name, action_id, create_time, update_time")
			for _, actionId := range actionIds {
				ib = ib.Values(cs.ID, u.ID, u.UserName, actionId, timestamp, timestamp)
			}
			_, err = ib.Exec()
			if err != nil {
				return err
			}
		}

		// 扣除用户的课程次数
		//err = models.ReservationCourse(cp.ID, trainerId, g.ID, u.ID, remainTimes)
		//if err != nil {
		//	return err
		//}

		// 教练预计收入
		trainerIncome := price * 30 / 100
		ti := &schedule.TrainerIncome{
			GymId:       g.ID,
			ScheduleId:  cs.ID,
			TrainerId:   trainerId,
			TrainerName: trainerName,
			Income:      trainerIncome,
			CreateTime:  timestamp,
			UpdateTime:  timestamp}
		err = models.AddTrainerIncome(ti)
		if err != nil {
			return err
		}

		return nil
	})

	return
}

func (ss *ScheduleService) CancelSchedule(ctx context.Context, scheduleId int64, rs []*schedule.Reservation) (bool, error) {
	err := dao.Transaction(func() error {
		// 取消排期
		err := models.CancelCourseSchedule(scheduleId)
		if err != nil {
			return err
		}

		// 批量取消排期已预约用户
		err = models.CancelReservationByScheduleId(scheduleId)
		if err != nil {
			return err
		}

		// 批量恢复用户小团课／私教课次数
		if len(rs) != 0 {
			return models.BatchPutUserCourseRemainTimes(rs)
		}
		return nil
	})
	if err != nil {
		return false, fmt.Errorf("scheduleId %d, rs %v, err %s", scheduleId, rs, err)
	}

	return true, nil
}

func (ss *ScheduleService) PutTrainer(ctx context.Context, scheduleId int64, trainerId int64, trainerName string, assistantTrainerId int64, assistantTrainerName string) (bool, error) {
	err := dao.Transaction(func() (err error) {
		err = models.PutCourseSchedule1Trainer(scheduleId, trainerId, trainerName)
		if err != nil {
			return err
		}

		return models.PutCourseSchedule2Trainer(scheduleId, trainerId, assistantTrainerName)
	})
	if err != nil {
		return false, fmt.Errorf("scheduleId %d, trainerId %d, trainerName %s, assistantTrainerId %d, assistantTrainerName %s, err %s", scheduleId, trainerId, trainerName, assistantTrainerId, assistantTrainerName, err)
	}

	return true, nil
}

func (ss *ScheduleService) CancelReservation(ctx context.Context, scheduleId, userId, userCourseId int64, remainTimes int32) (bool, error) {
	err := dao.Transaction(func() error {
		err := models.CancelReservation(scheduleId, userId)
		if err != nil {
			return err
		}

		nowTime := time.Now().Unix()
		rf := &schedule.ReservationFlow{
			ScheduleId:   scheduleId,
			FlowStatus:   models.FlowStatusCancel,
			OperatorId:   userId,
			OperatorType: models.OperatorTypeUser,
			CreateTime:   nowTime,
			UpdateTime:   nowTime}

		err = models.AddReservationFlow(rf)
		if err != nil {
			return err
		}

		return models.PutUserCourseRemainTimes(userCourseId, remainTimes)
	})
	if err != nil {
		return false, fmt.Errorf("scheduleId %d, userId %d, userCourseId %d, remainTimes %d, err %s", scheduleId, userId, userCourseId, remainTimes, err)
	}

	return true, nil
}
