package main

import "fmt"

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	User
	VIPLevel int
	Price    int
	Name     string
}

func main() {

	user := User{"송하나", "hana", 23}
	vip := VIPUser{
		User{"화랑", "hwarang", 48},
		3,
		250,
		"아무개",
	}

	fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s ID: %s VIP 레벨: %d VIP 가격: %d만원\n", vip.User.Name, vip.ID, vip.VIPLevel, vip.Price)

}
