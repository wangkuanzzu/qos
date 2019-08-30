# test case of qoscli tx create-qsc`

> `qoscli tx create-qsc` 初始化联盟币

---

## 情景说明

1. 初始化联盟币使用的证书需要从kepler进行证书申请，安装证书后，初始化联盟币。查看联盟币信息，确认是否操作成功。
2. 证书来源并非kepler，进行证书安装，验证是否可以成功。
3. 使用同一个证书进行多个联盟币创建操作，验证是否可行。做幂等检查。
4. 初始化（创建）联盟币的tx操作，有基础的gas费用消耗（1.8QOS=1800000GAS），再加上访问存储所消耗的GAS费用。

## 数据准备

banker公私钥：

```text
    "privKey": {
      "type": "tendermint/PrivKeyEd25519",
      "value": "lSKZZy/8X75/gBMCFqba8wspl0wPIcvFvvHkVVktduh3hkQsH80LaGe8kGEL5nwA7CkQ68HGoY77rSF+MJLM6A=="
    },
    "pubKey": {
      "type": "tendermint/PubKeyEd25519",
      "value": "d4ZELB/NC2hnvJBhC+Z8AOwpEOvBxqGO+60hfjCSzOg="
    }
```

kepler颁发的qsc证书：

```text
{"csr":{"subj":{"type":"certificate/QSCSubject","value":{"chain_id":"aquarius-1000","name":"ZZU","banker":{"type":"tendermint/PubKeyEd25519","value":"d4ZELB/NC2hnvJBhC+Z8AOwpEOvBxqGO+60hfjCSzOg="}}},"is_ca":false,"not_before":"2019-08-30T03:33:40.345319721Z","not_after":"2020-08-30T03:33:40.345347129Z","public_key":{"type":"tendermint/PubKeyEd25519","value":"ioD1Dv36Mf1SfGPkmNmiH8jBr5SanFC2eAKuQSCg2tM="}},"ca":{"subj":null,"public_key":{"type":"tendermint/PubKeyEd25519","value":"w+UlkkcrHKKwAmEEl76rO5xHHj3quoxLgN5rvE5yYQ0="}},"signature":"RYUObAIj68qK70GkYMh7+FeotnyIWtCH4/Nr0f0+7Y2KnH97mhsIlDU041m3K7xLRTETxAWUZWHQ9zGG65SyDg=="}
```

非kepler颁发的qsc证书：

```text
{"csr":{"subj":{"type":"certificate/QSCSubject","value":{"chain_id":"aquarius-1000","name":"ZZU","banker":{"type":"tendermint/PubKeyEd25519","value":"d4ZELB/NC2hnvJBhC+Z8AOwpEOvBxqGO+60hfjCSzOg="}}},"is_ca":false,"not_before":"2019-08-30T03:33:40.345319721Z","not_after":"2020-08-30T03:33:40.345347129Z","public_key":{"type":"tendermint/PubKeyEd25519","value":"ioD1Dv36Mf1SfGPkmNmiH8jBr5SanFC2eAKuQSCg2tM="}},"ca":{"subj":null,"public_key":{"type":"tendermint/PubKeyEd25519","value":"w+UlkkcrHKKwAmEEl76rO5xHHj3quoxLgN5rvE5yYQ0="}},"signature":"RYUObAIj68qK70GkYMh7+FeotnyIWtCH4/Nr0f0+7Y2KnH97mhsIlDU041m3SFSsfSETxAWUZWHQ9zGG65SyDg=="}
```

## 测试命令

```bash
//测试命令见测试结果
```

## 预测结果

```text

```

## 测试结果

