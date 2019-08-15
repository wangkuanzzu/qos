# Description

```text
    正常查询抵押
```

## Input

```bash
    //传入正确参数
    qoscli query deposit 2 abc
```

## Output

```bash
    //传入正确参数
    qoscli query deposit 2 abc
    {"depositor":"address10xwx06gnrt3dlz7hfrx6a8wx3gyeghxm54rv7a","proposal_id":"2","amount":"66480000"}
```

ps：
    1. 在查询结果中amount数量是所有抵押次数的合计amount数量。
    2. **是否可以增加查询某一账号在所有提议上的抵押情况？**
