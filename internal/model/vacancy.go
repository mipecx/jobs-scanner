// Package model contains Vacancy model
package model

import (
	"time"
)

type Vacancy struct {
	Title        string    `json:"title"`
	Company      string    `json:"company,omitempty"`
	Salary       string    `json:"salary,omitempty"`
	URL          string    `json:"url,omitempty"`
	Source       string    `json:"source"`
	NormalizedAt time.Time `json:"published_at"`
	RawDate      string    `json:"raw_date,omitempty"`
	Skills       []string  `json:"skills,omitempty"`
	Description  string    `json:"description,omitempty"`
	RawText      string    `json:"raw_text,omitempty"`
}
