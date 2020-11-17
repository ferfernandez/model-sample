package main

import (
  "fmt"
  "strconv"

  "./domain/customer"
  "./domain/drink"
  "./mocks"
)

func performDI() {
  drink.RegisterRepositories(mocks.NewMockCustomerRepository(), mocks.NewMockDrinkRepository())
}

func main() {
  var custIn, drinkIn string
  fmt.Print("ID cliente: ")
  fmt.Scanln(&custIn)
  custId, _ := strconv.Atoi(custIn)

  fmt.Print("ID bebida: ")
  fmt.Scanln(&drinkIn)

  performDI()
  e := drink.Sell(customer.ID(custId), drink.ID(drinkIn))

  if e != nil {
    fmt.Println("ups! error:", e)
  } else {
    fmt.Println("enjoy your drink!")
  }
}
