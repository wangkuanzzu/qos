# test case of qoscli tx transfer

> `qoscli tx transfer` 转账操作

---

## 情景说明

1. 单对单账户转账：转出账户金额的减少，转入账户金额的增加，变化是否一致。
2. 单对多账户转账：转出账户金额的减少，转入的多个账户金额的变化。
3. 多对单账户转账：多个转出账户金额的减少是否正确，转入的账户金额的变化是否等于多个转出合计。
4. 多对多账户转账：多个转出账户金额的减少是否正确，多个转入账户金额的增加是否正确，

## 测试命令

```bash
    //单对单
    qoscli query account jlgy05
    qoscli query account jlgy06
    qoscli tx transfer --senders jlgy05,50000QOS  --receivers jlgy06,50000QOS
    qoscli query account jlgy05
    qoscli query account jlgy06

    //单对多
    qoscli query account jlgy05
    qoscli query account jlgy06
    qoscli query account jlgy07
    qoscli tx transfer --senders jlgy05,50000QOS  --receivers jlgy06,10000QOS;jlgy07,40000QOS
    qoscli query account jlgy05
    qoscli query account jlgy06
    qoscli query account jlgy07

    //多对单
    qoscli query account jlgy05
    qoscli query account jlgy06
    qoscli query account jlgy07
    qoscli tx transfer --senders jlgy05,50000QOS;jlgy06,10000QOS  --receivers jlgy07,60000QOS
    qoscli query account jlgy05
    qoscli query account jlgy06
    qoscli query account jlgy07


    //多对多
    qoscli query account jlgy05
    qoscli query account jlgy06
    qoscli query account jlgy07
    qoscli query account jlgy08
    qoscli tx transfer --senders jlgy05,50000QOS;jlgy06,10000QOS  --receivers jlgy07,20000QOS;jlgy08,40000QOS
    qoscli query account jlgy05
    qoscli query account jlgy06
    qoscli query account jlgy07
    qoscli query account jlgy08

```

## 预测结果

```bash
1.账户jlgy05减少50000qos，同时会扣除gas费用，账户jlgy06增加50000qos。

```

## 测试结果

```bash

```
