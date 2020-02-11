# Status Key Utility

## Installation
- `go get github.com/status-im/security-utils`
- `cd key-util`
- `go install`

## Command Usage
### legacySeedToKey
Will generate the `m/44'/60'/0'/0'/1` Status wallet key using the legacy (pre March 2018) generation method

usage:
- `key-util legacySeedToKey "your seed phrase here"`

## Contributing
The CLI was built using [cobra](https://github.com/spf13/cobra) so follow its command line arguments for adding functionality.
