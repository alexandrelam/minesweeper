# message

## request

```json
{
  "action": "CREATE_GAME"
}
```

```json
{
  "action": "DIG", // FLAG, UNFLAG
  "row": "0",
  "column": "0"
}
```

```json
{
  "action": "JOIN",
  "name": "alexandre"
}
```

action: CREATE_GAME, FLAG, UNFLAG, DIG, JOIN

## response

```
byte: board[][]

```
