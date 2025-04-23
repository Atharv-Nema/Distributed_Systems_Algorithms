package ExampleAlgo

import (
	D "distributed_systems/distributed_systems_simulator"
)

type DummyAlgoNode struct {
	// Fill in params here
}

func (d DummyAlgoNode) ReceiveMessage( /*takes in a node message of a predefined form*/ ) {
	// I think that the connection will be opened for the object in the main function(probably by having an)
}

type DummyMessage struct {
	key   int
	value int
}

func main() {
	// Takes in a few parameters from the command line. Creates a DummyAlgoNode object and loads the persistent storage. Same as the start node function. Perhaps some form of composition to hide the hidden TCP connection will be needed
	// something like DummyAlgoNode.StartServer() that handles all of the server
	// creation stuff
	// The startserver will implicitly call HandleRequests in the listener. Probably need a not implemented error for the interface
	// The information about the other port needs to be hidden from the user. Perhaps, handlerequest needs to be wrapped up
	comm_handler := D.CreateCommunicationHandler[DummyMessage]()
	// now that all the CLI arguments have been handled, we can just load the persistent storage.
}
