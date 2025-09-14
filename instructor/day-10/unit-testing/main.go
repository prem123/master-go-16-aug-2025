package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	err := addContact("first", "last", "phone")
	if err != nil {
		fmt.Println("error while adding contact: " + err.Error())
	} else {
		fmt.Println("contact added successfully")
	}
}

func addContact(firstName string, lastName string, phone string) error {
	if firstName == "" {
		return errors.New("first name cannot be blank")
	}

	if lastName == "" {
		return errors.New("last name cannot be blank")
	}

	if phone == "" {
		return errors.New("phone cannot be blank")
	} else {
		if len(phone) < 10 {
			return errors.New("phone number should be of atleast 10 digits")
		}
		firstChar := phone[0:1]
		if firstChar != "0" {
			return errors.New("phone number should start with zero")
		} else {
			_, err := strconv.Atoi(phone)
			if err != nil {
				return errors.New("phone number should only be numeric: " + err.Error())
			}
		}
	}
	return nil
}
