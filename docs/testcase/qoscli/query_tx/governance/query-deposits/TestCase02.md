# Description

```text
    参数`--proposal-id`不合法。前提条件：QOS网络中提议小于100个
```

## Input

```bash
    //传入错误参数
    qoscli query deposits 0

    qoscli query deposits 100
```

## Output

```bash
    //传入错误参数
    qoscli query deposits 0
    ERROR: {"codespace":"sdk","code":1,"message":"proposal id 0 not exsits"}
    qoscli query deposits 100
    ERROR: {"codespace":"sdk","code":1,"message":"proposal id 100 not exsits"}
```
