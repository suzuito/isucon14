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
# isu_a はサーバーのグループ
ansible isu_a -m ping -i inventory.ini
```

サーバーホストへのplaybook

```bash
ansible-playbook -i inventory.ini playbook_isu_a.yaml
```
