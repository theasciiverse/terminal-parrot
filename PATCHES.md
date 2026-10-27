# Changes from upstream

This is the Asciiverse packaging of [terminal-parrot](https://github.com/jmhobbs/terminal-parrot)
by John Hobbs and contributors: the party parrot dancing in your terminal
in 256 colours (Go, termbox-go). Upstream is the `upstream` remote; the
`asciiverse` branch carries the changes below.

Built from upstream `master` at `2472c23` (2026-03-29), three commits
after the 1.2.1 tag: Windows line endings in animation files, and two
README edits. The version here is 1.2.1+20260329.

## Licence

MIT, "Copyright (c) 2016 John Hobbs" in LICENSE (checked 2026-10-18). Its
modules: termbox-go (MIT), go-isatty (MIT), go-runewidth (MIT); the
golang.org/x/sys listed in go.mod is not compiled in (no build here,
Windows and macOS included, downloads it).

The parrot frames are a jp2a conversion of the Party Parrot GIF (the
README thanks jp2a), and that GIF was made by fans from BBC footage of
Sirocco the kakapo; nobody grants rights to it either, but it has been
redistributed freely for a decade, packaged in nixpkgs, Homebrew and
Snap, and it is what the project is.

## Code

- **Pedro the raccoon is left out.** Upstream's second built-in
  animation, `animations/pedro.animation` (added 2024 after issue #25,
  which names no source), is an ASCII rendering of the 2024 "Pedro"
  raccoon meme, frames from someone else's video set to a commercial
  song, with no grant of rights. The file is removed, so the binary
  embeds only the parrot. `-path` still loads animations from files.
- **Any key quits.** Esc, Ctrl+C and q did; now every key does.
- **SIGTERM and SIGHUP quit cleanly.** termbox does not catch them, so a
  SIGTERM ended it in raw mode and left the terminal without echo. They
  now interrupt the event loop and termbox closes as it does for a key
  (exit status 0).
- **The parrot stands centred on the bottom row.** Upstream drew every
  frame from the top left corner. Now the box that holds all the frames
  is centred across the terminal with its last row on the terminal's
  last row, the same box for every frame so the parrot keeps its own
  bob, and it follows a resize. A terminal smaller than the parrot cuts
  off the top and both sides evenly.

## Packaging

- Upstream's workflows (`release.yaml`, goreleaser on every pushed tag,
  ours included, and `test-snap-can-build.yml`) are removed. Its
  `.goreleaser.yaml`, `Dockerfile` and `snap/` stay, unused here.
- `.github/workflows/release.yml`: on one Linux runner, `go vet`, the
  tests, then static builds (cgo off) for Linux, macOS and Windows,
  x86-64 and ARM64 each, packed as `terminal-parrot-<target>` with
  LICENSE, README and this file, plus SHA256SUMS. Runs on our tags only
  (`*-av.*`).
- `asciiverse/record-preview.sh` records `asciiverse/terminal-parrot.cast`,
  the catalog preview: 12 seconds of `terminal-parrot` at 100x30, ended by
  SIGTERM.

av starts it as `terminal-parrot` with no arguments.

## Checked

- Built and tested (go vet, go test) in golang:1-alpine with cgo off;
  cross-built for macOS and Windows, x86-64 and ARM64.
- Linux, in tmux at 100x30: q, x, Space, Esc and Ctrl+C quit with status
  0, SIGTERM too after the change (before it: status 143, echo and line
  editing left off); resized to 20x6 and 120x40 it keeps running;
  `-list` shows only the parrot.
- Windows: the x86-64 build run under Wine draws the parrot and quits on q
  with status 0.
- macOS: built, not run.

## Ours

- **How to play it, at the top of README.md.** A short note on installing and
  starting it with av, or from a release archive; the original text follows
  unchanged.
