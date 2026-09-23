package coursego

import "fmt"

type User struct {
	Name        string // not null
	Age         uint8  // >0 && <150
	PhoneNumber string // not null
	IsClose     bool
	Rating      float64 // >= 0 && <= 10
}

func NewUser(name string, age uint8,
	phoneNumber string, isClose bool,
	rating float64) User {
	if name == "" {
		return User{}
	}
	if age >= 150 && age <= 0 {
		return User{}
	}
	if phoneNumber == "" {
		return User{}
	}
	if rating > 10 && rating < 0 {
		return User{}
	}

	return User{
		name, age, phoneNumber, isClose, rating,
	}
}

func (u *User) ChangeName(NewName string) {
	if NewName != "" {
		u.Name = NewName
	}
}

func (u *User) ChangeAge(NewAge uint8) {
	if NewAge > 0 && NewAge < 150 {
		u.Age = NewAge
	}
}

func (u *User) ChangePhoneNumber(NewPhoneNumber string) {
	if NewPhoneNumber != "" {
		u.PhoneNumber = NewPhoneNumber
	}
}

func (u *User) CloseAccount() {
	u.IsClose = false
}

func (u *User) OpenAccount() {
	u.IsClose = true
}

func (u *User) ChangeRating(NewRating float64) {
	if NewRating >= 0 && NewRating <= 10 {
		u.Rating = NewRating
	}
}

func (u *User) RatingUP(rat float64) {
	if u.Rating+rat <= 10.0 && u.Rating+rat >= 0.0 {
		u.Rating += rat
	} else {
		fmt.Println("to mutch rat")
	}

}

func (u *User) RatingDOWN(rat float64) {
	if u.Rating-rat <= 10.0 && u.Rating-rat >= 0.0 {
		u.Rating -= rat
	} else {
		fmt.Println("to mutch rat")
	}

}

func structs() {
	user := NewUser(
		"Serg",
		64,
		"782319319",
		true,
		3,
	)
	user.RatingUP(8.1)

}
