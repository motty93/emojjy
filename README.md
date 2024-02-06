# open-emoji-api
## Get Start
こちらでAPI Keyを取得しましょう。

https://api-ninjas.com/api/emoji

```bash
$ go install

$ go run main.go
```

## How to use
常時標準出力でemoji名を入力するとunicodeを含むjsonが返ってきます。

取得したemojiは煮るなり焼くなりしてください。

### 終了方法
`Enter emoji name: `に対して `exit`と入力するか `ctrl-c`でプログラムを終了させてください。

## Build
```bash
$ go build -o bin/main main.go

$ ./bin/main
```
