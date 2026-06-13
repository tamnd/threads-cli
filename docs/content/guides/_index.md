---
title: "Guides"
linkTitle: "Guides"
description: "Task-oriented walkthroughs for the things people do with th."
weight: 20
featured: true
---

Each guide is built around a job rather than a command. They assume you have run
the [quick start](/getting-started/quick-start/).

## Walk a profile's whole recent feed

```bash
th profile zuck --posts -n 100 -o jsonl > zuck-posts.jsonl
```

Each line is one post with its text, media, timestamp, and engagement counts.
Pipe it into `jq` to slice, or into `th db build` to store.

## Pull a post and its replies

```bash
th post <url> -o json                 # the post itself
th post <url> --replies -n 50 -o jsonl  # its reply thread
```

## Build and query a local dataset

```bash
th db build zuck --posts -n 200 --db zuck.db
th db query --db zuck.db \
  "select strftime('%Y-%m', timestamp) m, count(*), sum(like_count) \
   from posts group by 1 order by 1"
```
