package datatable

import (
	"embed"

	"github.com/livetemplate/components/base"
)

// templateFS contains all datatable template files embedded at compile time.
//
//go:embed templates/*.tmpl
var templateFS embed.FS

// Templates returns the datatable component's template set for registration
// with the LiveTemplate framework.
//
// Example usage in main.go:
//
//	import "github.com/livetemplate/components/datatable"
//
//	tmpl, err := livetemplate.New("app",
//	    livetemplate.WithComponentTemplates(datatable.Templates()),
//	)
//
// Available templates:
//   - "lvt:datatable:default:v1" - Data table with sorting/pagination
func Templates() *base.TemplateSet {
	return base.NewTemplateSet(templateFS, "templates/*.tmpl", "datatable")
}
