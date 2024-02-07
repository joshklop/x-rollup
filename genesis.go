package rollup

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joshklop/x-rollup/keeper"
	"github.com/joshklop/x-rollup/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k *keeper.Keeper, gs types.GenesisState) {

}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k *keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{}
}
