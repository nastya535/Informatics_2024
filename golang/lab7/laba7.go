package lab7

import (
	"fmt"
)

func CalculationSumProduct(listProducts []Product) string {
	var sum float64 = 0
	for _, product := range listProducts {
		sum += product.GetPrice()
	}
	s := fmt.Sprintf("%.2f", sum)
	return s
}

func Lab7Run() {
	var lipstick Product = &Cosmetic{name: "помада", price: 400, brand: "influence"}
	var ship Product = &Ship{name: "яхта", material: "алюминий", price: 9000000}
	var doll Product = &Toy{name: "барби", price: 1500, material: "пластик"}

	doll.SetPrice(3000)
	ship.SetMaterial("сталь")
	lipstick.SetBrand("NYX")

	listProducts := []Product{lipstick, ship, doll}
	fmt.Printf("Сумма товаров, без учета скидки, равна: %v рублей \n", CalculationSumProduct(listProducts))
	lipstick.ApplyDiscount(5)
	ship.ApplyDiscount(15)
	doll.ApplyDiscount(10)
	fmt.Printf("Сумма товаров, без учета скидки, равна: %v рублей \n", CalculationSumProduct(listProducts))
}
