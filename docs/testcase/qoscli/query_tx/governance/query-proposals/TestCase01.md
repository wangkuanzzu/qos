# Description

```text
    查询所有提议
```

## Input

```bash
    //查询所有提议
    qoscli query proposals

    //带有筛选条件的查询
    qoscli query proposals --status (DepositPeriod|VotingPeriod|Passed|Rejected).
```

## Output

```bash
    //查询所有提议
    qoscli query proposals
    [{"proposal_content":{"type":"gov/TextProposal","value":{"title":"myproposal","description":"the first proposal","deposit":"33400000"}},"proposal_id":"1","proposal_status":"Deposit","final_tally_result":{"yes":"0","abstain":"0","no":"0","no_with_veto":"0"},"submit_time":"2019-08-01T07:16:17.516684138Z","deposit_end_time":"2019-08-08T07:16:17.516684138Z","total_deposit":"33400000","voting_start_time":"0001-01-01T00:00:00Z","voting_start_height":"0","voting_end_time":"0001-01-01T00:00:00Z"},{"proposal_content":{"type":"gov/TextProposal","value":{"title":"myproposal","description":"the first proposal","deposit":"66300000"}},"proposal_id":"2","proposal_status":"Deposit","final_tally_result":{"yes":"0","abstain":"0","no":"0","no_with_veto":"0"},"submit_time":"2019-08-01T08:53:30.889401314Z","deposit_end_time":"2019-08-08T08:53:30.889401314Z","total_deposit":"66300000","voting_start_time":"0001-01-01T00:00:00Z","voting_start_height":"0","voting_end_time":"0001-01-01T00:00:00Z"},{"proposal_content":{"type":"gov/TextProposal","value":{"title":"myproposal","description":"the first proposal","deposit":"66300000"}}]

    qoscli query proposals --status rejected
    [{"proposal_content":{"type":"gov/TextProposal","value":{"title":"myproposal","description":"the first proposal","deposit":"750000000"}},"proposal_id":"5","proposal_status":"Rejected","final_tally_result":{"yes":"2000000000","abstain":"0","no":"100000","no_with_veto":"0"},"submit_time":"2019-08-01T09:00:38.798681467Z","deposit_end_time":"2019-08-08T09:00:38.798681467Z","total_deposit":"750000000","voting_start_time":"2019-08-01T09:00:38.798681467Z","voting_start_height":"516054","voting_end_time":"2019-08-08T09:00:38.798681467Z"}]

```
