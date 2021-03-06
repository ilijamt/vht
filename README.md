# Vault helper tool

A small tool written in go when you need to search your vault.

Requires that you have **VAULT_ADDR** set in your environment. For the token you can either have **VAULT_TOKEN** set or if not set it reads from **~/.vault-token** for the token.

Take care when using the **tree** and **search** functionality as if there are a lot of paths the process can take quite a while.

**IMPORTANT**:  If you want to delete from a KV v2 backend make sure you prefix your path with **secret/metadata**, where **secret** is the name of the backend you want to delete in.

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
  completion  Generates bash completion scripts
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

## Install

### Pre-compiled binary

#### manually

Download the pre-compiled binaries from the [releases](https://github.com/ilijamt/vht/releases) page and copy to the desired location.

#### macos (homebrew)

```bash
brew tap ilijamt/tap
brew install vht
```

#### snapcraft

```bash
snap install vht
```

## Example 

Create a docker instance
```shell script
docker run -d  --rm --cap-add=IPC_LOCK -p 1234:1234 --name vault -e 'VAULT_DEV_ROOT_TOKEN_ID=myroot' -e 'VAULT_DEV_LISTEN_ADDRESS=0.0.0.0:1234' vault
```

Now let's pre fill it with some data
```shell script

export VAULT_ADDR='http://127.0.0.1:1234'
export VAULT_TOKEN=myroot
vault secrets enable -version=2 -path kv2 kv
vault secrets enable -version=1 -path kv1 kv
export KV_DATE=$(date +%s)

vht search -r kv1
vht search -r kv2
 
vault kv put kv1/abba timestamp=$KV_DATE value=abba
vault kv put kv1/abba/caab/qaz timestamp=$KV_DATE value=qaz
vault kv put kv1/qaz/abba/caab timestamp=$KV_DATE value=caab
vault kv put kv1/abba/overwrite timestamp=$KV_DATE value=overwrite
vault kv put kv1/abba/overwrite timestamp=$KV_DATE value=overriden

vault kv put kv2/abba timestamp=$KV_DATE value=abba
vault kv put kv2/abba/caab/qaz timestamp=$KV_DATE value=qaz
vault kv put kv2/qaz/abba/caab timestamp=$KV_DATE value=caab
vault kv put kv2/abba/overwrite timestamp=$KV_DATE value=overwrite
vault kv put kv2/abba/overwrite timestamp=$KV_DATE value=overriden
```

Let's see it in action

```
$ vht search -r kv1 -k "q.z"
kv1/abba/caab/qaz
-----------------
timestamp = 1584114901
value = qaz

kv1/qaz/abba/caab
-----------------
timestamp = 1584114901
value = caab

$ vht search -r kv2 -k "q.z" -f "c*b"
kv2/qaz/abba/caab
-----------------
timestamp = 1584114901
value = caab

$ vht search -r kv2 -k "q.z"
kv2/abba/caab/qaz
-----------------
timestamp = 1584114901
value = qaz

kv2/qaz/abba/caab
-----------------
timestamp = 1584114901
value = caab

$ vht search -r kv1 -k "q.z" -f "c*b"
kv1/qaz/abba/caab
-----------------
timestamp = 1584114901
value = caab

```

