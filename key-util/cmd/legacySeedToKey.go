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
	extkeys "github.com/status-im/security-utils/key-util/legacyExtKeys"
	"github.com/status-im/status-go/eth-node/crypto"
)

// legacySeedToKeyCmd represents the legacySeedToKey command
var legacySeedToKeyCmd = &cobra.Command{
	Use:   "legacySeedToKey",
	Short: "Convert a legacy seed phrase into the first generated wallet private key",
	Long: `The way Status generated wallet addresses from a given seed phrase was changed 
	after March 2018.  In specific, the salt that was used was changed, so if upgrading
	an account generated before that, the seed phrase would create a different wallet
	account and funds would not be accessible.`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		legacyFunc(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(legacySeedToKeyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// legacySeedToKeyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// legacySeedToKeyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func legacyFunc(rawSeed string, password string) {
	// trim whitespace
	seedPhrase := strings.TrimSpace(rawSeed)
	// convert the seed phrase to a private key
	mn := extkeys.NewMnemonic("")
	masterKey, err := extkeys.NewMaster(mn.MnemonicSeed(seedPhrase, password), []byte(extkeys.Salt))
	if err != nil {
		fmt.Printf("cannot create master extended key: %v\n", err)
		os.Exit(1)
	}

	walletExtKey, err := masterKey.BIP44Child(60, 0)
	// walletKey, err := walletExtKey.Child(1)
	if err != nil {
		fmt.Printf("cannot extract chat key from master key: %v\n", err)
		os.Exit(1)
	}
	walletKey := walletExtKey.ToECDSA()
	walletPubKey := walletKey.Public().(*ecdsa.PublicKey)
	walletPubKeyBytes := crypto.FromECDSAPub(walletPubKey)
	address := crypto.PubkeyToAddress(walletKey.PublicKey).Hex()

	// print legacy private key of m/44'/60'/0'/0'/0'
	fmt.Printf("private: %#x\n", crypto.FromECDSA(walletKey))
	fmt.Printf("public:  %x\n", walletPubKeyBytes)
	fmt.Printf("address: %+v\n\n", address)
}
