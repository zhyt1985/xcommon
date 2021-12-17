package validatorx

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	uni := ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")
	validate = validator.New()
	//注册一个函数，获取struct tag里自定义的label作为字段名
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})
	//注册翻译器
	err := zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println(err)
	}
}
func Struct(v interface{}) error {
	var (
		msg string
	)
	err := validate.Struct(v)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			msg = v.Translate(trans)
			break
		}
	}
	if msg == "" {
		return nil
	}
	return errors.New(msg)
}
