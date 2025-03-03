// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package fields

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/KutoUI/UI/components"
	"github.com/KutoUI/UI/style"
)

type Button struct {
	BaseField
	Class     string
	Palette   style.PaletteClass
	Content   templ.Component
	Disabled  bool
	Type      string
	IconLeft  templ.Component
	IconRight templ.Component
}

func (b Button) GetBaseField() BaseField {
	return b.BaseField
}

func (b *Button) Configure() {
	b.BaseField.Configure()
}

func (b Button) Display() templ.Component {
	return components.Button(components.ButtonArgs{
		Class:     b.Class,
		Palette:   b.Palette,
		Content:   b.Content,
		Disabled:  b.Disabled,
		Type:      b.Type,
		IconLeft:  b.IconLeft,
		IconRight: b.IconRight,
	})
}

var _ = templruntime.GeneratedTemplate
