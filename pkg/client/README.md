# Bhojpur Suga - Client-side Framework

## How to use it

Here is a code example to see how to use it.

```golang
var information map[string]interface{}
client, err := NewClient("localhost:8080", true, &information)
if err != nil {
	panic(err)
}

defer client.Close()

client.SendMessage("Hello Suga!")
```
