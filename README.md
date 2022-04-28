# LUKSO CLI


## Install Cobra

    go install github.com/spf13/cobra-cli@latest

    cobra init --pkg-name luksocli

## Add A New Command

Execute

    cobra-cli add [COMMAND_NAME] # lukso-cli

If it is a subcommand of **examplecmd** then rename the file to

    mv cmd/[command_name].go cmd/[examplecmd]_[command_name].go

!!!Document the command in this README.md
    


##  Commands

    lukso network describe -ip [xxx.xxx.xxx.xxx]   # shows peers and enode of a geth node (should show the enr of beacon)

## Download
Use the following command to download `lukso`
```bash
curl https://raw.githubusercontent.com/lukso-network/lukso-cli/main/cli_downloader.sh | bash
```

Or use go installer to install `lukso-cli` into your GOPATH
```bash
go install github.com/lukso-network/lukso-cli@v0.0.1-dev
```