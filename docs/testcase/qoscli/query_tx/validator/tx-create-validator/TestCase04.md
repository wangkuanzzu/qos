# test case of qoscli create validator*

> `qoscli create validator*` 创建验证节点

---

## 情景说明

1. 创建验证节点：前提是节点需要满足官方文档中的要求，使用之前在网路中创建的账户进行创建，创建后查询验证节点信息。
2. 节点创建后根据状态可以区分为三种，此处关注成为验证节点，可以参与网络中打块与投票，获取挖矿收益，检测创建人账户变化。

## 测试命令

```bash
    //验证create语句
    qoscli query validator jlgy01

    qoscli tx create-validator --moniker jlgy666 --owner jlgy01 --tokens 2000000000

    //验证create语句
    qoscli query validator jlgy01

```

## 预测结果

```bash
1.创建验证节点成功后，可以查询到。
```

## 测试结果

```bash
    qoscli query validator jlgy01
    ERROR: owner does not have validator

    qoscli tx create-validator --moniker jlgy666 --owner jlgy01 --tokens 2000000000
    Password to sign with 'jlgy01':
    {"check_tx":{"gasWanted":"100000","gasUsed":"8916"},"deliver_tx":{"gasWanted":"100000","gasUsed":"45720","tags":[{"key":"YWN0aW9u","value":"Y3JlYXRlLXZhbGlkYXRvcg=="},{"key":"dmFsaWRhdG9y","value":"YWRkcmVzczFkZWNuNjhldWVjNWRzZ3hyanB2N3Q1eWR5OHR5ZDc1dzhncnlhZg=="},{"key":"b3duZXI=","value":"YWRkcmVzczFubnZkcWVmdmE4OXh3cHB6czQ2dnVza2NrcjdrbHZ6azhyNXVhYQ=="},{"key":"ZGVsZWdhdG9y","value":"YWRkcmVzczFubnZkcWVmdmE4OXh3cHB6czQ2dnVza2NrcjdrbHZ6azhyNXVhYQ=="}]},"hash":"924D7AD4B02BBD32AE0C6F1228BE02802F2B6A098C55EE3FBAE88D6217B6C4FF","height":"617422"}

    qoscli query validator jlgy01
    {"owner":"address1nnvdqefva89xwppzs46vuskckr7klvzk8r5uaa","validatorAddress":"6E713D1F3CCE28D820C39059E5D08D21D646FA8E","validatorPubkey":{"type":"tendermint/PubKeyEd25519","value":"exGS/yWJthwY8za4dlrPRid2I9KE4G15nlJwO/+Off8="},"bondTokens":"2000000000","description":{"moniker":"jlgy","logo":"","website":"","details":""},"status":"active","InactiveDesc":"","inactiveTime":"0001-01-01T00:00:00Z","inactiveHeight":"0","bondHeight":"617422"}
```
