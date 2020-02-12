# Status Key Utility

## Installation
- `go get github.com/status-im/security-utils`
- `cd key-util`
- `go install`

## Command Usage
### legacySeedToKey
Will generate the `m/44'/60'/0'/0'/0` Status wallet key using the legacy (pre March 2018) generation method.  You WILL need your account password to correctly generate the account. 

usage:
- `key-util legacySeedToKey "your seed phrase here" password`

## Contributing
The CLI was built using [cobra](https://github.com/spf13/cobra) so follow its command line arguments for adding functionality.
