# th

A delightful, scriptable command line for [Threads](https://www.threads.com).
One binary that resolves a profile to a rich record, streams its recent posts
and replies, pulls a single post's whole thread, and shapes everything into
clean structured data you can pipe anywhere, with no login and no browser.

```
th profile zuck
```

```
ID           USERNAME  NAME             FOLLOWERS  FOLLOWING  VERIFIED  BIO                                     URL
63055343223  zuck      Mark Zuckerberg  5530027    0          true      Mostly superintelligence and MMA takes  https://www.threads.com/@zuck
```

Full documentation: [threads-cli.tamnd.com](https://threads-cli.tamnd.com).

## Why

Pulling data out of Threads usually means a headless browser, a brittle pile of
selectors, or the official API with its app review and tokens. `th` takes a
different route: it reads the public pages Threads serves to search engines,
parses the server-rendered payload into typed records, and renders them in the
output format you ask for. One static binary, no login, no browser, no API key.

Because it reads the public crawler surface, `th` sees what a search engine
sees: full public profiles, posts, and reply threads, the most recent posts
rather than the entire history. Private content stays private, and `th` is
explicit about the wall rather than silently returning nothing.

## Install

```sh
go install github.com/tamnd/threads-cli/cmd/th@latest
```

Or grab a prebuilt binary from the [releases page](https://github.com/tamnd/threads-cli/releases).
The binary is pure Go with no runtime dependencies.

Build from source:

```sh
git clone https://github.com/tamnd/threads-cli
cd threads-cli
make build      # produces ./bin/th
```

## Quick start

```sh
th profile zuck                    # a profile's full record
th profile zuck --posts -n 20      # its twenty most recent posts
th post <url> --replies            # a post and its reply thread
th id <anything>                   # classify any Threads handle, id, or URL
```

## How it reads Threads

`th` reads anonymously, as a web crawler, with no login and no cookie. It asks
Threads for the same server-rendered pages a search engine gets and parses what
comes back.

```sh
th whoami        # reports the access mode and user agent
```

This works on any public profile or post. When a target is private, or behind a
login wall, `th` exits `4` with a one-line hint so scripts can tell that apart
from a real error. The trade-off is depth: a profile exposes the most recent
posts rather than the full history, and a post carries the visible reply window.

To go deeper on your own account, set a Graph API token (`THREADS_TOKEN`) or a
logged-in session (`THREADS_SESSION` and `THREADS_CSRF`). These are off by
default; the anonymous crawler is the floor everything else builds on.

## How it works

`th` resolves any handle, id, shortcode, or URL to a typed identity first
(`th id` shows exactly what it sees), then fetches the matching profile or post
page with a crawler user agent. That page carries the post text, the engagement
counts, the media, and a window of replies inside its embedded JSON; `th` walks
that tree into records with the standard library alone. Past the rendered
window, a set of logged-out persisted queries extend the walk by cursor while
they remain current. Responses are cached on disk by URL, so re-running a
command is instant and polite to Threads.

Every record is a plain struct with JSON tags, so `-o json` gives you the full
shape and `--fields` narrows it. Nothing is invented: a field that Threads does
not surface anonymously simply stays empty rather than being guessed.

## Commands

| Command | What it does |
| --- | --- |
| `profile` | A profile, fully resolved; `--posts` and `--replies` walk its feeds |
| `post` | A single post in full; `--replies` streams its thread, `--raw` the source |
| `replies` | Stream replies to a post as their own records |
| `feed` | Walk a profile's most-recent posts (alias for `profile --posts`) |
| `search` | Keyword search across public posts |
| `id` | Classify any Threads handle, id, shortcode, or URL, no network needed |
| `db` | Build a local SQLite dataset from a profile, then query it |
| `whoami` | Report how `th` is accessing Threads |
| `config` | Show resolved configuration and paths |
| `cache` | Inspect and clear the on-disk cache |
| `completion` | Generate a shell completion script |
| `version` | Print version, commit, and build date |

Run `th <command> --help` for the full flag list on any command.

## Recipes

Pull a profile's last 50 posts as JSON Lines:

```sh
th profile zuck --posts -n 50 -o jsonl
```

Get a post and its whole reply thread:

```sh
th post <url> --replies -o jsonl
```

Collect every media URL from a profile's posts:

```sh
th profile zuck --posts -n 200 --fields media_urls -o jsonl
```

Rewrite a threads.net link to its canonical id:

```sh
th id "https://www.threads.net/@zuck/post/ABC123" -o json
```

Build a dataset: crawl a profile's posts into SQLite, then query it:

```sh
th db build zuck --posts -n 200 --db zuck.db
th db query --db zuck.db "select username, sum(like_count) from posts group by 1"
```

## Output formats

Every command renders through the same formatter. Pick a format with `-o`, or
let `th` choose: a table when writing to a terminal, JSON Lines when piped.

```sh
th profile zuck --posts -o table   # aligned columns for reading
th profile zuck --posts -o jsonl   # one JSON object per line, for piping
th profile zuck --posts -o json    # a single JSON array
th profile zuck --posts -o csv     # spreadsheet friendly (tsv too)
th profile zuck --posts -o yaml    # YAML documents
th profile zuck --posts -o url     # just the permalink
```

Narrow the columns with `--fields`, or template each row:

```sh
th profile zuck --posts --fields permalink,like_count,reply_count
th profile zuck --posts --template '{{.Permalink}} {{.LikeCount}}'
```

`--raw` (on `post`) prints the upstream HTML untouched, for when you want to
parse it yourself.

## Configuration

`th` keeps its cache and data under the standard XDG paths (`~/.cache/th` and
`~/.local/share/th` by default; honor `XDG_CACHE_HOME` and `XDG_DATA_HOME` to
move them). See the resolved paths and settings any time:

```sh
th config show
th config path
```

Useful global flags (all have sensible defaults):

| Flag | Meaning |
| --- | --- |
| `-o, --output` | Output format (default auto) |
| `-n, --limit` | Maximum records (`0` means unlimited) |
| `--delay` | Minimum delay between requests, to stay polite (default 1s) |
| `--retries` | Retry attempts on 429/5xx |
| `--no-cache` | Bypass the on-disk cache |
| `--token` | Official Graph API token (or `THREADS_TOKEN`) |
| `--session` / `--csrf` | Logged-in session (or `THREADS_SESSION` / `THREADS_CSRF`) |

## Exit codes

| Code | Meaning |
| --- | --- |
| `0` | Success |
| `1` | Generic error |
| `2` | Usage error |
| `3` | Content not found or unavailable |
| `4` | Login wall: the content is not public |
| `5` | Rate limited |
| `6` | Network error |

## Development

```sh
make test    # run the test suite
make vet     # go vet
make build   # build ./bin/th
make smoke   # run every command and assert it works or walls cleanly
```

The code is layered. `cli/` is the command tree built on Cobra. `threads/` is
the library it sits on: the HTTP client, cache, the server-rendered page
parsers, the logged-out GraphQL queries, and the SQLite store. `pkg/thid/` is a
standalone handle/id/shortcode/URL classifier with no other dependencies,
importable on its own.

## License

Apache-2.0. See [LICENSE](LICENSE).
