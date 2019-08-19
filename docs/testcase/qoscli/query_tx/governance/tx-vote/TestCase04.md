# test case of qoscli tx vote

> `qoscli tx vote` 提议投票

---

## 情景说明

1. 对某一提议完成投票后，可以查询到，验证投票操作成功。
2. 在投票阶段可以多次进行投票，记录最后一次投票。

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
```

## 测试结果

```bash

```
