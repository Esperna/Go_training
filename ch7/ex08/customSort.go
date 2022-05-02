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

type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type sortKeys struct {
	first  string
	second string
}

func main() {
	key1 := flag.String("k1", "Title", "first sort key")
	key2 := flag.String("k2", "Artist", "second sort key")
	var keys sortKeys
	flag.Parse()
	keys.first = *key1
	keys.second = *key2

	if !(*key1 == "Title" || *key1 == "Artist" || *key1 == "Album" || *key1 == "Year" || *key1 == "Length") {
		fmt.Print("invalid first key")
		os.Exit(1)
	}
	if !(*key1 == "Title" || *key1 == "Artist" || *key1 == "Album" || *key1 == "Year" || *key1 == "Length") {
		fmt.Print("invalid second key")
		os.Exit(1)
	}
	tracks := CustomSortBy1st2ndKey(*key1, *key2, tracks1)
	printTracks(tracks)
}

func CustomSortBy1st2ndKey(key1, key2 string, tracks []*Track) []*Track {
	lessBy1st2ndKey := func(x, y *Track) bool {
		m := map[string]func(x, y *Track) bool{
			"Title":  lessByTitle,
			"Artist": lessByArtist,
			"Album":  lessByAlbum,
			"Year":   lessByYear,
			"Length": lessByLength,
		}
		if m[key1](x, y) {
			return m[key1](x, y)
		}
		if m[key1](y, x) {
			return m[key1](x, y)
		}
		return m[key2](x, y)
	}
	sort.Sort(customSortBy2Key{tracks, lessBy1st2ndKey})
	return tracks
}

func SortStable(key1, key2 string, tracks []*Track) []*Track {
	lessBy1st2ndKey := func(x, y *Track) bool {
		m := map[string]func(x, y *Track) bool{
			"Title":  lessByTitle,
			"Artist": lessByArtist,
			"Album":  lessByAlbum,
			"Year":   lessByYear,
			"Length": lessByLength,
		}
		if m[key1](x, y) {
			return m[key1](x, y)
		}
		if m[key1](y, x) {
			return m[key1](x, y)
		}
		return m[key2](x, y)
	}
	sort.Stable(customSortBy2Key{tracks, lessBy1st2ndKey})
	return tracks
}

type customSortBy2Key struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSortBy2Key) Len() int           { return len(x.t) }
func (x customSortBy2Key) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSortBy2Key) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func lessByTitle(x, y *Track) bool  { return x.Title < y.Title }
func lessByArtist(x, y *Track) bool { return x.Artist < y.Artist }
func lessByAlbum(x, y *Track) bool  { return x.Album < y.Album }
func lessByYear(x, y *Track) bool   { return x.Year < y.Year }
func lessByLength(x, y *Track) bool { return x.Length < y.Length }
