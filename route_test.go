package gofast

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRouteParams(t *testing.T) {
	app := New()
	app.Get("/projects/:projectID", func(ctx *Context) Page {
		return ctx.HTMLPage("Project", HTML("<h1>"+ctx.Param("projectID")+"</h1>"))
	}).Name("project.show")

	req := httptest.NewRequest(http.MethodGet, "/projects/gofast", nil)
	res := httptest.NewRecorder()
	app.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", res.Code, http.StatusOK)
	}
	if !strings.Contains(res.Body.String(), "<h1>gofast</h1>") {
		t.Fatalf("body does not contain route param:\n%s", res.Body.String())
	}
}

func TestNamedRouteURL(t *testing.T) {
	app := New()
	app.Get("/projects/:projectID/issues/:issueID", func(*Context) Page {
		return Page{Title: "Issue", Body: HTML("<h1>Issue</h1>")}
	}).Name("issue.show")

	got, err := app.URL("issue.show", Params{"projectID": "go fast", "issueID": "a/b"})
	if err != nil {
		t.Fatal(err)
	}
	want := "/projects/go%20fast/issues/a%2Fb"
	if got != want {
		t.Fatalf("URL = %q, want %q", got, want)
	}
}

func TestNamedRouteURLRequiresParams(t *testing.T) {
	app := New()
	app.Get("/projects/:projectID", func(*Context) Page {
		return Page{Title: "Project", Body: HTML("<h1>Project</h1>")}
	}).Name("project.show")

	_, err := app.URL("project.show", nil)
	if err == nil {
		t.Fatal("expected missing param error")
	}
}

func TestContextRenderUsesViews(t *testing.T) {
	app := New().WithViews(MustViewsText("views", `{{ define "pages/home" }}<h1>{{ .Name }}</h1>{{ end }}`))
	app.Get("/", func(ctx *Context) Page {
		return ctx.Render("Home", "pages/home", map[string]string{"Name": "Gofast"})
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	app.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", res.Code, http.StatusOK)
	}
	if !strings.Contains(res.Body.String(), "<h1>Gofast</h1>") {
		t.Fatalf("body does not contain rendered template:\n%s", res.Body.String())
	}
}
