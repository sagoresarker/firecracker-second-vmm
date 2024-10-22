package database

import (
	"context"
	"log"

	"github.com/sagoresarker/firecracker-second-vmm/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetVMDetails retrieves the VM details for the given userID
func GetVMDetails(userID string) (types.VMDetails, error) {
	if mongoClient == nil {
		log.Fatal("MongoDB client not initialized.")
		return types.VMDetails{}, nil
	}

	collection := mongoClient.Database("firecrackerdb").Collection("vm-info")

	filter := bson.M{"userID": userID}
	result := collection.FindOne(context.Background(), filter, options.FindOne())

	if result.Err() != nil {
		log.Println("Error fetching bridge details from MongoDB:", result.Err())
		return types.VMDetails{}, result.Err()
	}

	var document bson.M
	if err := result.Decode(&document); err != nil {
		log.Println("Error decoding bridge details:", err)
		return types.VMDetails{}, err
	}

	bridgeName, _ := document["bridgeName"].(string)
	tapName1, _ := document["tapName1"].(string)
	tapName2, _ := document["tapName2"].(string)
	vm1Eth0IP, _ := document["vm1_eth0_ip"].(string)
	vm2Eth0IP, _ := document["vm2_eth0_ip"].(string)
	macAddress1, _ := document["mac_address1"].(string)
	macAddress2, _ := document["mac_address2"].(string)
	bridgeIPAddress, _ := document["Bridge_ipAddress"].(string)
	bridgeGatewayIPAddress, _ := document["bridge_gateway_ip"].(string)

	return types.VMDetails{
		BridgeName:      bridgeName,
		TapName1:        tapName1,
		TapName2:        tapName2,
		VM1Eth0IP:       vm1Eth0IP,
		VM2Eth0IP:       vm2Eth0IP,
		MacAddress1:     macAddress1,
		MacAddress2:     macAddress2,
		BridgeIPAddress: bridgeIPAddress,
		BridgeGatewayIP: bridgeGatewayIPAddress,
	}, nil
}
