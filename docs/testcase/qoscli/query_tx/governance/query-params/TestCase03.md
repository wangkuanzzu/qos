# test case of qoscli query params

> `qoscli query params` 查询参数

---

## 情景说明

1. 查询`系统`所配置的参数信息，全部参数查询。
2. 查询`模块`所配置的参数信息，全部参数查询。
3. 查询`模块`所配置的参数信息，`单个`参数查询。用户可以通过提议的方式对单个参数进行调整，前提是该提议必须完成投票得到qos网络中所有验证人的认可。

## 测试命令

```bash
    qoscli query params

    qoscli query params --module gov

    qoscli query params --module stake

    qoscli query params --module distribution

    qoscli query params --module gov --key min_deposit
```

## 预测结果

```bash
    得到对应的系统的参数设置。
```

## 测试结果

```bash
    qoscli query params
    [{"type":"stake/params","value":{"max_validator_cnt":10,"voting_status_len":100,"voting_status_least":50,"survival_secs":600,"unbond_return_height":10}},{"type":"distribution/params","value":{"proposer_reward_rate":{"value":"0.040000000000000000"},"community_reward_rate":{"value":"0.010000000000000000"},"validator_commission_rate":{"value":"0.010000000000000000"},"delegator_income_period_height":"10","gas_per_unit_cost":"10"}},{"type":"gov/params","value":{"min_deposit":"10","min_proposer_deposit_rate":"0.334000000000000000","max_deposit_period":"172800000000000","voting_period":"172800000000000","quorum":"0.334000000000000000","threshold":"0.500000000000000000","veto":"0.334000000000000000","penalty":"0.000000000000000000","burn_rate":"0.500000000000000000"}}]

    qoscli query params --module gov
    {"type":"gov/params","value":{"min_deposit":"10","min_proposer_deposit_rate":"0.334000000000000000","max_deposit_period":"172800000000000","voting_period":"172800000000000","quorum":"0.334000000000000000","threshold":"0.500000000000000000","veto":"0.334000000000000000","penalty":"0.000000000000000000","burn_rate":"0.500000000000000000"}}

    qoscli query params --module stake
    {"type":"stake/params","value":{"max_validator_cnt":10,"voting_status_len":100,"voting_status_least":50,"survival_secs":600,"unbond_return_height":10}}

    qoscli query params --module distribution
    {"type":"distribution/params","value":{"proposer_reward_rate":{"value":"0.040000000000000000"},"community_reward_rate":{"value":"0.010000000000000000"},"validator_commission_rate":{"value":"0.010000000000000000"},"delegator_income_period_height":"10","gas_per_unit_cost":"10"}}

    qoscli query params --module gov --key min_deposit
    "10"
```
