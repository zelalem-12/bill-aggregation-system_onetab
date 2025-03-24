package service

import (
	"context"
	"log"

	"github.com/mehdihadeli/go-mediatr"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/client"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/command/syncbills"
)

func ScheduleBillFetcher(
	scheduler client.Scheduler,
) {
	log.Println("Initializing scheduled bill sync job")

	err := scheduler.ScheduleJob("*/5 * * * *", func() {
		log.Println("Starting scheduled bill sync job")

		cmd := &syncbills.SyncAllBillsCommand{}
		_, err := mediatr.Send[*syncbills.SyncAllBillsCommand, *syncbills.SyncAllBillsCommandResponse](context.Background(), cmd)
		if err != nil {
			log.Printf("Error executing scheduled bill sync: %v", err)
			return
		}

		log.Println("Scheduled bill sync completed successfully")
	})

	if err != nil {
		log.Printf("Failed to schedule bill sync job: %v", err)
	}
}
