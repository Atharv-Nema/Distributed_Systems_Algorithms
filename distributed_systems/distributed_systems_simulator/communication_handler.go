package DistributedSimulator

// type Node struct {
// 	id int
// 	totalNodes int // Perhaps we should not have this. There could be a discovery phase
// 	static_storage_file_path string// Nodes can handle it any way they want
// }
// Perhaps a communication handler is all we need(node is not needed)

type message[T any] struct {
	size    int // The total size of the message
	from    int
	to      int
	payload T
}

type CommunicationHandler[T any] struct {
	// This object handles the communication between the different nodes
	// T is the message type. So, we can send any message over tcp.
	// Note that this communication handler sends this message to the simulator,
	// so perhaps we do not need the ports of all the thingies
	NodeNumber int
	// Perhaps a node port number
	// Perhaps a connection object. Details can be decided later
	nodePort       string
	simulatorPort  string                          // It may not be string
	messageHandler func(nodeNumber int, payload T) // Unpacks the message and calls the messageHandler
}

func (c CommunicationHandler[T]) SendMessage(receiverNode int, payload T) {
	// 1. Prepare the message
	// 2. Send it to the server via the connection
}

func (c CommunicationHandler[T]) RestartConnection( /*again, probably need a different datatype*/ ) {
	// Starts a tcp server at a particular port number
	// TCP is a two way connection. Perhaps we do not need this?
}

func CreateCommunicationHandler[T any]() CommunicationHandler[T] {
	// Uses the command line arguments to create the communication handler
}

// We will need a function called create communication handler. Note that
// this is independent of the specific node type

func main() {
	// Create a node using the CLI arguments(which have the node number)??
}
