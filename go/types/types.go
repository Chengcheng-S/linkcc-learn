package types

import (
	"github.com/consensys/gnark/frontend"
)

// x**3 +x +5 = y
type Groth16_Circuit struct {
	X frontend.Variable `gnark:"x"`
	Y frontend.Variable `gnark:",public"`
}

func (m *Groth16_Circuit) Define(api frontend.API) error {
	// x**3 +x +5 = y
	x3 := api.Mul(m.X, m.X, m.X)
	x3 = api.Add(x3, m.X, 5)
	api.AssertIsEqual(m.Y, x3)
	return nil
}

// Y == X**E
type Plonk_Circuit struct {
	X frontend.Variable `gnark:",public"`
	Y frontend.Variable `gnark:",public"`
	E frontend.Variable
}

func (pc *Plonk_Circuit) Define(api frontend.API) error {
	// number of bits of exponent
	const bitSize = 4000

	output := frontend.Variable(1)
	bits := api.ToBinary(pc.E, bitSize)

	for i := 0; i < len(bits); i++ {
		if i != 0 {
			output = api.Mul(output, output)
		}
		multiply := api.Mul(output, pc.X)
		output = api.Select(bits[len(bits)-1-i], multiply, output)
	}

	api.AssertIsEqual(pc.Y, output)
	return nil
}
