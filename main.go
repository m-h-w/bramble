package main

import (
    "errors"
    "fmt"
    "net/http"
    "strings"
)

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
            
            url,err := response.Location()
            if (err != nil){
            	fmt.Println (err)
            }
	        fmt.Println(response)
            fmt.Println(url.Host)
            fmt.Println(url.Path)
            s:=strings.Split (url.Path,"/")
            fmt.Println(s[1])	//s[1] contains the token
            
            
            req, err := http.NewRequest("GET", "http://bramblemet.co.uk/" + s[1] + "/default.aspx", nil) // 
            
            // header contents gleaned form looking at the network view in chrome  developer mode. Seems to be sensative to the referer header setting.
            
            req.Header.Set("Accept-Encoding","gzip, deflate")
            req.Header.Set("Accept-Language","en-GB,en-US;q=0.9,en;q=0.8")
            req.Header.Set("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
            req.Header.Set("Accept","image/webp,image/apng,image/*,*/*;q=0.8")
            req.Header.Set("Referer","http://bramblemet.co.uk/" + s[1] + "/default.aspx")
		    if err != nil {
		        panic(err)
		    }
            response, err := client.Do(req)
            if err != nil {
		        panic(err)
		    }
            
            fmt.Println(response)
            
        } else {
            panic(err)
        }
    }

}	