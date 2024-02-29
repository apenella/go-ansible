# Example ansibleplaybook-signals-and-cancellation

```sh
❯ go run ansibleplaybook-signals-and-cancellation.go
Go-ansible example ──
Go-ansible example ── PLAY [all] *********************************************************************
Go-ansible example ──
Go-ansible example ── TASK [Gathering Facts] *********************************************************
Go-ansible example ── ok: [127.0.0.1]
Go-ansible example ──
Go-ansible example ── TASK [ansibleplaybook-signals-and-cancellation] ********************************
Go-ansible example ── ok: [127.0.0.1] => {
Go-ansible example ──     "msg": "Your are running 'ansibleplaybook-signals-and-cancellation' example"
Go-ansible example ── }
Go-ansible example ──
Go-ansible example ── PLAY RECAP *********************************************************************
Go-ansible example ── 127.0.0.1                  : ok=2    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0
Go-ansible example ──
```