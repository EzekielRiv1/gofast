package gofast

import (
	_ "embed"
	"html/template"
	"io"
)

//go:embed assets/gofast.js
var browserScript string

// Layout renders a full HTML document for normal browser requests.
type Layout struct {
	template *template.Template
}

// DefaultLayout returns the built-in document template.
func DefaultLayout() Layout {
	return MustLayout(`<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>{{ .Title }}</title>
  <script type="module">` + browserScript + `</script>
</head>
<body>
  <main id="gofast-app">{{ .Body }}</main>
</body>
</html>`)
}

// BrowserScriptForLayout returns the built-in navigation script for custom layouts.
func BrowserScriptForLayout() string {
	return browserScript
}

// MustLayout parses a document template or panics.
func MustLayout(markup string) Layout {
	return Layout{template: template.Must(template.New("gofast-layout").Parse(markup))}
}

// Execute renders the layout.
func (l Layout) Execute(w io.Writer, page Page) error {
	return l.template.Execute(w, page)
}
