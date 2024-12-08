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
ansible-playbook -i inventory.deploy.ini playbook.deploy.yaml && ansible-playbook -i inventory.deploy.ini playbook.restart.yaml
ansible-playbook -i inventory.metrics.ini --extra-vars "dt=$(date +%s)" playbook.metrics.yaml
```
