# Description

```text
    参数`address`，`deleted-by`不合法
```

## Input

```bash
    //不存在账户defg
    qoscli tx delete-guardian --address hij --deleted-by defg

    //hij账户不是guardian账户
    qoscli tx delete-guardian --address hij --deleted-by defg
```

## Output

```bash
    //不存在账户defg
    qoscli tx delete-guardian --address hij --deleted-by defg
    null
    ERROR: Name: defg not found

    //hij账户不是guardian账户
    qoscli tx delete-guardian --address hij --deleted-by def
    Password to sign with 'def':
    {"check_tx":{"code":1,"log":"{\"codespace\":\"sdk\",\"code\":1,\"message\":\"TxStd's ITx ValidateData error:  ERROR:\\nCodespace: guardian\\nCode: 603\\nMessage: \\\"unknown guardian\\\"\\n\"}","gasWanted":"100000","gasUsed":"1000"},"deliver_tx":{},"hash":"2942FF6550BED360D240654B0E39446386B9DE85D74FEB9E1E7B49968250F6BF","height":"0"}
    ERROR: {"codespace":"sdk","code":1,"message":"TxStd's ITx ValidateData error:  ERROR:\nCodespace: guardian\nCode: 603\nMessage: \"unknown guardian\"\n"}
```
