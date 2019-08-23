# test case of qoscli tx create-approve 

> `qoscli tx create-approve` 创建授权

---

## 情景说明

1. 分别查询在创建预授权之前和创建预授权之后的授权信息，校验是否创建预授权成功。
2. 创建授权操作，监测授权人和被授权人的账户数额变化。
3. 创建授权相关的tx操作，无基础的gas费用消耗，但是有访问存储消耗的gas。

## 测试命令

```bash
//未创建授权前查询
qoscli query approve --from jlgy05 --to jlgy06
//未创建授权前查询账户数额
qoscli query account jlgy05
qoscli query account jlgy06
//创建授权
qoscli tx create-approve --from jlgy05 --to jlgy06 --coins 1000000qos
//创建授权后查询账户数额
qoscli query account jlgy05
qoscli query account jlgy06
//创建授权后查询
qoscli query approve --from jlgy05 --to jlgy06
```

## 预测结果

```bash
    1.创建预授权之前，查询授权信息不存在
    2.执行预授权创建交易后，可以查询到授权信息。
    3.授权的意思：仅仅是一个授权，在没有实质对授权使用情况下，授权人和被授权人的账户数额是不会发生变化的。
    4.对于授权人执行创建预授权交易时候，需要支付gas来完成交易，故授权人账户只会扣除相应的gas。
```

## 测试结果

```bash
    qoscli query approve --from jlgy05 --to jlgy06
    ERROR: approve does not exist

    qoscli query account jlgy05
    {"type":"qos/types/QOSAccount","value":{"base_account":{"account_address":"address1y4qgt89vfujkh7sc0wfnfqn2q5qmaerlunmnax","public_key":null,"nonce":"0"},"qos":"1000000000","qscs":null}}
    qoscli query account jlgy06
    {"type":"qos/types/QOSAccount","value":{"base_account":{"account_address":"address1zegv02xua87qsehmrmclrykqqdwnwkuf9ek5gt","public_key":null,"nonce":"0"},"qos":"1000000000","qscs":null}}

    qoscli tx create-approve --from jlgy05 --to jlgy06 --coins 1000000qos
    Password to sign with 'jlgy05':
    {"check_tx":{"gasWanted":"100000","gasUsed":"6550"},"deliver_tx":{"gasWanted":"100000","gasUsed":"12500","tags":[{"key":"YWN0aW9u","value":"Y3JlYXRlLWFwcHJvdmU="},{"key":"YXBwcm92ZS1mcm9t","value":"YWRkcmVzczF5NHFndDg5dmZ1amtoN3NjMHdmbmZxbjJxNXFtYWVybHVubW5heA=="},{"key":"YXBwcm92ZS10bw==","value":"YWRkcmVzczF6ZWd2MDJ4dWE4N3FzZWhtcm1jbHJ5a3FxZHdud2t1ZjllazVndA=="}]},"hash":"DD3B510948755F658D62D827ABFB4B0014FB907B32F85BAA0A2FFEF11BFFB6F7","height":"783939"}

    qoscli query account jlgy05
    {"type":"qos/types/QOSAccount","value":{"base_account":{"account_address":"address1y4qgt89vfujkh7sc0wfnfqn2q5qmaerlunmnax","public_key":{"type":"tendermint/PubKeyEd25519","value":"x54mlQpo3XDIfPevRQEqHgVS9/6cG85a+VjUDbPteEU="},"nonce":"1"},"qos":"999998750","qscs":null}}
    qoscli query account jlgy06
    {"type":"qos/types/QOSAccount","value":{"base_account":{"account_address":"address1zegv02xua87qsehmrmclrykqqdwnwkuf9ek5gt","public_key":null,"nonce":"0"},"qos":"1000000000","qscs":null}}


    qoscli query approve --from jlgy05 --to jlgy06
    {"from":"address1y4qgt89vfujkh7sc0wfnfqn2q5qmaerlunmnax","to":"address1zegv02xua87qsehmrmclrykqqdwnwkuf9ek5gt","qos":"1000000","qscs":null}
```
