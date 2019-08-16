# test case of qoscli tx submit-proposal

> `qoscli tx submit-proposal` 提交提议

---

## 情景说明

1. 提议类型为：Text，提交提议的质押（deposit）不小于规定的MinDeposit的三分之一。前提条件：qos网络上有账户abc；最小质押数为MinDeposit=100000000，三分之一设置为：33400000
2. 提议类型为：ParameterChange ，提交提议的proposer质押（deposit）超未超出本身账户拥有的qos数量，且足够支付gas。前提条件：对于账户abc的qos数量要大于50100000（质押+gasWanted）
3. 提议类型为：TaxUsage，提交提议的目标地址为guardian账号，社区费池提取比例介于0~1。
前提条件：qos网络中存在账户adas，且有guardian地址为address14k430znxuf83ruta9qq8hzyrxe8x7mkya4x60a
4. 当进行一笔交易（以提交提议为例）时候，所消耗的gas大于系统默认的100000.前提条件：账户abc中拥有的qos数量须大于750200000

## 测试命令

```bash
    //1
    // 等于最小质押
    qoscli tx submit-proposal --title "myproposal" --proposal-type "Text" --proposer abc --deposit 33400000 --description "the first proposal"

    // 大于最小质押
    qoscli tx submit-proposal --title "myproposal" --proposal-type "Text" --proposer abc --deposit 40000000 --description "the first proposal"

    //2
    qoscli tx submit-proposal --title "myproposal" --proposal-type "ParameterChange" --proposer abc --deposit 50000000 --description "the first proposal for update qos"  --params gov:min_deposit:1000

    //3
    qoscli tx submit-proposal --title 'update qos' --proposal-type TaxUsage --proposer adas --deposit 50000000 --description 'this is the description' --dest-address address14k430znxuf83ruta9qq8hzyrxe8x7mkya4x60a --percent 0.5

    //4 指定gasWanted大于gasUsed
    qoscli tx submit-proposal --title "myproposal" --proposal-type "Text" --proposer abc --deposit 750000000 --description "the first proposal" --max-gas 200000
```

## 测试结果

```bash
    //1
    qoscli tx submit-proposal --title "myproposal" --proposal-type "Text" --proposer abc --deposit 33400000 --description "the first proposal"
    Password to sign with 'abc':
    {"check_tx":{"gasWanted":"100000","gasUsed":"16093"},"deliver_tx":{"gasWanted":"100000","gasUsed":"61020","tags":[{"key":"YWN0aW9u","value":"c3VibWl0LXByb3Bvc2Fs"},{"key":"cHJvcG9zYWwtaWQ=","value":"MQ=="},{"key":"cHJvcG9zZXI=","value":"YWRkcmVzczEweHd4MDZnbnJ0M2RsejdoZnJ4NmE4d3gzZ3llZ2h4bTU0cnY3YQ=="},{"key":"cHJvcG9zYWwtdHlwZQ==","value":"VGV4dA=="}]},"hash":"2760C03AE0CF8C8603449F9F6E8DAB49BC39F1E4404F372E443B286AFA238951","height":"514967"}

    qoscli tx submit-proposal --title "myproposal" --proposal-type "Text" --proposer abc --deposit 40000000 --description "the first proposal"
    Password to sign with 'abc':
    {"check_tx":{"gasWanted":"100000","gasUsed":"16093"},"deliver_tx":{"gasWanted":"100000","gasUsed":"61020","tags":[{"key":"YWN0aW9u","value":"c3VibWl0LXByb3Bvc2Fs"},{"key":"cHJvcG9zYWwtaWQ=","value":"MQ=="},{"key":"cHJvcG9zZXI=","value":"YWRkcmVzczEweHd4MDZnbnJ0M2RsejdoZnJ4NmE4d3gzZ3llZ2h4bTU0cnY3YQ=="},{"key":"cHJvcG9zYWwtdHlwZQ==","value":"VGV4dA=="}]},"hash":"2760C03AE0CF8C8603449F9F6E8DAB49BC39F1E4404F372E443B286AFA238951","height":"514967"}

    //2
    qoscli tx submit-proposal --title "myproposal" --proposal-type "ParameterChange" --proposer abc --deposit 50000000 --description "the first proposal"  --params gov:min_deposit:1000
    Password to sign with 'abc':
    {"check_tx":{"gasWanted":"100000","gasUsed":"16093"},"deliver_tx":{"gasWanted":"100000","gasUsed":"61020","tags":[{"key":"YWN0aW9u","value":"c3VibWl0LXByb3Bvc2Fs"},{"key":"cHJvcG9zYWwtaWQ=","value":"MQ=="},{"key":"cHJvcG9zZXI=","value":"YWRkcmVzczEweHd4MDZnbnJ0M2RsejdoZnJ4NmE4d3gzZ3llZ2h4bTU0cnY3YQ=="},{"key":"cHJvcG9zYWwtdHlwZQ==","value":"VGV4dA=="}]},"hash":"2760C03AE0CF8C8603449F9F6E8DAB49BC39F1E4404F372E443B286AFA238951","height":"514967"}

    //3
    qoscli tx submit-proposal --title 'update qos' --proposal-type TaxUsage --proposer adas --deposit 50000000 --description 'this is the description' --dest-address address14k430znxuf83ruta9qq8hzyrxe8x7mkya4x60a --percent 0.5
    Password to sign with 'adas':
    {"check_tx":{"gasWanted":"100000","gasUsed":"17219"},"deliver_tx":{"gasWanted":"100000","gasUsed":"65240","tags":[{"key":"YWN0aW9u","value":"c3VibWl0LXByb3Bvc2Fs"},{"key":"cHJvcG9zYWwtaWQ=","value":"Ng=="},{"key":"cHJvcG9zZXI=","value":"YWRkcmVzczFsNmp1YXF5OWZrMGRwczBmbjVkY2c0ZnB5MzZ6bXJ5cDhteTR1eA=="},{"key":"cHJvcG9zYWwtdHlwZQ==","value":"VGF4VXNhZ2U="}]},"hash":"BC1015EE299133FA2EBC0851D507AAF0CBAEB9BD7E6B820763723A102EC71B3F","height":"517822"}

    //4 指定gasWanted大于gasUsed
    qoscli tx submit-proposal --title "myproposal" --proposal-type "Text" --proposer abc --deposit 750000000 --description "the first proposal" --max-gas 200000
    Password to sign with 'abc':
    {"check_tx":{"gasWanted":"200000","gasUsed":"16303"},"deliver_tx":{"gasWanted":"200000","gasUsed":"109230","tags":[{"key":"YWN0aW9u","value":"c3VibWl0LXByb3Bvc2Fs"},{"key":"cHJvcG9zYWwtaWQ=","value":"NQ=="},{"key":"cHJvcG9zZXI=","value":"YWRkcmVzczEweHd4MDZnbnJ0M2RsejdoZnJ4NmE4d3gzZ3llZ2h4bTU0cnY3YQ=="},{"key":"cHJvcG9zYWwtdHlwZQ==","value":"VGV4dA=="}]},"hash":"0F0995DB73BCB85975320F0AF10F4DBD06019F71C4B4DC147D4E47F4BD0CB749","height":"516054"}
```
