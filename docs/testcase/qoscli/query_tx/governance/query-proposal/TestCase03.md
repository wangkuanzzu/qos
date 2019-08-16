# test case of qoscli query proposal

> `qoscli query proposal` 查询提议

---

## 情景说明

需要查询提议的具体内容，以及提议目前的阶段或是进度。

## 测试命令

```bash
    //查询指定的提议
    qoscli query proposal 2
```

## 测试结果

```bash
    //查询指定提议
    qoscli query proposal 2
    {"proposal_content":{"type":"gov/TextProposal","value":{"title":"myproposal","description":"the first proposal","deposit":"66300000"}},"proposal_id":"2","proposal_status":"Deposit","final_tally_result":{"yes":"0","abstain":"0","no":"0","no_with_veto":"0"},"submit_time":"2019-08-01T08:53:30.889401314Z","deposit_end_time":"2019-08-08T08:53:30.889401314Z","total_deposit":"66300000","voting_start_time":"0001-01-01T00:00:00Z","voting_start_height":"0","voting_end_time":"0001-01-01T00:00:00Z"}
```
