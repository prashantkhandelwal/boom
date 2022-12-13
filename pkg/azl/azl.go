package azl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type Lyrics struct {
	Artistname string `json:"artist"`
	Featartist string `json:"feat"`
	Songname   string `json:"song"`
	Lyrics     string `json:"lyrics"`
}

func format(text string) string {
	t := strings.ToLower(text)
	t = strings.ReplaceAll(t, " ", "")
	return t
}

func FetchLyrics(artistname, songname string) (response *Lyrics) {

	l := &Lyrics{}
	c := colly.NewCollector()

	c.OnHTML(".col-xs-12.col-lg-8.text-center", func(e *colly.HTMLElement) {

		var depth int = 8

		artistname := e.DOM.Find(".lyricsh").Text()
		l.Artistname = strings.ReplaceAll(strings.TrimSpace(artistname), "Lyrics", "")

		songname := strings.ReplaceAll(e.DOM.Find(".col-xs-12.col-lg-8.text-center > b").Text(), "\"", "")
		l.Songname = strings.TrimSpace(songname)

		feat := e.DOM.Find(".feat")

		if feat.Length() > 0 {
			depth = 10
			l.Featartist = feat.Text()
		}

		lyrics := e.DOM.Find(".col-xs-12.col-lg-8.text-center > div:nth-child(" + strconv.Itoa(depth) + ")").Text()
		l.Lyrics = strings.TrimSpace(lyrics)
	})

	url := fmt.Sprintf("https://www.azlyrics.com/lyrics/%s/%s.html", format(artistname), format(songname))
	c.Visit(url)
	return l
}
