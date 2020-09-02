package migration

import (
	reward1 "github.com/filecoin-project/specs-actors/actors/builtin/reward"

	"github.com/filecoin-project/specs-actors/v2/actors/abi"
	"github.com/filecoin-project/specs-actors/v2/actors/builtin/reward"
)

func MigrateReward(state reward1.State) (reward.State, error) {
	return reward.State{
		CumsumBaseline:          bint(state.CumsumBaseline),
		CumsumRealized:          bint(state.CumsumRealized),
		EffectiveNetworkTime:    abi.ChainEpoch(state.EffectiveNetworkTime),
		EffectiveBaselinePower:  bint(state.EffectiveBaselinePower),
		ThisEpochReward:         bint(state.ThisEpochReward),
		ThisEpochRewardSmoothed: filter(state.ThisEpochRewardSmoothed),
		ThisEpochBaselinePower:  bint(state.ThisEpochBaselinePower),
		Epoch:                   abi.ChainEpoch(state.Epoch),
		TotalMined:              bint(state.TotalMined),
	}, nil
}

