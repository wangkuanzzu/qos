package stake

import (
	"fmt"

	"github.com/QOSGroup/qos/module/eco"

	"github.com/QOSGroup/qbase/context"
	btypes "github.com/QOSGroup/qbase/types"
	"github.com/QOSGroup/qos/module/eco/mapper"
	ecotypes "github.com/QOSGroup/qos/module/eco/types"
	"github.com/QOSGroup/qos/types"
	"github.com/tendermint/tendermint/crypto"
)

type GenesisState struct {
	Params                 ecotypes.StakeParams             `json:"params"`
	Validators             []ecotypes.Validator             `json:"validators"`            //validatorKey, validatorByOwnerKey,validatorByInactiveKey,validatorByVotePowerKey
	ValidatorsVoteInfo     []ValidatorVoteInfoState         `json:"val_votes_info"`        //validatorVoteInfoKey
	ValidatorsVoteInWindow []ValidatorVoteInWindowInfoState `json:"val_votes_in_window"`   //validatorVoteInfoInWindowKey
	DelegatorsInfo         []DelegationInfoState            `json:"delegators_info"`       //DelegationByDelValKey, DelegationByValDelKey
	DelegatorsUnbondInfo   []DelegatorUnbondState           `json:"delegator_unbond_info"` //DelegatorUnbondingQOSatHeightKey
	CurrentValidators      []ecotypes.Validator             `json:"current_validators"`    // currentValidatorsAddressKey
}

func NewGenesisState(params ecotypes.StakeParams,
	validators []ecotypes.Validator,
	validatorsVoteInfo []ValidatorVoteInfoState,
	validatorsVoteInWindow []ValidatorVoteInWindowInfoState,
	delegatorsInfo []DelegationInfoState,
	delegatorsUnbondInfo []DelegatorUnbondState,
	currentValidators []ecotypes.Validator) GenesisState {
	return GenesisState{
		Params:                 params,
		Validators:             validators,
		ValidatorsVoteInfo:     validatorsVoteInfo,
		ValidatorsVoteInWindow: validatorsVoteInWindow,
		DelegatorsInfo:         delegatorsInfo,
		DelegatorsUnbondInfo:   delegatorsUnbondInfo,
		CurrentValidators:      currentValidators,
	}
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		Params: ecotypes.DefaultStakeParams(),
	}
}

func InitGenesis(ctx context.Context, data GenesisState) {
	validatorMapper := mapper.GetValidatorMapper(ctx)

	if len(data.CurrentValidators) > 0 {
		validatorMapper.Set(ecotypes.BuildCurrentValidatorsAddressKey(), data.CurrentValidators)
	}

	initValidators(ctx, data.Validators)
	initParams(ctx, data.Params)
	initValidatorsVotesInfo(ctx, data.ValidatorsVoteInfo, data.ValidatorsVoteInWindow)
	initDelegatorsInfo(ctx, data.DelegatorsInfo, data.DelegatorsUnbondInfo)
}

func initValidators(ctx context.Context, validators []ecotypes.Validator) {
	validatorMapper := mapper.GetValidatorMapper(ctx)
	for _, v := range validators {

		if validatorMapper.Exists(v.ValidatorPubKey.Address().Bytes()) {
			panic(fmt.Errorf("validator %s already exists", v.ValidatorPubKey.Address()))
		}
		if validatorMapper.ExistsWithOwner(v.Owner) {
			panic(fmt.Errorf("owner %s already bind a validator", v.Owner))
		}
		validatorMapper.CreateValidator(v)
		if !v.IsActive() {
			validatorMapper.MakeValidatorInactive(v.GetValidatorAddress(), v.InactiveHeight, v.InactiveTime, v.InactiveCode)
		}
	}
}

func initValidatorsVotesInfo(ctx context.Context, voteInfos []ValidatorVoteInfoState, voteWindowInfos []ValidatorVoteInWindowInfoState) {
	voteMapper := mapper.GetVoteInfoMapper(ctx)
	for _, voteInfo := range voteInfos {
		voteMapper.SetValidatorVoteInfo(btypes.Address(voteInfo.ValidatorPubKey.Address()), voteInfo.VoteInfo)
	}

	for _, voteWindowInfo := range voteWindowInfos {
		voteMapper.SetVoteInfoInWindow(btypes.Address(voteWindowInfo.ValidatorPubKey.Address()), voteWindowInfo.Index, voteWindowInfo.Vote)
	}
}

