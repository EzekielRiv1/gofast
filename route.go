package gofast

import (
	"fmt"
	"net/url"
	"strings"
)

// Params stores route and URL generation parameters.
type Params map[string]string

// Get returns a parameter value by name.
func (p Params) Get(name string) string {
	if p == nil {
		return ""
	}
	return p[name]
}

// Clone returns a shallow copy of the params.
func (p Params) Clone() Params {
	if len(p) == 0 {
		return Params{}
	}
	clone := make(Params, len(p))
	for key, value := range p {
		clone[key] = value
	}
	return clone
}

// Route is a registered application route.
type Route struct {
	app     *App
	name    string
	path    string
	parts   []routePart
	handler Handler
}

type routePart struct {
	value string
	param bool
}

func newRoute(app *App, path string, handler Handler) *Route {
	route := &Route{app: app, path: path, handler: handler}
	route.parts = parseRouteParts(path)
	return route
}

// Name assigns a name used for URL generation.
func (r *Route) Name(name string) *Route {
	if name == "" {
		panic("gofast: route name cannot be empty")
	}
	if existing, ok := r.app.named[name]; ok && existing != r {
		panic("gofast: route name already registered: " + name)
	}
	if r.name != "" {
		delete(r.app.named, r.name)
	}
	r.name = name
	r.app.named[name] = r
	return r
}

// URL builds a path for this route.
func (r *Route) URL(params Params) (string, error) {
	if len(r.parts) == 0 {
		return "/", nil
	}

	var b strings.Builder
	for _, part := range r.parts {
		b.WriteByte('/')
		if !part.param {
			b.WriteString(part.value)
			continue
		}

		value, ok := params[part.value]
		if !ok || value == "" {
			return "", fmt.Errorf("gofast: missing route parameter %q for %s", part.value, r.path)
		}
		b.WriteString(url.PathEscape(value))
	}
	return b.String(), nil
}

func (r *Route) match(path string) (Params, bool) {
	requestParts := splitPath(path)
	if len(requestParts) != len(r.parts) {
		return nil, false
	}

	params := Params{}
	for i, routePart := range r.parts {
		requestPart := requestParts[i]
		if routePart.param {
			value, err := url.PathUnescape(requestPart)
			if err != nil {
				return nil, false
			}
			params[routePart.value] = value
			continue
		}
		if routePart.value != requestPart {
			return nil, false
		}
	}
	return params, true
}

func parseRouteParts(path string) []routePart {
	parts := splitPath(path)
	routeParts := make([]routePart, 0, len(parts))
	seen := map[string]bool{}
	for _, part := range parts {
		if strings.HasPrefix(part, ":") {
			name := strings.TrimPrefix(part, ":")
			if name == "" {
				panic("gofast: route parameter name cannot be empty")
			}
			if seen[name] {
				panic("gofast: duplicate route parameter: " + name)
			}
			seen[name] = true
			routeParts = append(routeParts, routePart{value: name, param: true})
			continue
		}
		routeParts = append(routeParts, routePart{value: part})
	}
	return routeParts
}

func splitPath(path string) []string {
	trimmed := strings.Trim(path, "/")
	if trimmed == "" {
		return nil
	}
	return strings.Split(trimmed, "/")
}
