#!/bin/bash

args=(
    -x
    --audio-format mp3
    --audio-quality 0
    --embed-thumbnail
    --embed-metadata
    --split-chapters
    'https://www.youtube.com/watch?v=jnm2Lf1TiB4'
)
yt-dlp "${args[@]}"
