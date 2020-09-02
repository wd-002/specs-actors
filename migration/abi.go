package migration

import (
	big1 "github.com/filecoin-project/specs-actors/actors/abi/big"
	smoothing1 "github.com/filecoin-project/specs-actors/actors/util/smoothing"

	"github.com/filecoin-project/specs-actors/v2/actors/abi/big"
	"github.com/filecoin-project/specs-actors/v2/actors/util/smoothing"
)

func bint(i big1.Int) big.Int {
	return big.NewFromBig(*i.Int)
}

func filter(f *smoothing1.FilterEstimate) *smoothing.FilterEstimate {
	return &smoothing.FilterEstimate{
		PositionEstimate: bint(f.PositionEstimate),
		VelocityEstimate: bint(f.VelocityEstimate),
	}
}
