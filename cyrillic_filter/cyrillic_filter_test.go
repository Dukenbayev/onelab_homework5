package cyrillic_filter

import (
	"reflect"
	"testing"
)

type Customer struct {
	User *User
	Email *string
	Phone string
	Age int
}

type User struct {
	Nickname string
	Surname string
}

var (
	email1 = "exampleпример@gmail.com"
	ptr_email1 = &email1
	expected_email2 ="example@gmail.com"
	ptr_email2 = &expected_email2
)

func TestFilter(t *testing.T) {
	testTable := []struct{
		myStruct Customer
		expected Customer
	}{
		{
			myStruct: Customer{},
			expected: Customer{},
		},
		{
			myStruct: Customer{&User{"MarlenМарлен","Dukenбаев"} ,ptr_email1,"8702911м0036",1},
			expected: Customer{&User{"Marlen","Duken"}, ptr_email2,"87029110036",1},
		},
		{
			myStruct: Customer{&User{"Николай-","Фамилия"},ptr_email1,"ЯЗИВД87029110036",18},
			expected: Customer{&User{"-",""},ptr_email2,"87029110036",18},
		},
	}

	for _,testCase :=range testTable{
		if err := Filter(&testCase.myStruct); err!=nil{
			t.Errorf("%v",err)
		}
		if !reflect.DeepEqual(testCase.myStruct,testCase.expected){
			t.Errorf("Incorrect value: %v expected: %v",testCase.myStruct, testCase.expected)
		}
	}
}