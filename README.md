# LUKSO CLI


## Install Cobra

    go install github.com/spf13/cobra-cli@latest

    cobra init --pkg-name luksocli

## Add A New Command

Execute

    cobra-cli add [COMMAND_NAME] # lukso-cli

If it is a subcommand of **examplecmd** then rename the file to

    mv cmd/[command_name].go cmd/[examplecmd]_[command_name].go

Generate the doc with

    go run main.go docs
## Download
Use the script to download the binary

    ...

or use

    go install ...

##  Commands

The cli distinguishes subcommands. You can always 

### Network
The network subcommand contains commands relevant for running a node or adding a validator. 

    lukso network 

### Wallet
The wallet subcommand deals with the properties of a single wallet 

    lukso wallet

### UP
The up subcommand deals with the universal profile smart contracts

    lukso up