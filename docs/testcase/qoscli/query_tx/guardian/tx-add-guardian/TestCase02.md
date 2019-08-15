# Description

```text
    参数`address`，`creator`，`description`不合法
```

## Input

```bash
    qoscli tx add-guardian --address def --creator abcd --description "this is the description"

    qoscli tx add-guardian --address aef --creator abc --description "this is the description"
```

## Output

```bash
    qoscli tx add-guardian --address def --creator abcd --description "this is the description"
    null
    ERROR: Name: abcd not found

    qoscli tx add-guardian --address aef --creator abc --description "this is the description"
    null
    ERROR: Name: aef not found
```
