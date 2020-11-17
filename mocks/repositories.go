package mocks

import (
  "fmt"
  "time"

  "../domain/customer"
  "../domain/drink"
)

type mockCustomerRepository struct {

}

func NewMockCustomerRepository() customer.Repository {
  return &mockCustomerRepository{}
}

func (r *mockCustomerRepository)Get(id customer.ID) (*customer.Customer, error) {
  var date time.Time
  var accumulatedAlcohol float32 = 0
  if id < 5 {
    date = time.Date(2006, 12, 5, 0, 0, 0, 0, time.UTC)
  } else {
    date = time.Date(2000, 12, 5, 0, 0, 0, 0, time.UTC)

    if id > 10 {
      accumulatedAlcohol = 6
    }
  }

  return &customer.Customer{
    ID:id,
    Birth:date,
    Name:"Fake customer",
    AccumulatedAlcohol: accumulatedAlcohol,
  }, nil
}

func (r *mockCustomerRepository)Save(c *customer.Customer) error {
  return nil
}

type mockDrinkRepository struct {

}

func NewMockDrinkRepository() drink.Repository {
  return &mockDrinkRepository{}
}

func (r *mockDrinkRepository)Get(id drink.ID) (*drink.Drink, error) {
  if id == "licor" {
    return nil, nil
  } else {
    if id == "gin" {
      return nil, fmt.Errorf("fake network error")
    }
  }

  return &drink.Drink{
    ID:id, Name:"Beer", Abv:4.5,
  }, nil
}
