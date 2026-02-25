package validator

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	once     sync.Once //pastikan object validator hanya di create satu instance object
	instance *validator.Validate
)

// GetValidator mengembalikan singleton instance validator
func GetValidator() *validator.Validate {
	once.Do(func() {
		instance = validator.New()
		// Register custom validations di sini (jika ada)
		registerCustomValidations(instance)
	})
	return instance
}

// registerCustomValidations untuk menambahkan validasi kustom global
func registerCustomValidations(v *validator.Validate) {
	// Contoh custom validation
	v.RegisterValidation("phone", validatePhone)
}

// Contoh custom validator global
func validatePhone(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	// Implementasi validasi phone number
	return len(phone) >= 10 && len(phone) <= 15
}
