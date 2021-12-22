package helper

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func validate() (*validator.Validate, ut.Translator) {
	validate := validator.New()
	readableErr := en.New()
	uni := ut.New(readableErr, readableErr)
	trans, found := uni.GetTranslator("en")
	if !found {
		PrintErrorMessage("404", "validator translator not found")
	}
	if err := en_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		PrintErrorMessage("404", err.Error())
	}
	return validate, trans
}

func translateError(err error, trans ut.Translator) (fieldErrors []error) {
	if err == nil {
		return nil
	}
	for _, e := range err.(validator.ValidationErrors) {
		translatedErr := fmt.Errorf(e.Translate(trans))
		fieldErrors = append(fieldErrors, translatedErr)
	}
	return fieldErrors
}

func ValidateStruct(data interface{}) error {
	validate, trans := validate()
	err := validate.Struct(data)
	fieldErrors := translateError(err, trans)
	return ErrorArrayToError(fieldErrors)
}
