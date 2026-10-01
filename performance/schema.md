The files in this directory each represent a single performance as identified
by a YouTube URL, and include metadata about where each track within the
performance lies.

## Example file

```
url = "https://www.youtube.com/watch?v=jnm2Lf1TiB4"
title = "Malcriada - Full Performance (Live on KEXP)'
date = "2026-06-17"
tracks = [
    { time = "00:55", name = "Under" }
    { time = "03:37", name = "Six Seven" }
    { time = "06:33", name = "Acid Lover" }
    { time = "10:29", name = "Miénteme" }
]
```

## Schema

Files are TOML with this schema.

Top level:
- url: string
- title: string, taken from the YouTube page
- date: string, in YYYY-MM-DD format
- tracks: array of tracks

Tracks:
- time, string, MM:SS of when the track starts (omit if unknown)
- name, string, track name
