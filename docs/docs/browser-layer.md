---
id: browser-layer
title: Browser Layer
---

Gofast includes a small browser script in the default layout. It does three jobs:

- Intercepts clicks on same-origin links.
- Fetches the next page with the `X-Gofast-Navigate: true` header.
- Replaces the contents of `#gofast-app` with the rendered HTML fragment.

Links still behave like normal links when users open them in a new tab, use modifier keys, download files, or navigate to another origin.

Opt out for a single link with `data-gofast-ignore`:

```html
<a href="/report.csv" data-gofast-ignore>Download CSV</a>
```

If a navigation request fails, the browser falls back to a normal page load.
