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

action: CREATE_GAME, FLAG, UNFLAG, DIG

## response

```json
{
  "type": "UPDATE_BOARD",
  "data": "board[][]"
}

{
  "type": "CONNECTED_USERS",
  "data": "users[]"
}

{
  "type": "USERS_MOUSE",
  "data": {
    "id": "id",
    "name": "name",
    "mouseX": 0,
    "mouseY": 0
  }
}
```

a update always contains a type and a data field
