package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetVMDetails(userID string) (string, string, string, string, string, string, string, string, string, string, error) {
	if mongoClient == nil {
		log.Fatal("MongoDB client not initialized.")
		return "", "", "", "", "", "", "", "", "", "", nil
	}

	collection := mongoClient.Database("firecrackerdb").Collection("vm-info")

	filter := bson.M{"userID": userID}
	result := collection.FindOne(context.Background(), filter, options.FindOne())

	if result.Err() != nil {
		log.Println("Error fetching bridge details from MongoDB:", result.Err())
		return "", "", "", "", "", "", "", "", "", "", result.Err()
	}

	var document bson.M
	if err := result.Decode(&document); err != nil {
		log.Println("Error decoding bridge details:", err)
		return "", "", "", "", "", "", "", "", "", "", err
	}

	bridgeName := document["bridgeName"].(string)
	tapName1 := document["tapName1"].(string)
	tapName2 := document["tapName2"].(string)
	vm1Eth0IP := document["vm1_eth0_ip"].(string)
	vm2Eth0IP := document["vm2_eth0_ip"].(string)
	macAddress1 := document["mac_address1"].(string)
	macAddress2 := document["mac_address2"].(string)
	bridgeIPAddress := document["Bridge_ipAddress"].(string)
	bridgeGatewayIPAddress := document["Bridge_gateway_ip"].(string)

	return bridgeName, tapName1, tapName2, vm1Eth0IP, vm2Eth0IP, macAddress1, macAddress2, bridgeIPAddress, bridgeGatewayIPAddress, userID, nil
}
