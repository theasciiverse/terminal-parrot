#!/bin/sh
# Records the catalog preview at 100x30 the way av starts terminal-parrot
# (no arguments: the party parrot), ended by SIGTERM, which it answers by
# giving the terminal back. Needs asciinema 3.
#
#   asciiverse/record-preview.sh out/terminal-parrot asciiverse/terminal-parrot.cast [seconds]
set -eu
binary=$(cd "$(dirname "$1")" && pwd)/$(basename "$1")
out=$2
seconds=${3:-12}
home=$(mktemp -d)
trap 'rm -rf "$home"' EXIT
mkdir "$home/bin" && cp "$binary" "$home/bin/terminal-parrot"
cd "$home"  # nothing in the recording names a local path
HOME="$home" PATH="$home/bin:$PATH" TERM=xterm-256color COLORTERM=truecolor LANG=C.UTF-8 \
asciinema rec --headless --overwrite -q \
  --window-size 100x30 --idle-time-limit 2 \
  --capture-env "TERM,COLORTERM,LANG" --title "terminal-parrot" \
  --command "timeout --foreground -s TERM $seconds terminal-parrot" \
  "$OLDPWD/$out"
