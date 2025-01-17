package distribution

import (
	"github.com/QOSGroup/qbase/context"
	btypes "github.com/QOSGroup/qbase/types"
	"github.com/QOSGroup/qos/module/eco/mapper"
	"github.com/QOSGroup/qos/module/eco/types"
	qtypes "github.com/QOSGroup/qos/types"
	"github.com/tendermint/tendermint/crypto"
)

type GenesisState struct {
	CommunityFeePool         btypes.BigInt                 `json:"community_fee_pool"`
	LastBlockProposer        btypes.Address                `json:"last_block_proposer"`
	PreDistributionQOSAmount btypes.BigInt                 `json:"pre_distribute_amount"`
	ValidatorHistoryPeriods  []ValidatorHistoryPeriodState `json:"validators_history_period"`
	ValidatorCurrentPeriods  []ValidatorCurrentPeriodState `json:"validators_current_period"`
	DelegatorEarningInfos    []DelegatorEarningStartState  `json:"delegators_earning_info"`
	DelegatorIncomeHeights   []DelegatorIncomeHeightState  `json:"delegators_income_height"`
	ValidatorEcoFeePools     []ValidatorEcoFeePoolState    `json:"validator_eco_fee_pools"`
	Params                   types.DistributionParams      `json:"params"`
}

func NewGenesisState(communityFeePool btypes.BigInt,
	lastBlockProposer btypes.Address,
	preDistributionQOSAmount btypes.BigInt,
	validatorHistoryPeriods []ValidatorHistoryPeriodState,
	validatorCurrentPeriods []ValidatorCurrentPeriodState,
	delegatorEarningInfos []DelegatorEarningStartState,
	delegatorIncomeHeights []DelegatorIncomeHeightState,
	validatorEcoFeePools []ValidatorEcoFeePoolState,
	params types.DistributionParams) GenesisState {
	return GenesisState{
		CommunityFeePool:         communityFeePool,
		LastBlockProposer:        lastBlockProposer,
		PreDistributionQOSAmount: preDistributionQOSAmount,
		ValidatorHistoryPeriods:  validatorHistoryPeriods,
		ValidatorCurrentPeriods:  validatorCurrentPeriods,
		DelegatorEarningInfos:    delegatorEarningInfos,
		DelegatorIncomeHeights:   delegatorIncomeHeights,
		ValidatorEcoFeePools:     validatorEcoFeePools,
		Params:                   params,
	}
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		CommunityFeePool:         btypes.ZeroInt(),
		PreDistributionQOSAmount: btypes.ZeroInt(),
		Params:                   types.DefaultDistributionParams(),
	}
}

func InitGenesis(ctx context.Context, data GenesisState) {

	distributionMapper := mapper.GetDistributionMapper(ctx)

	feePool := data.CommunityFeePool
	distributionMapper.SetCommunityFeePool(feePool.NilToZero())

	proposer := data.LastBlockProposer
	if !proposer.Empty() {
		distributionMapper.SetLastBlockProposer(proposer)
	}

	distributionMapper.SetPreDistributionQOS(data.PreDistributionQOSAmount.NilToZero())
	distributionMapper.SetParams(ctx, data.Params)

	for _, validatorHistoryPeriodState := range data.ValidatorHistoryPeriods {
		key := types.BuildValidatorHistoryPeriodSummaryKey(btypes.Address(validatorHistoryPeriodState.ValidatorPubKey.Address()), validatorHistoryPeriodState.Period)
		distributionMapper.Set(key, validatorHistoryPeriodState.Summary)
	}

	for _, validatorCurrentPeriodState := range data.ValidatorCurrentPeriods {
		key := types.BuildValidatorCurrentPeriodSummaryKey(btypes.Address(validatorCurrentPeriodState.ValidatorPubKey.Address()))
		distributionMapper.Set(key, validatorCurrentPeriodState.CurrentPeriodSummary)
	}

	for _, delegatorEarningInfoState := range data.DelegatorEarningInfos {
		key := types.BuildDelegatorEarningStartInfoKey(btypes.Address(delegatorEarningInfoState.ValidatorPubKey.Address()), delegatorEarningInfoState.DeleAddress)
		distributionMapper.Set(key, delegatorEarningInfoState.DelegatorEarningsStartInfo)
	}

	for _, delegatorIncomeHeightState := range data.DelegatorIncomeHeights {
		key := types.BuildDelegatorPeriodIncomeKey(btypes.Address(delegatorIncomeHeightState.ValidatorPubKey.Address()), delegatorIncomeHeightState.DeleAddress, delegatorIncomeHeightState.Height)
		distributionMapper.Set(key, true)
	}

	for _, validatorFeePoolState := range data.ValidatorEcoFeePools {
		key := types.BuildValidatorEcoFeePoolKey(validatorFeePoolState.ValidatorAddress)
		distributionMapper.Set(key, validatorFeePoolState.EcoFeePool)
	}

}

