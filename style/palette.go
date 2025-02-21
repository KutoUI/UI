package style

type PaletteCssVars struct {
	Color           string
	BackgroundColor string
}

type PaletteClass string

type paletteClassStruct struct {
	IsBlank   PaletteClass
	IsPrimary PaletteClass
	IsLink    PaletteClass
	IsInfo    PaletteClass
	IsSuccess PaletteClass
	IsWarning PaletteClass
	IsDanger  PaletteClass
}

var Palette = paletteClassStruct{
	IsBlank:   "is-blank",
	IsPrimary: "is-primary",
	IsLink:    "is-link",
	IsInfo:    "is-info",
	IsSuccess: "is-success",
	IsWarning: "is-warning",
	IsDanger:  "is-danger",
}
