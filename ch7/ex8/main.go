package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

// Track は音楽プレイリストの要素
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func tracks() []*Track {
	return []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go", "Aoby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // 列幅を計算して表を印字する
}

type lessFunc func(x, y *Track) bool

type byColumns struct {
	tracks    []*Track
	lessFuncs []lessFunc
}

func (x byColumns) Len() int { return len(x.tracks) }
func (x byColumns) Swap(i, j int) {
	x.tracks[i], x.tracks[j] = x.tracks[j], x.tracks[i]
}

func (x byColumns) Less(i, j int) bool {
	if len(x.lessFuncs) < 1 {
		panic("set less function")
	}
	for _, less := range x.lessFuncs {
		switch {
		case less(x.tracks[i], x.tracks[j]):
			return true
		case less(x.tracks[j], x.tracks[i]):
			return false
		}
	}
	return false
}

func byArtist(x, y *Track) bool {
	return x.Artist < y.Artist
}

func byYear(x, y *Track) bool {
	return x.Year < y.Year
}

func byTitle(x, y *Track) bool {
	return x.Title < y.Title
}

func byAlbum(x, y *Track) bool {
	return x.Album < y.Album
}

func byLength(x, y *Track) bool {
	return x.Length < y.Length
}

func sortByColumns(t []*Track, f ...lessFunc) *byColumns {
	return &byColumns{
		t,
		f,
	}
}

func main() {
	ts := tracks()
	d := sortByColumns(ts, byYear, byArtist)
	sort.Sort(d)
	printTracks(ts)
	fmt.Println()

	// ts = tracks()
	// d = sortByColumns(ts, byYear)
	// sort.Stable(d)
	// d.lessFuncs[0] = byArtist
	// sort.Stable(d)
	// printTracks(ts)
}
