// Copyright © 2017 Benjamin Martensson <benjamin.martensson@nrk.no>
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

	"github.com/spf13/cobra"
)

var (
	name      string
	username  string
	password  string
	projectID string
	email     string
	tags      string
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create a new password.",
	Long:  "Adds a new password to tpm.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Printf("%s\n\n", cmd.Short)
			fmt.Println(cmd.UsageString())
			os.Exit(1)
		} else {
			name = args[0]
		}
		if username == "" {
			fmt.Println("flag --username not provided, aborting.")
			os.Exit(1)
		}
		if password == "" {
			fmt.Println("flag --password not provided, aborting.")
			os.Exit(1)
		}
		if projectID == "" {
			fmt.Println("flag --project not provided, aborting.")
			os.Exit(1)
		}
		uri := "api/v4/passwords.json"
		var payload = []byte(`{"name":"` + name + `","project_id":` + projectID + `,"username":"` + username + `","password":"` + password + `","email":"` + email + `","tags":"` + tags + `"}`)
		resp := postTpm(uri, payload)
		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		var status map[string]interface{}
		err = json.Unmarshal(body, &status)
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode == 201 {
			fmt.Printf("Password created with id: %s\n", status["id"].(string))
		} else {
			fmt.Println(status["message"].(string))
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&username, "username", "u", "", "username (required)")
	createCmd.Flags().StringVarP(&password, "password", "p", "", "password (required)")
	createCmd.Flags().StringVarP(&projectID, "project", "i", "", "project_id (required)")
	createCmd.Flags().StringVarP(&email, "email", "e", "", "e-mail")
	createCmd.Flags().StringVarP(&tags, "tags", "t", "", "tags (list of comma separated strings)")
}
