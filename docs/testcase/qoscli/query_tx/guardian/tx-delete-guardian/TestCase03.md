# Description

```text
    正常删除特权账户
```

## Input

```bash
    //账户abc是配置在genesis.json文件中的特权账户
    qoscli tx delete-guardian --address def --deleted-by abc
```

## Output

```bash
    qoscli tx delete-guardian --address def --deleted-by abc
    Password to sign with 'abc':
    {"check_tx":{"gasWanted":"100000","gasUsed":"8069"},"deliver_tx":{"gasWanted":"100000","gasUsed":"9069","tags":[{"key":"YWN0aW9u","value":"ZGVsZXRlLWd1YXJkaWFu"},{"key":"ZGVsZXRlLWJ5","value":"YWRkcmVzczEweHd4MDZnbnJ0M2RsejdoZnJ4NmE4d3gzZ3llZ2h4bTU0cnY3YQ=="},{"key":"Z3VhcmRpYW4=","value":"YWRkcmVzczFqajQ5NGE0dWd0NDhzeTgwbjNhbWc2ZHZoejB5M3lwOTRhM3B4dA=="}]},"hash":"0D5FE4776B02D8B7D6479FD38FDA954DCDEA8CB70C0E38CFD38D01B98C17EB15","height":"1856"}
```
