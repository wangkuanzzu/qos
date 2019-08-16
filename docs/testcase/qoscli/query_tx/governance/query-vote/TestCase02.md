# Description

```text
    参数`--id`,`--voter`不合法，查询指定提议的投票信息：提议的序号不存在或不符合规则。
```

## Input

```bash
    //提议编号错误
    qoscli query vote 0 abc

    qoscli query vote 100 abc

    //提议voter错误
    qoscli query vote 5 abcdefg

    qoscli query vote 5 def
```

## Output

```bash
    //提议编号错误
    qoscli query vote 0 abc
    ERROR: {"codespace":"sdk","code":1,"message":"proposal id 0 not exsits"}

    qoscli query vote 100 abc
    ERROR: {"codespace":"sdk","code":1,"message":"proposal id 100 not exsits"}

    //提议voter错误
    qoscli query vote 5 abcdefg
    ERROR: voter abcdefg is not a valid address value

    qoscli query vote 5 abcd
    ERROR: {"codespace":"sdk","code":1,"message":"voter \ufffd\t\ufffd9\ufffd\ufffd\ufffd\ufffdZ\ufffd\u0008\ufffd\ufffd\u0026\ufffd\ufffd\ufffd%!!(MISSING)�(MISSING)\" is not vote on proposal 5"}


```
