# Example ansibleplaybook-cobra-cmd

```sh
$ go run ansibleplaybook-cobra-cmd.go --connection-local  --playbook site.yml --inventory 127.0.0.1, --extra-var example="Hi! There"
cobra-cmd-ansibleplaybook example ──
cobra-cmd-ansibleplaybook example ── PLAY [all] *********************************************************************
cobra-cmd-ansibleplaybook example ──
cobra-cmd-ansibleplaybook example ── TASK [Gathering Facts] *********************************************************
cobra-cmd-ansibleplaybook example ── ok: [127.0.0.1]
cobra-cmd-ansibleplaybook example ──
cobra-cmd-ansibleplaybook example ── TASK [ansibleplaybook-cobra-cmd] ***********************************************
cobra-cmd-ansibleplaybook example ── ok: [127.0.0.1] => {
cobra-cmd-ansibleplaybook example ──     "msg": "Your are running 'ansibleplaybook-cobra-cmd' example"
cobra-cmd-ansibleplaybook example ── }
cobra-cmd-ansibleplaybook example ──
cobra-cmd-ansibleplaybook example ── TASK [Print extra variable 'example'] ******************************************
cobra-cmd-ansibleplaybook example ── ok: [127.0.0.1] => {
cobra-cmd-ansibleplaybook example ──     "msg": "Example value is Hi! There"
cobra-cmd-ansibleplaybook example ── }
cobra-cmd-ansibleplaybook example ──
cobra-cmd-ansibleplaybook example ── PLAY RECAP *********************************************************************
cobra-cmd-ansibleplaybook example ── 127.0.0.1                  : ok=3    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0
cobra-cmd-ansibleplaybook example ──
```