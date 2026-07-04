---
id: browser-layer
title: Browser Navigation
---

Gofast includes a small browser script in the default layout. It improves internal navigation without changing the server-side routing model.

## What it does

- Intercepts clicks on same-origin links.
- Fetches the next page with the `X-Gofast-Navigate: true` header.
- Replaces the contents of `#gofast-app` with the rendered HTML fragment.

Links still behave like normal links when users open them in a new tab, use modifier keys, download files, or navigate to another origin.

## Opt out for one link

Opt out for a single link with `data-gofast-ignore`:

```html
<a href="/report.csv" data-gofast-ignore>Download CSV</a>
```

If a navigation request fails, the browser falls back to a normal page load.

## Custom layouts

If you replace the default layout, include the browser script and keep a `#gofast-app` container.

```go
app := gofast.New().WithLayout(gofast.MustLayout(`<!doctype html>
<html lang="en">
<head>
  <title>{{ .Title }}</title>
  <script type="module">` + gofast.BrowserScriptForLayout() + `</script>
</head>
<body>
  <main id="gofast-app">{{ .Body }}</main>
</body>
</html>`))
```

## What it does not do

- It does not move routing into JavaScript.
- It does not make server routes optional.
- It does not manage client-side application state.
- It does not intercept external links or download links.
