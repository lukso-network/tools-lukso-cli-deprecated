# LUKSO CLI


## Install Cobra

    go install github.com/spf13/cobra-cli@latest

    cobra init --pkg-name luksocli

## Add A New Command

Execute

    cobra-cli add [COMMAND_NAME] # lukso-cli

If it is a subcommand of **examplecmd** then rename the file to

    mv cmd/[command_name].go cmd/[examplecmd]_[command_name].go

    

