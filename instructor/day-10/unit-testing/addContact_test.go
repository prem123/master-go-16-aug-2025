package main

import (
	"testing"
)

func Test_first_name_cannot_be_blank(t *testing.T) {

	err := addContact("", "lastname", "0987654321")

	if err == nil {
		t.Error("First name cannot be empty")
	}

	if err.Error() != "first name cannot be blank" {
		t.Error("unexpected error message")
	}
}

func Test_last_name_cannot_be_blank(t *testing.T) {

	err := addContact("firstname", "", "0987654321")

	if err == nil {
		t.Error("Last name cannot be empty")
	}

	if err.Error() != "last name cannot be blank" {
		t.Error("unexpected error message")
	}
}

func Test_phone_number_cannot_be_blank(t *testing.T) {

	// AAA -> code test structure
	// A -> Arrange
	// A -> Act
	// A -> Assert

	// Arrange -> setup/pre-condition/background/input things that are required for running your test
	phoneNumber := ""

	// Act
	err := addContact("firstname", "lastname", phoneNumber)

	// Assert
	if err == nil {
		t.Error("phone name cannot be empty")
	}

	// Assert
	if err.Error() != "phone cannot be blank" {
		t.Error("unexpected error message")
	}
}

func Test_phone_number_should_be_more_than_10_characters(t *testing.T) {

	// Arrange -> setup/pre-condition/background/input things that are required for running your test
	phoneNumber := "junk"

	// Act
	err := addContact("firstname", "lastname", phoneNumber)

	// Assert
	if err == nil {
		t.Error("phone number should be of atleast 10 digits")
	}

	// Assert
	if err.Error() != "phone number should be of atleast 10 digits" {
		t.Error("unexpected error message")
	}
}

func Test_phone_number_should_start_with_zero(t *testing.T) {

	// Arrange -> setup/pre-condition/background/input things that are required for running your test
	phoneNumber := "junkvalue0"

	// Act
	err := addContact("firstname", "lastname", phoneNumber)

	// Assert
	if err == nil {
		t.Error("phone number should start with zero")
	}

	// Assert
	if err.Error() != "phone number should start with zero" {
		t.Errorf("unexpected error message: %s", err.Error())
	}
}
