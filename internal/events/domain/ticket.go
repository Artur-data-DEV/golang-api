package domain

import "errors"

type TicketType string

const (
	TicketTypeHalf TicketType = "half"
	TicketTypeFull TicketType = "full"
)

var (
	ErrTicketPriceZero = errors.New("ticket price must be greater than zero")
)

type Ticket struct {
	ID         string
	EventID    string
	Spot       *Spot
	TicketType TicketType
	Price      float64
}

func IsValidTicketType(ticketType TicketType) bool {
	switch ticketType {
	case TicketTypeHalf, TicketTypeFull:
		return true
	}
	return false
}

func (t *Ticket) CalculatePrice() float64 {
	if t.TicketType == TicketTypeHalf {
		return t.Price / 2
	}
	return t.Price
}

func (t *Ticket) Validate() error {
	if t.Price <= 0 {
		return ErrTicketPriceZero
	}
	return nil
}
