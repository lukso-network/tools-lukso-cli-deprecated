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
docker-compose --upgrade
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

To start the node and the syncing process run 

```shell
lukso network start
```

You can always stop it by hitting

```shell
lukso network stop
```

Once the network is started log your nodes with

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

Run

```shell
lukso network validator setup
```

It will prompt a couple of questions. Most importantly how many validators you want to
run in this node. Mnemonics will be created and you will need to select a password for the keystore.
!!!The password will not be stored but the Mnemonic.

The setup command will also create a **transaction wallet**. This wallet is used to pay
for the transaction fee and the staking amount. The initial balance will be 0. So you need to 
make sure that a sufficient amount of LYX is stored on that wallet (Faucet, ...). Be aware that
10 validators also means that you need 10 * MIN_STAKING_AMOUNT of LYX + around 0.5 LYX for the transaction fee.

You can always see the state of your setup with

```shell
lukso network validator describe
```

## Deposit

Once the transaction wallet has enough LYX:

First start the validator

```shell
lukso network validator start
```

You can always stop it with

```shell
lukso network validator stop
```

then run

```shell
lukso network validator deposit --dry
```

The **dry** flag will ensure that the transactions will not be executed but you will get a picture
of what will happen.

Once you are sure that you everything looks correct run

```shell
lusko network validator deposit
```
The transactions will be executed sequentially.
Depending on the amount of validators this can take a moment of time.

When the command successfully finished you should see that all your validators are pending
when you run

```shell
lukso network validator describe
```
...and that your transaction wallet is almost empty :(.

The activation can take up to 10h. You can always check with the command above or by checking

```shell
https://explorer.consensus.[CHAIN_NAME].lukso.network
```

in the validator section

You should also check

```shell
lukso network log validator -f
```

to see if your node is running and your key(s) are accepted.

## Rewards

The CLI is not the greatest possibility to see your rewards (Explorer and Prometheus are doing a better job)

but to quickly check on your state you can always hit

```shell
lukso network validator describe
```

and once the validator is activated you will see a quick overview of your balance(s)

