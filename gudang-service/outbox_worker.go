/*package main

import (
	"encoding/json"
	"log"
	"time"
)

func StartOutboxWorker(repo *PackageRepository, mq *MQClient) {

	for {
		events, err := repo.GetPendingEvents()
		if err != nil {
			log.Println("outbox error:", err)
			time.Sleep(3 * time.Second)
			continue
		}

		for _, e := range events {

			err := mq.Publish("package.ready", e.Payload)
			if err != nil {
				log.Println("publish failed:", err)
				continue
			}

			_ = repo.MarkAsSent(e.ID)
			log.Println("event sent:", e.ID)
		}

		time.Sleep(2 * time.Second)
	}
} */

package main

import (
	"log"
	"time"

	"github.com/hawadwi/gudang-service/mq"
)

func StartOutboxWorker(repo *PackageRepository) {

	for {
		events, err := repo.GetPendingEvents()
		if err != nil {
			log.Println("outbox error:", err)
			time.Sleep(3 * time.Second)
			continue
		}

		for _, e := range events {

			err := mq.Publish("package.ready", e.Payload)
			if err != nil {
				log.Println("publish failed:", err)
				continue
			}

			// Kirim ke Tracking Service
			err = mq.Publish("tracking_queue", e.Payload)
			if err != nil {
				log.Println("publish tracking_queue failed:", err)
				continue
			}

			_ = repo.MarkAsSent(e.ID)
			log.Println("event sent:", e.ID)
		}

		time.Sleep(2 * time.Second)
	}
}
