// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 187.

// Sorting sorts a music playlist into a variety of orders.
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
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

type trackData struct {
	Tracks []*Track
}

var trackList = template.Must(template.New("tracklist").Parse(`
<h1>music list</h1>
<table>
<tr style = 'text-align: left'>
	<th>Title</th>
	<th>Artist</th>
	<th>Album</th>
	<th>Year</th>
	<th>Length</th>
{{range .Tracks}}
<tr>
	<td>{{.Title}}</td>
	<td>{{.Artist}}</td>
	<td>{{.Album}}</td>
	<td>{{.Year}}</td>
	<td>{{.Length}}</td>
</tr>
{{end}}
</table>
`))

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
	tw.Flush() // calculate column widths and print table
}

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
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	var keys sortKeys
	keys.first = "Title"
	keys.second = "Year"

	m := map[sortKeys]func(x, y *Track) bool{
		{"Title", "Artist"}:  lessByTitleArtist,
		{"Title", "Album"}:   lessByTitleAlbum,
		{"Title", "Year"}:    lessByTitleYear,
		{"Title", "Length"}:  lessByTitleLength,
		{"Artist", "Title"}:  lessByArtistTitle,
		{"Artist", "Album"}:  lessByArtistAlbum,
		{"Artist", "Year"}:   lessByArtistYear,
		{"Artist", "Length"}: lessByArtistLength,
	}
	if m[keys] == nil {
		fmt.Print("not implemented currently")
		os.Exit(1)
	}
	sort.Sort(customSortBy2Key{tracks, m[keys]})
	printTracks(tracks)
	var data trackData
	data.Tracks = append(data.Tracks, tracks...)
	if err := trackList.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}

type customSortBy2Key struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSortBy2Key) Len() int           { return len(x.t) }
func (x customSortBy2Key) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSortBy2Key) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func lessByTitle(x, y *Track) bool {
	if x.Title != y.Title {
		return x.Title < y.Title
	}
	return false
}

func lessByArtist(x, y *Track) bool {
	if x.Artist != y.Artist {
		return x.Artist < y.Artist
	}
	return false
}

func lessByAlbum(x, y *Track) bool {
	if x.Album != y.Album {
		return x.Album < y.Album
	}
	return false
}

func lessByYear(x, y *Track) bool {
	if x.Year != y.Year {
		return x.Year < y.Year
	}
	return false
}

func lessByLength(x, y *Track) bool {
	if x.Length != y.Length {
		return x.Length < y.Length
	}
	return false
}

func lessByTitleArtist(x, y *Track) bool {
	if x.Title != y.Title {
		return x.Title < y.Title
	}
	return lessByArtist(x, y)
}

func lessByTitleAlbum(x, y *Track) bool {
	if x.Title != y.Title {
		return x.Title < y.Title
	}
	return lessByAlbum(x, y)
}

func lessByTitleYear(x, y *Track) bool {
	if x.Title != y.Title {
		return x.Title < y.Title
	}
	return lessByYear(x, y)
}

func lessByTitleLength(x, y *Track) bool {
	if x.Title != y.Title {
		return x.Title < y.Title
	}
	return lessByYear(x, y)
}

func lessByArtistTitle(x, y *Track) bool {
	if x.Artist != y.Artist {
		return x.Artist < y.Artist
	}
	return lessByTitle(x, y)
}

func lessByArtistAlbum(x, y *Track) bool {
	if x.Artist != y.Artist {
		return x.Artist < y.Artist
	}
	return lessByAlbum(x, y)
}

func lessByArtistYear(x, y *Track) bool {
	if x.Artist != y.Artist {
		return x.Artist < y.Artist
	}
	return lessByYear(x, y)
}

func lessByArtistLength(x, y *Track) bool {
	if x.Artist != y.Artist {
		return x.Artist < y.Artist
	}
	return lessByLength(x, y)
}