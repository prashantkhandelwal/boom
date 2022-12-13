package id3

import (
	"log"

	"github.com/bogem/id3v2"
)

func Read(f string) (*id3v2.Tag, error) {
	tag, err := id3v2.Open(f, id3v2.Options{Parse: true})
	if err != nil {
		log.Fatal("Error while opening mp3 file: ", err)
	}
	defer tag.Close()
	return tag, nil

	//l := azl.FetchLyrics(tag.Artist(), tag.Title())
	//fmt.Println(l.Lyrics)
}

/*func boom() {
/*
	o, err := os.Open("Charlie Puth - One Call Away(128).mp3")
	defer o.Close()
	if err != nil {
		fmt.Println("ERROR: Cannot read file")
	}

	m, err := tag.ReadFrom(o)
	log.Println(m.Title())
	log.Println(m.AlbumArtist())
	log.Println(m.Artist())
	log.Println(m.Picture())*/

//tag.DeleteFrames(tag.CommonID("Attached picture"))
//tag.Save()

// artwork, err := ioutil.ReadFile("art.jpg")
// if err != nil {
// 	log.Fatal("Error while reading artwork file", err)
// }

// pic := id3v2.PictureFrame{
// 	Encoding:    id3v2.EncodingUTF8,
// 	MimeType:    "image/jpeg",
// 	PictureType: id3v2.PTFrontCover,
// 	Description: "Front cover",
// 	Picture:     artwork,
// }
// tag.AddAttachedPicture(pic)
// fmt.Println("Album art updated!")
// tag.Save()

// pictures := tag.GetFrames(tag.CommonID("Attached picture"))
// var c = 0
// for _, f := range pictures {
// 	pic, ok := f.(id3v2.PictureFrame)
// 	if !ok {
// 		log.Fatal("Couldn't assert picture frame")
// 	}

// 	fmt.Println(pic.Encoding.Name)
// 	// Do something with picture frame
// 	b := pic.Picture
// 	name := "art" + strconv.Itoa(c) + ".jpg"
// 	ff, err := os.Create(name)
// 	if err != nil {
// 		fmt.Println("ERROR")
// 	}
// 	defer ff.Close()

// 	_, err2 := ff.Write(b)
// 	if err2 != nil {
// 		fmt.Println("ERROR")
// 	}
// 	c++
// }

// Set tags
//tag.SetArtist("Aphex Twin")
//tag.SetTitle("Xtal")

/*comment := id3v2.CommentFrame{
	Encoding:    id3v2.EncodingUTF8,
	Language:    "eng",
	Description: "My opinion",
	Text:        "I like this song!",
}
tag.AddCommentFrame(comment)*/

// Write tag to file.mp3
/*if err = tag.Save(); err != nil {
	log.Fatal("Error while saving a tag: ", err)
}*/
