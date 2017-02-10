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
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Create login configuration",
	Long:  "Add your HMAC private/public keys and TPM domain",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("TPM domain (tpm.example.com): ")
		var domain string
		fmt.Scanln(&domain)
		config["domain"] = domain
		viper.Set("domain", domain)
		fmt.Print("HMAC public key: ")
		var pubkey string
		fmt.Scanln(&pubkey)
		config["pubkey"] = pubkey
		viper.Set("pubkey", pubkey)
		fmt.Print("HMAC private key: ")
		var privkey string
		fmt.Scanln(&privkey)
		config["privkey"] = privkey
		viper.Set("privkey", privkey)
		resp := getTpm("api/v4/version.json")
		defer resp.Body.Close()
		//body, err := ioutil.ReadAll(resp.Body)
		//if err != nil {
		//	log.Fatal(err)
		//}
		if resp.StatusCode != 200 {
			fmt.Println("Authentication failure.")
			os.Exit(1)
		} else {
			fmt.Println("Authentication Successful.")
			writeConfig()
		}
	},
}

func init() {
	RootCmd.AddCommand(loginCmd)
}
