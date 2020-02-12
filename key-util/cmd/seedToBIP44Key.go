/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"crypto/ecdsa"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/status-im/keycard-go/derivationpath"
	"github.com/status-im/status-go/eth-node/crypto"
	"github.com/status-im/status-go/extkeys"
)

// seedToBIP44KeyCmd represents the seedToBIP44Key command
var seedToBIP44KeyCmd = &cobra.Command{
	Use:   "seedToBIP44Key",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		seedToBIP44KeyFunc(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(seedToBIP44KeyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// seedToBIP44KeyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// seedToBIP44KeyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func seedToBIP44KeyFunc(rawSeed string, rawPath string) {
	// trim whitespace
	seedPhrase := strings.TrimSpace(rawSeed)
	// convert the seed phrase to a private key
	mn := extkeys.NewMnemonic()
	masterKey, err := extkeys.NewMaster(mn.MnemonicSeed(seedPhrase, ""))
	if err != nil {
		fmt.Printf("cannot create master extended key: %v\n", err)
		os.Exit(1)
	}
	_, path, err := derivationpath.Decode(rawPath)
	walletExtKey, err := masterKey.Derive(path)
	if err != nil {
		fmt.Printf("cannot drive key from given path: %v\n", err)
		os.Exit(1)
	}
	walletKey := walletExtKey.ToECDSA()
	walletPubKey := walletKey.Public().(*ecdsa.PublicKey)
	walletPubKeyBytes := crypto.FromECDSAPub(walletPubKey)
	address := crypto.PubkeyToAddress(walletKey.PublicKey).Hex()

	// print legacy private key of m/44'/60'/0'/0'/0'
	fmt.Printf("private: %#x\n", crypto.FromECDSA(walletKey))
	fmt.Printf("public:  %#x\n", walletPubKeyBytes)
	fmt.Printf("address: %+v\n\n", address)
}
