package utils

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin/binding"
	zh "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/pkg/errors"
)

var Trans ut.Translator

func InitValidator() {
	// 1. 取 Gin 默认 validator 引擎
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//  注册中文翻译
		zhCn := zh.New()
		uni := ut.New(zhCn, zhCn)
		Trans, _ = uni.GetTranslator("zh")

		err := zh_translations.RegisterDefaultTranslations(v, Trans)
		if err != nil {
			fmt.Println("翻译错误")
			return
		}
	}

}

// 翻译错误
func PareJSONError(e error) string {
	var translated []string
	var errs validator.ValidationErrors
	if errors.As(e, &errs) {
		//把错误依次翻译并插入切片中
		for _, e := range errs {
			translated = append(translated, e.Translate(Trans))
		}
	}
	return strings.Join(translated, ",")
}
