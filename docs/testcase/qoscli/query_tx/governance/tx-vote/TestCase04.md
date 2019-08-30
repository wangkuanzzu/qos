# test case of qoscli tx vote

> `qoscli tx vote` 提议投票

---

## 情景说明

1. 对某一提议完成投票后，可以查询到，验证投票操作成功。
2. 在投票阶段可以多次进行投票，记录最后一次投票。
3. 投票阶段完成后，提议的处理方式。通过，不通过，严重反对。校验质押人质押的tokens去向。
4. 投票的tx操作，无基础的gas费用消耗，但是有访问存储消耗的gas。

## 测试命令

```bash
    //查询 账户abc在提议编号为1的投票情况
    qoscli query vote 1 acc2

    //首次投票
    qoscli tx vote --proposal-id 1 --voter acc2 --option yes

    //查询 账户abc在提议编号为1的投票情况
    qoscli query vote 1 acc2

    //再次投票
    qoscli tx vote --proposal-id 1 --voter acc2 --option no

    //查询 账户abc在提议编号为1的投票情况
    qoscli query vote 1 acc2
    qoscli 
```

## 预测结果

```bash
1.初始查询投票没有结果
2.首次投票后，查询投票结果为yes
3.再次投票后，查询投票结果为no
4.投票前后，投票人账户的数额会减少，vote的tx会消耗访问存储的gas费用。
```

## 测试结果

