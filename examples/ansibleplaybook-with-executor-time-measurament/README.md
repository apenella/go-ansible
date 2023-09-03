# Example ansibleplaybook-with-executor-time-measurament

```sh
❯ go run ansibleplaybook-with-executor-time-measurament.go
2023-07-27 08:27:45     Go-ansible example ──
2023-07-27 08:27:45     Go-ansible example ── PLAY [all] *********************************************************************
2023-07-27 08:27:45     Go-ansible example ──
2023-07-27 08:27:45     Go-ansible example ── TASK [Gathering Facts] *********************************************************
2023-07-27 08:27:48     Go-ansible example ── ok: [127.0.0.1]
2023-07-27 08:27:48     Go-ansible example ──
2023-07-27 08:27:48     Go-ansible example ── TASK [ansibleplaybook-wit-executor-time-measurement] ***************************
2023-07-27 08:27:48     Go-ansible example ── ok: [127.0.0.1] => {
2023-07-27 08:27:48     Go-ansible example ──     "msg": "Your are running 'ansibleplaybook-wit-executor-time-measurement' example"
2023-07-27 08:27:48     Go-ansible example ── }
2023-07-27 08:27:48     Go-ansible example ──
2023-07-27 08:27:48     Go-ansible example ── PLAY RECAP *********************************************************************
2023-07-27 08:27:48     Go-ansible example ── 127.0.0.1                  : ok=2    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0
2023-07-27 08:27:48     Go-ansible example ──

        Duration: 3.397462474s

```
