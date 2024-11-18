# isucon14

## Installation

### Ansible

```bash
brew install ansible
```

## Setup

### Ansible

- `inventory.ini`
  - サーバーホスト名の一覧
  - sshの設定

サーバーホストへの接続確認

```bash
# isu_g1 はサーバーのグループ
ansible isu_g1 -m ping -i inventory.ini
```

サーバーホストへのplaybook

#### isu_g1

```bash
# deploy app(source code, nginx, mysql settings) for isu_g1
cd webapp/go/ && CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o a.out cmd/isuports/*.go && cd - && \
ansible-playbook -i inventory.ini playbook.isu_g1.deploy.yaml && \
ansible-playbook -i inventory.ini playbook.isu_g1.restart.yaml
# metrics
ansible-playbook -i inventory.ini --extra-vars "dt=$(date +%s)" playbook.isu_g1.metrics.yaml
```
