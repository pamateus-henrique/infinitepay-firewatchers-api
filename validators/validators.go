package validators

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
    validate = validator.New()
    // Register custom validation functions if necessary
}

// ValidateStruct validates a struct based on tags
func ValidateStruct(s interface{}) error {
    return validate.Struct(s)
}

// ValidationError is a custom error type for validation errors
type ValidationError struct {
    Err      error
    Messages []string
}

// Error implements the error interface
func (v *ValidationError) Error() string {
    return "validation failed"
}

// ErrorMessages returns the validation error messages
func (v *ValidationError) ErrorMessages() []string {
    if v.Messages != nil {
        return v.Messages
    }
    return ExtractValidationErrors(v.Err)
}

// ExtractValidationErrors formats validation errors into a slice of messages
func ExtractValidationErrors(err error) []string {
    var errors []string
    if err != nil {
        if validationErrors, ok := err.(validator.ValidationErrors); ok {
            for _, err := range validationErrors {
                errors = append(errors, formatErrorMessage(err))
            }
        } else {
            // Non-validation error occurred
            errors = append(errors, err.Error())
        }
    }
    return errors
}

// formatErrorMessage customizes the error message for each field error
func formatErrorMessage(fe validator.FieldError) string {
    field := fe.Field()
    tag := fe.Tag()
    switch tag {
    case "required":
        return field + " is required"
    case "email":
        return field + " must be a valid email address"
    case "gte":
        return field + " must be greater than or equal to " + fe.Param()
    case "lte":
        return field + " must be less than or equal to " + fe.Param()
    default:
        return field + " is not valid"
    }
}