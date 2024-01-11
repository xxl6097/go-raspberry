#!/bin/bash

function menu() {
  echo "1. 列出运行容器"
  echo "2. 列出所有容器"
  echo "3. 列出所有镜像"
  echo "4. 启动容器(up)"
  echo "5. 停止容器(down)"
  echo "6. 重置容器(down->up)"
  echo "7. 启动容器(up -d)"
  echo "请输入编号:"
  read index

  case "$index" in
  [1]) (list_container) ;;
  [2]) (list_container -a) ;;
  [3]) (list_images) ;;
  [4]) (docker_up) ;;
  [5]) (docker_down) ;;
  [6]) (docker_reset) ;;
  [7]) (docker_up_d) ;;
  *) echo "exit" ;;
  esac

}

function docker_up_d() {
  docker compose up -d
}

function docker_up() {
  docker compose up
}
function docker_down() {
  docker compose down
}

function docker_reset() {
  docker compose down && docker compose up
}

function list_images() {
  clear
  # 使用 IFS 和 while read 将 docker images 的输出存入数组
  IFS=$'\n' read -r -d '' -a docker_images_array < <(docker images | tail -n +2 && printf '\0')
  index=0
  for image in "${docker_images_array[@]}"; do
    echo "$index: $image"
    ((index++))
  done

  echo "请按输入要删除的序号:"
  read pos
  line=${docker_images_array[${pos}]}
  _sel=$(echo "$line" | awk '{print $3}')
  docker rmi ${_sel}
}

function list_container() {
  clear
  IFS=$'\n' read -r -d '' -a docker_ps_array < <(docker ps $1 --format "table {{.ID}}\t{{.Names}}\t{{.Image}}\t{{.Ports}}" | tail -n +2 && printf '\0')
  index=0
  for line in "${docker_ps_array[@]}"; do
    _pos=$(printf "%2d" "$index")
    echo "$_pos: $line"
    ((index++))
  done
  echo "请按照序号选择:"
  read pos
  line=${docker_ps_array[${pos}]}
  _sel=$(echo "$line" | awk '{print $1}')
  echo "1. 停止容器(stop)"
  echo "2. 删除容器(stop->rm)"
  echo "3. 重启容器(stop->start)"
  echo "4. 查看日志"
  echo "请输入编号:"
  read index
  case "$index" in
  [1]) (stop_container $_sel) ;;
  [2]) (rm_container $_sel) ;;
  [3]) (restart_container $_sel) ;;
  [4]) (docker_log $_sel) ;;
  *) echo "exit" ;;
  esac
}

function stop_container() {
  docker stop $1
}
function rm_container() {
  docker stop $1 && docker rm ${1}
}
function restart_container() {
  docker restart $1
}
function docker_log() {
  docker logs -f $1
}
menu
