#!/bin/bash

args=(
    -ss 00:55
    -to 03:37
    -i out.mp3
    -c copy
    '01 Under.mp3'
)
ffmpeg "${args[@]}"
