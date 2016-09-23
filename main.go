package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/heppu/insta-fetch/models"
	"github.com/valyala/fasthttp"
	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	URL = "https://www.instagram.com/%s/media?max_id=%s"
)

var (
	nicks    = kingpin.Flag("nicks", "Instagram nicks to fetch").Required().Short('n').Strings()
	filePath = kingpin.Flag("path", "Path where nick.hmtl will be saved").Short('p').String()
)

func main() {
	kingpin.Parse()

	wg := &sync.WaitGroup{}
	for _, nick := range *nicks {
		wg.Add(1)
		go process(nick, wg)
	}
	wg.Wait()
}

func process(nick string, wg *sync.WaitGroup) {
	defer wg.Done()

	ir := &models.InstaResponse{}
	pics := make([]models.Image, 0)
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
				if index := strings.Index(item.Images.StandardResolution.Image.URL, "?ig"); index != -1 {
					item.Images.StandardResolution.Image.URL = item.Images.StandardResolution.Image.URL[0:index]
					pics = append(pics, item.Images.StandardResolution.Image)
				}
			}
		}
		if !ir.MoreAvailable {
			break
		}
		minId = ir.Items[len(ir.Items)-1].ID
	}

	log.Printf("Found %d pictures for %s\n", len(pics), nick)

	file, err := os.Create(path.Join(*filePath, nick+".html"))
	if err != nil {
		log.Printf("Could not create file %s\n", err)
	}

	t, _ := template.New("photos").Parse(models.HTML)
	if err = t.Execute(file, pics); err != nil {
		log.Printf("Could not write file %s\n", err)
	}
}
