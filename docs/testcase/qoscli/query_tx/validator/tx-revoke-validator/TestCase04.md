# test case of qoscli revoke validator*

> `qoscli revoke validator*` 撤销验证节点

---

## 情景说明

1. 撤销验证节点：执行操作前，确认节点状态为active，执行操作后，查看状态是否更改为inactive。
2. 状态变更为inactive，验证节点创建人以及节点的委托人，不能再收到参与系统挖矿所获取的收益了，监测创建人以及委托人账户变化（在没有参与其他交易情况下）。
3. 等待系统设定的周期过后，查看验证节点是否关闭。关闭后对委托的tokens进行unbond操作，等待unbond周期结束后，绑定的tokens是否原数归还。
4. 验证节点的votingpower占全网1/3以下，撤销后对全网的影响。
5. 验证节点的votingpower占全网2/3以上，撤销后对全网的影响。
6. 撤销验证节点的tx操作，有基础的GAS费用消耗（18QOS=18000000GAS），再加上访问存储所消耗的GAS费用。验证操作人账户余额的扣除的GAS费用。

## 测试命令

```bash
//参照测试结果中命令
```

## 预测结果

```bash
1.验证节点的状态即时修改。
2.节点状态为inactive下，委托人的账户数额，在等待大于收益发放周期后，未发生变化。
3.节点从inactive的状态下，在等待存活时间后，验证节点关闭，无法查询到。此时委托人未收到先前委托的tokens，在unbonding中可以查询到。
4.等待"unbond_return_height": 259200  ~  15day时间后，委托的tokens才会返还至委托人账户中。
5.验证节点的votingpower占全网的1/3以下，对全网无影响，全网正常运作，只是少一个验证节点。
6.验证节点的votingpower占全网的2/3以上，全网崩溃。
7.消耗的gas费用大于18000000GAS
```

## 测试结果