```bash
[root@iz2zef57ni8z6ydrcloxy0z ~]# ll ./qosCA/
total 20
-rw-r--r-- 1 root root 550 Aug 30 12:07 qcp-zzu.crt
-rw-r--r-- 1 root root 550 Aug 30 17:29 qcp-zzu-err.crt
-rw-r--r-- 1 root root 636 Aug 30 12:08 qsc-zzu.crt
-rw-r--r-- 1 root root 636 Aug 30 17:42 qsc-zzu-err.crt
-rw-r--r-- 1 root root 153 Aug 30 12:06 zzubanker.json

[root@iz2zef57ni8z6ydrcloxy0z ~]# qoscli query qsc ZZU --indent
ERROR: ZZU not exists.

[root@iz2zef57ni8z6ydrcloxy0z ~]# qoscli tx create-qsc --creator jlgy --qsc.crt ./qosCA/qsc-zzu.crt
Password to sign with 'jlgy':
{"check_tx":{"gasWanted":"9223372036854775807","gasUsed":"189027","events":[]},"deliver_tx":{"gasWanted":"9223372036854775807","gasUsed":"196300","events":[{"type":"create-qsc","attributes":[{"key":"bmFtZQ==","value":"WlpV"},{"key":"Y3JlYXRvcg==","value":"YWRkcmVzczFucm0yNHZ4YXYwdTU4anB0MDR2N3Z0YXZyamp4d3c2MHg5eW03cA=="}]},{"type":"message","attributes":[{"key":"bW9kdWxl","value":"cXNj"},{"key":"Z2FzLnBheWVy","value":"YWRkcmVzczFucm0yNHZ4YXYwdTU4anB0MDR2N3Z0YXZyamp4d3c2MHg5eW03cA=="}]}]},"hash":"CC6F50D4DD12229CC8A1ECB216B210DFB149F0AFA9554C3A59683EE6B941C04D","height":"35250"}

[root@iz2zef57ni8z6ydrcloxy0z ~]# qoscli query qsc ZZU --indent
{
  "name": "ZZU",
  "chain_id": "aquarius-1000",
  "extrate": "1",
  "description": "",
  "banker": "address1hsurfl3kt334qwj45y3d2hsdwt6ufsfpzphuze",
  "total_amount": "0"
}

//重复创建联盟币
[root@iz2zef57ni8z6ydrcloxy0z ~]# qoscli tx create-qsc --creator jlgy --qsc.crt ./qosCA/qsc-zzu.crt
Password to sign with 'jlgy':
{"check_tx":{"code":1,"log":"{\"codespace\":\"sdk\",\"code\":1,\"message\":\"TxStd's ITx ValidateData error:  ERROR:\\nCodespace: qsc\\nCode: 306\\nMessage: \\\"qsc exists\\\"\\n\"}","gasWanted":"9223372036854775807","gasUsed":"182255","events":[]},"deliver_tx":{},"hash":"3313CB7067577614FA1D2FAAAE04F67B4A8202ECC05D3EEBD59243541F2E5D26","height":"0"}
ERROR: {"codespace":"sdk","code":1,"message":"TxStd's ITx ValidateData error:  ERROR:\nCodespace: qsc\nCode: 306\nMessage: \"qsc exists\"\n"}


[root@iz2zef57ni8z6ydrcloxy0z ~]# qoscli tx create-qsc --creator jlgy --qsc.crt ./qosCA/qsc-zzu-err.crt
Password to sign with 'jlgy':
{"check_tx":{"code":1,"log":"{\"codespace\":\"sdk\",\"code\":1,\"message\":\"TxStd's ITx ValidateData error:  ERROR:\\nCodespace: qsc\\nCode: 303\\nMessage: \\\"wrong qsc ca\\\"\\n\"}","gasWanted":"9223372036854775807","gasUsed":"181111","events":[]},"deliver_tx":{},"hash":"58E8B4C5756656BA7B7D7920713E9538C8C493E7D2F1CB309ADA5A7A5ADEB796","height":"0"}
ERROR: {"codespace":"sdk","code":1,"message":"TxStd's ITx ValidateData error:  ERROR:\nCodespace: qsc\nCode: 303\nMessage: \"wrong qsc ca\"\n"}

```

ps：
    1、消耗的gas费用为196300，大于基础gas消耗180000，基础gas消耗正常。
