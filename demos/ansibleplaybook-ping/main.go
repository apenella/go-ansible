package main

import (
	"context"
	"fmt"

	"github.com/apenella/go-ansible/modules/builtin/oping"
	"github.com/apenella/go-ansible/modules/enumtipe"
	"github.com/apenella/go-ansible/modules/helper"
	"github.com/apenella/go-ansible/modules/play"
)

func main() {

	op := oping.AnsibleBuiltinPing{
		Name: "ping  " + "localhost" + "[\"" + "127.0.0.1" + "\", ]",
	}

	pb := play.Playbook{
		Hosts:       []string{"127.0.0.1"},
		GatherFacts: enumtipe.CostomBoolFalse,
		Tasks:       []play.ITaskMaker{&op},
	}

	// dire, err := os.MkdirTemp("", "demos")
	// if err != nil {
	// 	panic(err)
	// }
	// file_path := filepath.Join(dire, "playbook.yaml")
	// err = os.WriteFile(file_path, []byte(pb_content), 0644)
	//
	//	if err != nil {
	//		panic(err)
	//	}

	ansible_result, pb_content, duration, err := pb.ExecPlaybook(context.TODO())
	// pb_content, err := pb.MakeAnsibleTask()
	fmt.Println(" === duration: ", duration)
	fmt.Println(" === pb_content: ", pb_content)
	if err != nil {
		panic(err)
	}

	helper.MarshalIndentToString(ansible_result, "", "  ")

}
