# test case of qoscli tx redelegate

> `qoscli tx redelegate` 委托

---

## 情景说明

1. 对于委托人验证是否绑定的tokens从一个validator转移到另一个validator，主要验证绑定的tokens变化，委托人账户会扣除gas费用。
2. 对于重新委托的tokens，生效周期，以及收益发放的时间验证。
3. 重新委托的tx操作，无基础的gas费用消耗，但是有访问存储消耗的gas。

## 测试命令

```bash

//转委托前查询账号委托信息
qoscli query delegations jlgy07
qoscli query validator address1nnvdqefva89xwppzs46vuskckr7klvzk8r5uaa
qoscli query validator address1f66wr25emjtp5urfcpd02epwg5ply3xzcv2u20
qoscli query account jlgy07

//转委托交易
qoscli tx redelegate --from-owner address1nnvdqefva89xwppzs46vuskckr7klvzk8r5uaa --to-owner address1f66wr25emjtp5urfcpd02epwg5ply3xzcv2u20 --delegator jlgy07 --tokens 1000000 --max-gas 200000

//转委托后查询账号委托信息
qoscli query account jlgy07
qoscli query validator address1nnvdqefva89xwppzs46vuskckr7klvzk8r5uaa
qoscli query validator address1f66wr25emjtp5urfcpd02epwg5ply3xzcv2u20
qoscli query delegations jlgy07

```

## 预测结果

```bash
1.可以查询到账户的委托信息发生变更。
2.转委托的两个验证节点绑定的tokens，转出的减少，转入的增加
3.账户的余额减少，需要支付gas费用，执行转委托交易。
```

## 测试结果

```bash
qoscli query delegations jlgy07
[{"delegator_address":"address1p5z8a9u2mmce9qx77knewt6pk4g9hrtdgrwpmf","owner_address":"address1nnvdqefva89xwppzs46vuskckr7klvzk8r5uaa","validator_pub_key":{"type":"tendermint/PubKeyEd25519","value":"exGS/yWJthwY8za4dlrPRid2I9KE4G15nlJwO/+Off8="},"delegate_amount":"1000000","is_compound":false}]
qoscli query validator address1nnvdqefva89xwppzs46vuskckr7klvzk8r5uaa
qoscli query validator address1f66wr25emjtp5urfcpd02epwg5ply3xzcv2u20
qoscli query account jlgy07


qoscli tx redelegate --from-owner address1nnvdqefva89xwppzs46vuskckr7klvzk8r5uaa --to-owner address1f66wr25emjtp5urfcpd02epwg5ply3xzcv2u20 --delegator jlgy07 --tokens 1000000 --max-gas 200000
Password to sign with 'jlgy07':
{"check_tx":{"gasWanted":"200000","gasUsed":"11852"},"deliver_tx":{"gasWanted":"200000","gasUsed":"91590","tags":[{"key":"YWN0aW9u","value":"Y3JlYXRlLXJlZGVsZWdhdGlvbg=="},{"key":"dmFsaWRhdG9y","value":"YWRkcmVzczFkZWNuNjhldWVjNWRzZ3hyanB2N3Q1eWR5OHR5ZDc1dzhncnlhZg=="},{"key":"bmV3LXZhbGlkYXRvcg==","value":"YWRkcmVzczFoeHl1dDJkeXZydnh1bGZ1OGZsYXl0MDl3eWhxN3IwNG05OGx2Ng=="},{"key":"ZGVsZWdhdG9y","value":"YWRkcmVzczFsMHduNjZnaDQ1bmZ0YTJyNHZxOHo1NHd1OWhnYXJzczI5OGU5Zw=="}]},"hash":"F9A859D542D2F2F28E6A6828750A1FA699BDABB3B457040994AD5DD66E92B7FB","height":"577817"}


qoscli query account jlgy07
qoscli query validator address1nnvdqefva89xwppzs46vuskckr7klvzk8r5uaa
qoscli query validator address1f66wr25emjtp5urfcpd02epwg5ply3xzcv2u20
qoscli query delegations jlgy07
[{"delegator_address":"address1p5z8a9u2mmce9qx77knewt6pk4g9hrtdgrwpmf","owner_address":"address1f66wr25emjtp5urfcpd02epwg5ply3xzcv2u20","validator_pub_key":{"type":"tendermint/PubKeyEd25519","value":"exGS/yWJthwY8za4dlrPRid2I9KE4G15nlJwO/+Off8="},"delegate_amount":"1000000","is_compound":false}]

```
