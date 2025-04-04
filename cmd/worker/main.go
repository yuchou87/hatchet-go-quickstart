package main

import (
	"hatchet-go-quickstart/workflows"

	v1 "github.com/hatchet-dev/hatchet/pkg/v1"
	"github.com/hatchet-dev/hatchet/pkg/v1/worker"
)

func main() {

	hatchet, err := v1.NewHatchetClient()

	if err != nil {
		panic(err)
	}

	w, err := hatchet.Worker(
		worker.CreateOpts{
			Name: "first-workflow-worker",
		},
		worker.WithWorkflows(workflows.FirstWorkflow(&hatchet)),
	)

	if err != nil {
		panic(err)
	}

	err = w.StartBlocking()

	if err != nil {
		panic(err)
	}
}
