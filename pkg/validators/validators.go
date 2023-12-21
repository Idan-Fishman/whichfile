package validators

import "github.com/go-playground/validator/v10"

var (
	V *validator.Validate = InitializeValidator() // Global validator instance.
)

// InitializeValidator initializes a validator and returns it.
func InitializeValidator() *validator.Validate {
	v := validator.New()

	return v
}
