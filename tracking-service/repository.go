package main

type TrackingRepository interface {
	Insert(event TrackingEvent) error
}