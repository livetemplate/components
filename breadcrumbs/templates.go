package breadcrumbs

import (
	"embed"

	"github.com/livetemplate/components/base"
)

//go:embed templates/*.tmpl
var templateFS embed.FS

// Templates returns the breadcrumbs template set.
func Templates() *base.TemplateSet {
	return base.NewTemplateSet(templateFS, "templates/*.tmpl", "breadcrumbs")
}
