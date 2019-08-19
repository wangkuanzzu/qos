# test case of qoscli tx create-approve 

> `qoscli tx create-approve` 创建授权

---

## 情景说明

1. 减少预授权后，查询预授权信息，验证前后的预授权信息变化是否正确。

## 测试命令

```bash
//减少预授权前查询授权信息
qoscli query approve --from jlgy05 --to jlgy06
//执行减少预授权交易
qoscli tx decrease-approve --from jlgy05 --to jlgy06 --coins 200000qos
//减少预授权后查询授权信息
qoscli query approve --from jlgy05 --to jlgy06
```

## 预测结果

```bash
1.授权的额度在完成减少预授权交易后减少。
```

## 测试结果

```bash
qoscli query approve --from jlgy05 --to jlgy06
{"from":"address1y4qgt89vfujkh7sc0wfnfqn2q5qmaerlunmnax","to":"address1zegv02xua87qsehmrmclrykqqdwnwkuf9ek5gt","qos":"1000000","qscs":null}

qoscli tx decrease-approve --from jlgy05 --to jlgy06 --coins 200000qos
Password to sign with 'jlgy05':
{"check_tx":{"gasWanted":"100000","gasUsed":"6799"},"deliver_tx":{"gasWanted":"100000","gasUsed":"11510","tags":[{"key":"YWN0aW9u","value":"ZGVjcmVhc2UtYXBwcm92ZQ=="},{"key":"YXBwcm92ZS1mcm9t","value":"YWRkcmVzczF5NHFndDg5dmZ1amtoN3NjMHdmbmZxbjJxNXFtYWVybHVubW5heA=="},{"key":"YXBwcm92ZS10bw==","value":"YWRkcmVzczF6ZWd2MDJ4dWE4N3FzZWhtcm1jbHJ5a3FxZHdud2t1ZjllazVndA=="}]},"hash":"C0663DC38E55504CC14C3F12EC1355E001016B5055EFC308A4FC624FFB53AC6C","height":"784112"}

qoscli query approve --from jlgy05 --to jlgy06
{"from":"address1y4qgt89vfujkh7sc0wfnfqn2q5qmaerlunmnax","to":"address1zegv02xua87qsehmrmclrykqqdwnwkuf9ek5gt","qos":"800000","qscs":null}
```
