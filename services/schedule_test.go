/**
 * schedule service Test (Basic test)
 *
 * Basic test:
 * 		go test -v -run=?
 * Benchmark test:
 * 		go test -v -run=none -bench=? -count=3
 * Benchmem test:
 * 		go test -v -run=none -test.bench=? -test.benchmem -count=3
 * Coverage test:
 *		go test -run=? -coverprofile=c.out
 *		go tool cover -html=c.out
 */
package services

import (
	"comma-rpc-services/dao"
	"testing"
	"comma-rpc-services/thrift/schedule"
	"context"
	"fmt"
)

var (
	ctx context.Context
	ss  = ScheduleService{}
)

func init() {
	err := dao.Init("comma-admin:comma-admin@tcp(120.79.101.51:3306)/comma?charset=utf8")
	if err != nil {
		return
	}
}

// go test -v -run=GetUser
func TestGetUser(t *testing.T) {
	tests := []struct {
		id           int64
		exceptResult *schedule.User
	}{
		{1, &schedule.User{ID: 1}},
	}

	for _, test := range tests {
		r, err := ss.GetUser(ctx, test.id)
		if err != nil {
			t.Errorf("GetUser(%d) fail, err %s, result %+v", test.id, err, r)
			return
		}

		if r.ID != test.exceptResult.ID {
			t.Errorf("GetUser(%d) = %+v, except %+v", test.id, r, test.exceptResult)
			return
		}
	}
}

// go test -v -run=GetGym
func TestGetGym(t *testing.T) {
	tests := []struct {
		id           int64
		exceptResult *schedule.Gym
	}{
		{1, &schedule.Gym{ID: 1}},
		{9, &schedule.Gym{ID: 9}},
	}

	for _, test := range tests {
		r, err := ss.GetGym(ctx, test.id)
		if err != nil {
			t.Errorf("GetGym(%d) fail, err %s, result %+v", test.id, err, r)
			return
		}

		if r.ID != test.exceptResult.ID {
			t.Errorf("GetGym(%d) = %+v, except %+v", test.id, r, test.exceptResult)
			return
		}
	}
}

// go test -v -run=GetCourseTeam
func TestGetCourseTeam(t *testing.T) {
	tests := []struct {
		id           int64
		exceptResult *schedule.CourseTeam
	}{
		{1, &schedule.CourseTeam{ID: 1}},
		{2, &schedule.CourseTeam{ID: 2}},
	}

	for _, test := range tests {
		r, err := ss.GetCourseTeam(ctx, test.id)
		if err != nil {
			t.Errorf("GetCourseTeam(%d) fail, err %s, result %+v", test.id, err, r)
			return
		}

		if r.ID != test.exceptResult.ID {
			t.Errorf("GetCourseTeam(%d) = %+v, except %+v", test.id, r, test.exceptResult)
			return
		}
	}
}

// go test -v -run=GetUserCourse
func TestGetUserCourse(t *testing.T) {
	tests := []struct {
		id           int64
		exceptResult *schedule.UserCourse
	}{
		{9, &schedule.UserCourse{ID: 9}},
		{10, &schedule.UserCourse{ID: 10}},
	}

	for _, test := range tests {
		r, err := ss.GetUserCourse(ctx, test.id)
		if err != nil {
			t.Errorf("GetUserCourse(%d) fail, err %s, result %+v", test.id, err, r)
			return
		}

		if r.ID != test.exceptResult.ID {
			t.Errorf("GetUserCourse(%d) = %+v, except %+v", test.id, r, test.exceptResult)
			return
		}
	}
}

// go test -v -run=GetReservation
func TestGetReservation(t *testing.T) {
	tests := []struct {
		id           int64
		exceptResult *schedule.Reservation
	}{
		{9, &schedule.Reservation{ID: 9}},
		{10, &schedule.Reservation{ID: 10}},
	}

	for _, test := range tests {
		r, err := ss.GetReservation(ctx, test.id)
		if err != nil {
			t.Errorf("GetReservation(%d) fail, err %s, result %+v", test.id, err, r)
			return
		}

		if r.ID != test.exceptResult.ID {
			t.Errorf("GetReservation(%d) = %+v, except %+v", test.id, r, test.exceptResult)
			return
		}
	}
}

