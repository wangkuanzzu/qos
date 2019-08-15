# Description

```text
    缺失必须参数`--proposal-id`，`--depositor`
```

## Input

```bash
    qoscli query deposit 5

    qoscli query deposit abc

    qoscli query deposit
```

## Output

```bash
    qoscli query deposit 5
    ERROR: accepts 2 arg(s), received 1

    qoscli query deposit abc
    ERROR: accepts 2 arg(s), received 1

    qoscli query deposit
    ERROR: accepts 2 arg(s), received 0
```
