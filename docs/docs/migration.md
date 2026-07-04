---
id: migration
title: Migration and Compatibility
---

Use this page to check assumptions when upgrading Gofast.

## Current compatibility expectations

- Every page should remain directly loadable from the server.
- Custom layouts should include `main#gofast-app` if they use SPA-style navigation.
- Browser behavior should enhance server-rendered HTML instead of replacing it.

## Before upgrading an early version

When upgrading between early versions:

1. Run your app's tests.
2. Check route definitions and named routes.
3. Check templates rendered through `ctx.Render`.
4. Build the documentation or example app if you depend on those patterns.

## Public API areas to watch

- Routing and URL generation.
- Template rendering helpers.
- Browser navigation headers.
- Layout requirements.
