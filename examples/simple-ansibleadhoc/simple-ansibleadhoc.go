package main

import ansibler "github.com/apenella/go-ansible"

func main() {

	ansibleConnectionOptions := &ansibler.AnsibleConnectionOptions{
		Connection: "local",
	}

	ansibleAdhocOptions := &ansibler.AnsibleAdhocOptions{
		Inventory:  "127.0.0.1,",
		ModuleName: "ping",
	}

	adhoc := &ansibler.AnsibleAdhocCmd{
		Pattern:           "all",
		Options:           ansibleAdhocOptions,
		ConnectionOptions: ansibleConnectionOptions,
	}

	err := adhoc.Run()
	if err != nil {
		panic(err)
	}
}
