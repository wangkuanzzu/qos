# Description

```text
    参数不合法。
```

## Input

```bash
//账号jlgy02未曾在账户jlgy01创建的validator进行过委托
qoscli query delegator-income --owner jlgy01 --delegator jlgy02 --indent

//账户jlgy02未曾创建过validator
qoscli query delegator-income --owner jlgy02 --delegator jlgy02 --indent
```

## Output

```bash
qoscli query delegator-income --owner jlgy01 --delegator jlgy02 --indent
ERROR: {"codespace":"sdk","code":1,"message":"delegator income info not exsits. delegator: address1l0wn66gh45nfta2r4vq8z54wu9hgarss298e9g , owner: address1nnvdqefva89xwppzs46vuskckr7klvzk8r5uaa"}

qoscli query delegator-income --owner jlgy02 --delegator jlgy02 --indent
ERROR: {"codespace":"sdk","code":1,"message":"validator not exsits. owner: address1l0wn66gh45nfta2r4vq8z54wu9hgarss298e9g"}
```
