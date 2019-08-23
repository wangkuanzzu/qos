# test case of qoscli tx add-guardian

> `qoscli tx add-guardian` 添加特权账户

---

## 情景说明

1. 添加特权账户，特权账户的添加只能是由特权账户来完成，也就是在genesis.json中配置的特权账户。前提条件：知晓在genesis.json文件中的特权账户的密码或是私钥。测试命令执行的前提是账户abc配置在genesis.json文件中，本地密钥库也保存账户abc的信息。
2. 添加特权账户的tx操作，无基础的gas费用消耗，也没有访问存储消耗的gas。

## 测试命令

```bash
    //查询所有系统账户
    qoscli query guardians

    //账户abc是配置在genesis.json文件中的特权账户
    qoscli tx add-guardian --address def --creator abc --description 'set def to be a guardian'

    //查询所有系统账户
    qoscli query guardians


    //账户def是通过账户abc添加的特权账户
    qoscli tx add-guardian --address hij --creator def --description 'set hij to be a guardian'


```

## 预测结果

```text
    1. 查询系统账户前后，可查到新增的系统账户。
    2. 添加特权账户成功，返回tx信息。
    3. 添加特权账户失败，def不是存在genesis文件中的系统账户，权限不足。
    4. 特权账户的数额不会变化。
```

## 测试结果

```bash
    qoscli tx add-guardian --address def --creator abc --description 'set def to be a guardian'
    Password to sign with 'abc':
    {"check_tx":{"gasWanted":"100000","gasUsed":"7856"},"deliver_tx":{"gasWanted":"100000","gasUsed":"12046","tags":[{"key":"YWN0aW9u","value":"YWRkLWd1YXJkaWFu"},{"key":"Y3JlYXRvcg==","value":"YWRkcmVzczEweHd4MDZnbnJ0M2RsejdoZnJ4NmE4d3gzZ3llZ2h4bTU0cnY3YQ=="},{"key":"Z3VhcmRpYW4=","value":"YWRkcmVzczFsNmp1YXF5OWZrMGRwczBmbjVkY2c0ZnB5MzZ6bXJ5cDhteTR1eA=="}]},"hash":"857BF0332E9FB1F0B89378833FCA1D06E1543464465E7DF4BF46AC417935CCEC","height":"203"}

    qoscli tx add-guardian --address hij --creator def --description 'set hij to be a guardian'
    Password to sign with 'def':
    {"check_tx":{"code":1,"log":"{\"codespace\":\"sdk\",\"code\":1,\"message\":\"TxStd's ITx ValidateData error:  ERROR:\\nCodespace: guardian\\nCode: 602\\nMessage: \\\"Creator not exists or not init from genesis\\\"\\n\"}","gasWanted":"100000","gasUsed":"2213"},"deliver_tx":{},"hash":"A935D93832DC2AEE23EC00813BA77D5C0280ECB0C8C7B50E34E6CAB75030360F","height":"0"}
    ERROR: {"codespace":"sdk","code":1,"message":"TxStd's ITx ValidateData error:  ERROR:\nCodespace: guardian\nCode: 602\nMessage: \"Creator not exists or not init from genesis\"\n"}

```
