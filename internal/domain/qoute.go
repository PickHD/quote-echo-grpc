package domain

import (
	"time"

	"github.com/google/uuid"
)

type (
	Quote struct {
		ID        uuid.UUID `db:"id"`
		Body      string    `db:"body"`
		CreatedAt time.Time
		UpdatedAt *time.Time
	}

	CreateOrUpdateQuoteRequest struct {
		Body string
	}

	QuoteResponse struct {
		ID        uuid.UUID
		Body      string
		CreatedAt time.Time
		UpdatedAt *time.Time
	}
)

func NewQuote(body string) Quote {
	return Quote{
		Body: body,
	}
}
