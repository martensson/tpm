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

type Newpassword struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	ProjectID int    `json:"project_id"`
	Email     string `json:"email"`
	Tags      string `json:"tags"`
	Notes     string `json:"notes"`
}

var newpassword Newpassword

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
			newpassword.Name = args[0]
		}
		if newpassword.Username == "" {
			fmt.Println("flag --username not provided, aborting.")
			os.Exit(1)
		}
		if newpassword.Password == "" {
			fmt.Println("flag --password not provided, aborting.")
			os.Exit(1)
		}
		if newpassword.ProjectID == -1 {
			fmt.Println("flag --project not provided, aborting.")
			os.Exit(1)
		}
		uri := "api/v4/passwords.json"
		payload, _ := json.Marshal(newpassword)
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
	createCmd.Flags().StringVarP(&newpassword.Username, "username", "u", "", "username (required)")
	createCmd.Flags().StringVarP(&newpassword.Password, "password", "p", "", "password (required)")
	createCmd.Flags().IntVarP(&newpassword.ProjectID, "project", "i", -1, "project_id (required)")
	createCmd.Flags().StringVarP(&newpassword.Email, "email", "e", "", "e-mail")
	createCmd.Flags().StringVarP(&newpassword.Tags, "tags", "t", "", "tags (list of comma separated strings)")
	createCmd.Flags().StringVarP(&newpassword.Notes, "notes", "n", "", "notes")
}
