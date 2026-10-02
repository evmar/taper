Read ./schema.md to understand the expected file schema for new entries.

To add a new performance given a URL, use `yt-dlp` to fetch the YouTube page
and extract the relevant information.  Do not fetch the actual video stream.

Example:

```
$ yt-dlp --dump-json --skip-download 'https://www.youtube.com/watch?v=ferZnZ0_rSM'
```

The track list and times are often found in the video description or
comments.  If the times aren't in the first few comments, it's fine to leave them out.
