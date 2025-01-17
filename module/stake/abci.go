package stake

import (
	"github.com/QOSGroup/qbase/context"
	btypes "github.com/QOSGroup/qbase/types"
	"github.com/QOSGroup/qos/module/eco"
	ecomapper "github.com/QOSGroup/qos/module/eco/mapper"
	ecotypes "github.com/QOSGroup/qos/module/eco/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

//1. 统计validator投票信息, 将不活跃的validator转成Inactive状态
func BeginBlocker(ctx context.Context, req abci.RequestBeginBlock) {

	validatorMapper := ecomapper.GetValidatorMapper(ctx)

	votingWindowLen := uint64(validatorMapper.GetParams(ctx).ValidatorVotingStatusLen)
	minVotingCounter := uint64(validatorMapper.GetParams(ctx).ValidatorVotingStatusLeast)

	for _, signingValidator := range req.LastCommitInfo.Votes {
		valAddr := btypes.Address(signingValidator.Validator.Address)
		voted := signingValidator.SignedLastBlock
		handleValidatorValidatorVoteInfo(ctx, valAddr, voted, votingWindowLen, minVotingCounter)
	}
}

//1. 将所有Inactive到一定期限的validator删除
//2. 统计新的validator
func EndBlocker(ctx context.Context) (res abci.ResponseEndBlock) {

	validatorMapper := ecomapper.GetValidatorMapper(ctx)
	survivalSecs := validatorMapper.GetParams(ctx).ValidatorSurvivalSecs
	maxValidatorCount := uint64(validatorMapper.GetParams(ctx).MaxValidatorCnt)

	closeExpireInactiveValidator(ctx, survivalSecs)
	res.ValidatorUpdates = GetUpdatedValidators(ctx, maxValidatorCount)
	return
}

//unbond的token返还至delegator账户中
func EndBlockerByReturnUnbondTokens(ctx context.Context) {
	height := uint64(ctx.BlockHeight())
	e := eco.GetEco(ctx)
	log := e.Context.Logger()

	prePrefix := ecotypes.BuildUnbondingDelegationByHeightPrefix(height)

	iter := btypes.KVStorePrefixIterator(e.DelegationMapper.GetStore(), prePrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		k := iter.Key()
		e.DelegationMapper.Del(k)

		var amount uint64
		e.DelegationMapper.BaseMapper.DecodeObject(iter.Value(), &amount)

		_, deleAddr := ecotypes.GetUnbondingDelegationHeightAddress(k)
		returnQOSAmount := amount
		log.Debug("stake end return unbond token", "height", height, "deleAddr", deleAddr, "tokens", returnQOSAmount)
		eco.IncrAccountQOS(ctx, deleAddr, btypes.NewInt(int64(returnQOSAmount)))
	}
}

func closeExpireInactiveValidator(ctx context.Context, survivalSecs uint32) {
	log := ctx.Logger()
	e := eco.GetEco(ctx)

	blockTimeSec := uint64(ctx.BlockHeader().Time.UTC().Unix())
	lastCloseValidatorSec := blockTimeSec - uint64(survivalSecs)

	iterator := e.ValidatorMapper.IteratorInactiveValidator(uint64(0), lastCloseValidatorSec)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		valAddress := btypes.Address(key[9:])
		log.Info("close validator", "height", ctx.BlockHeight(), "validator", valAddress.String())
		e.RemoveValidator(ctx, valAddress)
	}
	iterator.Close()
}

func closeAllInactiveValidator(ctx context.Context) {
	log := ctx.Logger()
	e := eco.GetEco(ctx)

	lastCloseValidatorSec := uint64(ctx.BlockHeader().Time.UTC().Unix()) + 1

	iterator := e.ValidatorMapper.IteratorInactiveValidator(uint64(0), lastCloseValidatorSec)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		valAddress := btypes.Address(key[9:])
		log.Info("close validator", "height", ctx.BlockHeight(), "validator", valAddress.String())
		e.RemoveValidator(ctx, valAddress)
	}
	iterator.Close()
}

