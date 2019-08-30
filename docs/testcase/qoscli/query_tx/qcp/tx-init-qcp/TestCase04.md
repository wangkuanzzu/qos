# test case of qoscli tx init-qcp`

> `qoscli tx init-qcp` 初始化联盟链

---

## 情景说明

1. 初始化联盟链使用的证书需要从kepler进行证书申请，安装证书后，初始化联盟链。查看联盟链信息，确认是否操作成功。
2. 证书来源并非kepler，进行证书安装，验证是否可以成功。
3. 使用同一个证书进行多个联盟链创建操作，验证是否可行。
4. 初始化联盟链的tx操作，有基础的gas费用消耗（1.8QOS=1800000GAS），再加上访问存储所消耗的GAS费用。

## 数据准备

qcp公钥与私钥（来自kepler）：

```text
    "privKey": {
      "type": "tendermint/PrivKeyEd25519",
      "value": "z1LC/iJDY6PDbKsfGrUORCIvkfjmcQypNCJJhFnfLxuKgPUO/fox/VJ8Y+SY2aIfyMGvlJqcULZ4Aq5BIKDa0w=="
    },
    "pubKey": {
      "type": "tendermint/PubKeyEd25519",
      "value": "ioD1Dv36Mf1SfGPkmNmiH8jBr5SanFC2eAKuQSCg2tM="
    }
```

kepler颁发的证书：qcp.crt

```text
{"csr":{"subj":{"type":"certificate/QCPSubject","value":{"chain_id":"aquarius-1000","qcp_chain":"qcptest-1000"}},"is_ca":false,"not_before":"2019-08-30T03:26:25.401819115Z","not_after":"2020-08-30T03:26:25.401839097Z","public_key":{"type":"tendermint/PubKeyEd25519","value":"ioD1Dv36Mf1SfGPkmNmiH8jBr5SanFC2eAKuQSCg2tM="}},"ca":{"subj":null,"public_key":{"type":"tendermint/PubKeyEd25519","value":"w+UlkkcrHKKwAmEEl76rO5xHHj3quoxLgN5rvE5yYQ0="}},"signature":"G/Wew0m3uf+uyHURh1/fOKi4duK6Rr+BBI29TqO31j7z9ODrECBV39NFhLfNv2Tmx2eSvzxj9x/XRf5oaNzTAQ=="}
```

非kepler颁发的证书：qcp.crt

```text
{"csr":{"subj":{"type":"certificate/QCPSubject","value":{"chain_id":"aquarius-1000","qcp_chain":"qcptest-1000"}},"is_ca":false,"not_before":"2019-08-30T03:26:25.401819115Z","not_after":"2020-08-30T03:26:25.401839097Z","public_key":{"type":"tendermint/PubKeyEd25519","value":"ioD1Dv36Mf1SfGPkmNmiH8jBr5SanFC2eAKuQSCg2tM="}},"ca":{"subj":null,"public_key":{"type":"tendermint/PubKeyEd25519","value":"w+UlkkcrHKKwAmEEl76rO5xHHj3quoxLgN5rvE5yYQ0="}},"signature":"G/Wew0m3uf+uyHURh1/fOKi4duK6Rr+BBI29TqO31j7z9ODrECBV39NFhLfNv2Tmx2eadexfdx/XRf5oaNzTAQ=="}
```

## 测试命令

```bash
//见测试结果
```

## 预测结果

```text
1、初始化联盟链使用正确证书，初始化成功。
2、使用非kepler颁发的证书，提示ca不正确。
3、同一份证书再次去初始化联盟链，返回联盟链已经存在。
4、gas的消耗大于1800000GAS
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

//首次创建联盟链成功
[root@iz2zef57ni8z6ydrcloxy0z ~]# qoscli tx init-qcp --creator jlgy --qcp.crt ./qosCA/qcp-zzu.crt --indent
Password to sign with 'jlgy':
{
  "check_tx": {
    "gasWanted": "9223372036854775807",
    "gasUsed": "189027",
    "events": []
  },
  "deliver_tx": {
    "gasWanted": "9223372036854775807",
    "gasUsed": "196100",
    "events": [
      {
        "type": "init-qcp",
        "attributes": [
          {
            "key": "Y2hhaW4taWQ=",
            "value": "cWNwdGVzdC0xMDAw"
          },
          {
            "key": "Y3JlYXRvcg==",
            "value": "YWRkcmVzczFucm0yNHZ4YXYwdTU4anB0MDR2N3Z0YXZyamp4d3c2MHg5eW03cA=="
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "bW9kdWxl",
            "value": "cWNw"
          },
          {
            "key": "Z2FzLnBheWVy",
            "value": "YWRkcmVzczFucm0yNHZ4YXYwdTU4anB0MDR2N3Z0YXZyamp4d3c2MHg5eW03cA=="
          }
        ]
      }
    ]
  },
  "hash": "57C1A4ACBFC08AD64E94CDC5BA9358C24C4B45977160FA98B225A99FA6CFE892",
  "height": "34417"
}

//第二次创建联盟链
[root@iz2zef57ni8z6ydrcloxy0z ~]# qoscli tx init-qcp --creator jlgy --qcp.crt ./qosCA/qcp-zzu.crt
Password to sign with 'jlgy':
{"check_tx":{"code":1,"log":"{\"codespace\":\"sdk\",\"code\":1,\"message\":\"TxStd's ITx ValidateData error:  ERROR:\\nCodespace: qcp\\nCode: 405\\nMessage: \\\"qcp exists\\\"\\n\"}","gasWanted":"9223372036854775807","gasUsed":"183465","events":[]},"deliver_tx":{},"hash":"C8EDBA2F057CE4EB19A9814EC75703E9DE7932EDC345794C5D43F4DB50B711C9","height":"0"}
ERROR: {"codespace":"sdk","code":1,"message":"TxStd's ITx ValidateData error:  ERROR:\nCodespace: qcp\nCode: 405\nMessage: \"qcp exists\"\n"}

//错误的qcp证书
[root@iz2zef57ni8z6ydrcloxy0z ~]# qoscli tx init-qcp --creator jlgy --qcp.crt ./qosCA/qcp-zzu-err.crt
Password to sign with 'jlgy':
{"check_tx":{"code":1,"log":"{\"codespace\":\"sdk\",\"code\":1,\"message\":\"TxStd's ITx ValidateData error:  ERROR:\\nCodespace: qcp\\nCode: 403\\nMessage: \\\"wrong qcp ca\\\"\\n\"}","gasWanted":"9223372036854775807","gasUsed":"182354","events":[]},"deliver_tx":{},"hash":"82272B282A9AE115FECB557A2F628BCA2C734F0163AA841A83201736FE59E448","height":"0"}
ERROR: {"codespace":"sdk","code":1,"message":"TxStd's ITx ValidateData error:  ERROR:\nCodespace: qcp\nCode: 403\nMessage: \"wrong qcp ca\"\n"}


```

ps ：
    消耗的gas为196100  大于180000gas，基础消耗的gas正常。
