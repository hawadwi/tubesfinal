package main

import (
	"database/sql"
)

type MySQLRepository struct {
	DB *sql.DB
}

func (r MySQLRepository) Insert(event TrackingEvent) error {

	query := `
	INSERT INTO tracking_events
	(resi, lokasi, event, timestamp)
	VALUES (?, ?, ?, ?)
	`

	_, err := r.DB.Exec(
		query,
		event.Resi,
		event.Lokasi,
		event.Event,
		event.Timestamp,
	)

	return err
}