func GetUpdatedValidators(ctx context.Context, maxValidatorCount uint64) []abci.ValidatorUpdate {
	log := ctx.Logger()
	validatorMapper := ctx.Mapper(ecotypes.ValidatorMapperName).(*ecomapper.ValidatorMapper)

	//获取当前的validator集合
	var currentValidators []ecotypes.Validator
	validatorMapper.Get(ecotypes.BuildCurrentValidatorsAddressKey(), &currentValidators)

	currentValidatorMap := make(map[string]ecotypes.Validator)
	for _, curValidator := range currentValidators {
		curValidatorAddrString := curValidator.GetValidatorAddress().String()
		currentValidatorMap[curValidatorAddrString] = curValidator
	}

	//返回更新的validator
	updateValidators := make([]abci.ValidatorUpdate, 0, len(currentValidatorMap))

	i := uint64(0)
	newValidatorsMap := make(map[string]ecotypes.Validator)
	newValidators := make([]ecotypes.Validator, 0, len(currentValidators))

	iterator := validatorMapper.IteratorValidatorByVoterPower(false)
	defer iterator.Close()

	var key []byte
	for ; iterator.Valid(); iterator.Next() {
		key = iterator.Key()
		valAddr := btypes.Address(key[9:])

		if i >= maxValidatorCount {
			//超出MaxValidatorCnt的validator修改为Inactive状态
			if validator, exsits := validatorMapper.GetValidator(valAddr); exsits {
				validatorMapper.MakeValidatorInactive(validator.GetValidatorAddress(), uint64(ctx.BlockHeight()), ctx.BlockHeader().Time.UTC(), ecotypes.MaxValidator)
			}
		} else {
			if validator, exsits := validatorMapper.GetValidator(valAddr); exsits {
				if !validator.IsActive() {
					continue
				}
				i++
				//保存数据
				newValidatorAddressString := validator.GetValidatorAddress().String()
				newValidatorsMap[newValidatorAddressString] = validator
				newValidators = append(newValidators, validator)

				//新增或修改
				curValidator, exsits := currentValidatorMap[newValidatorAddressString]
				if !exsits || (validator.BondTokens != curValidator.BondTokens) {
					updateValidators = append(updateValidators, validator.ToABCIValidatorUpdate(false))
				}
			}
		}
	}

	//删除
	for curValidatorAddr, curValidator := range currentValidatorMap {
		if _, ok := newValidatorsMap[curValidatorAddr]; !ok {
			// curValidator.Power = 0
			updateValidators = append(updateValidators, curValidator.ToABCIValidatorUpdate(true))
		}
	}

	if len(newValidators) == 0 {
		panic("consens error. no validator exsits")
	}

	//存储新的validator
	validatorMapper.Set(ecotypes.BuildCurrentValidatorsAddressKey(), newValidators)

	log.Info("update Validators", "len", len(updateValidators))

	return updateValidators
}

func handleValidatorValidatorVoteInfo(ctx context.Context, valAddr btypes.Address, isVote bool, votingWindowLen, minVotingCounter uint64) {

	log := ctx.Logger()
	height := uint64(ctx.BlockHeight())
	validatorMapper := ecomapper.GetValidatorMapper(ctx)
	voteInfoMapper := ecomapper.GetVoteInfoMapper(ctx)

	validator, exsits := validatorMapper.GetValidator(valAddr)
	if !exsits {
		log.Info("validatorVoteInfo", valAddr.String(), "not exsits,may be closed")
		return
	}

	//非Active状态不处理
	if !validator.IsActive() {
		log.Info("validatorVoteInfo", valAddr.String(), "is Inactive")
		return
	}

	voteInfo, exsits := voteInfoMapper.GetValidatorVoteInfo(valAddr)
	if !exsits {
		voteInfo = ecotypes.NewValidatorVoteInfo(height, 0, 0)
	}

	index := voteInfo.IndexOffset % votingWindowLen
	voteInfo.IndexOffset++

	previousVote := voteInfoMapper.GetVoteInfoInWindow(valAddr, index)

	switch {
	case previousVote && !isVote:
		voteInfoMapper.SetVoteInfoInWindow(valAddr, index, false)
		voteInfo.MissedBlocksCounter++
	case !previousVote && isVote:
		voteInfoMapper.SetVoteInfoInWindow(valAddr, index, true)
		voteInfo.MissedBlocksCounter--
	default:
		//nothing
	}

	if !isVote {
		log.Info("validatorVoteInfo", "height", height, valAddr.String(), "not vote")
	}

	// minHeight := voteInfo.StartHeight + votingWindowLen
	maxMissedCounter := votingWindowLen - minVotingCounter

	// if height > minHeight && voteInfo.MissedBlocksCounter > maxMissedCounter
	if voteInfo.MissedBlocksCounter > maxMissedCounter {
		log.Info("validator gets inactive", "height", height, "validator", valAddr.String(), "missed counter", voteInfo.MissedBlocksCounter)
		validatorMapper.MakeValidatorInactive(valAddr, uint64(ctx.BlockHeight()), ctx.BlockHeader().Time.UTC(), ecotypes.MissVoteBlock)
	}

	voteInfoMapper.SetValidatorVoteInfo(valAddr, voteInfo)
}
