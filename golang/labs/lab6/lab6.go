package lab6

import "fmt"

type City struct {
	Country    string
	Population int
}

func NewCity(country string, population int) *City {

	return &City{
		Country:    country,
		Population: population,
	}
}

func (d *City) Bark() string {
	return "Woof"
}

func (d *City) GetCountry() string {
	return d.Country
}

func (d *City) SetCountry(age string) {
	d.Country = age
}

func RunLab6() {

	city := NewCity("Reih", 865345)
	fmt.Println(city.Bark())
	fmt.Println(city.GetCountry())
	city.SetCountry("Russland")
	fmt.Println(city.GetCountry())
}
