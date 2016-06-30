// Copyright © 2016 Benjamin Martensson <benjamin.martensson@nrk.no>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/user"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

type Password struct {
	ID           int    `json:"id,omitempty" yaml:"id,omitempty"`
	Name         string `json:"name,omitempty" yaml:"name,omitempty"`
	NotesSnippet string `json:"notes_snippet,omitempty" yaml:"notes_snippet,omitempty"`
	Tags         string `json:"tags,omitempty" yaml:"tags,omitempty"`
	AccessInfo   string `json:"access_info,omitempty" yaml:"access_info,omitempty"`
	Username     string `json:"username,omitempty" yaml:"username,omitempty"`
	Email        string `json:"email,omitempty" yaml:"email,omitempty"`
	Password     string `json:"password,omitempty" yaml:"password,omitempty"`
	Notes        string `json:"notes,omitempty" yaml:"notes,omitempty"`
	UpdatedOn    string `json:"updated_on,omitempty" yaml:"updated_on,omitempty"`
	CreatedOn    string `json:"created_on,omitempty" yaml:"created_on,omitempty"`
	Project      struct {
		Name string `json:"name,omitempty" yaml:"name,omitempty"`
	} `json:"project,omitempty" yaml:"project,omitempty"`
}

var config = make(map[string]string)

func hmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

func reqTpm(uri string) *http.Response {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	time := strconv.FormatInt(time.Now().Unix(), 10)
	unhash := uri + time
	hash := hmac256(unhash, viper.GetString("privkey"))
	req, err := http.NewRequest("GET", viper.GetString("base")+uri, nil)
	req.Header.Add("X-Public-Key", viper.GetString("pubkey"))
	req.Header.Add("X-Request-Hash", hash)
	req.Header.Add("X-Request-Timestamp", time)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

func writeConfig() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	configFile := usr.HomeDir + "/.tpm.json"
	configJson, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		fmt.Println("error:", err)
	}
	err = ioutil.WriteFile(configFile, configJson, 0600)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Written to", configFile)
}
