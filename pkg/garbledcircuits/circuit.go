package garbledcircuit

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
