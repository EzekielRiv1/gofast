package gofast

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAppRendersFullDocument(t *testing.T) {
	app := New()
	app.Get("/", func(*Context) Page {
		return Page{Title: "Home", Body: HTML("<h1>Hello</h1>")}
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	app.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", res.Code, http.StatusOK)
	}
	body := res.Body.String()
	for _, want := range []string{"<title>Home</title>", `<main id="gofast-app">`, "<h1>Hello</h1>"} {
		if !strings.Contains(body, want) {
			t.Fatalf("body does not contain %q:\n%s", want, body)
		}
	}
}

func TestAppRendersNavigationFragment(t *testing.T) {
	app := New()
	app.Get("/about", func(*Context) Page {
		return Page{Title: "About", Body: HTML("<h1>About</h1>")}
	})

	req := httptest.NewRequest(http.MethodGet, "/about", nil)
	req.Header.Set(navigateHeader, "true")
	res := httptest.NewRecorder()
	app.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", res.Code, http.StatusOK)
	}
	if got := res.Header().Get("X-Gofast-Title"); got != "About" {
		t.Fatalf("X-Gofast-Title = %q, want %q", got, "About")
	}
	if got := res.Body.String(); got != "<h1>About</h1>" {
		t.Fatalf("body = %q", got)
	}
}

func TestGetRejectsPost(t *testing.T) {
	app := New()
	app.Get("/", func(*Context) Page {
		return Page{Title: "Home", Body: HTML("<h1>Hello</h1>")}
	})

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	res := httptest.NewRecorder()
	app.ServeHTTP(res, req)

	if res.Code != http.StatusMethodNotAllowed {
		t.Fatalf("status = %d, want %d", res.Code, http.StatusMethodNotAllowed)
	}
	if got := res.Header().Get("Allow"); got != http.MethodGet {
		t.Fatalf("Allow = %q, want %q", got, http.MethodGet)
	}
}

func TestAppUsesCustomNotFoundPage(t *testing.T) {
	app := New()
	app.NotFound(func(*Context) Page {
		return Page{
			Title:  "Missing",
			Body:   HTML("<h1>Missing</h1>"),
			Status: http.StatusNotFound,
		}
	})

	req := httptest.NewRequest(http.MethodGet, "/missing", nil)
	res := httptest.NewRecorder()
	app.ServeHTTP(res, req)

	if res.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want %d", res.Code, http.StatusNotFound)
	}
	if !strings.Contains(res.Body.String(), "<h1>Missing</h1>") {
		t.Fatalf("body does not contain custom not found page:\n%s", res.Body.String())
	}
}
