# test case of qoscli tx increase-approve 

> `qoscli tx increase-approve` 增加授权

---

## 情景说明

1. 增加预授权后，查询预授权信息，验证前后的预授权信息变化是否正确。
2. 增加授权的tx操作，无基础的gas费用消耗，但是有访问存储消耗的gas。

## 测试命令

```bash
//增加预授权前查询授权信息,账户jlgy05账户余额
qoscli query approve --from jlgy05 --to jlgy06
qoscli query account jlgy05
//执行增加预授权交易
qoscli tx increase-approve --from jlgy05 --to jlgy06 --coins 200000qos
//增加预授权后查询授权信息,账户jlgy05账户余额
qoscli query approve --from jlgy05 --to jlgy06
qoscli query account jlgy05
```

## 预测结果

```bash
1.授权的额度在完成增加预授权交易后增加。
2.账户jlgy05会扣除消耗的gas费用。
```

## 测试结果

```bash

qoscli query approve --from jlgy05 --to jlgy06
{"from":"address1y4qgt89vfujkh7sc0wfnfqn2q5qmaerlunmnax","to":"address1zegv02xua87qsehmrmclrykqqdwnwkuf9ek5gt","qos":"800000","qscs":null}
qoscli query account jlgy05
{"type":"qos/types/QOSAccount","value":{"base_account":{"account_address":"address1y4qgt89vfujkh7sc0wfnfqn2q5qmaerlunmnax","public_key":null,"nonce":"0"},"qos":"1000000000","qscs":null}}

qoscli tx increase-approve --from jlgy05 --to jlgy06 --coins 200000qos
Password to sign with 'jlgy05':
{"check_tx":{"gasWanted":"100000","gasUsed":"6799"},"deliver_tx":{"gasWanted":"100000","gasUsed":"12500","tags":[{"key":"YWN0aW9u","value":"ZGVjcmVhc2UtYXBwcm92ZQ=="},{"key":"YXBwcm92ZS1mcm9t","value":"YWRkcmVzczF5NHFndDg5dmZ1amtoN3NjMHdmbmZxbjJxNXFtYWVybHVubW5heA=="},{"key":"YXBwcm92ZS10bw==","value":"YWRkcmVzczF6ZWd2MDJ4dWE4N3FzZWhtcm1jbHJ5a3FxZHdud2t1ZjllazVndA=="}]},"hash":"C0663DC38E55504CC14C3F12EC1355E001016B5055EFC308A4FC624FFB53AC6C","height":"784112"}

qoscli query approve --from jlgy05 --to jlgy06
{"from":"address1y4qgt89vfujkh7sc0wfnfqn2q5qmaerlunmnax","to":"address1zegv02xua87qsehmrmclrykqqdwnwkuf9ek5gt","qos":"1000000","qscs":null}
qoscli query account jlgy05
{"type":"qos/types/QOSAccount","value":{"base_account":{"account_address":"address1y4qgt89vfujkh7sc0wfnfqn2q5qmaerlunmnax","public_key":null,"nonce":"0"},"qos":"999998750","qscs":null}}
```
