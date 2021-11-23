package response
//
//import (
//	ut "github.com/go-playground/universal-translator"
//	"github.com/go-playground/validator/v10"
//	"github.com/gofiber/fiber/v2"
//	"gitlab.com/shoesmart2.1/backend/packages/str"
//	"net/http"
//	"reflect"
//	"strings"
//)
//
////
////import (
////	ut "github.com/go-playground/universal-translator"
////	"github.com/go-playground/validator/v10"
////	"github.com/gofiber/fiber/v2"
////	"gitlab.com/shoesmart2.1/backend/packages/str"
////	"net/http"
////	"reflect"
////	"strings"
////)
////
////type SuccessResponseWithOutMetaVm struct {
////	Data interface{} `json:"data"`
////}
////
////type SuccessResponseWithMeta struct {
////	Data interface{} `json:"data"`
////	Meta interface{} `json:"meta"`
////}
////
////type ErrorResponseVm struct {
////	Message interface{}
////}
////
//const (
//	FactoryResponseWithMeta              = `FactoryResponseWithMeta`
//	ResponseWithOutMeta           = `ResponseWithOutMeta`
//	ResponseErrorValidationStruct = `ResponseErrorValidationStruct`
//)
//
//type IResponseBuilder interface {
//	SetFiberContext(ctx *fiber.Ctx) Response
//
//	SetResponseType(responseType string) Response
//
//	SetData(data interface{}) Response
//
//	SetMeta(meta interface{}) Response
//
//	SetErr(err error) Response
//
//	SetErrorValidator(error validator.ValidationErrors) Response
//
//	SetStatusCode(code int) Response
//
//	SetTranslator(translator ut.Translator) Response
//
//	SendResponse() error
//}
//
//type Response struct {
//	fiberCtx       *fiber.Ctx
//	responseType   string
//	data           interface{}
//	meta           interface{}
//	err            interface{}
//	errorValidator map[string][]string
//	statusCode     int
//	Translator     ut.Translator
//}
//
//func NewResponse() IResponseBuilder {
//	return &Response{}
//}
//
//func (r Response) SetFiberContext(ctx *fiber.Ctx) Response {
//	r.fiberCtx = ctx
//	return r
//}
//
//func (r Response) SetResponseType(responseType string) Response {
//	r.responseType = responseType
//
//	return r
//}
//
//func (r Response) SetData(data interface{}) Response {
//	r.data = data
//
//	return r
//}
//
//func (r Response) SetMeta(meta interface{}) Response {
//	r.meta = meta
//
//	return r
//}
//
//func (r Response) SetErr(err error) Response {
//	if err != nil {
//		r.err = err.Error()
//		if r.statusCode == 0 {
//			r.statusCode = http.StatusUnprocessableEntity
//		}
//	} else {
//		r.statusCode = http.StatusOK
//	}
//
//	return r
//}
//
//func (r Response) SetStatusCode(code int) Response {
//	r.statusCode = code
//
//	return r
//}
//
//func (r Response) SetTranslator(translator ut.Translator) Response {
//	r.Translator = translator
//
//	return r
//}
//
//func (r Response) SendResponse() error {
//	if r.err != nil {
//		return r.fiberCtx.Status(r.statusCode).JSON(ErrorResponseVm{Message: r.err})
//	}
//
//	switch r.responseType {
//	case FactoryResponseWithMeta:
//		return r.fiberCtx.Status(http.StatusOK).JSON(SuccessResponseWithMeta{Data: r.data, Meta: r.meta})
//	case ResponseWithOutMeta:
//		return r.fiberCtx.Status(http.StatusOK).JSON(SuccessResponseWithOutMetaVm{Data: r.data})
//	}
//
//	return nil
//}
//
//func (r Response) SetErrorValidator(error validator.ValidationErrors) Response {
//	errorMessage := map[string][]string{}
//	errorTranslation := error.Translate(r.Translator)
//
//	reflect.TypeOf(errorMessage).String()
//
//	for _, err := range error {
//		errKey := str.Underscore(err.StructField())
//		errorMessage[errKey] = append(
//			errorMessage[errKey],
//			strings.Replace(errorTranslation[err.Namespace()], err.StructField(), err.StructField(), -1),
//		)
//	}
//
//	if r.statusCode == 0 {
//		r.statusCode = http.StatusBadRequest
//	}
//	r.err = errorMessage
//
//	return r
//}