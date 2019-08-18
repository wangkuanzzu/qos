# test case of qoscli tx delegate

> `qoscli tx delegate` 委托

---

## 情景说明

1. 变更委托的账户没有代理验证节点。
2. 发起变更委托的账户在当前委托人验证节点绑定的tokens小于变更委托中指定的tokens数量。（这种情况下，是否可以使用发起变更委托账户的持有qos数量来填补不够的tokens？）

## 测试命令

```bash
    qoscli tx redelegate --from-owner address1nnvdqefva89xwppzs46vuskckr7klvzk8r5uaa --to-owner address14syavwd5lnedsk4hpztwwf46smqjt63z0wd0uz --delegator jlgy02 --tokens 100 --max-gas 200000

    qoscli tx redelegate --from-owner address1nnvdqefva89xwppzs46vuskckr7klvzk8r5uaa --to-owner address1f66wr25emjtp5urfcpd02epwg5ply3xzcv2u20 --delegator jlgy02 --tokens 2000000000 --max-gas 200000
```

## 测试结果

```bash
    qoscli tx redelegate --from-owner address1nnvdqefva89xwppzs46vuskckr7klvzk8r5uaa --to-owner address14syavwd5lnedsk4hpztwwf46smqjt63z0wd0uz --delegator jlgy02 --tokens 100 --max-gas 200000
    Password to sign with 'jlgy02':
    {"check_tx":{"code":1,"log":"{\"codespace\":\"sdk\",\"code\":1,\"message\":\"TxStd's ITx ValidateData error:  ERROR:\\nCodespace: stake\\nCode: 506\\nMessage: \\\"address14syavwd5lnedsk4hpztwwf46smqjt63z0wd0uz does't have validator.\\\"\\n\"}","gasWanted":"200000","gasUsed":"3375"},"deliver_tx":{},"hash":"9FB995ABEE88C4B8652B40029234DD4BA9BA3B4578036932E19C8D7BE430A735","height":"0"}
    ERROR: {"codespace":"sdk","code":1,"message":"TxStd's ITx ValidateData error:  ERROR:\nCodespace: stake\nCode: 506\nMessage: \"address14syavwd5lnedsk4hpztwwf46smqjt63z0wd0uz does't have validator.\"\n"}

    qoscli tx redelegate --from-owner address1nnvdqefva89xwppzs46vuskckr7klvzk8r5uaa --to-owner address1f66wr25emjtp5urfcpd02epwg5ply3xzcv2u20 --delegator jlgy02 --tokens 2000000000 --max-gas 200000
    Password to sign with 'jlgy02':
    {"check_tx":{"code":1,"log":"{\"codespace\":\"sdk\",\"code\":1,\"message\":\"TxStd's ITx ValidateData error:  ERROR:\\nCodespace: stake\\nCode: 501\\nMessage: \\\"delegator does't have enough amount of QOS\\\"\\n\"}","gasWanted":"200000","gasUsed":"6212"},"deliver_tx":{},"hash":"227211E9BC802CE1D4C83A89EDF7DC29B6F5322DE291ADCFF94A9D7DE2463E80","height":"0"}
    ERROR: {"codespace":"sdk","code":1,"message":"TxStd's ITx ValidateData error:  ERROR:\nCodespace: stake\nCode: 501\nMessage: \"delegator does't have enough amount of QOS\"\n"}
```
