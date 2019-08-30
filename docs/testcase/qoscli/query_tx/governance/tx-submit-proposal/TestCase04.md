# test case of qoscli tx submit-proposal

> `qoscli tx submit-proposal` 提交提议

---

## 情景说明

1. 提议成功后，是否可以查询到。
2. 校验提议的状态，质押等其他信息是否正确。
3. 提议的状态在创建提议时候会依据抵押数额，直接进入某一状态，验证不同数额下状态是否正确。
4. 提议的tx操作，无基础的gas费用消耗，但是有访问存储消耗的gas。

## 测试命令

```bash
    //查询目前存在的提议
    qoscli query proposals
    //创建提议
    qoscli tx submit-proposal --title "myproposal" --proposal-type "Text" --proposer jlgy07 --deposit 40000000 --description "the first proposal"
    //提议创建后查询
    qoscli query proposals

    //验证信息是否正确，阶段
    qoscli query proposal <新增提议的编号>

    //创建提议，直接进入voting阶段
    qoscli tx submit-proposal --title "myproposal" --proposal-type "Text" --proposer jlgy07 --deposit 140000000 --description "the first proposal"

```

## 预测结果

```bash
1.提议创建后，查询所有提议时候可以查到
2.质押tokens较低时候提议阶段进入deposit阶段。
3.质押tokens超过限定值时候提议阶段进入voting阶段。
```

## 测试结果

```bash
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query proposals
ERROR: no matching proposals found
[vagrant@vagrant-192-168-1-200 ~]$ qoscli tx submit-proposal --title "myproposal" --proposal-type "Text" --proposer acc1 --deposit 8 --description "the first proposal"
Password to sign with 'acc1':
{"check_tx":{"gasWanted":"9223372036854775807","gasUsed":"16180","events":[]},"deliver_tx":{"gasWanted":"9223372036854775807","gasUsed":"82600","events":[{"type":"submit-proposal","attributes":[{"key":"cHJvcG9zYWwtaWQ=","value":"MQ=="},{"key":"cHJvcG9zZXI=","value":"YWRkcmVzczFmOWtrMzVwZ250ZnlnMGp2czdlODRtdWVtbXRuZzA1d2tlOTBoOA=="},{"key":"ZGVwb3NpdG9y","value":"YWRkcmVzczFmOWtrMzVwZ250ZnlnMGp2czdlODRtdWVtbXRuZzA1d2tlOTBoOA=="},{"key":"cHJvcG9zYWwtdHlwZQ==","value":"VGV4dA=="}]},{"type":"message","attributes":[{"key":"bW9kdWxl","value":"Z292"},{"key":"Z2FzLnBheWVy","value":"YWRkcmVzczFmOWtrMzVwZ250ZnlnMGp2czdlODRtdWVtbXRuZzA1d2tlOTBoOA=="}]}]},"hash":"54BFC7164D48B3B64BFB825685A51334892713B576F388D2BE91170D09066E7F","height":"1057"}
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query proposals
[{"proposal_content":{"type":"gov/TextProposal","value":{"title":"myproposal","description":"the first proposal","deposit":"6"}},"proposal_id":"2","proposal_status":"Deposit","final_tally_result":{"yes":"0","abstain":"0","no":"0","no_with_veto":"0"},"submit_time":"2019-08-28T11:52:36.881898341Z","deposit_end_time":"2019-08-28T11:54:36.881898341Z","total_deposit":"6","voting_start_time":"0001-01-01T00:00:00Z","voting_start_height":"0","voting_end_time":"0001-01-01T00:00:00Z"}]
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query proposal 1

{"proposal_content":{"type":"gov/TextProposal","value":{"title":"myproposal","description":"the first proposal","deposit":"6"}},"proposal_id":"2","proposal_status":"Deposit","final_tally_result":{"yes":"0","abstain":"0","no":"0","no_with_veto":"0"},"submit_time":"2019-08-28T11:52:36.881898341Z","deposit_end_time":"2019-08-28T11:54:36.881898341Z","total_deposit":"6","voting_start_time":"0001-01-01T00:00:00Z","voting_start_height":"0","voting_end_time":"0001-01-01T00:00:00Z"}

[vagrant@vagrant-192-168-1-200 ~]$ qoscli tx submit-proposal --title "myproposal" --proposal-type "Text" --proposer acc1 --deposit 60 --description "the first proposal"
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query proposal 2
{"proposal_content":{"type":"gov/TextProposal","value":{"title":"myproposal","description":"the first proposal","deposit":"60"}},"proposal_id":"1","proposal_status":"Voting","final_tally_result":{"yes":"0","abstain":"0","no":"0","no_with_veto":"0"},"submit_time":"2019-08-28T11:50:21.28546855Z","deposit_end_time":"2019-08-28T11:52:21.28546855Z","total_deposit":"60","voting_start_time":"2019-08-28T11:50:21.28546855Z","voting_start_height":"1057","voting_end_time":"2019-08-28T11:52:21.28546855Z"}

```
