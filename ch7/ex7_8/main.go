// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 187.

// Sorting sorts a music playlist into a variety of orders.
package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

//!+main
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

//!-main

//!+printTracks
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

//!-printTracks

//!+artistcode
type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//!-artistcode

//!+yearcode
type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type sortKeys struct {
	first  string
	second string
}

//!-yearcode
func main() {
	fmt.Println("\nCustom:")
	key1 := flag.String("k1", "Title", "first sort key")
	key2 := flag.String("k2", "Artist", "second sort key")
	var keys sortKeys
	flag.Parse()
	keys.first = *key1
	keys.second = *key2

	m := map[sortKeys]func(x, y *Track) bool{
		{"Title", "Artist"}: lessByTitleArtist,
		{"Title", "Album"}:  lessByTitleAlbum,
		{"Title", "Year"}:   lessByTitleYear,
		{"Title", "Length"}: lessByTitleLength,
	}

	sort.Sort(customSortBy2Key{tracks, m[keys]})
	printTracks(tracks)
}

func lessByTitleArtist(x, y *Track) bool {
	if x.Title != y.Title {
		return x.Title < y.Title
	}
	if x.Artist != y.Artist {
		return x.Artist < y.Artist
	}
	return false
}

func lessByTitleAlbum(x, y *Track) bool {
	if x.Title != y.Title {
		return x.Title < y.Title
	}
	if x.Album != y.Album {
		return x.Album < y.Album
	}
	return false
}

func lessByTitleYear(x, y *Track) bool {
	if x.Title != y.Title {
		return x.Title < y.Title
	}
	if x.Year != y.Year {
		return x.Year < y.Year
	}
	return false
}

func lessByTitleLength(x, y *Track) bool {
	if x.Title != y.Title {
		return x.Title < y.Title
	}
	if x.Length != y.Length {
		return x.Length < y.Length
	}
	return false
}

//!+customcode
type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

//!-customcode

type customSortBy2Key struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSortBy2Key) Len() int           { return len(x.t) }
func (x customSortBy2Key) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSortBy2Key) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }
