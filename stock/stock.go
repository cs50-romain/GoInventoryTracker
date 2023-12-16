package stock

import "fmt"

type Stock struct {
	Name     string
	Quantity int
}

type Stocks struct {
	Data	map[string]*Stock	
}

func Init() *Stocks {
	return &Stocks{make(map[string]*Stock)}
}

func (s *Stocks) Add(item string, quantity int) {
	if _,ok := s.Data[item]; !ok {
		s.Data[item] = &Stock{item, quantity}
	} else {
		s.Data[item].Quantity += quantity
	}
}

func (s *Stocks) Remove(item string, quantity int) bool {
	if _,ok := s.Data[item]; !ok {
		return false
	} else {
		if s.Data[item].Quantity - quantity >= 0 {
			s.Data[item].Quantity -= quantity
		} else {
			return false
		}
		return true
	}
}

func (s *Stocks) Check(item string) *Stock {
	return s.Data[item]
}

func (s *Stocks) Display() {
	for _, val := range s.Data {
		fmt.Printf("Item: %s | Quantity: %d\n", val.Name, val.Quantity)
	}
}