```bash
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query account acc2
{"type":"qos/types/QOSAccount","value":{"base_account":{"account_address":"address1mxyz3rld8q25hx3m63vr4dt90z77n62684c737","public_key":{"type":"tendermint/PubKeyEd25519","value":"PXmD9F6HJa6KTKFV/1Tz4cavOcpqEinuOy/WP6IgIMw="},"nonce":"2"},"qos":"4000000000000","qscs":null}}


//acc2在验证节点没有委托过，投票无效
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query tally 1
{"yes":"0","abstain":"0","no":"0","no_with_veto":"0"}
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query vote 1 acc2
ERROR: {"codespace":"sdk","code":1,"message":"voter و(\ufffd\ufffd8\u0015K\ufffd;\ufffdX:\ufffdex\ufffd\ufffd\ufffdZ is not vote on proposal 1"}
[vagrant@vagrant-192-168-1-200 ~]$ qoscli tx vote --proposal-id 1 --voter acc2 --option yes
Password to sign with 'acc2':
{"check_tx":{"gasWanted":"9223372036854775807","gasUsed":"6958","events":[]},"deliver_tx":{"gasWanted":"9223372036854775807","gasUsed":"9700","events":[{"type":"vote-proposal","attributes":[{"key":"cHJvcG9zYWwtaWQ=","value":"MQ=="},{"key":"dm90ZXI=","value":"YWRkcmVzczFteHl6M3JsZDhxMjVoeDNtNjN2cjRkdDkwejc3bjYyNjg0YzczNw=="}]},{"type":"message","attributes":[{"key":"bW9kdWxl","value":"Z292"},{"key":"Z2FzLnBheWVy","value":"YWRkcmVzczFteHl6M3JsZDhxMjVoeDNtNjN2cjRkdDkwejc3bjYyNjg0YzczNw=="}]}]},"hash":"D83213770C4EC32AD4C4BAD37DF5C99007A79989AB4DA372576F98F966678A87","height":"232"}
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query vote 1 acc2
{"voter":"address1mxyz3rld8q25hx3m63vr4dt90z77n62684c737","proposal_id":"1","option":"Yes"}
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query tally 1
{"yes":"0","abstain":"0","no":"0","no_with_veto":"0"}
[vagrant@vagrant-192-168-1-200 ~]$ qoscli tx vote --proposal-id 1 --voter acc2 --option no
Password to sign with 'acc2':
{"check_tx":{"gasWanted":"9223372036854775807","gasUsed":"7081","events":[]},"deliver_tx":{"gasWanted":"9223372036854775807","gasUsed":"9800","events":[{"type":"vote-proposal","attributes":[{"key":"cHJvcG9zYWwtaWQ=","value":"MQ=="},{"key":"dm90ZXI=","value":"YWRkcmVzczFteHl6M3JsZDhxMjVoeDNtNjN2cjRkdDkwejc3bjYyNjg0YzczNw=="}]},{"type":"message","attributes":[{"key":"bW9kdWxl","value":"Z292"},{"key":"Z2FzLnBheWVy","value":"YWRkcmVzczFteHl6M3JsZDhxMjVoeDNtNjN2cjRkdDkwejc3bjYyNjg0YzczNw=="}]}]},"hash":"815DDB3D1CF957C8437F603038004D28E412078245F2149092F1C2B04FC3962E","height":"241"}
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query vote 1 acc2
{"voter":"address1mxyz3rld8q25hx3m63vr4dt90z77n62684c737","proposal_id":"1","option":"No"}
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query tally 1
{"yes":"0","abstain":"0","no":"0","no_with_veto":"0"}

[vagrant@vagrant-192-168-1-200 ~]$ qoscli query account acc2
{"type":"qos/types/QOSAccount","value":{"base_account":{"account_address":"address1mxyz3rld8q25hx3m63vr4dt90z77n62684c737","public_key":{"type":"tendermint/PubKeyEd25519","value":"PXmD9F6HJa6KTKFV/1Tz4cavOcpqEinuOy/WP6IgIMw="},"nonce":"2"},"qos":"3999999999805","qscs":null}}


//acc0在验证节点委托过，投票有效
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query tally 1
{"yes":"0","abstain":"0","no":"0","no_with_veto":"0"}
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query proposal 1
{"proposal_content":{"type":"gov/TextProposal","value":{"title":"myproposal","description":"the first proposal","deposit":"7"}},"proposal_id":"1","proposal_status":"Voting","final_tally_result":{"yes":"0","abstain":"0","no":"0","no_with_veto":"0"},"submit_time":"2019-08-29T02:41:47.604696459Z","deposit_end_time":"2019-08-29T02:51:47.604696459Z","total_deposit":"10","voting_start_time":"2019-08-29T02:44:03.131511973Z","voting_start_height":"190","voting_end_time":"2019-08-29T02:54:03.131511973Z"}
[vagrant@vagrant-192-168-1-200 ~]$ qoscli tx vote --proposal-id 1 --voter acc0 --option yes
Password to sign with 'acc0':
{"check_tx":{"gasWanted":"9223372036854775807","gasUsed":"7081","events":[]},"deliver_tx":{"gasWanted":"9223372036854775807","gasUsed":"9800","events":[{"type":"vote-proposal","attributes":[{"key":"cHJvcG9zYWwtaWQ=","value":"MQ=="},{"key":"dm90ZXI=","value":"YWRkcmVzczFuNjRoNnByeHo2bGQ1dmw4ZDBycnpzbnkyNW5xN3ZnY25yaHU5OA=="}]},{"type":"message","attributes":[{"key":"bW9kdWxl","value":"Z292"},{"key":"Z2FzLnBheWVy","value":"YWRkcmVzczFuNjRoNnByeHo2bGQ1dmw4ZDBycnpzbnkyNW5xN3ZnY25yaHU5OA=="}]}]},"hash":"1770DBC4886BD3D1611AC4F5F280C422676BC63D40FF736A823BE4910C7414F3","height":"275"}
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query tally 1
{"yes":"4000000000000","abstain":"0","no":"0","no_with_veto":"0"}

//等待投票期结束，提议结果如下
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query proposal 1
{"proposal_content":{"type":"gov/TextProposal","value":{"title":"myproposal","description":"the first proposal","deposit":"7"}},"proposal_id":"1","proposal_status":"Passed","final_tally_result":{"yes":"4000000000000","abstain":"0","no":"0","no_with_veto":"0"},"submit_time":"2019-08-29T02:41:47.604696459Z","deposit_end_time":"2019-08-29T02:51:47.604696459Z","total_deposit":"10","voting_start_time":"2019-08-29T02:44:03.131511973Z","voting_start_height":"190","voting_end_time":"2019-08-29T02:54:03.131511973Z"}

```
