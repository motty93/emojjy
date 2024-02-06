# emojjy
<img src="https://private-user-images.githubusercontent.com/25767163/302523416-0d8b493a-2f3e-4a8a-9b60-73cb01accdb3.gif?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MDcyMDE0NjQsIm5iZiI6MTcwNzIwMTE2NCwicGF0aCI6Ii8yNTc2NzE2My8zMDI1MjM0MTYtMGQ4YjQ5M2EtMmYzZS00YThhLTliNjAtNzNjYjAxYWNjZGIzLmdpZj9YLUFtei1BbGdvcml0aG09QVdTNC1ITUFDLVNIQTI1NiZYLUFtei1DcmVkZW50aWFsPUFLSUFWQ09EWUxTQTUzUFFLNFpBJTJGMjAyNDAyMDYlMkZ1cy1lYXN0LTElMkZzMyUyRmF3czRfcmVxdWVzdCZYLUFtei1EYXRlPTIwMjQwMjA2VDA2MzI0NFomWC1BbXotRXhwaXJlcz0zMDAmWC1BbXotU2lnbmF0dXJlPWY0MGM1Nzg5NGE1YzRiNTdlMDY3MmJjYzQxNDE2YzlmYWQ1OWFmNGExMmFjZGJhODcyYjUzMjg0NzU3YmI5YmEmWC1BbXotU2lnbmVkSGVhZGVycz1ob3N0JmFjdG9yX2lkPTAma2V5X2lkPTAmcmVwb19pZD0wIn0.ZpXUyb0pmOSUY-5aFBP4Cr3RwVlhQBWuq3-Ij-LBAKw" width="400">

## Get Start
こちらでAPI Keyを取得しましょう。

https://api-ninjas.com/api/emoji

環境変数の設定はenv_sampleを見てね。設定したら.envにrenameしてください。

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