func ExportGenesis(ctx context.Context) GenesisState {

	distributionMapper := mapper.GetDistributionMapper(ctx)
	validatorMapper := mapper.GetValidatorMapper(ctx)

	feePool := distributionMapper.GetCommunityFeePool()
	lastBlockProposer := distributionMapper.GetLastBlockProposer()
	preDistributionQOS := distributionMapper.GetPreDistributionQOS()
	params := distributionMapper.GetParams(ctx)

	var validatorHistoryPeriods []ValidatorHistoryPeriodState
	distributionMapper.IteratorValidatorsHistoryPeriod(func(valAddr btypes.Address, period uint64, frac qtypes.Fraction) {

		validator, exsits := validatorMapper.GetValidator(valAddr)
		if exsits {
			vhps := ValidatorHistoryPeriodState{
				ValidatorPubKey: validator.ValidatorPubKey,
				Period:          period,
				Summary:         frac,
			}
			validatorHistoryPeriods = append(validatorHistoryPeriods, vhps)
		}
	})

	var validatorCurrentPeriods []ValidatorCurrentPeriodState
	distributionMapper.IteratorValidatorsCurrentPeriod(func(valAddr btypes.Address, vcps types.ValidatorCurrentPeriodSummary) {

		validator, exsits := validatorMapper.GetValidator(valAddr)
		if exsits {
			vcpsState := ValidatorCurrentPeriodState{
				ValidatorPubKey:      validator.ValidatorPubKey,
				CurrentPeriodSummary: vcps,
			}
			validatorCurrentPeriods = append(validatorCurrentPeriods, vcpsState)
		}
	})

	var delegatorEarningInfos []DelegatorEarningStartState
	distributionMapper.IteratorDelegatorsEarningStartInfo(func(valAddr btypes.Address, deleAddr btypes.Address, desi types.DelegatorEarningsStartInfo) {

		validator, exsits := validatorMapper.GetValidator(valAddr)
		if exsits {
			dess := DelegatorEarningStartState{
				ValidatorPubKey:            validator.ValidatorPubKey,
				DeleAddress:                deleAddr,
				DelegatorEarningsStartInfo: desi,
			}
			delegatorEarningInfos = append(delegatorEarningInfos, dess)
		}
	})

	var delegatorIncomeHeights []DelegatorIncomeHeightState
	distributionMapper.IteratorDelegatorsIncomeHeight(func(valAddr btypes.Address, deleAddr btypes.Address, height uint64) {

		validator, exsits := validatorMapper.GetValidator(valAddr)
		if exsits {
			dihs := DelegatorIncomeHeightState{
				ValidatorPubKey: validator.ValidatorPubKey,
				DeleAddress:     deleAddr,
				Height:          height,
			}
			delegatorIncomeHeights = append(delegatorIncomeHeights, dihs)
		}
	})

	var validatorEcoFeePools []ValidatorEcoFeePoolState
	distributionMapper.IteratorValidatorEcoFeePools(func(validatorAddr btypes.Address, pool types.ValidatorEcoFeePool) {
		validatorEcoFeePools = append(validatorEcoFeePools, ValidatorEcoFeePoolState{
			ValidatorAddress: validatorAddr,
			EcoFeePool:       pool,
		})
	})

	return NewGenesisState(feePool,
		lastBlockProposer,
		preDistributionQOS,
		validatorHistoryPeriods,
		validatorCurrentPeriods,
		delegatorEarningInfos,
		delegatorIncomeHeights,
		validatorEcoFeePools,
		params,
	)
}

type ValidatorHistoryPeriodState struct {
	ValidatorPubKey crypto.PubKey   `json:"validator_pubkey"`
	Period          uint64          `json:"period"`
	Summary         qtypes.Fraction `json:"summary"`
}

type ValidatorCurrentPeriodState struct {
	ValidatorPubKey      crypto.PubKey                       `json:"validator_pub_key"`
	CurrentPeriodSummary types.ValidatorCurrentPeriodSummary `json:"current_period_summary"`
}

type DelegatorEarningStartState struct {
	ValidatorPubKey            crypto.PubKey                    `json:"validator_pub_key"`
	DeleAddress                btypes.Address                   `json:"delegator_address"`
	DelegatorEarningsStartInfo types.DelegatorEarningsStartInfo `json:"earning_start_info"`
}

type DelegatorIncomeHeightState struct {
	ValidatorPubKey crypto.PubKey  `json:"validator_pub_key"`
	DeleAddress     btypes.Address `json:"delegator_address"`
	Height          uint64         `json:"height"`
}

type ValidatorEcoFeePoolState struct {
	ValidatorAddress btypes.Address            `json:"validator_address"`
	EcoFeePool       types.ValidatorEcoFeePool `json:"eco_fee_pool"`
}
