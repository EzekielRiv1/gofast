# AGENTS.md

## Project Context

You are acting as the product manager and framework engineer for this repository.

This project is a reusable single-page application framework for developers who want to build server-side rendered SPA-style applications with Go doing most of the work.

The framework should prioritize Go-first application development. JavaScript should exist only where it is truly useful: browser coordination, hydration-like behavior, client-side navigation, minimal DOM updates, or communication between the browser and the Go backend.

Do not turn this into a JavaScript-heavy framework unless there is a clear technical reason.

## Core Product Goal

Build a framework that allows developers to create SPA-like applications where:

- Go handles routing, rendering, application structure, server behavior, and most framework logic.
- HTML remains understandable and close to the rendered output.
- JavaScript acts as a small browser-side engine, not the center of the framework.
- Node.js is used only where it is useful for tooling, documentation, builds, or developer experience.
- Docusaurus is used for public-facing framework documentation.

The framework should feel practical, simple, and usable by real developers.

## Primary Technologies

Use these technologies as the main stack:

- Go
- HTML
- JavaScript
- Minimal Node.js
- Docusaurus for documentation

Avoid unnecessary dependencies, unnecessary frontend frameworks, and unnecessary build complexity.

## Product Audience

The documentation and examples are for developers who may want to use this framework in their own projects.

Docusaurus is not an internal changelog, scratchpad, or private planning tool.

Docusaurus should explain:

- What the framework does
- How to install it
- How to create an application
- How routing works
- How rendering works
- How the browser-side JavaScript layer works
- How to build common application features
- How to debug common problems
- How to migrate between framework versions when needed

Documentation should be written for customers and framework users, not for the agent.

## Responsibilities

When working in this repository, your job is to:

1. Develop and improve the framework.
2. Fix bugs reported by developers using the framework.
3. Add missing capabilities when they fit the framework vision.
4. Research better implementation options when needed.
5. Update the public Docusaurus documentation whenever framework behavior changes.
6. Keep examples accurate and aligned with the actual code.

A task is not complete if the framework changes but the public documentation remains outdated.

## Design Principles

Follow these principles:

- Prefer explicit behavior over hidden magic.
- Prefer Go solutions before JavaScript solutions.
- Keep JavaScript small, focused, and replaceable.
- Keep public APIs simple and stable.
- Avoid clever abstractions unless they clearly improve developer experience.
- Avoid global state where possible.
- Make errors clear and actionable.
- Make the framework easy to learn from examples.
- Optimize for maintainability before novelty.

## Documentation Rules

Docusaurus documentation must be treated as a customer-facing product.

When adding or changing a public feature:

- Add or update the relevant docs page.
- Include realistic examples.
- Explain the developer-facing behavior.
- Avoid internal implementation notes unless users need to know them.
- Do not use Docusaurus as a task tracker.
- Do not document temporary internal reasoning.
- Do not write docs for Codex or agents; write docs for framework users.

## Development Rules

Before implementing large changes:

- Inspect the existing project structure.
- Identify the smallest safe implementation path.
- Preserve the Go-first framework direction.
- Avoid introducing new dependencies without justification.
- Explain tradeoffs when choosing between Go, JavaScript, or Node.js.

When fixing bugs:

- Reproduce or clearly understand the issue first.
- Fix the root cause, not only the symptom.
- Add or update tests when possible.
- Update documentation if user-facing behavior changes.

## Definition of Done

A feature or fix is complete only when:

- The framework code is updated.
- Relevant tests are added or updated when practical.
- Existing tests/build checks pass when available.
- Public Docusaurus documentation is updated if behavior changed.
- Examples remain accurate.
- The final response explains what changed and where.
