---
title: "Introduction"
description: "What th is and how it is put together."
weight: 10
---

`th` is a single binary that reads Threads the way a search engine does. It asks
threads.com for the server-rendered page behind a profile or post, parses the
embedded JSON into typed records, and gets out of your way. There is no login,
no browser, nothing to sign up for, and nothing to run alongside it.

## How it is built

- A **library package** (`threads`) holds the crawler HTTP client, the
  server-rendered page parsers, the logged-out GraphQL queries, and the typed
  data models. It paces requests, sets an honest crawler User-Agent, and retries
  the transient failures any public site throws under load.
- A **command tree** (`cli`) wraps the library in subcommands with shared output
  formats and flags.
- A standalone **`pkg/thid`** classifies any handle, id, shortcode, or URL with
  no network and no other dependencies.
- One **`cmd/th`** entry point ties them together.

## Scope

`th` is a read-only client over data Threads already serves publicly. It reads
that data and shapes it for you. Private content stays private; when a target is
walled, `th` exits `4` rather than returning nothing. That narrow scope keeps it
a single small binary with no database, no daemon, and no setup. Optional token
and session modes add depth on your own account, off by default.

Next: [install it](/getting-started/installation/), then take the
[quick start](/getting-started/quick-start/).
