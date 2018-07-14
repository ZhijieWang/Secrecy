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

// referencing implmentation from https://github.com/JoelOtter/go-garbled/blob/master/circuit.go

//Circuit struct for
type Circuit struct {
	Name      string            // canonical name for the circuit
	Inputs    map[string]*Input // a map of input names to input gates
	Outputs   map[string]*Wire  // a map of output names to output gates
	Encryptor CryptoFunc        // the function used for encryption, takes a number and a key
	Decryptor CryptoFunc        // the function used for decryption, takes the ciphertext and a key
}

// NewCircuit is the constructor function, takes string argument for circuit name
// returns the pointer to the newly creater circuit struct
// this uses a basic onetime XOR pad for both encryption and decryption
func NewCircuit(name string) *Circuit {
	c :=
		Circuit{
			Name:    name,
			Inputs:  make(map[string]*Input),
			Outputs: make(map[string]*Wire),
			Encryptor: func(m, k uint32) uint32 {
				return m ^ k
			},
			Decryptor: func(c, k uint32) uint32 {
				return c ^ k
			},
		}
	return &c
}

// AddInput registers a new Input in the Circuit  with the provided name and value.
// returns a pointer to the Input's output wire.
func (c *Circuit) AddInput(name string) *Input {
	i := new(Input)
	i.Output = NewWire(i)
	i.circuit = c
	c.Inputs[name] = i
	return i
}

// AddOutput registers a ew new Output in Output in the circuit
// with the provided name

func (c *Circuit) addOutput(name string, g Gate) {
	c.Outputs[name] = g.GetOutput()
}

// Evaluate wille valuate a whole circuite for the inputs specifiged
// in the map 'inputs'. Returns a map of outputs to their values.NewCircuit
//
// E.g. for a circuit containing a single AND gate with
// inputs A and B, and one output 0, the map:
// {"A" :0, "B":1}
// will evaluate to :
// {'0":0'}

func (c *Circuit) Evaluate(inputs map[string]uint32) map[string]uint32 {
	for k, v := range inpputs {
		c.Inputs[k].Value = v
	}
	outputs := make(map[string]uinst32)
	for k, v := range c.Outputs {
		_, outputP := v.Evaluate()
		outputs[k] = outputP ^ v.P
	}
	return outputs
}
