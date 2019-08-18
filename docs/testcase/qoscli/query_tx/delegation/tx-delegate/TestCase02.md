# test case of qoscli tx delegate

> `qoscli tx delegate` 委托

---

## 情景说明

1. 委托的账户没有代理验证节点。前提条件：在qos网络中hij不是代理验证节点账户
2. 发起委托的账户委托的tokens大于自身所持有的qos数量。前提条件：网络中abc是代理验证节点账户，def账户中持有的qos数量小于50000000000

## 测试命令

```bash
    qoscli tx delegate --owner hij --delegator def --tokens 50000

    qoscli tx delegate --owner abc --delegator def --tokens 50000000000
```

## 测试结果

```bash
    qoscli tx delegate --owner hij --delegator def --tokens 50000
    Password to sign with 'def':
    {"check_tx":{"code":1,"log":"{\"codespace\":\"sdk\",\"code\":1,\"message\":\"TxStd's ITx ValidateData error:  ERROR:\\nCodespace: stake\\nCode: 506\\nMessage: \\\"address1jxjl9mcghl60s6lu5s2mrtrppf5t2h5mjdum20 does't have validator.\\\"\\n\"}","gasWanted":"100000","gasUsed":"1000"},"deliver_tx":{},"hash":"5201428DBF952D55C678CCD6A95A3BBA1B4093BFE1208D867DBD0D8CAC824408","height":"0"}
    ERROR: {"codespace":"sdk","code":1,"message":"TxStd's ITx ValidateData error:  ERROR:\nCodespace: stake\nCode: 506\nMessage: \"address1jxjl9mcghl60s6lu5s2mrtrppf5t2h5mjdum20 does't have validator.\"\n"}

    qoscli tx delegate --owner abc --delegator def --tokens 50000000000
    Password to sign with 'def':
    {"check_tx":{"code":1,"log":"{\"codespace\":\"sdk\",\"code\":1,\"message\":\"TxStd's ITx ValidateData error:  ERROR:\\nCodespace: stake\\nCode: 503\\nMessage: \\\"No enough QOS in account: address16xd8tzrm6f4jfrmtvy6sjafuy80lgj0gjwu8zt\\\"\\n\"}","gasWanted":"100000","gasUsed":"3600"},"deliver_tx":{},"hash":"8F5F68D351D4D1898F697A4E39FDB5EB1D3A28E161DFE9E5A9C00D0158CABC56","height":"0"}
    ERROR: {"codespace":"sdk","code":1,"message":"TxStd's ITx ValidateData error:  ERROR:\nCodespace: stake\nCode: 503\nMessage: \"No enough QOS in account: address16xd8tzrm6f4jfrmtvy6sjafuy80lgj0gjwu8zt\"\n"}
```
