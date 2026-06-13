---
title: "CLI"
description: "Every command and subcommand, with the flags that matter."
weight: 10
---

```
th <command> [subcommand] [flags]
```

Run `th <command> --help` for the full flag list on any command.

## Commands

| Command | What it does |
|---|---|
| `profile <@handle\|id\|url>` | A profile's full record |
| `profile <h> --posts` | Walk the profile's recent posts |
| `profile <h> --replies` | Walk the profile's replies feed |
| `post <url\|shortcode\|id>` | A single post in full |
| `post <url> --replies` | The post's reply thread |
| `post <url> --raw` | The upstream HTML, untouched |
| `replies <url>` | Stream replies to a post as their own records |
| `feed <@handle>` | Walk a profile's recent posts (alias for `profile --posts`) |
| `search <query>` | Keyword search across public posts |
| `id <input>` | Classify any handle, id, shortcode, or URL (offline) |
| `db build <@handle> --db F` | Crawl a profile's posts into SQLite |
| `db query --db F "<sql>"` | Query the local dataset |
| `whoami` | Report how `th` is accessing Threads |
| `config show\|path` | Show resolved configuration and paths |
| `cache dir\|clear` | Inspect or clear the on-disk cache |
| `completion bash\|zsh\|fish` | Generate a shell completion script |
| `version` | Print version, commit, and build date |

## Global flags

These apply to every command.

| Flag | Default | Meaning |
|---|---|---|
| `-o, --output` | `auto` | `table`, `json`, `jsonl`, `csv`, `tsv`, `yaml`, `url`, `raw` |
| `--fields` | | Comma-separated columns to keep and order |
| `--no-header` | `false` | Omit the header row (table/csv/tsv) |
| `--template` | | Go `text/template` applied per record |
| `-n, --limit` | `0` | Max records emitted (`0` is unlimited) |
| `--delay` | `1s` | Minimum delay between requests |
| `--retries` | `4` | Retry attempts on 429/5xx |
| `--timeout` | `30s` | Per-request timeout |
| `--no-cache` | `false` | Bypass the on-disk cache |
| `--cache-ttl` | `1h` | Cache freshness window |
| `--lang` | `en-US` | `Accept-Language` / locale |
| `-q, --quiet` | `false` | Suppress progress on stderr |
| `-v, --verbose` | | Increase verbosity (repeatable) |
| `--proxy` | | HTTP/SOCKS proxy URL |
| `--user-agent` | | Override the default crawler UA |
| `--token` | | Official Graph API token (or `THREADS_TOKEN`) |
| `--session` | | Logged-in session id (or `THREADS_SESSION`) |
| `--csrf` | | Session CSRF token (or `THREADS_CSRF`) |

## Access modes

By default `th` reads anonymously as a crawler. Two optional modes add depth on
your own account and are off until you set their credentials:

- **token** — set `--token` or `THREADS_TOKEN` to use the official Graph API.
- **session** — set `--session`/`--csrf` (or `THREADS_SESSION`/`THREADS_CSRF`)
  to attach a logged-in session.

`th whoami` always reports which mode is active.

## Exit codes

| Code | Meaning |
|---|---|
| `0` | Success |
| `1` | Generic error |
| `2` | Usage error |
| `3` | Content not found or unavailable |
| `4` | Login wall: the content is not public |
| `5` | Rate limited |
| `6` | Network error |
