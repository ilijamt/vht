# Vault helper tool
[![Go Report Card](https://goreportcard.com/badge/github.com/ilijamt/vht)](https://goreportcard.com/report/github.com/ilijamt/vht)
[![Codecov](https://img.shields.io/codecov/c/gh/ilijamt/vht)](https://app.codecov.io/gh/ilijamt/vht)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ilijamt/vht)](go.mod)
[![GitHub](https://img.shields.io/github/license/ilijamt/vht)](LICENSE)
[![Release](https://img.shields.io/github/release/ilijamt/vht.svg)](https://github.com/ilijamt/vht/releases/latest)

A small tool written in go when you need to search your vault.

Requires that you have **VAULT_ADDR** set in your environment. 

Token will be set from: 
- **VAULT_TOKEN**
- **~/.vault-token**
- Using a token helper as defined in **~/.vault**

If you want to use the token helper then you have to login first by using `vault login` so you are logged into the system. You can check use this [vault-token-helper](https://github.com/ilijamt/vault-token-helper)

All the environment variables that work with the Vault client will work with this tool as well.

Take care when using the **tree** and **search** functionality as if there are a lot of paths the process can take quite a while.

**IMPORTANT**: If you want to delete from a KV v2 backend make sure you prefix your path with **secret/metadata**, where **secret** is the name of the backend you want to delete in.

**IMPORTANT**: The input for searching is based on https://pkg.go.dev/regexp/syntax.

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

## Example 

Create a docker instance
```shell script
docker run -d  --rm --cap-add=IPC_LOCK -p 1234:1234 --name vault -e 'VAULT_DEV_ROOT_TOKEN_ID=myroot' -e 'VAULT_DEV_LISTEN_ADDRESS=0.0.0.0:1234' vault
```

Now let's pre-fill it with some data
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

### Case-insensitive searching

The input is based on https://pkg.go.dev/regexp/syntax so you can define what ever input you want/need to do the searching.

```shell
❯ docker run -d  --rm --cap-add=IPC_LOCK -p 8200:8200 --name vault -e 'VAULT_DEV_ROOT_TOKEN_ID=token' -e 'VAULT_DEV_LISTEN_ADDRESS=0.0.0.0:8200' vault

❯ export VAULT_ADDR='http://127.0.0.1:8200'
❯ export VAULT_TOKEN=token

❯ vault secrets enable -version=2 -path kv2 kv
Success! Enabled the kv secrets engine at: kv2/

❯ export KV_DATE=$(date +%s) 
❯ vault kv put kv2/aBbA timestamp=$KV_DATE value=abba
❯ vault kv put kv2/abba/cAAb/qaz timestamp=$KV_DATE value=qaz
❯ vault kv put kv2/new value=aBbA
❯ vault kv put kv2/nEw value=abba

❯ vht tree -r kv2
kv2/aBbA
kv2/nEw
kv2/new
kv2/abba/cAAb/qaz

❯ vht tree -r kv2 -k 'abba'
kv2/abba/cAAb/qaz

❯ vht tree -r kv2 -k '(?i)abba'
kv2/aBbA
kv2/abba/cAAb/qaz

❯ vht search -r kv2 -d -f '(?i)abba'
kv2/aBbA
kv2/nEw
kv2/new

❯ vht search -r kv2 -d -f 'abba'
kv2/aBbA
kv2/nEw
```
