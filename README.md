# taper

Convert YouTube live performances into MP3s.

## Running

Run like `go run . path/to/toml` .  It will download and split tracks based on the
timings found in the file.  Requires `yt-dlp` and `ffmpeg`.

## Adding performances

Write a file by hand, or:

```
$ cd performance
$ codex "add an entry for https://youtube-url-here, following instructions in agents.md"
```

and then edit the file.  In particular note that unnamed entries in the track list are omitted
from the output, which can be used to remove interview chat or breaks etc.
