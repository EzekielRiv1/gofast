---
id: debugging
title: Debugging
---

If navigation does a full page reload, check these common causes:

- The link points to another origin.
- The link has `target`, `download`, or `data-gofast-ignore`.
- Your custom layout is missing `id="gofast-app"`.
- The server returned a non-success status for the navigation request.

If a page renders correctly on reload but not during client-side navigation, inspect the response for the `X-Gofast-Navigate: true` request. The server should return only the page body fragment and the `X-Gofast-Title` header.

When in doubt, reload the page directly. Gofast routes are normal server routes, so every page should work without JavaScript.
