# test case of qoscli tx delegate

> `qoscli tx delegate` 委托

---

## 情景说明

1. 对于委托人账户查询前后的复投方式，验证是否修改成功。
2. 验证在不同复投方式下的委托人账户金额的收益变化。
3. 验证在不同复投方式下，委托人在验证节点的委托tokens变化。

## 测试命令

```bash
qoscli query delegations jlgy07

qoscli tx modify-compound --owner jlgy01 --delegator jlgy07 --compound

qoscli query delegations jlgy07

`情景2，3暂时没有测试。`

```

## 预测结果

```bash
1.查询委托信息，通过参数is_compound：true或是false，验证修改成功。
2.在参数is_compound：false情况下：委托人账户余额会随着收益发放增加，委托的验证节点绑定的tokens不会发生变化，前提不存在其他委托操作。
3.在参数is_compound：true情况下：委托人账户余额会随着收益发放不变，委托的验证节点绑定的tokens会增加，前提不存在其他委托操作。
```

## 测试结果

```bash
qoscli query delegations jlgy07
[{"delegator_address":"address1p5z8a9u2mmce9qx77knewt6pk4g9hrtdgrwpmf","owner_address":"address1nnvdqefva89xwppzs46vuskckr7klvzk8r5uaa","validator_pub_key":{"type":"tendermint/PubKeyEd25519","value":"exGS/yWJthwY8za4dlrPRid2I9KE4G15nlJwO/+Off8="},"delegate_amount":"1000000","is_compound":false}]

qoscli tx modify-compound --owner jlgy01 --delegator jlgy07 --compound

qoscli query delegations jlgy07
[{"delegator_address":"address1p5z8a9u2mmce9qx77knewt6pk4g9hrtdgrwpmf","owner_address":"address1nnvdqefva89xwppzs46vuskckr7klvzk8r5uaa","validator_pub_key":{"type":"tendermint/PubKeyEd25519","value":"exGS/yWJthwY8za4dlrPRid2I9KE4G15nlJwO/+Off8="},"delegate_amount":"1000000","is_compound":true}]
```
