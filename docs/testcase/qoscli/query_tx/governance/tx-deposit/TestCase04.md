# test case of qoscli tx deposit

> `qoscli tx deposit` 提议抵押

---

## 情景说明

1. 验证质押前后账户余额的变化，扣除抵押tokens，扣除gas费用。
2. 验证质押前后提议的抵押token变化。
3. 校验提议所处的阶段是否随着抵押的token数量增加而变更。
4. 抵押的tx操作，无基础的gas费用消耗，但是有访问存储消耗的gas。

## 测试命令

```bash
    //验证抵押成功
    qoscli query account acc0
    qoscli query proposal 1

    qoscli tx deposit --proposal-id 1 --depositor acc0 --amount 1

    qoscli query account acc1
    qoscli query proposal 1


    //验证提议阶段变化
    qoscli tx deposit --proposal-id 1 --depositor acc0 --amount 1
    ...
    qoscli tx deposit --proposal-id 1 --depositor acc0 --amount 1

    qoscli query proposal 1
```

## 预测结果

```text
1.抵押的账户余额减少，10000和消耗的gas费用。
2.提议的抵押tokens增加。10000
3.提议的阶段会在不断增加抵押tokens后，达到最小抵押限制，进入voting阶段。
```

## 测试结果

```bash
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query proposal 1
{"proposal_content":{"type":"gov/TextProposal","value":{"title":"myproposal","description":"the first proposal","deposit":"7"}},"proposal_id":"1","proposal_status":"Deposit","final_tally_result":{"yes":"0","abstain":"0","no":"0","no_with_veto":"0"},"submit_time":"2019-08-29T02:41:47.604696459Z","deposit_end_time":"2019-08-29T02:51:47.604696459Z","total_deposit":"7","voting_start_time":"0001-01-01T00:00:00Z","voting_start_height":"0","voting_end_time":"0001-01-01T00:00:00Z"}
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query account acc0
{"type":"qos/types/QOSAccount","value":{"base_account":{"account_address":"address1n64h6prxz6ld5vl8d0rrzsny25nq7vgcnrhu98","public_key":{"type":"tendermint/PubKeyEd25519","value":"m4lqygnU2mG19Fpf3vj2K618G2e2WMwtxu6GANARIVY="},"nonce":"1"},"qos":"5000000000000","qscs":null}}
[vagrant@vagrant-192-168-1-200 ~]$ qoscli tx deposit --proposal-id 1 --depositor acc0 --amount 1
Password to sign with 'acc0':
{"check_tx":{"gasWanted":"9223372036854775807","gasUsed":"8324","events":[]},"deliver_tx":{"gasWanted":"9223372036854775807","gasUsed":"33500","events":[{"type":"deposit-proposal","attributes":[{"key":"cHJvcG9zYWwtaWQ=","value":"MQ=="},{"key":"ZGVwb3NpdG9y","value":"YWRkcmVzczFuNjRoNnByeHo2bGQ1dmw4ZDBycnpzbnkyNW5xN3ZnY25yaHU5OA=="}]},{"type":"message","attributes":[{"key":"bW9kdWxl","value":"Z292"},{"key":"Z2FzLnBheWVy","value":"YWRkcmVzczFuNjRoNnByeHo2bGQ1dmw4ZDBycnpzbnkyNW5xN3ZnY25yaHU5OA=="}]}]},"hash":"D7A7BA97D9F0C4C54B478DB15D37831B392318C4284F6EF08FA4F3888AB5A7D0","height":"182"}
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query proposal 1
{"proposal_content":{"type":"gov/TextProposal","value":{"title":"myproposal","description":"the first proposal","deposit":"7"}},"proposal_id":"1","proposal_status":"Deposit","final_tally_result":{"yes":"0","abstain":"0","no":"0","no_with_veto":"0"},"submit_time":"2019-08-29T02:41:47.604696459Z","deposit_end_time":"2019-08-29T02:51:47.604696459Z","total_deposit":"8","voting_start_time":"0001-01-01T00:00:00Z","voting_start_height":"0","voting_end_time":"0001-01-01T00:00:00Z"}
[vagrant@vagrant-192-168-1-200 ~]$ qoscli tx deposit --proposal-id 1 --depositor acc0 --amount 1
Password to sign with 'acc0':
{"check_tx":{"gasWanted":"9223372036854775807","gasUsed":"8324","events":[]},"deliver_tx":{"gasWanted":"9223372036854775807","gasUsed":"33600","events":[{"type":"deposit-proposal","attributes":[{"key":"cHJvcG9zYWwtaWQ=","value":"MQ=="},{"key":"ZGVwb3NpdG9y","value":"YWRkcmVzczFuNjRoNnByeHo2bGQ1dmw4ZDBycnpzbnkyNW5xN3ZnY25yaHU5OA=="}]},{"type":"message","attributes":[{"key":"bW9kdWxl","value":"Z292"},{"key":"Z2FzLnBheWVy","value":"YWRkcmVzczFuNjRoNnByeHo2bGQ1dmw4ZDBycnpzbnkyNW5xN3ZnY25yaHU5OA=="}]}]},"hash":"BFDF41DB19E935CE908DE2E555ED1016C1528E948A2B7E34AAA52BDBF5702F00","height":"187"}
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query proposal 1
{"proposal_content":{"type":"gov/TextProposal","value":{"title":"myproposal","description":"the first proposal","deposit":"7"}},"proposal_id":"1","proposal_status":"Deposit","final_tally_result":{"yes":"0","abstain":"0","no":"0","no_with_veto":"0"},"submit_time":"2019-08-29T02:41:47.604696459Z","deposit_end_time":"2019-08-29T02:51:47.604696459Z","total_deposit":"9","voting_start_time":"0001-01-01T00:00:00Z","voting_start_height":"0","voting_end_time":"0001-01-01T00:00:00Z"}
[vagrant@vagrant-192-168-1-200 ~]$ qoscli tx deposit --proposal-id 1 --depositor acc0 --amount 1
Password to sign with 'acc0':
{"check_tx":{"gasWanted":"9223372036854775807","gasUsed":"8324","events":[]},"deliver_tx":{"gasWanted":"9223372036854775807","gasUsed":"55400","events":[{"type":"deposit-proposal","attributes":[{"key":"cHJvcG9zYWwtaWQ=","value":"MQ=="},{"key":"ZGVwb3NpdG9y","value":"YWRkcmVzczFuNjRoNnByeHo2bGQ1dmw4ZDBycnpzbnkyNW5xN3ZnY25yaHU5OA=="}]},{"type":"message","attributes":[{"key":"bW9kdWxl","value":"Z292"},{"key":"Z2FzLnBheWVy","value":"YWRkcmVzczFuNjRoNnByeHo2bGQ1dmw4ZDBycnpzbnkyNW5xN3ZnY25yaHU5OA=="}]}]},"hash":"B5FF2684159BED297AC44A48EA266ACA4E39AAED135B6988B11E377C5E91FD18","height":"190"}
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query proposal 1
{"proposal_content":{"type":"gov/TextProposal","value":{"title":"myproposal","description":"the first proposal","deposit":"7"}},"proposal_id":"1","proposal_status":"Voting","final_tally_result":{"yes":"0","abstain":"0","no":"0","no_with_veto":"0"},"submit_time":"2019-08-29T02:41:47.604696459Z","deposit_end_time":"2019-08-29T02:51:47.604696459Z","total_deposit":"10","voting_start_time":"2019-08-29T02:44:03.131511973Z","voting_start_height":"190","voting_end_time":"2019-08-29T02:54:03.131511973Z"}
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query account acc0
{"type":"qos/types/QOSAccount","value":{"base_account":{"account_address":"address1n64h6prxz6ld5vl8d0rrzsny25nq7vgcnrhu98","public_key":{"type":"tendermint/PubKeyEd25519","value":"m4lqygnU2mG19Fpf3vj2K618G2e2WMwtxu6GANARIVY="},"nonce":"4"},"qos":"4999999998772","qscs":null}}

```
