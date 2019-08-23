# test case of qoscli tx deposit

> `qoscli tx deposit` 提议抵押

---

## 情景说明

1. 验证质押前后账户余额的变化，扣除抵押tokens，扣除gas费用。
2. 验证质押前后提议的抵押token变化。
3. 校验提议所处的阶段是否随着抵押的token数量增加而变更。
4. 抵押的tx操作，无基础的gas费用消耗，但是有访问存储消耗的gas。

## 测试命令

```bash
    //验证抵押成功
    qoscli query account jlgy07
    qoscli query proposal 1

    qoscli tx deposit --proposal-id 1 --depositor jlgy07 --amount 10000

    qoscli query account jlgy07
    qoscli query proposal 1


    //验证提议阶段变化
    qoscli tx deposit --proposal-id 1 --depositor jlgy07 --amount 10000
    ...
    qoscli tx deposit --proposal-id 1 --depositor jlgy07 --amount 10000

    qoscli query proposal 1
```

## 预测结果

```text
1.抵押的账户余额减少，10000和消耗的gas费用。
2.提议的抵押tokens增加。10000
3.提议的阶段会在不断增加抵押tokens后，达到最小抵押限制，进入voting阶段。
```

## 测试结果

```bash

```
