# telebot
telegram bot (Golang)

## Init project

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

5. generate kbot.go

```bash
cobra-cli add kbot 
``` 
