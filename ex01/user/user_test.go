package user_test

import (
	. "ex01/user"
	"fmt"
	"testing"
)

func TestCall_01(t *testing.T) {
	bill := User{"Bill", "bill@email.com"}
	bill.Notify()
	fmt.Print(bill.Name)

	jill := &User{"Jill", "bill@email.com"}
	jill.Notify()
	fmt.Print(jill.Name)
}

func TestCall_02(t *testing.T) {
	user := &User{
		Name:  "Pedro Silva",
		Email: "pedro@email.com",
	}

	SendNotification(user)
}

func TestCall_03(t *testing.T) {
	admin := &Admin{
		User: User{
			Name:  "Pedro Silva",
			Email: "pedro@email.com",
		},
		Level: "super",
	}

	SendNotification(admin)

}
