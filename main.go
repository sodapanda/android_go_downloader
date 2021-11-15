package main

import (
	"fmt"
	"strings"

	"fit.soda/youtubecrawler/goandroid"
	"github.com/pborman/uuid"
)

func main() {
	uuidWithHyphen := uuid.NewRandom()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	fmt.Println(uuid)

	rst := goandroid.Greetings("good", nil)
	fmt.Println(rst)

	// goandroid.ExampleClient("video.mp4", "JHX4Ntq1T5U")
	formats := goandroid.FormatList("JHX4Ntq1T5U")
	fmt.Println(formats)
}
