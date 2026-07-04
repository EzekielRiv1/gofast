package gofast

import (
	"bytes"
	"html/template"
	"io/fs"
)

// Views renders named html/template definitions.
type Views struct {
	template *template.Template
}

// MustViews parses templates from filesystem patterns or panics.
func MustViews(patterns ...string) *Views {
	if len(patterns) == 0 {
		panic("gofast: MustViews requires at least one pattern")
	}
	t := template.New("gofast-views")
	for _, pattern := range patterns {
		template.Must(t.ParseGlob(pattern))
	}
	return &Views{template: t}
}

// MustViewsFS parses templates from an fs.FS or panics.
func MustViewsFS(files fs.FS, patterns ...string) *Views {
	if len(patterns) == 0 {
		panic("gofast: MustViewsFS requires at least one pattern")
	}
	return &Views{template: template.Must(template.ParseFS(files, patterns...))}
}

// MustViewsText parses template text or panics. It is useful for tests and embedded views.
func MustViewsText(name, text string) *Views {
	return &Views{template: template.Must(template.New(name).Parse(text))}
}

// Render renders a named template definition as trusted HTML.
func (v *Views) Render(name string, data any) (template.HTML, error) {
	var b bytes.Buffer
	if err := v.template.ExecuteTemplate(&b, name, data); err != nil {
		return "", err
	}
	return template.HTML(b.String()), nil
}
