# How To initialise A Node


## Prerequisites

You need to have docker and docker-compose installed on your system. Check out [Docker Installation](https://docs.docker.com/get-docker/) and [Docker Compose Installation](https://docs.docker.com/compose/install/)
for more instructions.

Example script for Debian

```shell
# install dependencies
sudo apt-get -y update
sudo apt-get -y install curl

# install docker
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh

# install docker-compose
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
docker-compose --version
```

You will also need [curl](tecmint.com/install-curl-in-linux/)

## Install the LUKSO CLI

Run

```shell
sudo curl https://raw.githubusercontent.com/lukso-network/lukso-cli/main/install.sh | sudo bash
```

The CLI binary will be moved to

```shell
/usr/local/bin
```

This will make the CLI accessible from everywhere

You can check the installation location by

```shell
which lukso
```

### Version
Check the version if you are not sure which version you are running

```shell
lukso -v
```

## Initialise A Node

You will need a node name that will be used to represent your node in the stats services.
Pick a name, e.g. "MySuperNode" and run

```shell
lukso network init --nodeName MySuperNode [--chain l16beta]
```

This command will 
* download the docker and config files
* update the bootnodes
* create .env

The **.env** is auto generated and you should never modify it by yourself. The configuration of your node
is specified in the file ./node_conf.yaml in the top directory of your node. Don't change the values
there unless you really no what you are doing or you have been asked so.

You can specify a network with the *--chain*. The default network is currently l16beta if you don't specify
anything further.

### Overview

To get an overview of the network run

```shell
lukso network describe
```

It will show you the participation rate, the last finalised block and other useful metrics

You can also run

```shell
lukso network validator describe deposits
```
to see how many deposits in total (all participants in the network) were already triggered.

## Start the node

To start the node and the syncing process run:

```shell
lukso network start
```

You can always stop it by hitting

```shell
lukso network stop
```

Once the network is started log your nodes

```shell
lukso network log consensus -f     // beacon client
lukso network log execution -f     // execution client
lukso network log validator -f     // validator client
```

You can also visits the stats pages to see if your nodes connect properly to them.
Open in a browser

```shell
https://stats.consensus.[CHAIN_NAME].lukso.network
https://stats.execution.[CHAIN_NAME].lukso.network
```

## Setup Your Validator




