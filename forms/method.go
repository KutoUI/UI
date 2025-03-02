package forms

type MethodInstance string

type MethodStruct struct {
	GET  MethodInstance
	POST MethodInstance
}

var Method = MethodStruct{
	GET:  "GET",
	POST: "POST",
}
