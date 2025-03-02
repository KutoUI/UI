package fields

import (
	"math"
	"math/rand/v2"
	"strconv"

	"github.com/a-h/templ"
)

type BaseField struct {
	Name        string
	Id          string
	Label       string
	Value       string
	Description string
	Required    bool
	NoWrap      bool
	MinLength   int
	MaxLength   int
}

type Field interface {
	Configure()
	Display() templ.Component
	GetBaseField() BaseField
}

func (f *BaseField) Configure() {
	if f.Name == "" {
		f.Name = "error_" + strconv.Itoa(rand.IntN(10000))
	}

	if f.Id == "" {
		f.Id = f.Name
	}

	if f.MaxLength == 0 {
		f.MaxLength = math.MaxInt
	}
}
