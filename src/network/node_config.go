package network

type DataVolume struct {
	Volume string `yaml:",omitempty"`
}

type NodeDetails struct {
	IP   string `yaml:",omitempty"`
	Name string `yaml:",omitempty"`
}

type ClientDetails struct {
	StatsAddress string `yaml:",omitempty"`
	Verbosity    string `yaml:",omitempty"`
	Etherbase    string `yaml:",omitempty"`
	DataVolume   string `yaml:",omitempty"`
	NetworkId    string `yaml:",omitempty"`
	Bootnode     string `yaml:",omitempty"`
	Version      string `yaml:",omitempty"`
}

type PortDescription struct {
	HttpPort string `yaml:",omitempty"`
	PeerPort string `yaml:",omitempty"`
}

type NodeConfigs struct {
	Configs   *DataVolume    `yaml:",omitempty"`
	Keystore  *DataVolume    `yaml:",omitempty"`
	Node      *NodeDetails   `yaml:",omitempty"`
	Execution *ClientDetails `yaml:",omitempty"`
	Consensus *ClientDetails `yaml:",omitempty"`
	Validator *ClientDetails `yaml:",omitempty"`

	Ports map[string]PortDescription `yaml:",omitempty"`
}

func (d *DataVolume) getVolume() string {
	return d.Volume
}

func (node *NodeDetails) getIP() string {
	return node.IP
}

func (node *NodeDetails) getName() string {
	return node.Name
}

func (cd *ClientDetails) getStatAddress() string {
	return cd.StatsAddress
}

func (cd *ClientDetails) getVerbosity() string {
	return cd.Verbosity
}

func (cd *ClientDetails) getEtherbase() string {
	return cd.Etherbase
}

func (cd *ClientDetails) getDataVolume() string {
	return cd.DataVolume
}

func (cd *ClientDetails) getNetworkID() string {
	return cd.NetworkId
}

func (cd *ClientDetails) getBootnode() string {
	return cd.Bootnode
}

func (cd *ClientDetails) getVersion() string {
	return cd.Version
}

func (pd PortDescription) getHttpPort() string {
	return pd.HttpPort
}

func (pd *PortDescription) getPeerPort() string {
	return pd.PeerPort
}

func (nc *NodeConfigs) getConfigs() *DataVolume {
	return nc.Configs
}

func (nc *NodeConfigs) getKeystore() *DataVolume {
	return nc.Keystore
}

func (nc *NodeConfigs) getNode() *NodeDetails {
	return nc.Node
}

func (nc *NodeConfigs) getExecution() *ClientDetails {
	return nc.Execution
}

func (nc *NodeConfigs) getConsensus() *ClientDetails {
	return nc.Consensus
}

func (nc *NodeConfigs) getValidator() *ClientDetails {
	return nc.Validator
}

func (nc *NodeConfigs) getPort(portName string) *PortDescription {
	if nc.Ports == nil {
		return nil
	}
	portDesc := nc.Ports[portName]
	return &portDesc
}
