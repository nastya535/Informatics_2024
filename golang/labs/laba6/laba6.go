package labs

import (
	"fmt"
)

type Country struct {
	Name       string
	Population int
	Area       float64
	Capital    string
	Language   string
}

func (c Country) Info() string {
	return fmt.Sprintf("Страна: %s, Население: %d, Площадь: %.2f, Столица: %s, Язык: %s", c.Name, c.Population, c.Area, c.Capital, c.Language)
}

func (c *Country) UpdatePopulation(newPopulation int) {
	c.Population = newPopulation
}

func (c *Country) UpdateArea(newArea float64) {
	c.Area = newArea
}

func (c *Country) UpdateCapital(newCapital string) {
	c.Capital = newCapital
}

func (c *Country) UpdateLanguage(newLanguage string) {
	c.Language = newLanguage
}

func Pask() {
	country := Country{
		Name:       "Россия",
		Population: 145912025,
		Area:       17098242,
		Capital:    "Москва",
		Language:   "Русский",
	}

	fmt.Println(country.Info())

	country.UpdatePopulation(146000000)
	fmt.Println("Обновленное население:", country.Population)

	country.UpdateArea(17098250)
	fmt.Println("Обновленная площадь:", country.Area)

	country.UpdateCapital("Санкт-Петербург")
	fmt.Println("Обновленная столица:", country.Capital)

	country.UpdateLanguage("Русский и другие")
	fmt.Println("Обновленный язык:", country.Language)
}
