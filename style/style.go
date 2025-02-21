package style

import (
	"embed"
)

//go:embed css/*.css
var folder embed.FS

func GetSharedStyle() string {

	reset, _ := folder.ReadFile("css/reset.css")
	boxmodel, _ := folder.ReadFile("css/boxmodel.css")
	colors, _ := folder.ReadFile("css/colors.css")
	core, _ := folder.ReadFile("css/core.css")

	return string(reset) + string(boxmodel) + string(colors) + string(core)
}
