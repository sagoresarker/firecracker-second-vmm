package main

import (
	"fmt"
	"os"

	vm "github.com/sagoresarker/firecracker-second-vmm/container"
	"github.com/sagoresarker/firecracker-second-vmm/database"
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

	bridgeName, tapName1, tapName2, vm1Eth0IP, vm2Eth0IP, macAddress1, macAddress2, bridgeIPAddress, Bridge_gateway_ip, _, err := database.GetVMDetails(userID)
	if err != nil {
		fmt.Println("Error getting bridge details:", err)
		return
	}

	// Use the retrieved values as needed
	fmt.Println("Bridge Name:", bridgeName)
	fmt.Println("Tap Name 1:", tapName1)
	fmt.Println("Tap Name 2:", tapName2)
	fmt.Println("VM1 Eth0 IP:", vm1Eth0IP)
	fmt.Println("VM2 Eth0 IP:", vm2Eth0IP)
	fmt.Println("MAC Address 1:", macAddress1)
	fmt.Println("MAC Address 2:", macAddress2)
	fmt.Println("Bridge IP Address:", bridgeIPAddress)
	fmt.Println("Bridge Gateway IP:", Bridge_gateway_ip)
	fmt.Println("User ID:", userID)

	vm.LaunchSecondInstance(userID, bridgeName, tapName2, vm2Eth0IP, macAddress2, bridgeIPAddress, Bridge_gateway_ip)
}
