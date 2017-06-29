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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	yaml "gopkg.in/yaml.v2"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search for matching passwords.",
	Long:  "The returned data is the same as in the passwords lists (all active, archived, favorite and search) in the web interface.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Printf("%s\n\n", cmd.Short)
			fmt.Println(cmd.UsageString())
			os.Exit(1)
		}
		uri := "api/v4/passwords/search/" + args[0] + ".json"
		resp := getTpm(uri)
		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		var passwords PasswordList
		err = json.Unmarshal(body, &passwords)
		if err != nil {
			log.Fatal(err)
		}
		if len(passwords) == 0 {
			fmt.Println("No passwords found.")
		} else if len(passwords) == 1 {
			showCmd.Run(nil, []string{strconv.Itoa(passwords[0].ID)})
		} else {
			if outputFormat == "json" {
				jsonpass, _ := json.MarshalIndent(&passwords, "", "  ")
				fmt.Printf(string(jsonpass))
			} else {
				yamlpass, _ := yaml.Marshal(&passwords)
				fmt.Printf(string(yamlpass))
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringVarP(&outputFormat, "output", "o", "yaml", "Output format: json|yaml")
}
