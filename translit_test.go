package translit_test

import (
	"fmt"

	"github.com/mdigger/translit"
)

func ExampleRu() {
	tests := []string{
		"Проверочная СТРОКА для транслитерации",
		"ЧАЩА",
		"ЧаЩа",
		"Чаща",
		"чаЩА",
		"Наташа",
		"Игорь",
	}
	for _, text := range tests {
		fmt.Println(translit.Ru(text))
	}
	// Output:
	// Proverochnaja STROKA dlja transliteracii
	// CHASCHA
	// ChaScha
	// Chascha
	// chaSCHA
	// Natasha
	// Igor
}

func ExampleKz() {
	kzNames := []string{
		"Айсұлу",
		"Айханым",
		"Нұргүл",
		"Дәурен",
		"Қанат",
		"Қанаш",
		"Қасым",
		"Турғаным",
		"Мақсат",
		"Ұлболсын",
		"Ғани",
		"Ғалымжан",
		"Ігілік",
		"СұнҚар",
		"Біржан",
		"Мағжан",
		"Төрежан",
	}
	for _, text := range kzNames {
		fmt.Println(translit.Kz(text))
	}
	// Output:
	// Aiysulý
	// Aiyhanym
	// Nurgúl
	// Dáýren
	// Qanat
	// Qanash
	// Qasym
	// Týrǵanym
}
