package forms

type EnctypeInstance string

type enctypeStruct struct {
	UrlEncoded EnctypeInstance
	Multipart  EnctypeInstance
	TextPlain  EnctypeInstance
}

var Enctype = enctypeStruct{
	UrlEncoded: "application/x-www-form-urlencoded",
	Multipart:  "multipart/form-data",
	TextPlain:  "text/plain",
}
