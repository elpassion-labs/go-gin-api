package models

import (
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/qor/validations"
)

type Trip struct {
	gorm.Model
	ID                 uint      `json:"id"`
	StartAddress       string    `json:"start_address"`
	DestinationAddress string    `json:"destination_address"`
	Price              uint      "json:price"
	Date               time.Time "json:date"
}

func (trip Trip) Validate(db *gorm.DB) {
	if len(strings.TrimSpace(trip.StartAddress)) == 0 {
		db.AddError(validations.NewError(trip, "StartAddress", "start address must be present"))
	}
}
