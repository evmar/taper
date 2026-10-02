The files in this directory each represent a single performance as identified
by a YouTube URL, and include metadata about where each track within the
performance lies.

## Example file

```
url = "https://www.youtube.com/watch?v=jnm2Lf1TiB4"
artist = "Malcriada"
title = "Live on KEXP"
date = "2026-06-17"
tracks = [
    { start = "00:55", end = "03:27", name = "Under" },
    { start = "03:37", end = "06:27", name = "Six Seven" },
    { start = "06:33", end = "10:15", name = "Acid Lover" },
    { start = "10:29", end = "14:37", name = "Miénteme" },
]
]
```

## Schema

Files are TOML with this schema.

Top level:
- url: string
- artist: string
- title: string; from YouTube title but omit artist name or "full performance" annotations
- date: string, in YYYY-MM-DD format
- tracks: array of tracks

Tracks:
- start, string, MM:SS of when the track starts (omit if unknown)
- end, string, MM:SS of when the track ends
- name, string, track name
