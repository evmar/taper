package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/pelletier/go-toml/v2"
)

type Performance struct {
	URL        string `toml:"url"`
	Artist     string `toml:"artist"`
	Title      string `toml:"title"`
	Date       string `toml:"date"`
	parsedDate time.Time
	Tracks     []*Track `toml:"tracks"`
}

type Track struct {
	Start string `toml:"start"`
	End   string `toml:"end"`
	Name  string `toml:"name"`
}

func run() error {
	flag.Usage = func() {
		fmt.Printf("Usage: %s <input.toml>\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(2)
	}

	perf, err := load(flag.Arg(0))
	if err != nil {
		return err
	}

	outDir := filepath.Join("out", fmt.Sprintf("%s %s", perf.Title, perf.Date))
	if err := os.MkdirAll(outDir, 0777); err != nil {
		return err
	}
	if err := os.Chdir(outDir); err != nil {
		return err
	}

	_, err = os.Stat("out.opus")
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
		if err := download(perf); err != nil {
			return err
		}
	}

	if err := split(perf); err != nil {
		return err
	}

	return nil
}

func load(path string) (*Performance, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var perf Performance
	decoder := toml.NewDecoder(bytes.NewReader(data)).DisallowUnknownFields()
	if err := decoder.Decode(&perf); err != nil {
		return nil, err
	}

	if perf.Artist == "" {
		return nil, fmt.Errorf("missing artist")
	}
	if perf.Title == "" {
		return nil, fmt.Errorf("missing title")
	}

	perf.parsedDate, err = time.Parse("2006-01-02", perf.Date)
	if err != nil {
		return nil, fmt.Errorf("failed to parse date %v", perf.Date)
	}

	for i, track := range perf.Tracks {
		if track.Start == "" {
			return nil, fmt.Errorf("track %d (%q) missing start", i+1, track.Name)
		}
	}

	return &perf, nil
}

func download(perf *Performance) error {
	cmd := exec.Command("yt-dlp",
		"--extract-audio",
		"--write-thumbnail",
		"--output", "out",
		perf.URL,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func split(perf *Performance) error {
	trackNumber := 1
	for i, track := range perf.Tracks {
		if track.Name == "" {
			continue
		}

		// input 0: audio, seeked
		args := []string{
			"-ss", track.Start,
		}
		if track.End != "" {
			args = append(args, "-to", track.End)
		} else if i+1 < len(perf.Tracks) {
			args = append(args, "-to", perf.Tracks[i+1].Start)
		}
		args = append(args, "-i", "out.opus")

		// input 1: album art
		args = append(args, "-i", "out.webp")

		// output options after input filename
		args = append(args,
			"-y", // overwrite output
			// copy audio stream as mp3
			"-map", "0:a",
			"-c:a", "libmp3lame",
			"-q:a", "4", // quality 4

			// copy album art if available
			"-map", "1:v?",
			"-c:v", "mjpeg", // convert to jpeg
			"-disposition:v", "attached_pic", // attach as picture

			"-metadata", fmt.Sprintf("artist=%s", perf.Artist),
			"-metadata", fmt.Sprintf("album=%s %s", perf.Title, perf.Date),
			"-metadata", fmt.Sprintf("date=%s", perf.Date),
			"-metadata", fmt.Sprintf("track=%d", trackNumber),
			"-metadata", fmt.Sprintf("title=%s", track.Name),
			"-metadata", fmt.Sprintf("comment=%s", perf.URL),

			fmt.Sprintf("%02d %s.mp3", trackNumber, track.Name),
		)
		log.Println("ffmpeg", args)
		cmd := exec.Command("ffmpeg", args...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return err
		}
		trackNumber += 1
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
