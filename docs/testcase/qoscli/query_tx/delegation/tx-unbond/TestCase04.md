# test case of qoscli tx unbond

> `qoscli tx unbond` 委托

---

## 情景说明

1. 对于解绑的tokens，验证是否返还到委托人账户，委托的节点的tokens是否减少。
2. 对于委托人解绑的tokens，在到收益发放时候对比先前收益应该减少。
3. 委托的tx操作，有基础的gas费用（0.18QOS=180000GAS）消耗，再加上访问存储消耗的GAS，为该交易消耗的总GAS费用。

## 测试命令

```bash
    //验证测试结果命令
    qoscli query delegations jlgy07
    qoscli query validator jlgy01
    qoscli query account jlgy07

    //解除委托10000
    qoscli tx unbond --owner jlgy01 --delegator jlgy07 --tokens 10000

    //验证测试结果命令
    qoscli query delegations jlgy07
    qoscli query validator jlgy01
    qoscli query account jlgy07

    //解除所有委托的qos数量
    qoscli tx unbond --owner jlgy01 --delegator jlgy07 --tokens 10000 --all

    //验证测试结果命令
    qoscli query delegations jlgy07
    qoscli query validator jlgy01
    qoscli query account jlgy07
```

## 预测结果

```bash
1.解绑操作后，等待一个解绑周期结束后，委托人账户余额增加，会扣除交易花费gas费用，验证节点绑定的tokens减少。
2.收益会降低。
```

## 测试结果

