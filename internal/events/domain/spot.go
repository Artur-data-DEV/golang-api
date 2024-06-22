package domain

import (
	"errors"

	"github.com/google/uuid"
)

type SpotStatus string

const (
	SpotStatusAvaiable SpotStatus = "avaiable"
	SpotStatusSold     SpotStatus = "sold"
)

var (
	ErrInvalidSpotNumber       = errors.New("invalid spot number")
	ErrSpotNotFound            = errors.New("spot not found")
	ErrSpotAlreadyReserved     = errors.New("spot already reserved")
	ErrSpotNameTwoCharacters   = errors.New("spot name must be at least 2 characters")
	ErrSpotNameStartWithLetter = errors.New("spot name must start with an uppercase letter")
	ErrSpotNameEndWithNumber   = errors.New("spot name must end with a number")
)

type Spot struct {
	ID       string     `json:"id"`
	EventID  string     `json:"event_id"`
	Name     string     `json:"name"`
	Status   SpotStatus `json:"status"`
	TicketID string     `json:"ticket_id,omitempty"`
}

func NewSpot(event *Event, name string, status SpotStatus) (*Spot, error) {
	spot := &Spot{
		ID:      uuid.New().String(),
		EventID: event.ID,
		Name:    name,
		Status:  SpotStatusAvaiable,
	}

	if err := spot.Validate(); err != nil {
		return nil, err
	}
	return spot, nil
}

func (s *Spot) Validate() error {
	if len(s.Name) == 0 {
		return ErrSpotNameTwoCharacters
	}
	if len(s.Name) < 2 {
		return ErrSpotNameTwoCharacters
	}
	if s.Name[0] < 'A' || s.Name[0] > 'Z' {
		return ErrSpotNameStartWithLetter
	}
	if s.Name[1] < '0' || s.Name[1] > '9' {
		return ErrSpotNameEndWithNumber
	}
	if s.Status != SpotStatusAvaiable && s.Status != SpotStatusSold {
		return ErrInvalidSpotNumber
	}
	return nil
}

func (s *Spot) Reserve(ticketID string) error {
	if s.Status == SpotStatusSold {
		return ErrSpotAlreadyReserved
	}
	s.Status = SpotStatusSold
	s.TicketID = ticketID
	return nil
}
