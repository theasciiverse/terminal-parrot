# terminal-parrot (Asciiverse packaging)

John Hobbs' party parrot for the terminal (Go, termbox-go), MIT; packaged
for the Asciiverse catalog as a demo (kind `art`). Not our code: keep
changes small, list every one in PATCHES.md, and offer fixes upstream.
The Pedro raccoon animation is removed on purpose (third-party meme
footage, no rights); do not bring it back on an update.

- Remotes: `upstream` is github.com/jmhobbs/terminal-parrot (branch
  `master`), `origin` is github.com/theasciiverse/terminal-parrot.
- Branch `asciiverse` = upstream master plus our commits.
- Release: tag `v<version>-av.<n>` (now `v1.2.1+20260329-av.1`) and push
  only that tag; the workflow runs on `*-av.*` tags.
- av starts it as `terminal-parrot`.
- Build, in a copy of the tree: `docker run --rm --cpus 4 --memory 8g -e
  CGO_ENABLED=0 -e GOFLAGS=-p=2 -v "$PWD":/src -w /src golang:1-alpine go
  build -trimpath -ldflags="-s -w" -o out/terminal-parrot .`
- Preview: `asciiverse/record-preview.sh out/terminal-parrot asciiverse/terminal-parrot.cast`.
