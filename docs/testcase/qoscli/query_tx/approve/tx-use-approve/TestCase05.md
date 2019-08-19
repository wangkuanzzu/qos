# test case of qoscli tx create-approve

> `qoscli tx create-approve` 创建授权

---

## 情景说明

1. 使用预授权后，查询预授权信息，验证前后的预授权信息变化（授权数额）是否正确。
2. 使用预授权后，查询授权人账户信息，验证前后的账户余额信息变化是否正确。
3. 创建授权后，被授权账户可以进行交易，交易的数额在授权范围内，验证是否可以正常交易。
4. 交易的数额超出预授权的范围，验证是否可以进行交易。

## 测试命令

```bash
//使用授权前查询账户信息和授权信息
qoscli query account jlgy05
qoscli query account jlgy06
qoscli query approve --from jlgy05 --to jlgy06

//执行使用授权交易
qoscli tx use-approve --coins 500000qos --from jlgy05 --to jlgy06

//使用授权后查询账户信息和授权信息
qoscli query approve --from jlgy05 --to jlgy06
qoscli query account jlgy05
qoscli query account jlgy06

//超出使用授权额度
qoscli tx use-approve --coins 600000qos --from jlgy05 --to jlgy06
```

## 预测结果

```bash
1.使用授权时候，首先授权人账户扣除被授权人使用的授权数额。
2.被授权人账户增加使用的授权数额，但是在执行使用授权交易，需要支付gas费用。
3.对应的授权信息数额减少被授权人使用的数额。
4.超出使用授权额度，系统会报错。
```

## 测试结果

```bash
qoscli query account jlgy05
{"type":"qos/types/QOSAccount","value":{"base_account":{"account_address":"address1y4qgt89vfujkh7sc0wfnfqn2q5qmaerlunmnax","public_key":{"type":"tendermint/PubKeyEd25519","value":"x54mlQpo3XDIfPevRQEqHgVS9/6cG85a+VjUDbPteEU="},"nonce":"3"},"qos":"1000000000","qscs":null}}
qoscli query account jlgy06
{"type":"qos/types/QOSAccount","value":{"base_account":{"account_address":"address1zegv02xua87qsehmrmclrykqqdwnwkuf9ek5gt","public_key":null,"nonce":"0"},"qos":"1000000000","qscs":null}}
qoscli query approve --from jlgy05 --to jlgy06
{"from":"address1y4qgt89vfujkh7sc0wfnfqn2q5qmaerlunmnax","to":"address1zegv02xua87qsehmrmclrykqqdwnwkuf9ek5gt","qos":"1000000","qscs":null}

qoscli tx use-approve --coins 500000qos --from jlgy05 --to jlgy06
Password to sign with 'jlgy06':
{"check_tx":{"gasWanted":"100000","gasUsed":"7952"},"deliver_tx":{"gasWanted":"100000","gasUsed":"23980","tags":[{"key":"YWN0aW9u","value":"dXNlLWFwcHJvdmU="},{"key":"YXBwcm92ZS1mcm9t","value":"YWRkcmVzczF5NHFndDg5dmZ1amtoN3NjMHdmbmZxbjJxNXFtYWVybHVubW5heA=="},{"key":"YXBwcm92ZS10bw==","value":"YWRkcmVzczF6ZWd2MDJ4dWE4N3FzZWhtcm1jbHJ5a3FxZHdud2t1ZjllazVndA=="}]},"hash":"74F1765BF9E51877B4D4A9D201E353EEFDA53AE3E8BC73807AC96BA55A58F712","height":"784405"}

qoscli query approve --from jlgy05 --to jlgy06
{"from":"address1y4qgt89vfujkh7sc0wfnfqn2q5qmaerlunmnax","to":"address1zegv02xua87qsehmrmclrykqqdwnwkuf9ek5gt","qos":"500000","qscs":null}
qoscli query account jlgy05
{"type":"qos/types/QOSAccount","value":{"base_account":{"account_address":"address1y4qgt89vfujkh7sc0wfnfqn2q5qmaerlunmnax","public_key":{"type":"tendermint/PubKeyEd25519","value":"x54mlQpo3XDIfPevRQEqHgVS9/6cG85a+VjUDbPteEU="},"nonce":"3"},"qos":"999500000","qscs":null}}
qoscli query account jlgy06
{"type":"qos/types/QOSAccount","value":{"base_account":{"account_address":"address1zegv02xua87qsehmrmclrykqqdwnwkuf9ek5gt","public_key":{"type":"tendermint/PubKeyEd25519","value":"5QtNE7Tzr4nHU2MDq2D6T9iXTTOxyzbWxxJJDWBzjtw="},"nonce":"1"},"qos":"1000497602","qscs":null}}

qoscli tx use-approve --coins 600000qos --from jlgy05 --to jlgy06
Password to sign with 'jlgy06':
{"check_tx":{"code":1,"log":"{\"codespace\":\"sdk\",\"code\":1,\"message\":\"TxStd's ITx ValidateData error:  ERROR:\\nCodespace: approve\\nCode: 106\\nMessage: \\\"approve not enough\\\"\\n\"}","gasWanted":"100000","gasUsed":"1141"},"deliver_tx":{},"hash":"0E080482C9C716E9412AFA4ADB62925A98DC8E9D2B65D1099DFFD29065AD33C2","height":"0"}
ERROR: {"codespace":"sdk","code":1,"message":"TxStd's ITx ValidateData error:  ERROR:\nCodespace: approve\nCode: 106\nMessage: \"approve not enough\"\n"}
```
