package goandroid

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/kkdai/youtube/v2"
)

var callback SomeCallback

type SomeCallback interface {
	DoSomething(msg string)
}

type GoPrinter struct {
	brand string
}

func (goPrinter *GoPrinter) PrintMsg(msg string) string {
	return goPrinter.brand + msg
}

func GetGoPrinter() *GoPrinter {
	rst := new(GoPrinter)
	rst.brand = "jim green"
	return rst
}

func Greetings(name string, callback SomeCallback) string {
	callback.DoSomething(name)
	return fmt.Sprintf("Hello, %s", name)
}

func ExampleClient(savePath string, videoID string) {
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		panic(err)
	}

	stream, _, err := client.GetStream(video, &video.Formats[0])
	if err != nil {
		panic(err)
	}

	file, err := os.Create(savePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}
}

func FormatList(videoID string) string {
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		return err.Error()
	}

	rst := ""
	for _, element := range video.Formats {
		format := fmt.Sprintf("quality:%s audio:%d\n", element.Quality, element.AudioChannels)
		rst = rst + format
	}

	return rst
}

// Example usage for playlists: downloading and checking information.
func ExamplePlaylist() {
	playlistID := "PLQZgI7en5XEgM0L1_ZcKmEzxW1sCOVZwP"
	client := youtube.Client{}

	playlist, err := client.GetPlaylist(playlistID)
	if err != nil {
		panic(err)
	}

	/* ----- Enumerating playlist videos ----- */
	header := fmt.Sprintf("Playlist %s by %s", playlist.Title, playlist.Author)
	println(header)
	println(strings.Repeat("=", len(header)) + "\n")

	for k, v := range playlist.Videos {
		fmt.Printf("(%d) %s - '%s'\n", k+1, v.Author, v.Title)
	}

	/* ----- Downloading the 1st video ----- */
	entry := playlist.Videos[0]
	video, err := client.VideoFromPlaylistEntry(entry)
	if err != nil {
		panic(err)
	}
	// Now it's fully loaded.

	fmt.Printf("Downloading %s by '%s'!\n", video.Title, video.Author)

	stream, _, err := client.GetStream(video, &video.Formats[0])
	if err != nil {
		panic(err)
	}

	file, err := os.Create("video.mp4")

	if err != nil {
		panic(err)
	}

	defer file.Close()
	_, err = io.Copy(file, stream)

	if err != nil {
		panic(err)
	}

	println("Downloaded /video.mp4")
}
