# test case of qoscli query vote

> `qoscli query vote` 查询投票

---

## 情景说明

查询提议的投票信息：查询某一提议上某一账号的投票信息；前提条件：提议编号为5的提议存在.

## 测试命令

```bash
    //查询 账户abc在提议编号为5的投票情况
    qoscli query vote 5 abc
```

## 测试结果

```bash
    qoscli query vote 5 abc
    {"voter":"address1l0wn66gh45nfta2r4vq8z54wu9hgarss298e9g","proposal_id":"5","option":"No"}
```
