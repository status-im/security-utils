# How to use `key-util` binary on Windows

1. Download the binary for Windows [here](https://github.com/status-im/security-utils/releases/download/v0.1/key-util.exe)
2. Open the file explorer to the directory where you downloaded it
3. Open the command prompt in this folder via [this](https://www.thewindowsclub.com/how-to-open-command-prompt-from-right-click-menu) method
4. Execute the downloaded binary via the following command:
   - `key-util.exe legacySeedToKey "<Your Seedphrase Here>" <Your Password Here>`
     - where `<Your Seedphrase Here>` is your seedphrase
     - and `<Your Password Here>` is your account password
5. The program will out put the private key, public key, and address
6. Check the address to be correct via a blockchain explorer like Etherscan
7. Repeat as necessary