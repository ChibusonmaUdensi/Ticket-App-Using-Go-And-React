package models

import ("time"
    "context"
)

type Ticket struct{
	ID 			uint `json:"id" gorm:"primarykey"`
	EventID     uint `json:"event_id"`
	UserID      uint `json:"userID" gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Event       Event `json:"event" gorm:"foreignkey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Entered     bool `json:"entered" default:"false"`
	CreatedAt	time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`

	
}

type TicketRepository interface{
		GetMany (ctx context.Context, userId uint) ([]*Ticket, error)
		GetOne (ctx context.Context,userId uint, ticketId uint) (*Ticket, error)
		CreateOne (ctx context.Context, userId uint,ticket *Ticket) (*Ticket, error)
		UpdateOne(ctx context.Context,userId uint, ticketId uint, updateData map[string] interface{}) (*Ticket, error)	
}

type ValidateTicket struct{
	TicketId uint `json:"ticketId"`
	OwnerId uint `json:"ownerId"`
}