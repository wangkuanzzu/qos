# test case of qoscli tx submit-proposal

> `qoscli tx submit-proposal` 提交提议

---

## 情景说明

1. 提议成功后，是否可以查询到。
2. 校验提议的状态，质押等其他信息是否正确。
3. 提议的状态在创建提议时候会依据抵押数额，直接进入某一状态，验证不同数额下状态是否正确。

## 测试命令

```bash
    //查询目前存在的提议
    qoscli query proposals
    //创建提议
    qoscli tx submit-proposal --title "myproposal" --proposal-type "Text" --proposer jlgy07 --deposit 40000000 --description "the first proposal"
    //提议创建后查询
    qoscli query proposals

    //验证信息是否正确，阶段
    qoscli query proposal <新增提议的编号>

    //创建提议，直接进入voting阶段
    qoscli tx submit-proposal --title "myproposal" --proposal-type "Text" --proposer jlgy07 --deposit 140000000 --description "the first proposal"

```

## 预测结果

```bash
1.提议创建后，查询所有提议时候可以查到
2.质押tokens较低时候提议阶段进入deposit阶段。
3.质押tokens较低时候提议阶段进入voting阶段。
```

## 测试结果

```bash

```
