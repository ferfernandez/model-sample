package drink

import (
  "fmt"

  "../customer"
)

type drinkNotFoundError struct {
  id ID
}

func NewDrinkNotFoundError(id ID) *drinkNotFoundError {
  return &drinkNotFoundError{id}
}

type customerNotFoundError struct {
  id customer.ID
}

func NewCustomerNotFoundError(id customer.ID) *customerNotFoundError {
  return &customerNotFoundError{id}
}

func (e *drinkNotFoundError)Error() string {
  return fmt.Sprintf("Drink %s wasn't found", e.id)
}

func (e *customerNotFoundError)Error() string {
  return fmt.Sprintf("Customer %d wasn't found", int32(e.id))
}

type tooDrunkError struct {
}

func NewTooDrunkError() *tooDrunkError {
  return &tooDrunkError{}
}

func (e *tooDrunkError)Error() string {
  return "The customer may get drunk..."
}

type notAllowedForDrinkingError struct {
}

func NewNotAllowedForDrinkingError() *notAllowedForDrinkingError {
  return &notAllowedForDrinkingError{}
}

func (e *notAllowedForDrinkingError) Error() string{
  return "Tell him/her to come back when he/she is 18"
}
