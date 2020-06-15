package repository

import (
	"context"
	"errors"
	"time"
)

var (
	ErrTimeBusy = errors.New("event's time busy")
	ErrNotFound = errors.New("not found")
)

type EventID uint64

type Event struct {
	ID          EventID
	Title       string
	Datetime    time.Time
	Duration    time.Duration
	Description string
	OwnerID     uint64
}

type EventsRepo interface {
	Create(ctx context.Context, event Event) (Event, error)
	Update(ctx context.Context, event Event) (Event, error)
	Delete(ctx context.Context, id EventID) (EventID, error)
	ListOfDay(ctx context.Context, from time.Time) ([]Event, error)
	ListOfWeek(ctx context.Context, from time.Time) ([]Event, error)
	ListOfMonth(ctx context.Context, from time.Time) ([]Event, error)
}

type BaseRepo interface {
	Connect(ctx context.Context, dsn string) error
	Close() error
	EventsRepo
}