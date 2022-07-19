# LUKSO CLI


## Download

### Install
Use the following command to download `lukso` inside `/usr/local/bin` directory.  

```
sudo curl https://raw.githubusercontent.com/lukso-network/lukso-cli/main/install.sh | sudo bash
```
 

### RC

to only download the binary use

```
sudo curl https://raw.githubusercontent.com/lukso-network/lukso-cli/main/install-rc.sh | sudo bash
```
 

## Commands

[HERE](./docs/cli.md) you can find the documentation of the commands.

### Network
The network subcommand contains commands relevant for running a node or adding a validator. 

    lukso network 

### Wallet
The wallet subcommand deals with the properties of a single wallet 

    lukso wallet

### UP
The up subcommand deals with the universal profile smart contracts

    lukso up


## Development


### Install Cobra

    go install github.com/spf13/cobra-cli@latest

    cobra init --pkg-name luksocli

## Test

    

### Add A New Command

Execute

    cobra-cli add [COMMAND_NAME] # lukso-cli

If it is a subcommand of **examplecmd** then rename the file to

    mv cmd/[command_name].go cmd/[examplecmd]_[command_name].go

Generate the doc with

    go run main.go docs 

## Release

Release are created by creating a Git release. Make sure to 

1. modify **install.sh** with the desired version 
2. add updates to CHANGELOG.md
3. cmd/root.go  -> rootCmd.Version