func initDelegatorsInfo(ctx context.Context, delegatorsInfo []DelegationInfoState, delegatorsUnbondInfo []DelegatorUnbondState) {
	delegationMapper := mapper.GetDelegationMapper(ctx)

	for _, info := range delegatorsInfo {
		delegationMapper.SetDelegationInfo(ecotypes.DelegationInfo{
			DelegatorAddr: info.DelegatorAddr,
			ValidatorAddr: btypes.Address(info.ValidatorPubKey.Address()),
			Amount:        info.Amount,
			IsCompound:    info.IsCompound,
		})
	}

	for _, info := range delegatorsUnbondInfo {
		delegationMapper.SetDelegatorUnbondingQOSatHeight(info.Height, info.DeleAddress, info.Amount)
	}
}

func initParams(ctx context.Context, params ecotypes.StakeParams) {
	mapper := ctx.Mapper(ecotypes.ValidatorMapperName).(*mapper.ValidatorMapper)
	mapper.SetParams(ctx, params)
}

func ValidateGenesis(genesisAccounts []*types.QOSAccount, data GenesisState) error {
	err := validateValidators(genesisAccounts, data.Validators)
	if err != nil {
		return err
	}

	return nil
}

func validateValidators(genesisAccounts []*types.QOSAccount, validators []ecotypes.Validator) (err error) {
	addrMap := make(map[string]bool, len(validators))
	for i := 0; i < len(validators); i++ {
		val := validators[i]
		strKey := string(val.ValidatorPubKey.Bytes())
		if _, ok := addrMap[strKey]; ok {
			return fmt.Errorf("duplicate validator in genesis state: Name %v, Owner %v", val.Description.Moniker, val.Owner)
		}
		if val.Status != ecotypes.Active {
			return fmt.Errorf("validator is bonded and jailed in genesis state: Name %v, Owner %v", val.Description.Moniker, val.Owner)
		}
		addrMap[strKey] = true

		var ownerExists bool
		for _, acc := range genesisAccounts {
			if acc.AccountAddress.EqualsTo(val.Owner) {
				ownerExists = true
			}
		}

		if !ownerExists {
			return fmt.Errorf("owner of %s not exists", val.Description.Moniker)
		}
	}
	return nil
}

func ExportGenesis(ctx context.Context) GenesisState {

	validatorMapper := mapper.GetValidatorMapper(ctx)
	voteMapper := mapper.GetVoteInfoMapper(ctx)
	delegationMapper := mapper.GetDelegationMapper(ctx)

	var currentValidators []ecotypes.Validator
	validatorMapper.Get(ecotypes.BuildCurrentValidatorsAddressKey(), &currentValidators)

	params := validatorMapper.GetParams(ctx)

	var validators []ecotypes.Validator
	validatorMapper.IterateValidators(func(validator ecotypes.Validator) {
		validators = append(validators, validator)
	})

	var validatorsVoteInfo []ValidatorVoteInfoState
	voteMapper.IterateVoteInfos(func(valAddr btypes.Address, info ecotypes.ValidatorVoteInfo) {

		validator, exsits := validatorMapper.GetValidator(valAddr)
		if exsits {
			vvis := ValidatorVoteInfoState{
				ValidatorPubKey: validator.ValidatorPubKey,
				VoteInfo:        info,
			}
			validatorsVoteInfo = append(validatorsVoteInfo, vvis)
		}
	})

	var validatorsVoteInWindow []ValidatorVoteInWindowInfoState
	voteMapper.IterateVoteInWindowsInfos(func(index uint64, valAddr btypes.Address, vote bool) {

		validator, exsits := validatorMapper.GetValidator(valAddr)
		if exsits {
			validatorsVoteInWindow = append(validatorsVoteInWindow, ValidatorVoteInWindowInfoState{
				ValidatorPubKey: validator.ValidatorPubKey,
				Index:           index,
				Vote:            vote,
			})
		}
	})

	var delegatorsInfo []DelegationInfoState
	delegationMapper.IterateDelegationsInfo(btypes.Address{}, func(info ecotypes.DelegationInfo) {

		validator, exists := validatorMapper.GetValidator(info.ValidatorAddr)
		if !exists {
			panic(fmt.Sprintf("validator:%s not exsits", info.ValidatorAddr.String()))
		}

		delegatorsInfo = append(delegatorsInfo, DelegationInfoState{
			DelegatorAddr:   info.DelegatorAddr,
			ValidatorPubKey: validator.ValidatorPubKey,
			Amount:          info.Amount,
			IsCompound:      info.IsCompound,
		})
	})

	var delegatorsUnbondInfo []DelegatorUnbondState
	delegationMapper.IterateDelegationsUnbondInfo(func(deleAddr btypes.Address, height uint64, amount uint64) {
		delegatorsUnbondInfo = append(delegatorsUnbondInfo, DelegatorUnbondState{
			DeleAddress: deleAddr,
			Height:      height,
			Amount:      amount,
		})
	})

	return GenesisState{
		Params:                 params,
		Validators:             validators,
		ValidatorsVoteInfo:     validatorsVoteInfo,
		ValidatorsVoteInWindow: validatorsVoteInWindow,
		DelegatorsInfo:         delegatorsInfo,
		DelegatorsUnbondInfo:   delegatorsUnbondInfo,
		CurrentValidators:      currentValidators,
	}
}

