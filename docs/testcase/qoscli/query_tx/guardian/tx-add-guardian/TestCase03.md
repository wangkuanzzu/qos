# Description

```text
    正常添加特权账户
```

## Input

```bash
    //账户abc是配置在genesis.json文件中的特权账户
    qoscli tx add-guardian --address def --creator abc --description 'set def to be a guardian'
```

## Output

```bash
    qoscli tx add-guardian --address def --creator abc --description 'set def to be a guardian'
    Password to sign with 'abc':
    {"check_tx":{"gasWanted":"100000","gasUsed":"7856"},"deliver_tx":{"gasWanted":"100000","gasUsed":"12046","tags":[{"key":"YWN0aW9u","value":"YWRkLWd1YXJkaWFu"},{"key":"Y3JlYXRvcg==","value":"YWRkcmVzczEweHd4MDZnbnJ0M2RsejdoZnJ4NmE4d3gzZ3llZ2h4bTU0cnY3YQ=="},{"key":"Z3VhcmRpYW4=","value":"YWRkcmVzczFsNmp1YXF5OWZrMGRwczBmbjVkY2c0ZnB5MzZ6bXJ5cDhteTR1eA=="}]},"hash":"857BF0332E9FB1F0B89378833FCA1D06E1543464465E7DF4BF46AC417935CCEC","height":"203"}
```
