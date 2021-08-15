package models

import (
	"github.com/kamva/mgm/v3"
	"time"
)

// CovidData Covid Data Model Persisted in DB
type CovidData struct {
	mgm.DefaultModel `json:"-" bson:",inline"`
	Region           string    `json:"region" bson:"region"`
	ActiveCases      int64     `json:"activeCases,string,omitempty"`
	ConfirmedCases   int64     `json:"confirmedCases,string,omitempty"`
	Deaths           int64     `json:"deaths,string,omitempty"`
	Recovered        int64     `json:"recovered,string,omitempty"`
	RemoteSyncTime   time.Time `json:"remoteSyncTime"`
	CreatedAt        time.Time `bson:"created_at" json:"-"`
	UpdatedAt        time.Time `bson:"updated_at" json:"updatedAt"`
}
