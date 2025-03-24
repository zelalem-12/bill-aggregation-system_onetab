package client

import (
	"github.com/robfig/cron/v3"
	clientPort "github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/client"
)

type CronScheduler struct {
	scheduler *cron.Cron
}

func NewCronScheduler(cron *cron.Cron) clientPort.Scheduler {
	return &CronScheduler{
		scheduler: cron,
	}
}

func (c *CronScheduler) ScheduleJob(cronExpression string, job func()) error {
	_, err := c.scheduler.AddFunc(cronExpression, job)
	if err != nil {
		return err
	}
	return nil
}
