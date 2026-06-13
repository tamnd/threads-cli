---
title: "Troubleshooting"
description: "The handful of things that trip people up, and how to fix each one."
weight: 40
---

Most of these come down to network reality or how Threads serves its data, not a
bug.

## Requests start failing or returning 429

Threads rate-limits like any public site. `th` already paces requests and
retries the transient failures, but a hard limit still means backing off. Raise
the delay between requests with `--delay` (for example `--delay 3s`) and retry
later. A burst of 429 or 5xx responses is the site asking you to slow down, not
a defect; `th` exits `5` so a script can tell.

## Search returns nothing, or exits 3

The logged-out search and deep-pagination paths ride on persisted GraphQL query
ids that Threads rotates every few weeks. When the current id no longer answers
anonymously, `th` exits `3` with a clear note rather than guessing. The
server-rendered surface (`profile`, `post`, `replies`, `feed`) does not depend
on those ids and keeps working.

## Nothing is found for something you expected

The public surface is not the whole site. Some data sits behind a login, a
region, or a page that only renders with JavaScript, and that part is not
reachable without the right session. Check that the input is spelled the way the
site uses it, try a broader query, and see whether the same thing is visible in
a private browser window before assuming it is missing.

## A command needs a session

Where a surface is gated, th reads a cookie or token you supply
rather than logging in for you. Pass it on the command that needs it and keep it
out of your shell history. Commands that work without one stay anonymous.

## The binary is not on your PATH

`go install` puts the binary in `$(go env GOPATH)/bin` (usually `~/go/bin`), and
a release archive leaves it wherever you unpacked it. If your shell cannot find
`th`, add that directory to your `PATH`. See
[installation](/getting-started/installation/).

## Seeing what th actually did

When something behaves unexpectedly, `-v` adds per-request detail so you can see
the URLs it hit and the responses it got. That is usually enough to tell a rate
limit apart from a genuinely empty result.
