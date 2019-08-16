# Description

```text
    参数`--id`不合法，查询指定提议的投票信息：提议的序号不存在或不符合规则。
```

## Input

```bash
    //查询指定的提议：提议的序号不存在或不符合规则
    qoscli query tally 0
    qoscli query tally 9999
```

## Output

```bash
    //查询指定的提议：提议的序号不存在或不符合规则
    qoscli query tally 0
    ERROR: {"codespace":"sdk","code":1,"message":"proposal id 0 not exsits"}
    qoscli query tally 9999
    ERROR: {"codespace":"sdk","code":1,"message":"proposal id 9999 not exsits"}
```
