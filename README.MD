## 添加依赖

```shell
go get -u github.com/mdp/qrterminal/v3 v3.0.0

go get -u github.com/wechaty/go-wechaty

go get -u github.com/xxl6097/go-http@v0.0.7

go get -u github.com/gorilla/mux

go get -u github.com/google/uuid

go get -u github.com/dgrijalva/jwt-go

go get -u gopkg.in/yaml.v3

docker cp wxBotWebhook:/conf/app.yaml .

```


alias dps='docker ps $1 --format "table {{.ID}}\t{{.Names}}\t{{.Image}}\t{{.Status}}\t{{.Ports}}"'

alias dbash='docker exec -it $1 /bin/bash'
alias dsh='docker exec -it $1 /bin/sh'
alias drm='docker stop $1 && docker rm $1'
alias ddel='docker rmi $1'
alias dlog='docker logs -f $1'