```bash
[root@iz2zef57ni8z6ydrcloxy0z ~]# qoscli query validator jlgy 
{"owner":"address1nrm24vxav0u58jpt04v7vtavrjjxww60x9ym7p","validatorAddress":"1C97C563A804CB8B379622DA2610422807DFDCA9","validatorPubkey":{"type":"tendermint/PubKeyEd25519","value":"Hx/cQ562Bp8ZJpRBg0/r17+QvgEPxBptqjBoAeNYH2c="},"bondTokens":"610000000","description":{"moniker":"jlgy","logo":"","website":"","details":""},"commission":{"commission_rates":{"rate":"0.100000000000000000","max_rate":"0.200000000000000000","max_change_rate":"0.010000000000000000"},"update_time":"2019-08-29T06:16:18.033191341Z"},"status":"active","InactiveDesc":"Replaced","inactiveTime":"2019-08-29T06:21:06.822679715Z","inactiveHeight":"16404","bondHeight":"16348"}

[root@iz2zef57ni8z6ydrcloxy0z ~]# qoscli query account jlgy
{"type":"qos/types/QOSAccount","value":{"base_account":{"account_address":"address1nrm24vxav0u58jpt04v7vtavrjjxww60x9ym7p","public_key":{"type":"tendermint/PubKeyEd25519","value":"9fpQTirTSpYNHw4qu6CdN456oEZeb5AM0PxIMRZre2s="},"nonce":"4"},"qos":"1395327397","qscs":null}}

[root@iz2zef57ni8z6ydrcloxy0z ~]# qoscli tx revoke-validator --owner jlgy
Password to sign with 'jlgy':
{"check_tx":{"gasWanted":"9223372036854775807","gasUsed":"1808258","events":[]},"deliver_tx":{"gasWanted":"9223372036854775807","gasUsed":"1822700","events":[{"type":"revoke-validator","attributes":[{"key":"dmFsaWRhdG9y","value":"YWRkcmVzczFyanR1MmNhZ3FuOWNrZHVreXRkenZ5eno5cXJhbGg5ZnAzM2NrOQ=="},{"key":"b3duZXI=","value":"YWRkcmVzczFucm0yNHZ4YXYwdTU4anB0MDR2N3Z0YXZyamp4d3c2MHg5eW03cA=="}]},{"type":"message","attributes":[{"key":"bW9kdWxl","value":"c3Rha2U="},{"key":"Z2FzLnBheWVy","value":"YWRkcmVzczFucm0yNHZ4YXYwdTU4anB0MDR2N3Z0YXZyamp4d3c2MHg5eW03cA=="}]}]},"hash":"C5074E18F19615E65FB6007E25BA7F3D869AB33E1FED3FF5CF51F44AC14C6A22","height":"17422"}

[root@iz2zef57ni8z6ydrcloxy0z ~]# qoscli query validator jlgy
{"owner":"address1nrm24vxav0u58jpt04v7vtavrjjxww60x9ym7p","validatorAddress":"1C97C563A804CB8B379622DA2610422807DFDCA9","validatorPubkey":{"type":"tendermint/PubKeyEd25519","value":"Hx/cQ562Bp8ZJpRBg0/r17+QvgEPxBptqjBoAeNYH2c="},"bondTokens":"610000000","description":{"moniker":"jlgy","logo":"","website":"","details":""},"commission":{"commission_rates":{"rate":"0.100000000000000000","max_rate":"0.200000000000000000","max_change_rate":"0.010000000000000000"},"update_time":"2019-08-29T06:16:18.033191341Z"},"status":"inactive","InactiveDesc":"Revoked","inactiveTime":"2019-08-29T07:48:31.820138345Z","inactiveHeight":"17422","bondHeight":"16348"}

[root@iz2zef57ni8z6ydrcloxy0z ~]# qoscli query validators --indent | grep "moniker"
      "moniker": "星红安",
      "moniker": "出行链",
      "moniker": "jlgy",
      "moniker": "EPICOM",
      "moniker": "qos",
      "moniker": "缔联科技",
      "moniker": "瑞格钱包",
      "moniker": "snow capital",
[root@iz2zef57ni8z6ydrcloxy0z ~]# qoscli query validators --indent | grep "status"
    "status": "active",
    "status": "active",
    "status": "inactive",
    "status": "active",
    "status": "active",
    "status": "active",
    "status": "active",
    "status": "inactive",

//高度为16356，revoke高度为16348，成为validator高度为16539按正常收益发放下次收益应该为：16539+720+720=17979
[root@iz2zef57ni8z6ydrcloxy0z ~]# qoscli query account jlgy
{"type":"qos/types/QOSAccount","value":{"base_account":{"account_address":"address1nrm24vxav0u58jpt04v7vtavrjjxww60x9ym7p","public_key":{"type":"tendermint/PubKeyEd25519","value":"9fpQTirTSpYNHw4qu6CdN456oEZeb5AM0PxIMRZre2s="},"nonce":"5"},"qos":"1395309170","qscs":null}}

//在高度大于17979查询，已经没有挖矿奖励
[root@iz2zef57ni8z6ydrcloxy0z ~]# qoscli query account jlgy
{"type":"qos/types/QOSAccount","value":{"base_account":{"account_address":"address1nrm24vxav0u58jpt04v7vtavrjjxww60x9ym7p","public_key":{"type":"tendermint/PubKeyEd25519","value":"9fpQTirTSpYNHw4qu6CdN456oEZeb5AM0PxIMRZre2s="},"nonce":"5"},"qos":"1395309170","qscs":null}}

//从revoke的高度16348，等待28800s后，也就是8个小时后。查询验证人信息
[root@iz2zef57ni8z6ydrcloxy0z ~]# qoscli query validator jlgy
ERROR: owner does't have validator

//"unbond_return_height": 259200  ~  15day    "survival_secs": 28800 ~ 8hour
[root@iz2zef57ni8z6ydrcloxy0z ~]# qoscli query unbondings jlgy
[{"delegator_addr":"address1nrm24vxav0u58jpt04v7vtavrjjxww60x9ym7p","validator_addr":"address1rjtu2cagqn9ckdukytdzvyzz9qralh9fp33ck9","height":"9415","complete_height":"268615","delegate_amount":"100000000"}]
```
