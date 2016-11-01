package parsehub

import (
	"log"
	"os"
	"fmt"
)

func ExampleSetLogger() {
	logger := &log.Logger{}
	logger.SetOutput(os.Stdout)

	SetLogger(LogLevelDebug, logger)
}

func ExampleProject_Run() {
	parsehub := NewParseHub("__API_KEY__")

	if project, err := parsehub.GetProject("__PROJECT_TOKEN__"); err != nil {
		// handle error
	} else {
		// async run
		project.Run(ProjectRunParams{
			StartTemplate: "__START_TEMPLATE__",
			StartUrl: "__START_URL__",
		}, func(run *Run) error {
			val := map[string]interface{}{}

			if err := run.LoadData(&val); err != nil {
				log.Fatalf(err.Error())
			}

			fmt.Println("result", val)

			// delete after extract data
			if err := run.Delete(); err != nil {
				log.Fatalf(err.Error())
			}
			return nil
		})
	}

	// code for save main thread
}

func ExampleParseHub_GetProject() {
	parsehub := NewParseHub("__API_KEY__")

	if project, err := parsehub.GetProject("__PROJECT_TOKEN__"); err != nil {
		// handle error
	} else {
		fmt.Printf("%+v", project)
	}
}

func ExampleParseHub_GetAllProjects() {
	parsehub := NewParseHub("__API_KEY__")

	if projects, err := parsehub.GetAllProjects(); err != nil {
		// handle error
	} else {
		for _, project := range projects {
			fmt.Printf("%+v", project)
		}
	}
}

func ExampleParseHub_GetRun() {
	parsehub := NewParseHub("__API_KEY__")

	if run, err := parsehub.GetRun("__RUN_TOKEN__"); err != nil {
		// handle error
	} else {
		fmt.Printf("%+v", run)
	}
}
