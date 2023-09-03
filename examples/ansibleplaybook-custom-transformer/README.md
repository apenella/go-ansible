# Example ansibleplaybook-custom-transformer

```sh
❯ go run ansibleplaybook-custom-transformer.go
2023-07-27 08:24:30     Go-ansible example ──
2023-07-27 08:24:30     Go-ansible example ── PLAY [all] *********************************************************************
2023-07-27 08:24:30     Go-ansible example ──
2023-07-27 08:24:30     Go-ansible example ── TASK [Gathering Facts] *********************************************************
2023-07-27 08:24:33     Go-ansible example ── ok: [127.0.0.1]
2023-07-27 08:24:33     Go-ansible example ──
2023-07-27 08:24:33     Go-ansible example ── TASK [ansibleplaybook-custom-transformer] **************************************
2023-07-27 08:24:33     Go-ansible example ── ok: [127.0.0.1] => {
2023-07-27 08:24:33     Go-ansible example ──     "msg": "Your are running 'ansibleplaybook-custom-transformer' example"
2023-07-27 08:24:33     Go-ansible example ── }
2023-07-27 08:24:33     Go-ansible example ──
2023-07-27 08:24:33     Go-ansible example ── PLAY RECAP *********************************************************************
2023-07-27 08:24:33     Go-ansible example ── 127.0.0.1                  : ok=2    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0
2023-07-27 08:24:33     Go-ansible example ──
```