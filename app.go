package gofast

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

const navigateHeader = "X-Gofast-Navigate"

// Handler renders a page for the current request.
type Handler func(*Context) Page

// Context carries request-specific framework helpers.
type Context struct {
	Response http.ResponseWriter
	Request  *http.Request
	app      *App
	params   Params
}

// Page is the rendered result for a route.
type Page struct {
	Title  string
	Body   template.HTML
	Status int
}

// App is a small Go-first application router and renderer.
type App struct {
	routes   []*Route
	named    map[string]*Route
	layout   Layout
	views    *Views
	notFound Handler
}

// New creates an application with the default document layout.
func New() *App {
	app := &App{
		named:  make(map[string]*Route),
		layout: DefaultLayout(),
	}
	app.notFound = func(*Context) Page {
		return Page{
			Title:  "Not found",
			Body:   "<h1>Not found</h1><p>The page could not be found.</p>",
			Status: http.StatusNotFound,
		}
	}
	return app
}

// WithLayout replaces the document layout used for full page requests.
func (a *App) WithLayout(layout Layout) *App {
	a.layout = layout
	return a
}

// WithViews sets the template registry used by Context.Render.
func (a *App) WithViews(views *Views) *App {
	a.views = views
	return a
}

// NotFound replaces the default 404 page.
func (a *App) NotFound(handler Handler) {
	a.notFound = handler
}

// Route registers a path handler. Paths may include named parameters, such as /projects/:id.
func (a *App) Route(path string, handler Handler) *Route {
	if path == "" || !strings.HasPrefix(path, "/") {
		panic("gofast: route path must begin with /")
	}
	route := newRoute(a, path, handler)
	a.routes = append(a.routes, route)
	return route
}

// Get registers a GET route. Paths may include named parameters, such as /projects/:id.
func (a *App) Get(path string, handler Handler) *Route {
	return a.Route(path, func(ctx *Context) Page {
		if ctx.Request.Method != http.MethodGet {
			ctx.Response.Header().Set("Allow", http.MethodGet)
			return Page{
				Title:  "Method not allowed",
				Body:   "<h1>Method not allowed</h1>",
				Status: http.StatusMethodNotAllowed,
			}
		}
		return handler(ctx)
	})
}

// URL builds a path for a named route.
func (a *App) URL(name string, params Params) (string, error) {
	route, ok := a.named[name]
	if !ok {
		return "", fmt.Errorf("gofast: route %q is not registered", name)
	}
	return route.URL(params)
}

// MustURL builds a path for a named route or panics.
func (a *App) MustURL(name string, params Params) string {
	path, err := a.URL(name, params)
	if err != nil {
		panic(err)
	}
	return path
}

// ServeHTTP implements http.Handler.
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range a.routes {
		params, ok := route.match(r.URL.Path)
		if ok {
			a.render(w, r, route.handler, params)
			return
		}
	}
	a.render(w, r, a.notFound, nil)
}

// ListenAndServe starts an HTTP server for the application.
func (a *App) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, a)
}

// HTML marks a trusted HTML fragment for rendering.
func HTML(markup string) template.HTML {
	return template.HTML(markup)
}

// RequestContext returns the request context for convenience in route helpers.
func (c *Context) RequestContext() context.Context {
	return c.Request.Context()
}

// Param returns a matched route parameter by name.
func (c *Context) Param(name string) string {
	return c.params.Get(name)
}

// Params returns a copy of all matched route parameters.
func (c *Context) Params() Params {
	return c.params.Clone()
}

// Query returns the first query string value for name.
func (c *Context) Query(name string) string {
	return c.Request.URL.Query().Get(name)
}

// URL builds a path for a named route.
func (c *Context) URL(name string, params Params) (string, error) {
	return c.app.URL(name, params)
}

// MustURL builds a path for a named route or panics.
func (c *Context) MustURL(name string, params Params) string {
	return c.app.MustURL(name, params)
}

// HTMLPage returns a Page from a trusted HTML fragment.
func (c *Context) HTMLPage(title string, body template.HTML) Page {
	return Page{Title: title, Body: body}
}

// Render renders a named template from the app's Views registry.
func (c *Context) Render(title, templateName string, data any) Page {
	if c.app.views == nil {
		return Page{
			Title:  "Template error",
			Body:   HTML("<h1>Template error</h1><p>No template registry is configured.</p>"),
			Status: http.StatusInternalServerError,
		}
	}

	body, err := c.app.views.Render(templateName, data)
	if err != nil {
		return Page{
			Title:  "Template error",
			Body:   HTML("<h1>Template error</h1><p>The page template could not be rendered.</p>"),
			Status: http.StatusInternalServerError,
		}
	}
	return Page{Title: title, Body: body}
}

func (a *App) render(w http.ResponseWriter, r *http.Request, handler Handler, params Params) {
	page := handler(&Context{Response: w, Request: r, app: a, params: params})
	if page.Status == 0 {
		page.Status = http.StatusOK
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.Header.Get(navigateHeader) == "true" {
		w.Header().Set("X-Gofast-Title", page.Title)
		w.WriteHeader(page.Status)
		_, _ = w.Write([]byte(page.Body))
		return
	}

	w.WriteHeader(page.Status)
	_ = a.layout.Execute(w, page)
}
