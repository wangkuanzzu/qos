# test case of qoscli tx submit-proposal

> `qoscli tx submit-proposal` 提交提议

---

## 情景说明

1. 提议类型为：Text，提交提议的质押（deposit）小于规定的MinDeposit的三分之一。前提条件：qos网络上有账户；最小质押数为MinDeposit=100000000，三分之一设置为：33400000
2. 提议类型为：ParameterChange ，提交提议的proposer质押（deposit）超出本身账户拥有的qos数量。前提条件：在qos网络中存在账户abc，且abc账户中只有20000000qos。
3. 提议类型为：TaxUsage，提交提议的目标地址非guardian账号，社区费池提取比例介于0~1。前提条件：qos网络中存在账户，且有一定量qos，足以进行提交提议（大于提交提议最小质押数量）。
4. 提议类型为：TaxUsage，提交提议的目标地址为guardian账号，社区费池提取比例低于0，或高于1。前提条件：在QOS网络中存在guardian账号abc

## 测试命令

```bash
    //1
    qoscli tx submit-proposal --title "myproposal" --proposal-type "Text" --proposer abc --deposit 33300000 --description "the first proposal"

    //2
    qoscli tx submit-proposal --title "myproposal" --proposal-type "ParameterChange" --proposer abc --deposit 50000000 --description "the first proposal"  --params gov:min_deposit:1000

    //3
    qoscli tx submit-proposal --title 'update qos' --proposal-type TaxUsage --proposer adas --deposit 50000000 --description 'this is the description' --dest-address abc --percent 0.5

    //4
    qoscli tx submit-proposal --title 'update qos' --proposal-type TaxUsage --proposer adas --deposit 50000000 --description 'this is the description' --dest-address address14k430znxuf83ruta9qq8hzyrxe8x7mkya4x60a --percent -0.5

    qoscli tx submit-proposal --title 'update qos' --proposal-type TaxUsage --proposer adas --deposit 50000000 --description 'this is the description' --dest-address address14k430znxuf83ruta9qq8hzyrxe8x7mkya4x60a --percent 1.5

    //5 gas不够使用时候
    qoscli tx submit-proposal --title "myproposal" --proposal-type "Text" --proposer abc --deposit 750000000 --description "the first proposal"
```

## 测试结果

```bash
    //1
    qoscli tx submit-proposal --title "myproposal" --proposal-type "Text" --proposer abc --deposit 33300000 --description "the first proposal"
    Password to sign with 'abc':
    {"check_tx":{"code":1,"log":"{\"codespace\":\"sdk\",\"code\":1,\"message\":\"TxStd's ITx ValidateData error:  ERROR:\\nCodespace: gov\\nCode: 601\\nMessage: \\\"initial deposit is too small\\\"\\n\"}","gasWanted":"100000","gasUsed":"9423"},"deliver_tx":{},"hash":"9D7D6AAD8EBE6D3BC7DF77DF9D7AED62331EBF405D3EE0B957B3F667B0C9C55A","height":"0"}
    ERROR: {"codespace":"sdk","code":1,"message":"TxStd's ITx ValidateData error:  ERROR:\nCodespace: gov\nCode: 601\nMessage: \"initial deposit is too small\"\n"}

    //2
    qoscli tx submit-proposal --title "myproposal" --proposal-type "ParameterChange" --proposer abc --deposit 50000000 --description "the first proposal"  --params gov:min_deposit:1000
    Password to sign with 'abc':
    {"check_tx":{"code":1,"log":"{\"codespace\":\"sdk\",\"code\":1,\"message\":\"TxStd's ITx ValidateData error:  ERROR:\\nCodespace: gov\\nCode: 601\\nMessage: \\\"proposer has no enough qos\\\"\\n\"}","gasWanted":"100000","gasUsed":"10654"},"deliver_tx":{},"hash":"2C888042708638EBF0CABE8A81525A9BB1DBD7A764778B93292C37F13881DE40","height":"0"}
    ERROR: {"codespace":"sdk","code":1,"message":"TxStd's ITx ValidateData error:  ERROR:\nCodespace: gov\nCode: 601\nMessage: \"proposer has no enough qos\"\n"}

    //3
    qoscli tx submit-proposal --title 'update qos' --proposal-type TaxUsage --proposer adas --deposit 50000000 --description 'this is the description' --dest-address abc --percent 0.5
    Password to sign with 'adas':
    {"check_tx":{"code":1,"log":"{\"codespace\":\"sdk\",\"code\":1,\"message\":\"TxStd's ITx ValidateData error:  ERROR:\\nCodespace: gov\\nCode: 601\\nMessage: \\\"DestAddress must be guardian\\\"\\n\"}","gasWanted":"100000","gasUsed":"11543"},"deliver_tx":{},"hash":"93AFC6A01965089D359C7E7AF1DAE7B85AEB093EDD73A772C2E712FC3B460402","height":"0"}
    ERROR: {"codespace":"sdk","code":1,"message":"TxStd's ITx ValidateData error:  ERROR:\nCodespace: gov\nCode: 601\nMessage: \"DestAddress must be guardian\"\n"}

    //4
    qoscli tx submit-proposal --title 'update qos' --proposal-type TaxUsage --proposer adas --deposit 50000000 --description 'this is the description' --dest-address address14k430znxuf83ruta9qq8hzyrxe8x7mkya4x60a --percent -0.5
    null
    ERROR: deposit must be positive

    qoscli tx submit-proposal --title 'update qos' --proposal-type TaxUsage --proposer adas --deposit 50000000 --description 'this is the description' --dest-address address14k430znxuf83ruta9qq8hzyrxe8x7mkya4x60a --percent 1.5
    Password to sign with 'adas':
    {"check_tx":{"gasWanted":"100000","gasUsed":"17429"},"deliver_tx":{"gasWanted":"100000","gasUsed":"65380","tags":[{"key":"YWN0aW9u","value":"c3VibWl0LXByb3Bvc2Fs"},{"key":"cHJvcG9zYWwtaWQ=","value":"Nw=="},{"key":"cHJvcG9zZXI=","value":"YWRkcmVzczFsNmp1YXF5OWZrMGRwczBmbjVkY2c0ZnB5MzZ6bXJ5cDhteTR1eA=="},{"key":"cHJvcG9zYWwtdHlwZQ==","value":"VGF4VXNhZ2U="}]},"hash":"0392E7A11E656334F275AB0563FE6CD7DE571411B7ECA34875420BA0D4674A9C","height":"517914"}

    //5 gas不够使用时候
    qoscli tx submit-proposal --title "myproposal" --proposal-type "Text" --proposer abc --deposit 750000000 --description "the first proposal"
    Password to sign with 'abc':
    {"check_tx":{"gasWanted":"100000","gasUsed":"16303"},"deliver_tx":{"code":12,"log":"{\"codespace\":\"sdk\",\"code\":12,\"message\":\"deliverTxStd out of gas\"}","gasWanted":"100000","gasUsed":"105088","codespace":"sdk"},"hash":"570700D5062BACFEA902FAF8EDC53BBC45363AACAB11889AB60EF4AB958F205F","height":"516034"}
    ERROR: {"codespace":"sdk","code":12,"message":"deliverTxStd out of gas"}
```
