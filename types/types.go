package types

type VMDetails struct {
	BridgeName      string
	TapName1        string
	TapName2        string
	VM1Eth0IP       string
	VM2Eth0IP       string
	MacAddress1     string
	MacAddress2     string
	BridgeIPAddress string
	BridgeGatewayIP string
}

type Launcher struct {
	UserID          string
	BridgeName      string
	TapName2        string
	VM2Eth0IP       string
	MacAddress2     string
	BridgeIPAddress string
	BridgeGatewayIP string
}
