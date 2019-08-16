# test case of qoscli tx vote

> `qoscli tx vote` 提议投票

---

## 情景说明

* 情景1:
  
    投票的提议编号不存在，提议编号传入错误。前提条件：在QOS网络中存在提议，且数量小于100。对提议的状态没有要求。

* 情景2：

    投票的选项错误，不在指定范围：Yes,Abstain,No,NoWithVeto。前提条件：QOS网络中存在有提议，对提议的状态没有要求。

* 情景3：

    投票的提议编号存在，但是提议的状态不是voting。前提条件：提议编号为4状态不处于voting，而是其他状态例如：deposit

## 测试命令

```bash
    //情景1
    qoscli tx vote --proposal-id 0 --voter abc --option Yes

    qoscli tx vote --proposal-id 100 --voter abc --option Yes

    //情景2
    qoscli tx vote --proposal-id 5 --voter abc --option not

    qoscli tx vote --proposal-id 5 --voter abc --option ok

    //情景3
    qoscli tx vote --proposal-id 4 --voter abc --option No 
```

## 测试结果

```bash
    //情景1
    qoscli tx vote --proposal-id 0 --voter abc --option Yes
    null
    ERROR: proposal-id must be positive

    qoscli tx vote --proposal-id 100 --voter abc --option Yes
    Password to sign with 'abc':
    {"check_tx":{"code":1,"log":"{\"codespace\":\"sdk\",\"code\":1,\"message\":\"TxStd's ITx ValidateData error:  ERROR:\\nCodespace: gov\\nCode: 603\\nMessage: \\\"unknown proposal 100\\\"\\n\"}","gasWanted":"100000","gasUsed":"1000"},"deliver_tx":{},"hash":"645E8751CE78142AE15CE921FCB3742A1AD60C78238713CC2AE6BB3262B39635","height":"0"}
    ERROR: {"codespace":"sdk","code":1,"message":"TxStd's ITx ValidateData error:  ERROR:\nCodespace: gov\nCode: 603\nMessage: \"unknown proposal 100\"\n"}

    //情景2
    qoscli tx vote --proposal-id 5 --voter abc --option not
    null
    ERROR: invalid option

    qoscli tx vote --proposal-id 5 --voter abc --option ok
    null
    ERROR: invalid option

    //情景3
    qoscli tx vote --proposal-id 4 --voter abc --option No
    Password to sign with 'abc':
    {"check_tx":{"code":1,"log":"{\"codespace\":\"sdk\",\"code\":1,\"message\":\"TxStd's ITx ValidateData error:  ERROR:\\nCodespace: gov\\nCode: 607\\nMessage: \\\"wrong status of proposal 4\\\"\\n\"}","gasWanted":"100000","gasUsed":"1318"},"deliver_tx":{},"hash":"A1E3A2B437A87BEF8BB940A5F02B4731816D9B569C7D79AEA67A9080003CBB00","height":"0"}
    ERROR: {"codespace":"sdk","code":1,"message":"TxStd's ITx ValidateData error:  ERROR:\nCodespace: gov\nCode: 607\nMessage: \"wrong status of proposal 4\"\n"}
```
