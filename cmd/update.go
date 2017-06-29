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

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [id]",
	Short: "Update password with id.",
	Long:  "Only the fields that are included are updated, the other fields are left unchanged.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Printf("%s\n\n", cmd.Short)
			fmt.Println(cmd.UsageString())
			os.Exit(1)
		}
		uri := "api/v4/passwords/" + args[0] + ".json"
		payload, _ := json.Marshal(newpassword)
		resp := putTpm(uri, payload)
		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode == 204 {
			fmt.Printf("Password updated successfully")
		} else {
			var status map[string]interface{}
			err = json.Unmarshal(body, &status)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(status["message"].(string))
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&newpassword.Name, "name", "", "", "name")
	updateCmd.Flags().StringVarP(&newpassword.Username, "username", "u", "", "username")
	updateCmd.Flags().StringVarP(&newpassword.Password, "password", "p", "", "password")
	updateCmd.Flags().StringVarP(&newpassword.Email, "email", "e", "", "e-mail")
	updateCmd.Flags().StringVarP(&newpassword.Tags, "tags", "t", "", "tags (list of comma separated strings)")
	updateCmd.Flags().StringVarP(&newpassword.Notes, "notes", "n", "", "notes")
	updateCmd.Flags().StringVarP(&newpassword.AccessInfo, "access", "a", "", "access info (url)")
	updateCmd.Flags().StringVarP(&newpassword.ExpiryDate, "expiry", "", "", "in ISO 8601 format: yyyy-mm-dd")
}
