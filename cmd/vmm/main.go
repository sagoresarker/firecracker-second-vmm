package main

import (
	"fmt"
	"os"

	"github.com/sagoresarker/firecracker-second-vmm/internal/database"
	"github.com/sagoresarker/firecracker-second-vmm/internal/runner"
	"github.com/sagoresarker/firecracker-second-vmm/types"
)

func main() {
	fmt.Println("Hello Poridhians!")
	database.InitMongoDB()

	// Check if the userID is provided as a command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <userID>")
		return
	}

	userID := os.Args[1]

	vmDetails, err := database.GetVMDetails(userID)
	if err != nil {
		fmt.Println("Error getting bridge details:", err)
		return
	}

	// Use the retrieved values as needed
	fmt.Println("Bridge Name:", vmDetails.BridgeName)
	fmt.Println("Tap Name 1:", vmDetails.TapName1)
	fmt.Println("Tap Name 2:", vmDetails.TapName2)
	fmt.Println("VM1 Eth0 IP:", vmDetails.VM1Eth0IP)
	fmt.Println("VM2 Eth0 IP:", vmDetails.VM2Eth0IP)
	fmt.Println("MAC Address 1:", vmDetails.MacAddress1)
	fmt.Println("MAC Address 2:", vmDetails.MacAddress2)
	fmt.Println("Bridge IP Address:", vmDetails.BridgeIPAddress)
	fmt.Println("Bridge Gateway IP:", vmDetails.BridgeGatewayIP)
	fmt.Println("User ID:", userID)

	launcher := types.Launcher{
		UserID:          userID,
		BridgeName:      vmDetails.BridgeName,
		TapName2:        vmDetails.TapName2,
		VM2Eth0IP:       vmDetails.VM2Eth0IP,
		MacAddress2:     vmDetails.MacAddress2,
		BridgeIPAddress: vmDetails.BridgeIPAddress,
		BridgeGatewayIP: vmDetails.BridgeGatewayIP,
	}

	runner.LaunchVM(launcher)

}
