package drink

type ID string

type Drink struct {
  ID   ID
  Name string
  Abv  float32
}

type Repository interface {
  Get(ID) (*Drink, error)
}
