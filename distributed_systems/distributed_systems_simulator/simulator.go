package DistributedSimulator

import (
	"os/exec"
)

type NodeInfo struct {
	filePath string // I would ideally want more protection.
	// For example, the path of the file to run would be better. More security
	// can be added later
	NodeCrashRate    float32 // Rate at which a node crashes
	MessageDropRate  float32
	NodeRecoveryRate float32 // Rate at which a node recovers
}

type SimulationInfo []NodeInfo

type nodeProcessInfo struct {
	portNumber string
	alive      bool
	command    *exec.Cmd
}

func (info SimulationInfo) setupProcesses() *map[int]nodeProcessInfo {
	// Sets up the process. I guess we can give it info about its own port.
	// I do not think that main should be handling this. We put this functionality
	// into another function called at the start of main. After this initial thing

	// Start the simulator listener(at a predefined port)

	// Start the node process
	nodeInfoMap := new(map[int]nodeProcessInfo)
	for nodeNumber, nodeInfo := range info {
		// Calculate the port number TODO

		// Fill out the command correctly TODO
		cmd := exec.Command("go", "run", nodeInfo.filePath, "node_number=", "--server_port=", "--node_port=")
		processInfo := nodeProcessInfo{
			portNumber: "TODO: Add proper port number",
			command:    cmd,
		}
		(*nodeInfoMap)[nodeNumber] = processInfo

		// Some datastructure/goroutine for wakeup/kill calls

	}

	return nodeInfoMap
}

func (info SimulationInfo) RunSimulation() {
	// Runs the simulation based on SimulationInfo
	// Cannot just pass in a function I want to start independent processes
	// Perhaps we need to have a path to a file that implements starting and stopping?
	// Or perhaps some sort of dynamic import
	// Need to add a tcp event listener(messages pass via the simulator).

	nodeInfoMap := info.setupProcesses() // All processes have been setup

}