// go test -v -run=GetCoursePersonal
func TestGetCoursePersonal(t *testing.T) {
	tests := []struct {
		id           int64
		exceptResult *schedule.CoursePersonal
	}{
		{4, &schedule.CoursePersonal{ID: 4}},
		{5, &schedule.CoursePersonal{ID: 5}},
	}

	for _, test := range tests {
		r, err := ss.GetCoursePersonal(ctx, test.id)
		if err != nil {
			t.Errorf("GetCoursePersonal(%d) fail, err %s, result %+v", test.id, err, r)
			return
		}

		if r.ID != test.exceptResult.ID {
			t.Errorf("GetCoursePersonal(%d) = %+v, except %+v", test.id, r, test.exceptResult)
			return
		}
	}
}

// go test -v -run=GetReservationsBySchedule
func TestGetReservationsBySchedule(t *testing.T) {
	tests := []struct {
		scheduleId   int64
		exceptResult []*schedule.Reservation
	}{
		{126, []*schedule.Reservation{{ID: 63}}},
	}

	for _, test := range tests {
		r, err := ss.GetReservationsBySchedule(ctx, test.scheduleId)
		if err != nil {
			t.Errorf("GetReservationsBySchedule(%d) fail, err %s, result %+v", test.scheduleId, err, r)
			return
		}

		if r[0].ID != test.exceptResult[0].ID {
			t.Errorf("GetReservationsBySchedule(%d) = %+v, except %+v", test.scheduleId, r, test.exceptResult)
			return
		}
	}
}

// go test -v -run=GetUserCourseByType
func TestGetUserCourseByType(t *testing.T) {
	tests := []struct {
		gymId        int64
		trainerId    int64
		courseId     int64
		userId       int64
		courseType   int8
		exceptResult *schedule.UserCourse
	}{
		{1, 1, 49, 1, 3, &schedule.UserCourse{ID: 9}},
	}

	for _, test := range tests {
		r, err := ss.GetUserCourseByType(ctx, test.gymId, test.trainerId, test.courseId, test.userId, test.courseType)
		if err != nil {
			t.Errorf("GetUserCourseByType(%d, %d, %d, %d, %d) fail, err %s, result %+v", test.gymId, test.trainerId, test.courseId, test.userId, test.courseType, err, r)
			return
		}

		if r.ID != test.exceptResult.ID {
			t.Errorf("GetUserCourseByType(%d, %d, %d, %d, %d) = %+v, except %+v", test.gymId, test.trainerId, test.courseId, test.userId, test.courseType, r, test.exceptResult)
			return
		}
	}
}

// go test -v -run=GetTrainersByGym
func TestGetTrainersByGym(t *testing.T) {
	tests := []struct {
		gymId        int64
		trainerType  int8
		trainerName  string
		phone        string
		exceptResult []*schedule.Trainer
	}{
		{117, 2, "76", "76", []*schedule.Trainer{{ID: 1}}},
		//{2, 3, "1", "13", []*schedule.Trainer{{ID: 2}}},
	}

	for _, test := range tests {
		r, err := ss.GetTrainersByGym(ctx, test.gymId, test.trainerType, test.trainerName, test.phone)
		fmt.Printf("%v", r)
		if err != nil {
			t.Errorf("GetTrainersByGym(%d, %d) fail, err %s, result %+v", test.gymId, test.trainerType, err, r)
			return
		}

		if len(r) == 0 {
			t.Errorf("GetTrainersByGym(%d, %d) = %v, expect %v", test.gymId, test.trainerType, r, test.exceptResult)
			return
		}

		if r[0].ID != test.exceptResult[0].ID {
			t.Errorf("GetTrainersByGym(%d, %d) = %+v, except %+v", test.gymId, test.trainerType, r, test.exceptResult)
			return
		}
	}
}

