package screens

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"strings"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"

	extkeys "github.com/status-im/security-utils/key-util/legacyExtKeys"
	"github.com/status-im/status-go/eth-node/crypto"
)

// KEYINFO is a struct that contains all derived info from a seedphrase and password
type KEYINFO struct {
	Private   string `json:"private"`
	Public    string `json:"public"`
	Address   string `json:"address"`
	Etherscan string `json:"etherscan"`

	seedEntry     *widget.Entry
	passwordEntry *widget.Entry

	hyperlinks map[string]*widget.Hyperlink
	labels     map[string]*widget.Label
}

func (k *KEYINFO) newLabel(name string) *widget.Label {
	w := widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Monospace: true})
	k.labels[name] = w
	return w
}

func (k *KEYINFO) newHyperlink(name string) *widget.Hyperlink {
	w := widget.NewHyperlink("", parseURL("https://google.com"))
	k.hyperlinks[name] = w
	return w
}

// Parses a given string for a URL and returns it
func parseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return link
}

// NewKEYINFO returns a new keyinfo
func NewKEYINFO() *KEYINFO {
	rand.Seed(time.Now().UnixNano())
	return &KEYINFO{
		labels:     make(map[string]*widget.Label),
		hyperlinks: make(map[string]*widget.Hyperlink),
	}
}

func (k *KEYINFO) processSeedAndPass() {
	// trim whitespace
	seedPhrase := strings.TrimSpace(k.seedEntry.Text)
	// convert the seed phrase to a private key
	mn := extkeys.NewMnemonic("")
	// check validity of given mnemonic
	isValid := mn.ValidMnemonic(seedPhrase, 0)
	if !isValid {
		fmt.Printf("seedphrase has an error, please verify all information is spelled correctly.")
	}

	masterKey, err := extkeys.NewMaster(mn.MnemonicSeed(seedPhrase, k.passwordEntry.Text), []byte(extkeys.Salt))
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
	walletPrivKeyString := hex.EncodeToString(crypto.FromECDSA(walletKey))
	walletPubKey := walletKey.Public().(*ecdsa.PublicKey)
	walletPubKeyString := hex.EncodeToString(crypto.FromECDSAPub(walletPubKey))
	address := crypto.PubkeyToAddress(walletKey.PublicKey).Hex()

	k.labels["private"].SetText("0x" + walletPrivKeyString)
	k.labels["public"].SetText("0x" + walletPubKeyString)
	k.labels["address"].SetText(address)

	k.hyperlinks["etherscan"].SetText("Check it on Etherscan")
	k.hyperlinks["etherscan"].SetURL(parseURL("https://etherscan.io/address/" + address))
}

// Submit will derive keys and derivative information
func (k *KEYINFO) Submit() {
	// print legacy private key of m/44'/60'/0'/0'/0'
	go k.processSeedAndPass()
	return
}

func makeFormTab(k *KEYINFO) fyne.Widget {
	k.seedEntry = widget.NewEntry()
	k.seedEntry.SetPlaceHolder("friend margin letter stove assist retire anchor inherit replace height useful pass")
	k.passwordEntry = widget.NewPasswordEntry()
	k.passwordEntry.SetPlaceHolder("Password")

	form := &widget.Form{
		OnSubmit: func() {
			k.Submit()
		},
	}
	form.Append("Seedphrase", k.seedEntry)
	form.Append("Password", k.passwordEntry)

	form.Append("Private Key", k.newLabel("private"))
	form.Append("Public Key", k.newLabel("public"))
	form.Append("Address", k.newLabel("address"))

	outputBox := widget.NewHScrollContainer(
		k.newHyperlink("etherscan"),
	)
	form.Append("", outputBox)

	return form
}

// LegacyScreen shows a panel containing Legacy Key Derivation form
func LegacyScreen() fyne.CanvasObject {
	k := NewKEYINFO()
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, nil),
		widget.NewTabContainer(
			widget.NewTabItem("Keys", makeFormTab(k)),
		),
	)
}
