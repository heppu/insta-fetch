package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
	"path"

	"github.com/valyala/fasthttp"
)

const (
	URL   = "https://www.instagram.com/%s/media?max_id=%s"
	USAGE = `
	Usage:
		insta-fetch [nick] [path]
	`
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal(USAGE)
	}

	process(os.Args[1], os.Args[2])
}

func process(nick, filePath string) {
	ir := &InstaResponse{}
	pics := make([]Image, 0)
	minId := "0"

	for {
		log.Printf("Fetch data for %s with min_id %s\n", nick, minId)
		code, body, err := fasthttp.Get(nil, fmt.Sprintf(URL, nick, minId))
		if err != nil {
			log.Printf("Fetching data for %s failed: %s\n", nick, err)
			return
		}
		if code != 200 {
			log.Printf("Fetching data for %s failed with code: %d\n", nick, code)
			return
		}
		if err = json.Unmarshal(body, ir); err != nil {
			log.Printf("Unmarshaling data for %s failed: %s\n", nick, err)
			return
		}
		if len(ir.Items) == 0 {
			return
		}

		for _, item := range ir.Items {

			if item.Type == "image" {
				pics = append(pics, item.Images.StandardResolution.Image)
			}
		}
		if !ir.MoreAvailable {
			break
		}
		minId = ir.Items[len(ir.Items)-1].ID
	}

	log.Printf("Found %d pictures for %s\n", len(pics), nick)

	file, err := os.Create(path.Join(filePath, nick+".html"))
	if err != nil {
		log.Printf("Could not create file %s\n", err)
	}

	t, _ := template.New("photos").Parse(HTML)
	if err = t.Execute(file, pics); err != nil {
		log.Printf("Could not write file %s\n", err)
	}
}
