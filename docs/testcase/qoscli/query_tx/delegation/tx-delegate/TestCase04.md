# test case of qoscli tx delegate

> `qoscli tx delegate` 委托

---

## 情景说明

1. 对于委托人账户，验证委托之前和委托之后的账户余额变化，委托金额和手续费扣除。
2. 对于委托的地址，验证委托之前和委托之后的绑定tokens变化。
3. 查询委托的信息，校验是否委托成功。

## 测试命令

```bash
//查询委托前账户以及验证人信息
qoscli query delegations jlgy07
qoscli query validator jlgy01
qoscli query account jlgy07

//执行委托交易
qoscli tx delegate --owner jlgy01 --delegator jlgy07 --tokens 1000000

//查询委托后账户以及验证人信息
qoscli query account jlgy07
qoscli query validator jlgy01
qoscli query delegations jlgy07
```

## 预测结果

```bash
1.在账户没有委托操作前是没有可查的信息。
2.执行委托交易后，委托账户余额减少，同时扣除gas费用。
3.验证节点的绑定tokens增加。
4.委托操作交易完成后，可查询到委托账户的委托信息。
```

## 测试结果

```bash
qoscli query delegations jlgy07
null

qoscli query validator jlgy01
{"owner":"address1nnvdqefva89xwppzs46vuskckr7klvzk8r5uaa","validatorAddress":"6E713D1F3CCE28D820C39059E5D08D21D646FA8E","validatorPubkey":{"type":"tendermint/PubKeyEd25519","value":"exGS/yWJthwY8za4dlrPRid2I9KE4G15nlJwO/+Off8="},"bondTokens":"2000000000","description":{"moniker":"jlgy666","logo":"http://pic32.nipic.com/20130813/3347542_160503703000_2.jpg","website":"https://github.com/wangkuanzzu","details":"jlgy23333333333"},"status":"active","InactiveDesc":"Revoked","inactiveTime":"2019-08-08T04:05:39.975061439Z","inactiveHeight":"617984","bondHeight":"617422"}

qoscli query account jlgy07
{"type":"qos/types/QOSAccount","value":{"base_account":{"account_address":"address1p5z8a9u2mmce9qx77knewt6pk4g9hrtdgrwpmf","public_key":null,"nonce":"1"},"qos":"1000000000","qscs":null}}

qoscli tx delegate --owner jlgy01 --delegator jlgy07 --tokens 1000000
Password to sign with 'jlgy07':
{"check_tx":{"gasWanted":"100000","gasUsed":"9366"},"deliver_tx":{"gasWanted":"100000","gasUsed":"60630","tags":[{"key":"YWN0aW9u","value":"Y3JlYXRlLWRlbGVnYXRpb24="},{"key":"dmFsaWRhdG9y","value":"YWRkcmVzczFkZWNuNjhldWVjNWRzZ3hyanB2N3Q1eWR5OHR5ZDc1dzhncnlhZg=="},{"key":"ZGVsZWdhdG9y","value":"YWRkcmVzczFwNXo4YTl1Mm1tY2U5cXg3N2tuZXd0NnBrNGc5aHJ0ZGdyd3BtZg=="}]},"hash":"FE959884DC37BC39B9DFA2EE855D93144D8DE4C1D8280792834A8F5685922E6D","height":"784915"}

qoscli query account jlgy07 --indent
{"type":"qos/types/QOSAccount","value":{"base_account":{"account_address":"address1p5z8a9u2mmce9qx77knewt6pk4g9hrtdgrwpmf","public_key":{"type":"tendermint/PubKeyEd25519","value":"+RWFzvyUul7bTmcykEyFYzFaCA9+vuDJJwxi7ZmGGdI="},"nonce":"1"},"qos":"998994016","qscs":null}}

qoscli query validator jlgy01
{"owner":"address1nnvdqefva89xwppzs46vuskckr7klvzk8r5uaa","validatorAddress":"6E713D1F3CCE28D820C39059E5D08D21D646FA8E","validatorPubkey":{"type":"tendermint/PubKeyEd25519","value":"exGS/yWJthwY8za4dlrPRid2I9KE4G15nlJwO/+Off8="},"bondTokens":"2001000000","description":{"moniker":"jlgy666","logo":"http://pic32.nipic.com/20130813/3347542_160503703000_2.jpg","website":"https://github.com/wangkuanzzu","details":"jlgy23333333333"},"status":"active","InactiveDesc":"Revoked","inactiveTime":"2019-08-08T04:05:39.975061439Z","inactiveHeight":"617984","bondHeight":"617422"}

qoscli query delegations jlgy07
[{"delegator_address":"address1p5z8a9u2mmce9qx77knewt6pk4g9hrtdgrwpmf","owner_address":"address1nnvdqefva89xwppzs46vuskckr7klvzk8r5uaa","validator_pub_key":{"type":"tendermint/PubKeyEd25519","value":"exGS/yWJthwY8za4dlrPRid2I9KE4G15nlJwO/+Off8="},"delegate_amount":"1000000","is_compound":false}]

```
