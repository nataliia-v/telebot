# telebot
telegram bot (Golang)

## ⚙️ 🛠 Init project

1. Init go module

```bash
go mod init github.com/nataliia-v/telebot  
``` 

2. install codegenerator cobra-cli

```bash
go install github.com/spf13/cobra-cli@latest 
``` 

3. generate main.go and 'cmd' dir
```bash
cobra-cli init 
``` 

4. add version ( generate versionin.go )
```bash
cobra-cli add version 
``` 

5. generate telebot.go

```bash
cobra-cli add telebot 
``` 

6. set version
```bash
go build -ldflags "-X="github.com/nataliia-v/telebot/cmd.appVersion=v1.0.0
``` 

##  🍀 Telegram bot integration  ☘️

1. Add TELE_TOKEN and logic in Run func in [telebot.go](./cmd/telebot.go)

2. Download and install dependencies

```bash
go get 
``` 

3. Build and update version to v1.0.1
```bash
go build -ldflags "-X="github.com/nataliia-v/telebot/cmd.appVersion=v1.0.1
``` 
4. Test 'TELE_TOKEN'
 (return message "telebot v1.0.1 started2023/05/05 13:11:38 Please check TELE_TOKEN env variable. %stelegram: Not Found (404)")

```bash
./telebot start
``` 

5. Create new bot in BotFather, copy token and set token securely wit command: 
```bash
read -s TELE_TOKEN 
``` 
press Enter (or return) and paste token copied earlier.

check token:
```bash
echo $TELE_TOKEN
```

6. Export variable and run app
```bash
export TELE_TOKEN
./telebot start
```


## 💬 🤖 Chat in Telegram bot 

1. Open in Telegram chat @natalka_k_bot and send 

```bash
/start hello
```

