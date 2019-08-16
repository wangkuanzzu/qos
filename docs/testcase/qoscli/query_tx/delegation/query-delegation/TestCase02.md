# Description

```text
    参数不合法。
```

## Input

```bash
//jlgy02在jlgy01账户创建的validator没有进行过委托
qoscli query delegation --owner jlgy01 --delegator jlgy02 --indent

//jlgy02未创建过validator
qoscli query delegation --owner jlgy02 --delegator jlgy01 --indent
```

## Output

```bash
qoscli query delegation --owner jlgy01 --delegator jlgy02 --indent
ERROR: {"codespace":"sdk","code":1,"message":"delegationInfo not exsits. owner: address1nnvdqefva89xwppzs46vuskckr7klvzk8r5uaa , deleAddr: address1l0wn66gh45nfta2r4vq8z54wu9hgarss298e9g"}

qoscli query delegation --owner jlgy02 --delegator jlgy01 --indent
ERROR: {"codespace":"sdk","code":1,"message":"validator not exsits. owner: address1l0wn66gh45nfta2r4vq8z54wu9hgarss298e9g"}
```
