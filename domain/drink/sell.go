package drink

import (
  "../customer"
  "github.com/pkg/errors"
)

const (
  AlcoholLimit = 5
)

var (
  customerRepository customer.Repository
  drinkRepository Repository
)

func calculateAlcoholInBlood(c* customer.Customer, avb float32) float32 {
  return c.AccumulatedAlcohol + avb * 0.25
}

func RegisterRepositories(custRepo customer.Repository, drinkRepo Repository) error {
  if custRepo == nil || drinkRepo == nil {
    return errors.Errorf("arguments cannot be nil")
  }

  customerRepository = custRepo
  drinkRepository = drinkRepo

  return nil
}

func Sell(custID customer.ID, drinkID ID) error {
  // validar que no sean nulas...
  cust, e := customerRepository.Get(custID)

  if e != nil {
    return e
  }

  if cust == nil {
    return NewCustomerNotFoundError(custID)
  }

  drink, e := drinkRepository.Get(drinkID)

  if e != nil {
    return e
  }

  if drink == nil {
    return NewDrinkNotFoundError(drinkID)
  }

  if drink.Abv > 0 && !cust.HasLegalAge() {
    return NewNotAllowedForDrinkingError()
  }

  if cust.AccumulatedAlcohol >= AlcoholLimit {
    return NewTooDrunkError()
  }

  cust.AccumulatedAlcohol = calculateAlcoholInBlood(cust, drink.Abv)
  e = customerRepository.Save(cust)

  if e != nil {
    return e
  }

  return nil
}

