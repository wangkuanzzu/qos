# test case of qoscli tx unbond

> `qoscli tx unbond` 解除委托

---

## 情景说明

1. 解除委托时选择的代理验证节点owner错误，委托人未曾向该验证节点进行过委托。前提条件：账户abc在账户address1f66wr25emjtp5urfcpd02epwg5ply3xzcv2u20所创建的验证节点没有进行过委托。

2. 解除委托时解绑的tokens大于委托人在该验证节点委托的tokens数量。前提条件：账户def在账户abc创建的验证节点进行过amount为50000qos的委托。

## 测试命令

```bash
    qoscli tx unbond --owner address1f66wr25emjtp5urfcpd02epwg5ply3xzcv2u20 --delegator abc --tokens 100

    qoscli tx unbond --owner abc --delegator def --tokens 60000
```

## 测试结果

```bash
    qoscli tx unbond --owner address1f66wr25emjtp5urfcpd02epwg5ply3xzcv2u20 --delegator abc --tokens 100
    Password to sign with 'abc':
    {"check_tx":{"code":1,"log":"{\"codespace\":\"sdk\",\"code\":1,\"message\":\"TxStd's ITx ValidateData error:  ERROR:\\nCodespace: stake\\nCode: 501\\nMessage: \\\"delegator not delegate the owner's validator\\\"\\n\"}","gasWanted":"100000","gasUsed":"3687"},"deliver_tx":{},"hash":"2C0CC1B10ACB24FBAC3CD25F40D24448512EC3397770064CEE5292067858755C","height":"0"}
    ERROR: {"codespace":"sdk","code":1,"message":"TxStd's ITx ValidateData error:  ERROR:\nCodespace: stake\nCode: 501\nMessage: \"delegator not delegate the owner's validator\"\n"}

    qoscli tx unbond --owner abc --delegator def --tokens 60000
    Password to sign with 'def':
    {"check_tx":{"code":1,"log":"{\"codespace\":\"sdk\",\"code\":1,\"message\":\"TxStd's ITx ValidateData error:  ERROR:\\nCodespace: stake\\nCode: 501\\nMessage: \\\"delegator does't have enough amount of QOS\\\"\\n\"}","gasWanted":"100000","gasUsed":"3501"},"deliver_tx":{},"hash":"F1D406814264986735DE48CCF2447DD5CAEBD446BE052EA0A8AC95680FA80AB7","height":"0"}
    ERROR: {"codespace":"sdk","code":1,"message":"TxStd's ITx ValidateData error:  ERROR:\nCodespace: stake\nCode: 501\nMessage: \"delegator does't have enough amount of QOS\"\n"}
```
