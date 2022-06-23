# Change Log


## [v0.4.2] - L16 Release


### Fixed

- fixed bug in deposit (v0.4.1-rc)
### Added

- check to see if Docker is running
- pull & update configurations when **lukso network update** is run
### Changed

- root command use description from `cli` to `lukso`


## [v0.4.0] - L16 Release


### Added

- lukso network init without --chain param will error out
### Changed

- removed L16Beta references in code
- --nodeName renamed to --stats-name in **lukso network init**


### Fixed

- spelling of LYX
- contract address for L16

## [v0.3.6] - L16 Release


### Added

- added node params to l16 to change client versions on the fly
- command **lukso network validator setup range --from [FROM_POSITION] --to [TO_POSITION]
- added possibility to run l16 with **lukso network init --chain l16**
- added node_params for each chain
- added recovery mechanism


### Fixed

- spelling of LYX
- contract address for L16
 


## [v0.2.2] -


### Fixed

- amount in validator setup

## [v0.2.2] -


### Fixed

- amount in validator setup

## [0.2.1] - 2022-06-06

This update prepares the cli for multi chains, has a lot of improvements in std output and 
offers a new deposit functionality.

### Added
### Changed
- network balance show denominated balances
### Fixed

- command **lukso network block -n N** that will return the N-th execution block
- command **lukso network validator start** that will start a validator
- command **lukso network validator stop** that will stop a validator
- command **lukso network balance -a 0x...** to call balance of validator
- command **lukso network validator check.** to check the status of all deposited validators
- new chain env **local**. A network can be setup with **lukso network init --chain local**
- command **lukso network update** to update the bootnodes 
- new chain env **local**. A network can be setup with **lukso network init --chain local**
- command **lukso network validator deposits** reads out all deposits in the DepositContract
- command **lukso network validator byKey** to describe any validator key on any network
- dev environment for experimenting with nodes

### Changed

- command **lukso network start validator** became deprecated
- 2 new scripts installer script **install.sh** & **download.sh** to give more control on how to get the CLI binary
- the way node_conf.yaml is handled. It will not be auto-generated but **lukso network init** must be called to generate the file.
- deposit checks on past events to see if a key deposited already
- **lukos network validator describe --dry** differs between pending and observed and offers a dry run

### Fixed

- added Api entity to attach to the load balancer
- changed the **lukso network init** to receive configs from different directory
- cleaned up node_conf.yaml and introduced a new data structure
- .env has a comment that it is auto generated
- made **lukso network init** command to prompt for node name if not given
- fixed problem with startup time of cli -> runs faster

## [0.0.4] - 2022-05-13

This update mainly improves existing commands 

### Added

- command **lukso network describe** that will give information on the state of the network
- command **lukso network validator deposit** that will deposit new validators

### Changed

- command **lukso network validator describe** that will show additional information
- the installer script **cli_downloader.sh** will copy the cli binary to /usr/local/bin

### Fixed

- command **lukso network validator setup** it will create a local transaction wallet.