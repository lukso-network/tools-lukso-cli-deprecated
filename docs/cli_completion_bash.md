## cli completion bash

Generate the autocompletion script for bash

### Synopsis

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(cli completion bash)

To load completions for every new session, execute once:

#### Linux:

	cli completion bash > /etc/bash_completion.d/cli

#### macOS:

	cli completion bash > /usr/local/etc/bash_completion.d/cli

You will need to start a new shell for this setup to take effect.


```
cli completion bash
```

### Options

```
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions
```

### SEE ALSO

* [cli completion](cli_completion.md)	 - Generate the autocompletion script for the specified shell

###### Auto generated by spf13/cobra on 22-Apr-2022