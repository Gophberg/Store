package entity

import (
	"github.com/shopspring/decimal"
	"time"
)

type Ad struct {
	Photo        []string        `json:"photo"`
	Id           uint64          `json:"id"`
	Title        string          `json:"title"`
	Content      string          `json:"content"`
	Price        decimal.Decimal `json:"price"`
	CreationDate time.Time       `json:"datecreated"`
}

type Result struct {
	Id     uint64
	Status bool
	Reason string
}

type QueryCredentials struct {
	OrderBy   string `json:"orderby"`
	Direction string `json:"direction"`
	Limit     uint64 `json:"limit"`
	Offset    uint64 `json:"offset"`
	//Qc        QueryFields
	Required bool   `json:"required"`
	Fields   string `json:"fields"`
}

//type QueryFields struct {
//	Required bool   `json:"required"`
//	Fields   string `json:"fields"`
//}
