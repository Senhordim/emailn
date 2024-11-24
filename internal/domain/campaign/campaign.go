package campaign

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string
}

type Campaign struct {
	ID        string
	Name      string
	CreatedAt time.Time
	Content   string
	Contacts  []Contact
}

// Constructor
func NewCampaign(name string, content string, contacts []Contact) (*Campaign, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	if content == "" {
		return nil, errors.New("content is required")
	}

	if len(contacts) == 0 {
		return nil, errors.New("contacts is required")
	}

	contacts = make([]Contact, len(contacts))

	for i, contact := range contacts {
		contacts[i].Email = contact.Email
	}

	return &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedAt: time.Time(time.Now()),
		Content:   content,
		Contacts:  contacts,
	}, nil
}
