# Description

```text
    参数`--proposal-id`，`--depositor`不合法。前提条件：QOS网络中提议小于100个，账号abc在提议2上有抵押，账号abc在提议3没有抵押
```

## Input

```bash
    //账号abc在3号提议上没有抵押
    qoscli query deposit 3 abc

    //传入错误参数
    qoscli query deposit 0 abc

    qoscli query deposit 100 abc

    qoscli query deposits 0

    qoscli query deposits 100
```

## Output

```bash
    //账号abc在3号提议上没有抵押
    qoscli query deposit 3 abc
    ERROR: {"codespace":"sdk","code":1,"message":"depositer address10xwx06gnrt3dlz7hfrx6a8wx3gyeghxm54rv7a is not deposit on proposal 3"}

    //传入错误参数
    qoscli query deposit 0 abc
    ERROR: {"codespace":"sdk","code":1,"message":"proposal id 0 not exsits"}
    qoscli query deposit 100 abc
    ERROR: {"codespace":"sdk","code":1,"message":"proposal id 100 not exsits"}
    qoscli query deposits 0
    ERROR: {"codespace":"sdk","code":1,"message":"proposal id 0 not exsits"}
    qoscli query deposits 100
    ERROR: {"codespace":"sdk","code":1,"message":"proposal id 100 not exsits"}
```
