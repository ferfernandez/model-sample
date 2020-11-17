package customer

import (
  "math"
  "time"
)

const (
  LegalAge = 18
)

type ID int32

type Customer struct {
  ID   ID
  Name string
  Birth time.Time
  AccumulatedAlcohol float32
}

func (c *Customer) HasLegalAge() bool {
  today := time.Now()
  return math.Floor(today.Sub(c.Birth).Hours()/24/365) >= LegalAge
}

type Repository interface {
  Get(ID) (*Customer, error)
  Save(*Customer) error
}

