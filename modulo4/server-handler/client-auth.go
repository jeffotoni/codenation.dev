// Front-end in Go server
// @jeffotoni
// 2019-01-07

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func main() {

	LoginValid("xxxxxxxxx", "curso@gmail.com", "1234")

}

// Let's call the login endpoint
func LoginValid(token, email, password string) string {

	// Url and endpoint
	apiUrl := "http://localhost:8080"

	// only endpoint
	resource := "/v1/api/auth"

	// cancel if you pass the team
	ctx, cancel := context.WithCancel(context.TODO())
	afterFuncTimer := time.AfterFunc(2*time.Second, func() { cancel() })
	defer afterFuncTimer.Stop()

	data := url.Values{}
	data.Set("email", email)
	data.Set("password", password)

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()
	println(urlStr)

	req, err := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	req = req.WithContext(ctx)

	req.Header.Set("Authorization", "Bearer xxxxxxxxxx")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return string("error")
	}
	defer resp.Body.Close()

	fmt.Println(resp)

	if resp.Status == "200 OK" {
		type MessageJson struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		var mJson = MessageJson{}
		bodyMsg, _ := ioutil.ReadAll(resp.Body)
		println(string(bodyMsg))
		if err := json.Unmarshal(bodyMsg, &mJson); err != nil {
			return string("error")
		}
		// check if success
		if strings.ToLower(mJson.Status) == "success" {
			return string("success")
		} else {
			return string("error")
		}
	} else {

		// try again 2 times
		return string("error")
	}
}
