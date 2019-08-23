# test case of qoscli tx vote

> `qoscli tx vote` 提议投票

---

## 情景说明

1. 对某一提议完成投票后，可以查询到，验证投票操作成功。
2. 在投票阶段可以多次进行投票，记录最后一次投票。
3. 投票阶段完成后，提议的处理方式。通过，不通过，严重反对。校验质押人质押的tokens去向。
4. 投票的tx操作，无基础的gas费用消耗，但是有访问存储消耗的gas。

## 测试命令

```bash
    //查询 账户abc在提议编号为5的投票情况
    qoscli query vote 5 jlgy07

    //首次投票
    qoscli tx vote --proposal-id 5 --voter jlgy07 --option yes

    //查询 账户abc在提议编号为5的投票情况
    qoscli query vote 5 jlgy07

    //再次投票
    qoscli tx vote --proposal-id 5 --voter jlgy07 --option no

    //查询 账户abc在提议编号为5的投票情况
    qoscli query vote 5 jlgy07
```

## 预测结果

```bash
1.初始查询投票没有结果
2.首次投票后，查询投票结果为yes
3.再次投票后，查询投票结果为no
4.投票前后，投票人账户的数额会减少，vote的tx会消耗访问存储的gas费用。
```

## 测试结果

```bash

```