// go test -v -run=GetTrainersByScheduleTime
func TestGetTrainersByScheduleTime(t *testing.T) {
	tests := []struct {
		gymId        int64
		startTime    int64
		endTime      int64
		exceptResult []*schedule.Trainer
	}{
		{1, 1, 1, []*schedule.Trainer{{ID: 1}}},
		{2, 1, 1, []*schedule.Trainer{{ID: 2}}},
	}

	for _, test := range tests {
		r, err := ss.GetTrainersByScheduleTime(ctx, test.gymId, test.startTime, test.endTime)
		if err != nil {
			t.Errorf("GetTrainersByScheduleTime(%d, %d, %d) fail, err %s, result %+v", test.gymId, test.startTime, test.endTime, err, r)
			return
		}

		if len(r) == 0 {
			t.Errorf("GetTrainersByScheduleTime(%d, %d, %d) no data", test.gymId, test.startTime, test.endTime)
			return
		}

		if r[0].ID != test.exceptResult[0].ID {
			t.Errorf("GetTrainersByScheduleTime(%d, %d, %d) = %+v, except %+v", test.gymId, test.startTime, test.endTime, r, test.exceptResult)
			return
		}
	}
}

// go test -v -run=GetUserReservationByCourse
func TestGetUserReservationByCourse(t *testing.T) {
	tests := []struct {
		userId       int64
		startTime    int64
		courseType   int8
		exceptResult *schedule.Reservation
	}{
		{1, 3, 1, &schedule.Reservation{ID: 1}},
		{2, 3, 1, &schedule.Reservation{ID: 2}},
	}

	for _, test := range tests {
		r, err := ss.GetUserReservationByCourse(ctx, test.userId, test.startTime, test.courseType)
		if err != nil {
			t.Errorf("GetTrainersByGym(%d, %d, %d) fail, err %s, result %+v", test.userId, test.startTime, test.courseType, err, r)
			return
		}

		if r.ID != test.exceptResult.ID {
			t.Errorf("GetTrainersByGym(%d, %d, %d) = %+v, except %+v", test.userId, test.startTime, test.courseType, r, test.exceptResult)
			return
		}
	}
}

// go test -v -run=PutTrainer
func TestPutTrainer(t *testing.T) {
	tests := []struct {
		scheduleId           int64
		trainerId            int64
		trainerName          string
		assistantTrainerId   int64
		assistantTrainerName string
		exceptResult         bool
	}{
		{51, 3,"pb",1,"xx",true},
	}

	for _, test := range tests {
		r, err := ss.PutTrainer(ctx, test.scheduleId, test.trainerId, test.trainerName, test.assistantTrainerId, test.assistantTrainerName)
		if err != nil {
			t.Errorf("PutTrainer(%d, %d, %s %d, %s) fail, err %s, result %+v", test.scheduleId, test.trainerId, test.trainerName, test.assistantTrainerId, test.assistantTrainerName, err, r)
			return
		}

		if r != test.exceptResult {
			t.Errorf("PutTrainer(%d, %d, %s %d, %s) = %+v, except %+v", test.scheduleId, test.trainerId, test.trainerName, test.assistantTrainerId, test.assistantTrainerName, r, test.exceptResult)
			return
		}
	}
}


// go test -v -run=UserReservationTeamCourse
func TestUserReservationTeamCourse(t *testing.T) {
	tests := []struct {
		scheduleId   int64
		userId       int64
		userName     string
		userCourseId int64
		remainTimes  int32
	}{
		{51, 3, "pb", 10, 1},
	}

	for _, test := range tests {
		_, err := ss.UserReservationTeamCourse(ctx, test.scheduleId, test.userId, test.userName, test.userCourseId, test.remainTimes)
		if err != nil {
			t.Errorf("PutTrainer, err %s", err)
			return
		}
	}
}


// go test -v -run=IsEnReservation
func TestIsEnReservation(t *testing.T) {
	tests := struct {
		scheduleId           int64
	}{128}

	_, err := ss.IsEnReservation(ctx, tests.scheduleId)
	if err != nil {
		//t.Errorf("IsEnReservation(%d) fail, err %s, result %+v", tests.scheduleId, err, r)
		return
	}
}