type ValidatorVoteInfoState struct {
	ValidatorPubKey crypto.PubKey              `json:"validator_pub_key"`
	VoteInfo        ecotypes.ValidatorVoteInfo `json:"vote_info"`
}

type ValidatorVoteInWindowInfoState struct {
	ValidatorPubKey crypto.PubKey `json:"validator_pub_key"`
	Index           uint64        `json:"index"`
	Vote            bool          `json:"vote"`
}

type DelegationInfoState struct {
	DelegatorAddr   btypes.Address `json:"delegator_addr"`
	ValidatorPubKey crypto.PubKey  `json:"validator_pub_key"`
	Amount          uint64         `json:"delegate_amount"`
	IsCompound      bool           `json:"is_compound"`
}

type DelegatorUnbondState struct {
	DeleAddress btypes.Address `json:"delegator_address"`
	Height      uint64         `json:"height"`
	Amount      uint64         `json:"tokens"`
}

func PrepForZeroHeightGenesis(ctx context.Context) {

	e := eco.GetEco(ctx)

	// close all active validators
	var validators []ecotypes.Validator
	var delegations []ecotypes.DelegationInfo
	var vals = make(map[string]ecotypes.Validator)
	e.ValidatorMapper.IterateValidators(func(validator ecotypes.Validator) {
		val := validator.GetValidatorAddress()
		vals[validator.GetValidatorAddress().String()] = validator
		e.DelegationMapper.IterateDelegationsValDeleAddr(val, func(val btypes.Address, del btypes.Address) {
			delegation, exists := e.DelegationMapper.GetDelegationInfo(del, val)
			if !exists {
				panic(fmt.Sprintf("delegation from %s to %s should exists", del, val))
			}

			delegations = append(delegations, delegation)
		})

		if validator.Status == ecotypes.Active {
			e.ValidatorMapper.MakeValidatorInactive(val, uint64(ctx.BlockHeight()), ctx.BlockHeader().Time.UTC(), ecotypes.Revoke)
		}
	})

	// close all inactive validators
	closeAllInactiveValidator(ctx)

	for _, delegation := range delegations {
		var info ecotypes.DelegatorEarningsStartInfo
		e.DistributionMapper.Get(ecotypes.BuildDelegatorEarningStartInfoKey(delegation.ValidatorAddr, delegation.DelegatorAddr), &info)
		eco.BonusToDelegator(e.Context, delegation.DelegatorAddr, delegation.ValidatorAddr, info.HistoricalRewardFees, false)
		// eco.IncrAccountQOS(e.Context, delegation.DelegatorAddr, info.HistoricalRewardFees)
	}

	// return unbond tokens
	eco.ReturnAllUnbondTokens(ctx)

	// reinitialize all validators
	for _, validator := range validators {
		val := validator.GetValidatorAddress()
		e.VoteInfoMapper.DelValidatorVoteInfo(val)
		e.VoteInfoMapper.ClearValidatorVoteInfoInWindow(val)
		e.DistributionMapper.DeleteValidatorPeriodSummaryInfo(val)
		e.DistributionMapper.InitValidatorPeriodSummaryInfo(val)
		e.ValidatorMapper.CreateValidator(validator)
	}

	// reinitialize all delegations
	e.DistributionMapper.DeleteDelegatorsIncomeHeight()
	// reset block height
	ctx = ctx.WithBlockHeight(0)
	e = eco.GetEco(ctx)
	for _, delegation := range delegations {
		e.DistributionMapper.DelDelegatorEarningStartInfo(delegation.ValidatorAddr, delegation.DelegatorAddr)
		e.DelegationMapper.DelDelegationInfo(delegation.ValidatorAddr, delegation.DelegatorAddr)
		e.DelegateValidator(ctx, vals[delegation.ValidatorAddr.String()], delegation.DelegatorAddr, delegation.Amount, delegation.IsCompound, true)
	}
}
