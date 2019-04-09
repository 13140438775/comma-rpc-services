namespace go thrift.schedule
namespace php thrift.schedule

struct Gym {
	1:i64 id
	2:i32 provinceId
	3:string provinceName
	4:i32 cityId
	5:string cityName
	6:i32 districtId
	7:string districtName
}

struct Trainer {
    1:i64 id
    2:string name
    3:string phone
}

struct User {
	1:i64 id
	2:string userName
	3:string phone
}

struct UserCourse {
    1:i64 id
    2:i32 remainTimes
    3:i8 courseType
}

struct CourseTeam {
    1:i64 id
	2:string courseName
	3:i8 intensity
	4:i32 caloriePerHour
}

struct CoursePersonal {
	1:i64 id
	2:string courseName
	3:i8 intensity
	4:i32 caloriePerHour
}

struct CourseSchedule  {
	1:i64 id
	2:i64 courseId
	3:string courseName
	4:i8 courseIntensity
	5:i32 calorie
	6:i64 gymId
	7:i32 provinceId
	8:string provinceName
	9:i32 cityId
	10:string cityName
	11:i32 districtId
	12:string districtName
	13:i64 trainerId
	14:string trainerName
	15:i64 CourseDate
	16:i32 price
	17:i64 startTime
	18:i64 endTime
	19:i32 courseType
	20:i8 peopleNum
	21:i32 scheduleStatus
	22:string remark
	23:i64 createTime
	24:i64 updateTime
	25:i8 CourseStatus
	26:i8 ScheduleType
}

struct CourseScheduleLog{
	1:i64 id
	2:i64 courseId
	3:i64 scheduleId
	4:i64 operateId
	5:i8 OperateType
	6:i64 createTime
	7:i64 updateTime
}

struct Reservation {
    1:i64 id
    2:i64 scheduleId
    3:i64 userId
    4:string userName
    5:i64 userCourseId
    6:i8 reservationStatus
    7:i64 createTime
    8:i64 updateTime
}

struct ReservationFlow {
	1:i64 id
	2:i64 scheduleId
	3:i8 flowStatus
	4:i64 operatorId
	5:i8 operatorType
	6:i64 createTime
	7:i64 updateTime
}

struct TrainerIncome {
	1:i64 id
	2:i64 gymId
	3:i64 scheduleId
	4:i64 trainerId
	5:string trainerName
	6:i32 income
	7:i64 createTime
    8:i64 updateTime
}

service Schedule{
    User GetUser(1:i64 id)

    Gym GetGym(1:i64 id)

    CourseTeam GetCourseTeam(1:i64 id)

    UserCourse GetUserCourse(1:i64 id)

    UserCourse GetUserCourseByType(1:i64 gymId,2:i64 trainerId, 3:i64 courseId, 4:i64 userId, 5:i8 courseType)

    CoursePersonal GetCoursePersonal(1:i64 id)

    Reservation GetReservation(1:i64 id)

    Reservation GetUserReservationByCourse(1:i64 userId, 2:i64 startTime, 3:i8 courseType)

    list<Reservation> GetReservationsBySchedule(1:i64 scheduleId)

    list<Trainer> GetTrainersByGym(1:i64 gymId, 2:i8 trainerType, 3:string trainerName, 4:string phone)

    list<Trainer> GetTrainersByScheduleTime(1:i64 gymId, 2:i64 startTime, 3:i64 endTime)

    bool IsEnReservation(1:i64 scheduleId)

    i64 AddLittleTeamCourse(1:Gym g,2:CourseTeam ct,3:i64 startTime, 4:i64 endTime,5:i64 trainerId,6:string trainerName, 7:i64 assistantTrainerId, 8:string assistantTrainerName, 9:i8 num)

    bool PutLittleTeamCourse(1:Gym g,2:CourseTeam ct, 3:i64 scheduleId, 4:i64 startTime, 5:i64 endTime, 6:i64 trainerId, 7:string trainerName, 8:i64 scheduleTrainerId, 9:i64 assistantTrainerId, 10:string assistantTrainerName, 11:i8 num)

    bool UserReservationTeamCourse(1:i64 scheduleId, 2:i64 userId, 3:string userName, 4:i64 userCourseId, 5:i32 remainTimes)

    i64 UserReservationPersonalCourse(1:Gym g, 2:CoursePersonal cp, 3:User u, 4:i64 userCourseId, 5:i64 trainerId, 6:string trainerName, 7:i32 price, 8:i32 remainTimes, 9:i64 startTime, 10:i64 endTime, 11:i8 scheduleType)

    i64 TrainerReservationPersonalCourse(1:Gym g, 2:CoursePersonal cp, 3:User u, 4:i64 userCourseId, 5:i64 trainerId, 6:string trainerName, 7:i32 price, 8:i32 remainTimes, 9:i64 startTime, 10:i64 endTime, 11:string remark, 12:list<i64> actionIds)

    bool CancelSchedule(1:i64 scheduleId, 2:list<Reservation> rs)

    bool PutTrainer(1:i64 scheduleId, 2:i64 trainerId, 3:string trainerName, 4:i64 assistantTrainerId, 5:string assistantTrainerName)

    bool CancelReservation(1:i64 scheduleId, 2:i64 userId, 3:i64 userCourseId, 4:i32 remainTimes)
}