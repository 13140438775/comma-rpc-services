/**
 * ScheduleClient Test (Benchmem test)
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
package main

import (
	"testing"
	"context"
	"git.apache.org/thrift.git/lib/go/thrift"
	"comma-rpc-services/thrift/schedule"
	"fmt"
)

var (
	ctx       context.Context
	client    *schedule.ScheduleClient
	tr thrift.TTransport
)

func init() {
	var err error
	ctx = context.Background()

	tr, err = thrift.NewTSocket("120.79.101.51:19001")
	if err != nil {
		fmt.Println(err)
		return
	}
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	iProtocol := protocolFactory.GetProtocol(tr)
	oProtocol := protocolFactory.GetProtocol(tr)
	tClient := thrift.NewTStandardClient(iProtocol, oProtocol)

	client = schedule.NewScheduleClient(tClient)
}

// go test -v -run=GetTrainersByGym
func TestGetTrainersByGym(t *testing.T) {
	if err := tr.Open(); err != nil {
		t.Errorf("transport.Open err: %s", err)
		return
	}
	defer tr.Close()
	tests := []struct {
		gymId        int64
		trainerType  int8
		trainerName  string
		phone        string
		exceptResult []*schedule.Trainer
	}{
		{117, 2, "76", "76", []*schedule.Trainer{}},
		//{2, 3, "1", "13", []*schedule.Trainer{{ID: 2}}},
	}

	for _, test := range tests {
		r, err := client.GetTrainersByGym(ctx, test.gymId, test.trainerType, test.trainerName, test.phone)
		fmt.Printf("ssss%v", r)
		if err != nil {
			t.Errorf("GetTrainersByGym(%d, %d) fail, err %s", test.gymId, test.trainerType, err)
			return
		}

		//if len(r) == 0 {
		//	t.Errorf("GetTrainersByGym(%d, %d) no data", test.gymId, test.trainerType)
		//	return
		//}
		//
		//if r[0].ID != test.exceptResult[0].ID {
		//	t.Errorf("GetTrainersByGym(%d, %d) = %v, expect %v", test.gymId, test.trainerType, r, test.exceptResult)
		//	return
		//}
	}
}

// go test -v -run=GetTrainersByScheduleTime
func TestGetTrainersByScheduleTime(t *testing.T) {
	if err := tr.Open(); err != nil {
		t.Errorf("transport.Open err: %s", err)
		return
	}
	defer tr.Close()

	tests := []struct {
		gymId        int64
		startTime    int64
		endTime      int64
		exceptResult []*schedule.Trainer
	}{
		{22, 1566666668, 1566666668, []*schedule.Trainer{{ID: 1}}},
	}

	for _, test := range tests {
		r, err := client.GetTrainersByScheduleTime(ctx, test.gymId, test.startTime, test.endTime)
		if err != nil {
			t.Errorf("GetTrainersByScheduleTime(%d, %d, %d) fail, err %s", test.gymId, test.startTime, test.endTime, err)
			return
		}

		if len(r) == 0{
			t.Errorf("GetTrainersByScheduleTime(%d, %d, %d) no data", test.gymId, test.startTime, test.endTime)
			return
		}

		if r[0].ID != test.exceptResult[0].ID {
			t.Errorf("GetTrainersByGym(%d, %d, %d) = %v, expect %v", test.gymId, test.startTime, test.endTime, r, test.exceptResult)
			return
		}
	}
}

//// go test -v -run=none -bench=GetEnTrainers -count=3
//func BenchmarkGetEnTrainers(b *testing.B) {
//	client, transport, err := getScheduleClient()
//	if err != nil {
//		b.Errorf("getThriftScheduleClient err: %s", err)
//		return
//	}
//	if err := transport.Open(); err != nil {
//		b.Errorf("transport.Open err: %s", err)
//		return
//	}
//	defer transport.Close()
//
//	tests := struct {
//		gymId        int64
//		trainerType  int8
//		startTime    int64
//		endTime      int64
//		exceptResult []*schedule.Trainer
//	}{1, 3, 1566666668, 1566666668, []*schedule.Trainer{{ID: 1}}}
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		r, err := client.GetEnTrainers(ctx, tests.gymId, tests.trainerType, tests.startTime, tests.endTime)
//		if err != nil {
//			b.Errorf("client.GetEnTrainers err: %s", err)
//			return
//		}
//
//		if r[0].ID != tests.exceptResult[0].ID {
//			b.Errorf("BenchmarkScheduleClient client.GetEnTrainers result: %+v ", r)
//			b.Errorf("BenchmarkScheduleClient client.GetEnTrainers expect: %+v ", tests.exceptResult)
//			return
//		}
//	}
//}
//
//// go test -v -run=IsReservation
//func TestIsReservation(t *testing.T) {
//	client, transport, err := getScheduleClient()
//	if err != nil {
//		t.Errorf("getThriftScheduleClient err: %s", err)
//		return
//	}
//	if err := transport.Open(); err != nil {
//		t.Errorf("transport.Open err: %s", err)
//		return
//	}
//	defer transport.Close()
//
//	tests := []struct {
//		scheduleId   int64
//		exceptResult bool
//	}{
//		{511, true},
//	}
//
//	for _, test := range tests {
//		r, err := client.IsReservation(ctx, test.scheduleId)
//		if err != nil {
//			t.Errorf("client.IsReservation err: %s", err)
//			return
//		}
//
//		if r != test.exceptResult {
//			t.Errorf("client.IsReservation result: %t ", r)
//			t.Errorf("client.IsReservation expect: %t ", test.exceptResult)
//			return
//		}
//	}
//}
//
//// go test -v -run=UserTeamCourse
//func TestUserTeamCourse(t *testing.T) {
//	client, transport, err := getScheduleClient()
//	if err != nil {
//		t.Errorf("getThriftScheduleClient err: %s", err)
//		return
//	}
//	if err := transport.Open(); err != nil {
//		t.Errorf("transport.Open err: %s", err)
//		return
//	}
//	defer transport.Close()
//
//	test := struct {
//		scheduleId   int64
//		userId       int64
//		userName     string
//		exceptResult bool
//	}{1, 1, "pb", true}
//
//	r, err := client.UserReservationLittleTeamCourse(ctx, test.scheduleId, test.userId, test.userName)
//	if err != nil {
//		t.Errorf("client.UserTeamCourse err: %s", err)
//		return
//	}
//
//	if r != test.exceptResult {
//		t.Errorf("client.UserTeamCourse result: %t ", r)
//		t.Errorf("client.UserTeamCourse expect: %t ", test.exceptResult)
//		return
//	}
//}
//
//// go test -v -run=TrainerReservationPersonalCourse
//func TestTrainerReservationPersonalCourse(t *testing.T) {
//	client, transport, err := getScheduleClient()
//	if err != nil {
//		t.Errorf("getThriftScheduleClient err: %s", err)
//		return
//	}
//	if err := transport.Open(); err != nil {
//		t.Errorf("transport.Open err: %s", err)
//		return
//	}
//	defer transport.Close()
//
//	test := struct {
//		gymId       int64
//		courseId    int64
//		trainerId   int64
//		userId      int64
//		price       int32
//		remainTimes int32
//		startTime   int64
//		endTime     int64
//		remark      string
//		actionIds   []int64
//	}{1, 5, 1, 1, 100, 9, 1537513200, 1537516800, "你是的我", []int64{11, 12, 13, 14, 15, 16, 17, 18}}
//
//	_, err = client.TrainerReservationPersonalCourse(ctx, test.gymId, test.courseId, test.trainerId, test.userId, test.price, test.remainTimes, test.startTime, test.endTime, test.remark, test.actionIds)
//	if err != nil {
//		t.Errorf("client.TrainerReservationPersonalCourse err: %s", err)
//		return
//	}
//}
//
//func TestUserPersonalClient(t *testing.T) {
//	client, transport, err := getScheduleClient()
//	if err != nil {
//		t.Errorf("getThriftScheduleClient err: %s", err)
//		return
//	}
//	if err := transport.Open(); err != nil {
//		t.Errorf("transport.Open err: %s", err)
//		return
//	}
//	defer transport.Close()
//
//	test := struct {
//		gymId       int64
//		courseId    int64
//		trainerId   int64
//		userId      int64
//		price       int32
//		remainTimes int32
//		startTime   int64
//		endTime     int64
//	}{2, 5, 1, 10, 100, 10, 1560000000, 1560000001}
//
//	_, err = client.UserReservationPersonalCourse(ctx, test.gymId, test.courseId, test.trainerId, test.userId, test.price, test.remainTimes, test.startTime, test.endTime)
//	if err != nil {
//		t.Errorf("client.GetEnTrainers err: %s", err)
//		return
//	}
//}
// go test -v -run=UserReservationTeamCourse
func TestUserReservationTeamCourse(t *testing.T) {
	if err := tr.Open(); err != nil {
		t.Errorf("transport.Open err: %s", err)
		return
	}
	tests := []struct {
		scheduleId           int64
		userId            int64
		userName          string
		userCourseId	  int64
		remainTimes       int32
	}{
		{52, 3,"pb1",10,1},
	}

	for _, test := range tests {
		_, err := client.UserReservationTeamCourse(ctx, test.scheduleId, test.userId, test.userName, test.userCourseId, test.remainTimes)
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

	_, err := client.IsEnReservation(ctx, tests.scheduleId)
	if err != nil {
		//t.Errorf("IsEnReservation(%d) fail, err %s, result %+v", tests.scheduleId, err, r)
		return
	}
}