
.PHONY: target

create:
	rm -rf beta_genesis_wallets;go run main.go network createGenesis -m "circle foam employ kangaroo green shy real oblige hope canvas skate entry level submit hammer mean quality pulp kiwi avocado hat matter sample chat" -c beta

zip:
	for i in */; do zip -r "${i%/}.zip" "$i"; done




 Dear Beta Genesis Validator,

 Attached in this email you will find a .zip file with your personal genesis validator key. Although we are working on a different distribution method for the test network release. Please download the attached file.

 Please follow the provided steps.

 1. We have been working on a new LUKSO cli, in order to retrieve it, run the following command:

 curl ..../install.sh | sh

 It is recommended to put the cli binary into the PATH or in a bin directory that is already part of the PATH.

 2. You also have to create a directory and setup the network:

 mkdir MY_NODE_DIR
 cd MY_NODE_DIR
 lukso network init

 This will download the dependencies.

 3. Now, unzip the validator key archive beta_x.zip and copy the contents into the directory

 MY_NODE_DIR/keystore

 Your directory structure should look like this:

 MY_NODE_DIR
   configs
   data
   keystores
   	keys
   	lodestar-secrets
   	nimbus-keys
   	password.txt
   	prysm
   	pubkeys.json
   	secrets
   	teku-keys
   	teku-secrets
   docker-compose.yml
   .env
   node_config.yaml

 4. Start the node with:

 lukso network start validator

 Your node should be syncing and finally validating.

 NOTE: Do NOT modify the file .env. It is auto generated. Modify only node_config.yaml

 Thank you again for your support and have a lot of fun,
 LUKSO NETWORK TEAM