package helpers

import (
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func SetValidatorEngineToUseJSONTags() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}

func FillValidationErrorTag(err error, validationErrors map[string]string) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, fe := range ve {
			if _, exists := validationErrors[fe.Field()]; !exists {
				validationErrors[fe.Field()] = ValidationMessageForTag(fe.Tag(), fe.Param())
			}
		}
	}
}

func ValidationMessageForTag(tag string, param any) string {
	switch tag {
	case "required":
		return "Это поле обязательно к заполнению"
	case "email":
		return "Неправильный формат электронной почты"
	case "min":
		return "Длина данного поля не должна быть меньше " + param.(string) + " символов"
	case "max":
		return "Длина данного поля не должна быть больше " + param.(string) + " символов"
	case "len":
		return "Длина данного поля должна быть ровно " + param.(string) + " символов"
	case "gte":
		return "Значение этого поля должна быть больше или равно " + param.(string)
	case "lte":
		return "Значение этого поля должна быть меньше или равно " + param.(string)
	case "unique":
		return "Такое значение уже существует в БД"
	case "exists":
		return "Указанный ресурс не существует"
	case "date-format":
		return "Неправильный формат даты. Введите в формате гггг-мм-дд"
	case "time-format":
		return "Неправильный формат времени. Введите в формате чч:мм"
	case "numeric":
		return "Данное поле принимает только цифровое значение"
	case "oneof":
		return "Данное поле принимает значение из набора " + param.(string)
	case "max-number-of-elements":
		return "Данно поле принимает максимум " + strconv.Itoa(param.(int)) + " элемент(ов)"
	default:
		return ""
	}
}

//func SaveImage(path string, image *multipart.FileHeader, c *gin.Context) (string, error) {
//	newFileName := path + uuid.New().String() + filepath.Ext(image.Filename)
//	if err := c.SaveUploadedFile(image, newFileName); err != nil {
//		return "", err
//	}
//	return newFileName, nil
//}

func DeleteImage(url string) error {
	pathSlice := strings.TrimPrefix(url, "http://localhost:4000/")
	return os.RemoveAll(pathSlice)
}
