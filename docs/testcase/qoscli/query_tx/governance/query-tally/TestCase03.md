# test case of qoscli query tally

> `qoscli query tally` 查询投票

---

## 情景说明

查询提议的投票信息：查询某一提议上所有投票信息。前提条件：提议编号为5的提议存在.

## 测试命令

```bash
    //查询 提议编号为5的不同vote type的统计结果
    qoscli query tally 5
```

## 测试结果

```bash
    qoscli query tally 5
    {"yes":"1900000000","abstain":"0","no":"1900000100","no_with_veto":"0"}

```
