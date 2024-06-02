package groth16

import (
	"example.com/m/types"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

func Groth16_run() {
	// x**3 +x +5 = y
	// compiles our circuit into a R1CS
	var mycircuit types.Groth16_Circuit
	ccs, _ := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &mycircuit)

	// zk-snark setup
	pk, vk, _ := groth16.Setup(ccs)

	// witness definition
	assignment := types.Groth16_Circuit{X: 3, Y: 35}
	witness, _ := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
	publicWitness, _ := witness.Public()

	// groth16: Prove & Verify
	proof, _ := groth16.Prove(ccs, pk, witness)
	_ = groth16.Verify(proof, vk, publicWitness)
}
