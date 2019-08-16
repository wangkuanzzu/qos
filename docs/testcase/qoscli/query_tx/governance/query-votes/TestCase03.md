# test case of qoscli query votes

> `qoscli query votes` 查询投票

---

## 情景说明

查询提议的投票信息：查询某一提议上所有投票信息。前提条件：提议编号为5的提议存在.

## 测试命令

```bash
    //查询 所有账户在提议编号为5的投票情况
    qoscli query votes 5
```

## 测试结果

```bash
    qoscli query votes 5
    [{"voter":"address10xwx06gnrt3dlz7hfrx6a8wx3gyeghxm54rv7a","proposal_id":"5","option":"Yes"},{"voter":"address1l0wn66gh45nfta2r4vq8z54wu9hgarss298e9g","proposal_id":"5","option":"No"}]
```
