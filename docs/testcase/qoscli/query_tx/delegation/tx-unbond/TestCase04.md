# test case of qoscli tx unbond

> `qoscli tx unbond` 委托

---

## 情景说明

1. 对于解绑的tokens，验证是否返还到委托人账户，委托的节点的tokens是否减少。
2. 对于委托人解绑的tokens，在到收益发放时候对比先前收益应该减少。

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
1.解绑操作后，委托人账户余额增加，会扣除交易花费gas费用，验证节点绑定的tokens减少。
2.收益会降低。
```

## 测试结果

```bash

```
