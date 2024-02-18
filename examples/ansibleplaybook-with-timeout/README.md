# Example ansibleplaybook-with-timeout

```sh
❯ go run ansibleplaybook-with-timeout.go
Timeout: 15 seconds
Go-ansible example ──
Go-ansible example ── PLAY [all] *********************************************************************
Go-ansible example ──
Go-ansible example ── TASK [Gathering Facts] *********************************************************
Go-ansible example ── ok: [127.0.0.1]
Go-ansible example ──
Go-ansible example ── TASK [ansibleplaybook-with-timeout] ********************************************
Go-ansible example ── ok: [127.0.0.1] => {
Go-ansible example ──     "msg": "Your are running 'ansibleplaybook-with-timeout' example"
Go-ansible example ── }
Go-ansible example ──
Go-ansible example ── PLAY RECAP *********************************************************************
Go-ansible example ── 127.0.0.1                  : ok=2    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0
Go-ansible example ──
```
