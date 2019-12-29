/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"encoding/json"

	coap "github.com/moroen/gocoap"
	"github.com/spf13/cobra"
)

var ident, key string

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var req coap.RequestParams

		scheme, host, port, path, err := processURI(args[0])

		switch scheme {
		case "coap":
			break
		case "coaps":
			if ident == "" {
				fmt.Println("Missing IDENT for COAPS call")
				return
			}

			if key == "" {
				fmt.Println("Missing KEY for COAPS call")
				return
			}
			req = coap.RequestParams{Host: host, Port: port, Uri: path, Id: ident, Key: key}
			break
		}

		status := "ok"
		resp, err := coap.GetRequest(req)
		if err != nil {
			switch (err) {
			case coap.UriNotFound: status = "UriNotFound"
			case coap.ErrorHandshake: status = "HandshakeError"
			case coap.Unauthorized: status = "Unauthorized"
			}
		}

		res := Message{Status: status, Result: string(resp) }

		jsonObj, err := json.Marshal(res)

		fmt.Println(string(jsonObj))
	},
}

func init() {
	getCmd.Flags().StringVar(&ident, "ident", "", "Identity")
	getCmd.Flags().StringVar(&key, "key", "", "Pre-shared key")
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
