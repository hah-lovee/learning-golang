package coursego

import (
	"github.com/k0kubun/pp"
)

type User2 struct {
	Name    string
	Rating  float64
	Premium bool
}

func arr() {
	// greeting.SayHello()
	// arr := [5]int{5, 66, 7, 100, 1}

	// for i := 0; i < len(arr)-1; i++ {
	// 	if arr[i]%2 == 0 {
	// 		arr[i] *= 2
	// 	}
	// }

	// for i := 0; i < len(arr); i++ {
	// 	fmt.Printf("%d - %d", i, arr[i])
	// 	fmt.Println()
	// }
	user1 := User2{
		Name:    "Вася",
		Rating:  4,
		Premium: true,
	}

	user2 := User2{
		Name:    "kolz",
		Rating:  4,
		Premium: true,
	}

	user3 := User2{
		Name:    "vpiska",
		Rating:  5,
		Premium: false,
	}

	user4 := User2{
		Name:    "vitia",
		Rating:  5,
		Premium: true,
	}

	userArray := []*User2{&user1, &user2, &user3}

	// for i := 0; i < len(userArray); i++ {
	// 	if userArray[i].Premium {
	// 		userArray[i].Rating += 1
	// 	}
	// }
	userArray = append(userArray, &user4)
	for i, v := range userArray {
		if v.Premium {
			userArray[i].Rating += 1
		}
	}

	// for _, v := range userArray {
	// 	pp.Println(v)
	// }

	// pp.Println(user1)
	// pp.Println(user2)

	// pp.Println(user3)

	intSlice := make([]int, 0, 5)

	pp.Println(intSlice, cap(intSlice), len(intSlice))

	weather := map[int]int{
		11: 3.,
		12: -4,
	}

	for k, _ := range weather {
		weather[k] += 1
	}

	pp.Println(weather[11])

	tasteMap := make(map[int]int, 10)
	pp.Println(tasteMap)

}
