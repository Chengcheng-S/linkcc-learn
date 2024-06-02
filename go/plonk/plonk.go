package plonk

import (
	"log"

	"example.com/m/types"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/plonk"
	cs "github.com/consensys/gnark/constraint/bn254"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/scs"
	"github.com/consensys/gnark/test/unsafekzg"
)

func Plonk_run() {

	var mycircuit types.Plonk_Circuit

	ccs, _ := frontend.Compile(ecc.BN254.ScalarField(), scs.NewBuilder, &mycircuit)

	//create the necessary data for KZG.
	scs := ccs.(*cs.SparseR1CS)
	srs, srsLagrange, err := unsafekzg.NewSRS(scs)
	if err != nil {
		panic(err)
	}

	{
		// Witnesses instantiation. Witness is known only by the prover,
		// while public w is a public data known by the verifier.
		var w types.Plonk_Circuit
		w.X = 2
		w.E = 2
		w.Y = 4

		witnessFull, err := frontend.NewWitness(&w, ecc.BN254.ScalarField())
		if err != nil {
			log.Fatal(err)
		}

		witnessPublic, err := frontend.NewWitness(&w, ecc.BN254.ScalarField(), frontend.PublicOnly())
		if err != nil {
			log.Fatal(err)
		}

		// public data consists of the polynomials describing the constants involved
		// in the constraints, the polynomial describing the permutation ("grand
		// product argument"), and the FFT domains.
		pk, vk, err := plonk.Setup(ccs, srs, srsLagrange)
		//_, err := plonk.Setup(r1cs, kate, &publicWitness)
		if err != nil {
			log.Fatal(err)
		}

		proof, err := plonk.Prove(ccs, pk, witnessFull)
		if err != nil {
			log.Fatal(err)
		}

		err = plonk.Verify(proof, vk, witnessPublic)
		if err != nil {
			log.Fatal(err)
		}
	}
	// Wrong data: the proof fails
	{
		// Witnesses instantiation. Witness is known only by the prover,
		// while public w is a public data known by the verifier.
		var w, pW types.Plonk_Circuit
		w.X = 2
		w.E = 12
		w.Y = 4096

		pW.X = 3
		pW.Y = 4096

		witnessFull, err := frontend.NewWitness(&w, ecc.BN254.ScalarField())
		if err != nil {
			log.Fatal(err)
		}

		witnessPublic, err := frontend.NewWitness(&pW, ecc.BN254.ScalarField(), frontend.PublicOnly())
		if err != nil {
			log.Fatal(err)
		}

		// public data consists of the polynomials describing the constants involved
		// in the constraints, the polynomial describing the permutation ("grand
		// product argument"), and the FFT domains.
		pk, vk, err := plonk.Setup(ccs, srs, srsLagrange)
		//_, err := plonk.Setup(r1cs, kate, &publicWitness)
		if err != nil {
			log.Fatal(err)
		}

		proof, err := plonk.Prove(ccs, pk, witnessFull)
		if err != nil {
			log.Fatal(err)
		}

		err = plonk.Verify(proof, vk, witnessPublic)
		if err == nil {
			log.Fatal(err)
		}
	}
}
