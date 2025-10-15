package util

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error: %s path: %s error: %s", r.Method, r.URL.Path, err)
	WriteJSONError(w, http.StatusInternalServerError, "the server encountered a problem")
}

func BadRequestError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	// check validation error
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)
		for _, ve := range validationErrors {
			field := ve.Field()
			tag := ve.Tag()
			errors[field] = getValidationErrorMessage(field, tag, ve.Param())
		}
		WriteJSON(w, http.StatusBadRequest, map[string]any{
			"error":  "validation failed",
			"fields": errors,
		})
		return
	}
	WriteJSONError(w, http.StatusBadRequest, err.Error())
}

func NotFoundError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request error: %s path: %s error: %s", r.Method, r.URL.Path, err)
	WriteJSONError(w, http.StatusNotFound, "resource not found")
}

func getValidationErrorMessage(field, tag, param string) string {
	switch tag {
	case "required":
		return field + " is required"
	case "max":
		return field + " must be at most " + param + " characters long"
	case "min":
		return field + " must be at least " + param + " characters long"
	case "len":
		return field + " must be exactly " + param + " characters long"
	case "email":
		return field + " must be a valid email address"
	// Add more as needed
	default:
		return field + " is invalid"
	}
}
