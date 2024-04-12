# Ansible Playbook Embed Python

> **Disclaimer**: This example demonstrates how to use `go-ansible` library along with the `go-embed-python` solution. While this is not a feature of the library itself, its compatibility with various Ansible playbooks could vary depending on the requirements of the playbook. However, it serves as a source of inspiration for creating your solution.

- [Ansible Playbook Embed Python](#ansible-playbook-embed-python)
  - [Considerations](#considerations)
  - [Run the example](#run-the-example)

## Considerations

This example does not include the `go-embed-python` packages within the `data` folder. You need to install them manually by running the following command:

```sh
go generate ./...
```

As a result, the `data` folder will be created with the embedded Python packages.

## Run the example

This example provides a Dockerfile that encapsulates all the necessary steps to generate the embedded Python packages, create a Go binary containing the Ansible playbook, and the embedded Python packages.

The example installs the `ansible-core` package. That is the minimal package required to run the Ansible playbook embedded in the Go binary. You can include additional packages by adding them to the `internal/ansibleplaybook-embed-python/requirements.txt` file.

The Ansible playbook used in this example is defined in the `resources/ansible/site.yml` file. When the application is executed, the playbook is unpacked from the Go binary into a temporary directory. The embedded Python packages are also unpacked into another temporary directory. The playbook is then executed using the embedded Python interpreter.

You can run the example with the following command:

```sh
$ make run
[+] Creating 1/1
 ✔ Network go-ansible-ansibleplaybook-embed-python_default  Created                                                                                                  0.1s
[+] Building 52.2s (19/19) FINISHED                                                                                                                        docker:default
 => [app internal] load build definition from Dockerfile                                                                                                             0.0s
 => => transferring dockerfile: 812B                                                                                                                                 0.0s
 => [app internal] load metadata for docker.io/library/debian:bookworm-slim                                                                                          0.5s
 => [app internal] load metadata for docker.io/library/golang:1.19-bookworm                                                                                          0.5s
 => [app internal] load metadata for docker.io/library/python:3.12-bookworm                                                                                          0.0s
 => [app internal] load .dockerignore                                                                                                                                0.0s
 => => transferring context: 2B                                                                                                                                      0.0s
 => [app builder 1/9] FROM docker.io/library/python:3.12-bookworm                                                                                                    0.0s
 => [app internal] load build context                                                                                                                                0.2s
 => => transferring context: 852.70kB                                                                                                                                0.1s
 => [app stage-2 1/2] FROM docker.io/library/debian:bookworm-slim@sha256:3d5df92588469a4c503adbead0e4129ef3f88e223954011c2169073897547cac                            0.0s
 => [app golang 1/1] FROM docker.io/library/golang:1.19-bookworm@sha256:da9da58d86d106a5dda2ce249b00cf3b31cdd626ea41597e476de7b4eebad8c4                             0.0s
 => CACHED [app builder 2/9] RUN apt-get update     && apt-get install -y         openssh-client         git     && rm -rf /var/lib/apt/lists/*                      0.0s
 => CACHED [app builder 3/9] COPY --from=golang /usr/local/go /usr/local/go                                                                                          0.0s
 => CACHED [app builder 4/9] WORKDIR /app                                                                                                                            0.0s
 => CACHED [app builder 5/9] COPY go.mod go.sum ./                                                                                                                   0.0s
 => CACHED [app builder 6/9] RUN go mod download                                                                                                                     0.0s
 => [app builder 7/9] COPY . .                                                                                                                                       0.7s
 => [app builder 8/9] RUN go generate ./...                                                                                                                         47.0s
 => [app builder 9/9] RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /usr/local/bin/app ansibleplaybook-embed-python.go                                              3.8s
 => CACHED [app stage-2 2/2] COPY --from=builder /usr/local/bin/app /usr/local/bin/app                                                                               0.0s
 => [app] exporting to image                                                                                                                                         0.0s
 => => exporting layers                                                                                                                                              0.0s
 => => writing image sha256:9bac9ee8a88b138b87099e2d14a8b45d0b2313c59bb7f809ec27b529922235dc                                                                         0.0s
 => => naming to docker.io/library/go-ansible-ansibleplaybook-embed-python-app                                                                                       0.0s

PLAY [all] *********************************************************************

TASK [ansibleplaybook-simple] **************************************************
ok: [127.0.0.1] => {
    "msg": "Your are running 'ansibleplaybook-embed-python' example"
}

PLAY RECAP *********************************************************************
127.0.0.1                  : ok=1    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0
```
