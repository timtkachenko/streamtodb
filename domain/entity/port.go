package entity

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"time"
)

type Port struct {
	ID          uuid.UUID       `gorm:"primary_key;type:uuid;" json:"id"`
	Codename    string          `json:"codename"`
	Name        string          `json:"name"`
	City        string          `json:"city"`
	Country     string          `json:"country"`
	Alias       pq.StringArray  `gorm:"type:text[]" json:"alias"`
	Regions     pq.StringArray  `gorm:"type:text[]" json:"regions"`
	Coordinates pq.Float64Array `gorm:"type:decimal[]" json:"coordinates"`
	Province    string          `json:"province"`
	Timezone    string          `json:"timezone"`
	Unlocs      pq.StringArray  `gorm:"type:text[]" json:"unlocs"`
	Code        string          `json:"code"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (p *Port) BeforeSave() {
	p.CreatedAt = time.Now().UTC()
	p.UpdatedAt = p.CreatedAt
}

func (p *Port) BeforeUpdate() (err error) {
	p.UpdatedAt = time.Now().UTC()
	return
}

func (p *Port) Validate() error {
	// TODO validate the input here
	return nil
}
