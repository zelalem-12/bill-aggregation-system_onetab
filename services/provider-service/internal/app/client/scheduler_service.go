package client

type Scheduler interface {
	ScheduleJob(cronExpression string, job func()) error
}
