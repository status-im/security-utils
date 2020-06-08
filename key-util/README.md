# Status Key Utility

## Installation
- `go get github.com/status-im/security-utils`
- `cd key-util`
- `go install`

## Command Usage
### seedToBIP44Key
Will generate a associate [BIP44](https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki) key information for any given `path`

usage:
- `key-util seedToBIP44Key <seedphrase> <path>`
  - \<seedphrase\> (string): user's seed phrase.
  - \<path\> (string): BIP44 path surrounded in quotes

### legacySeedToKey
Will generate the `m/44'/60'/0'/0'/0` Status wallet key using the legacy (pre March 2018) generation method.  You WILL need your account password to correctly generate the account. 

usage:
- `key-util legacySeedToKey <seedphrase> <password>`
  - \<seedphrase\> (string): user's seed phrase.
  - \<password\> (string): user's BIP39 passphrase.  For legacy Status accounts, this was the account password.

### seedToStatusRandomName
TODO

### pubkeyToEthAddress
TODO

## Contributing
The CLI was built using [cobra](https://github.com/spf13/cobra) so follow its command line arguments for adding functionality.
