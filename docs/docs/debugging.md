---
id: debugging
title: Debugging
---

Use this page when a Gofast route, template, or browser navigation does not behave as expected.

## Internal links do a full page reload

Likely causes:

- The link points to another origin.
- The link has `target`, `download`, or `data-gofast-ignore`.
- Your custom layout is missing `id="gofast-app"`.
- The server returned a non-success status for the navigation request.

How to verify:

1. Open the page directly in the browser.
2. Click the link without modifier keys.
3. Check whether the request includes `X-Gofast-Navigate: true`.

## Page works on reload but not during navigation

If a page renders correctly on reload but not during client-side navigation, inspect the response for the `X-Gofast-Navigate: true` request. The server should return only the page body fragment and the `X-Gofast-Title` header.

Common causes:

- The handler returns different content depending on headers.
- A custom layout is wrapping content that should be a fragment.
- Browser-side code expects a full document after internal navigation.

## Template renders a 500 page

Likely causes:

- The app was not configured with `WithViews`.
- The template name passed to `ctx.Render` does not exist.
- The template expects fields that are not present in the data.

How to verify:

1. Check the template name in `ctx.Render`.
2. Check that the template file has `{{ define "name" }}`.
3. Confirm the app uses `gofast.MustViews` or `gofast.MustViewsFS`.

## Route parameter is empty

Likely causes:

- The route pattern does not include that parameter name.
- The requested URL does not match the route you expected.
- The parameter name differs by spelling or case.

Example:

```go
app.Get("/projects/:projectID", func(ctx *gofast.Context) gofast.Page {
	projectID := ctx.Param("projectID")
	// ...
})
```

When in doubt, reload the page directly. Gofast routes are normal server routes, so every page should work without JavaScript.
