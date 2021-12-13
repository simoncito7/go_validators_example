// ! Paquete github.com/go-playground/validator/v10

package main

import (
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type Person struct {
	Name     string `validate:"omitempty,min=2,max=30"`
	Nickname string `json:"nick_name,omitempty" validate:"min=2,max=50"`
	Surname  string `validate:"required_with=Name,min=2,max=30"`
	Adress   string `validate:"required,min=2,max=30"`
	DNI      string `validate:"allowed_length"`
	Age      string `validate:"omitempty,numeric,is_of_legal_age"`
	Country  string `validate:"iso3166_1_alpha2"`
	IP       string `validate:"ipv6"`
}

type Gamer struct {
	Name    string   `validate:"required,min=2,max=30"`
	Jugador string   `validate:"required,min=2,max=30"`
	Bot     string   `validate:"required,min=2,max=30"`
	Games   []string `validate:"max=3,dive,min=3,max=20"`
	Console string   `validate:"omitempty,oneof=Family Sega PC"`
	Level   string   `validate:"max=50"`
}

func main() {
	valid := validator.New()
	valid.RegisterValidation("is_of_legal_age", isOfLegalAge)     // register custom validation func
	valid.RegisterAlias("allowed_length", `min=8,max=10,numeric`) // register custom alias

	simon := Person{
		Name:     "Simón",
		Nickname: "ZORRO",
		Surname:  "Valenzuela",
		Adress:   "Barracas 123",
		DNI:      "12345678399",
		Age:      "18",
		Country:  "AR",
		IP:       "2001:db8::8a2e:370:7334",
	}

	gamer := Gamer{
		Name:    "Cirilo",
		Jugador: "Hombre",
		Bot:     "americaa",
		Games:   []string{"Counter Strike", "Age of Empires II", "Super"},
		Console: "",
		Level:   "123",
	}

	err := valid.Struct(simon)
	if err != nil {
		fmt.Println(err.Error())
	}

	validateErr := valid.Struct(gamer)
	if validateErr != nil {
		fmt.Println("Custom validator error: ", validateErr)
	}
	fmt.Println("DONE")

}

// isOfLegalAge returns true if the age is a numeric value and
// if it's greater or equal than 18
func isOfLegalAge(fl validator.FieldLevel) bool {
	age, err := strconv.Atoi(fl.Field().String())
	if err != nil {
		return false
	}

	if age >= 18 {
		return true
	}

	return false
}
