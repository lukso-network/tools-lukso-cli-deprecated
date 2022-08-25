# LUKSO CLI


## Installation

Use the following command to install `lukso` binary into `/usr/local/bin` directory:

```
sudo curl https://install.l16.lukso.network | sudo bash
```


If you want to install the **current release candidate**, use the following command:

```sh
sudo curl https://raw.githubusercontent.com/lukso-network/lukso-cli/main/install-rc.sh | sudo bash
```
 

## Documentation

See the [full documentation](./docs/cli.md) for more.

The network subcommand contains commands relevant for running a node or adding a validator. 

```sh
lukso network 
```

The wallet subcommand deals with the properties of a single wallet 

```sh
lukso wallet
```

The up subcommand deals with the universal profile smart contracts

```sh
lukso up
```

## Development


### Install Cobra

    go install github.com/spf13/cobra-cli@latest

    cobra init --pkg-name luksocli


## Tests
    

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