```bash
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query delegations acc0
[{"delegator_address":"address1hzyqw0d4nkaujk0lp2ef269fp2scm5ejq4jk55","owner_address":"address1hzyqw0d4nkaujk0lp2ef269fp2scm5ejq4jk55","validator_pub_key":{"type":"tendermint/PubKeyEd25519","value":"WCQPqO447eQxB/2EIv253UzHsadsK3CBmiROWSbVDwc="},"delegate_amount":"4000000000000","is_compound":false}]
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query validator acc0
{"owner":"address1hzyqw0d4nkaujk0lp2ef269fp2scm5ejq4jk55","validatorAddress":"99CDBE9B5675A5BB65376D227423C2448B016D33","validatorPubkey":{"type":"tendermint/PubKeyEd25519","value":"WCQPqO447eQxB/2EIv253UzHsadsK3CBmiROWSbVDwc="},"bondTokens":"4000000000000","description":{"moniker":"firstvalidator","logo":"","website":"","details":""},"commission":{"commission_rates":{"rate":"0.100000000000000000","max_rate":"0.200000000000000000","max_change_rate":"0.010000000000000000"},"update_time":"0001-01-01T00:00:00Z"},"status":"active","InactiveDesc":"","inactiveTime":"0001-01-01T00:00:00Z","inactiveHeight":"0","bondHeight":"0"}
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query account acc0
{"type":"qos/types/QOSAccount","value":{"base_account":{"account_address":"address1hzyqw0d4nkaujk0lp2ef269fp2scm5ejq4jk55","public_key":{"type":"tendermint/PubKeyEd25519","value":"E8ju2XSEevgHMYbU3cA5mvE+JiS6qC7mRihpc6kPMxM="},"nonce":"85"},"qos":"5003458450771","qscs":null}}
[vagrant@vagrant-192-168-1-200 ~]$ qoscli tx unbond --owner acc0 --delegator acc0 --tokens 1000000000000
Password to sign with 'acc0':
{"check_tx":{"gasWanted":"9223372036854775807","gasUsed":"27513","events":[]},"deliver_tx":{"gasWanted":"9223372036854775807","gasUsed":"84208","events":[{"type":"unbond-delegation","attributes":[{"key":"dmFsaWRhdG9y","value":"YWRkcmVzczFuOHhtYXg2a3dram1rZWZoZDUzOGdnN3pnajlzem1mbnJqbTI4Yw=="},{"key":"ZGVsZWdhdG9y","value":"YWRkcmVzczFoenlxdzBkNG5rYXVqazBscDJlZjI2OWZwMnNjbTVlanE0ams1NQ=="}]},{"type":"message","attributes":[{"key":"bW9kdWxl","value":"c3Rha2U="},{"key":"Z2FzLnBheWVy","value":"YWRkcmVzczFoenlxdzBkNG5rYXVqazBscDJlZjI2OWZwMnNjbTVlanE0ams1NQ=="}]}]},"hash":"DCC1816C05CB77A166B982F96804D79B52D582E19D91C88710DCD4AA9313388B","height":"3887"}
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query delegations acc0
[{"delegator_address":"address1hzyqw0d4nkaujk0lp2ef269fp2scm5ejq4jk55","owner_address":"address1hzyqw0d4nkaujk0lp2ef269fp2scm5ejq4jk55","validator_pub_key":{"type":"tendermint/PubKeyEd25519","value":"WCQPqO447eQxB/2EIv253UzHsadsK3CBmiROWSbVDwc="},"delegate_amount":"3000000000000","is_compound":false}]
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query validator acc0
{"owner":"address1hzyqw0d4nkaujk0lp2ef269fp2scm5ejq4jk55","validatorAddress":"99CDBE9B5675A5BB65376D227423C2448B016D33","validatorPubkey":{"type":"tendermint/PubKeyEd25519","value":"WCQPqO447eQxB/2EIv253UzHsadsK3CBmiROWSbVDwc="},"bondTokens":"3000000000000","description":{"moniker":"firstvalidator","logo":"","website":"","details":""},"commission":{"commission_rates":{"rate":"0.100000000000000000","max_rate":"0.200000000000000000","max_change_rate":"0.010000000000000000"},"update_time":"0001-01-01T00:00:00Z"},"status":"active","InactiveDesc":"","inactiveTime":"0001-01-01T00:00:00Z","inactiveHeight":"0","bondHeight":"0"}

//等待解绑返还周期结束
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query account acc0
{"type":"qos/types/QOSAccount","value":{"base_account":{"account_address":"address1hzyqw0d4nkaujk0lp2ef269fp2scm5ejq4jk55","public_key":{"type":"tendermint/PubKeyEd25519","value":"E8ju2XSEevgHMYbU3cA5mvE+JiS6qC7mRihpc6kPMxM="},"nonce":"86"},"qos":"5003458450771","qscs":null}}

[vagrant@vagrant-192-168-1-200 ~]$ qoscli tx unbond --owner acc0 --delegator acc0 --tokens 1000000000000 --all
Password to sign with 'acc0':
{"check_tx":{"gasWanted":"9223372036854775807","gasUsed":"27513","events":[]},"deliver_tx":{"gasWanted":"9223372036854775807","gasUsed":"84208","events":[{"type":"unbond-delegation","attributes":[{"key":"dmFsaWRhdG9y","value":"YWRkcmVzczFuOHhtYXg2a3dram1rZWZoZDUzOGdnN3pnajlzem1mbnJqbTI4Yw=="},{"key":"ZGVsZWdhdG9y","value":"YWRkcmVzczFoenlxdzBkNG5rYXVqazBscDJlZjI2OWZwMnNjbTVlanE0ams1NQ=="}]},{"type":"message","attributes":[{"key":"bW9kdWxl","value":"c3Rha2U="},{"key":"Z2FzLnBheWVy","value":"YWRkcmVzczFoenlxdzBkNG5rYXVqazBscDJlZjI2OWZwMnNjbTVlanE0ams1NQ=="}]}]},"hash":"DCC1816C05CB77A166B982F96804D79B52D582E19D91C88710DCD4AA9313388B","height":"3887"}

[vagrant@vagrant-192-168-1-200 ~]$ qoscli query delegations acc0
null
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query validator acc0
{"owner":"address1hzyqw0d4nkaujk0lp2ef269fp2scm5ejq4jk55","validatorAddress":"99CDBE9B5675A5BB65376D227423C2448B016D33","validatorPubkey":{"type":"tendermint/PubKeyEd25519","value":"WCQPqO447eQxB/2EIv253UzHsadsK3CBmiROWSbVDwc="},"bondTokens":"2000000000000","description":{"moniker":"firstvalidator","logo":"","website":"","details":""},"commission":{"commission_rates":{"rate":"0.100000000000000000","max_rate":"0.200000000000000000","max_change_rate":"0.010000000000000000"},"update_time":"0001-01-01T00:00:00Z"},"status":"active","InactiveDesc":"","inactiveTime":"0001-01-01T00:00:00Z","inactiveHeight":"0","bondHeight":"0"}
//等待解绑返还周期结束
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query account acc0
{"type":"qos/types/QOSAccount","value":{"base_account":{"account_address":"address1hzyqw0d4nkaujk0lp2ef269fp2scm5ejq4jk55","public_key":{"type":"tendermint/PubKeyEd25519","value":"E8ju2XSEevgHMYbU3cA5mvE+JiS6qC7mRihpc6kPMxM="},"nonce":"86"},"qos":"5003458450771","qscs":null}}

```
