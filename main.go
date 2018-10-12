package main

import (
	"errors"
	"fmt"
	"github.com/quirkey/magick"
	"io/ioutil"
	"net/http"
	"strings"
)

// Uses imagemagick to increase the image size by 50%
func convertIncrease50(img []byte) {

	image, err := magick.NewFromBlob(img, "gif")
	defer image.Destroy()
	err = image.Resize("resize 150%")
	if err != nil {
		panic(err)
	}
}

func main() {
	req, err := http.NewRequest("GET", "http://www.bramblemet.co.uk", nil)
	if err != nil {
		panic(err)
	}
	client := new(http.Client)
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return errors.New("Redirect")
	}

	response, err := client.Do(req)
	if err != nil {
		if response.StatusCode == http.StatusFound { //status code 302

			url, err := response.Location()
			if err != nil {
				fmt.Println(err)
			}

			fmt.Printf("response: ")
			fmt.Println(response)
			//fmt.Println(url.Host)
			//fmt.Println(url.Path)
			s := strings.Split(url.Path, "/")
			fmt.Printf("token: ")
			fmt.Println(s[1]) //s[1] contains the token

			req, err := http.NewRequest("GET", "http://bramblemet.co.uk/"+s[1]+"/GetImage.ashx?src=windreport.gif", nil) //
			if err != nil {
				panic(err)
			}

			// header contents gleaned form looking at the network view in chrome  developer mode. Server seems to be sensative to the referer header setting.
			req.Header.Set("Accept-Encoding", "gzip, deflate")
			req.Header.Set("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")
			req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
			req.Header.Set("Accept", "image/webp,image/apng,image/*,*/*;q=0.8")
			req.Header.Set("Referer", "http://bramblemet.co.uk/"+s[1]+"/default.aspx")

			fmt.Printf("req: ")
			fmt.Println(req)

			response, err := client.Do(req)
			if err != nil {
				panic(err)
			}

			fmt.Printf("response content length: ")
			fmt.Println(response.ContentLength)

			// read the image into a buffer
			b, err := ioutil.ReadAll(response.Body)
			response.Body.Close()
			if err != nil {
				panic(err)
			}

			fmt.Printf("Body: \n")
			fmt.Printf("%x", b)

		} else {
			panic(err)
		}
	}

}
