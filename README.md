# Vault helper tool

A small tool written in go when you need to search your vault.

Requires that you have **VAULT_ADDR** set in your environment. For the token you can either have **VAULT_TOKEN** set or if not set it reads from **~/.vault-token** for the token.

Take care when using the **tree** and **search** functionality as if there are a lot of paths the process can take quite a while.

## Badges

[![Release](https://img.shields.io/github/release/ilijamt/vht.svg?style=for-the-badge)](https://github.com/ilijamt/vht/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=for-the-badge)](/LICENSE.md)

## Help

The tool is quite simple to use, to get more details about the tool check the help command

```bash
$ vht help
A simple vault helper tool that simplifies the usage of Vault

Usage:
  vht [command]

Available Commands:
  delete      Delete a path recursively
  help        Help about any command
  search      Search in the secrets data
  tree        Print out a list of all the secrets in a path
  verify      Verify connection to Vault
  version     Shows the version of the application

Flags:
  -h, --help   help for vht

Use "vht [command] --help" for more information about a command.
```
