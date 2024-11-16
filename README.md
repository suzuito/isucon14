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
# setup
ansible-playbook -i inventory.ini playbook.isu_g1.setup.yaml
# deploy app(source code, nginx, mysql settings) for isu_g1
ansible-playbook -i inventory.ini playbook.isu_g1.deploy.yaml
```
