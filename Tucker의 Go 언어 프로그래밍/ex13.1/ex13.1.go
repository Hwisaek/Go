package main

import "fmt"

type House struct {
	Address  string
	Size     int
	Price    float64
	Category string
}

func main() {
	var house House
	house.Address = "서울시 강남구 ..."
	house.Size = 28
	house.Price = 10
	house.Category = "아파트"

	fmt.Printf("주소: %s, 사이즈: %d평, 가격: %f억원, 종류: %s\n",
		house.Address, house.Size, house.Price, house.Category)

	var house2 House = House{
		"서울시 강남구 ...",
		28,
		10,
		"아파트",
	}

	fmt.Printf("주소: %s, 사이즈: %d평, 가격: %f억원, 종류: %s\n",
		house2.Address, house2.Size, house2.Price, house2.Category)

	var house3 House = House{
		Size:     28,
		Category: "아파트",
	}

	fmt.Printf("주소: %s, 사이즈: %d평, 가격: %f억원, 종류: %s\n",
		house3.Address, house3.Size, house3.Price, house3.Category)
}
