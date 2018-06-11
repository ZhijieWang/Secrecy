package garbledcircuit

// A CryptoFunc is any function that takes a value and a key and
// returns an encrypted or decrypted value.
type CryptoFunc func(uint32, uint32) uint32

// A Wire represents the wire between two gates.
// currently assuming 2 parties, hosting 2 keys
type Wire struct {
	Input  Gate      // the input gate
	Output Gate      // the output gate
	Keys   [2]uint32 // the keys:one for 0, one for 1P
	uint32           // a randomised p-value
}

// Input 'gate', used to supply inputs to the circuit.
type Input struct {
	Value   uint32
	circuit *Circuit // a pointer back to the Circuit
	Output  *Wire    // the output wire
}

// A Gate represents any binary or unary gate.
type Gate interface {
	Evaluate() (uint32, uint32)
	Circuit() *Circuit
	GetOutput() *Wire
}
