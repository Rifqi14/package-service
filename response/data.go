package response

type IData interface {
	Create() interface{}
}

//with meta
type DataWithMeta struct {
	Data interface{}
	Meta interface{}
}

func newBodySuccessWithMeta(data, meta interface{}) IData {
	return DataWithMeta{
		Data: data,
		Meta: meta,
	}
}

func (body DataWithMeta) Create() interface{} {
	return newSuccessResponseWithMeta(body.Data, body.Meta)
}

//with out meta

type DataWithOutMeta struct {
	Data interface{}
}

func newBodyWithOutMeta(data interface{}) IData {
	return DataWithOutMeta{Data: data}
}

func (body DataWithOutMeta) Create() interface{} {
	return newSuccessResponseWithOutMeta(body.Data)
}