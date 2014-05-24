package builder

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/tedsuo/router"
	ProleBuilds "github.com/winston-ci/prole/api/builds"
	ProleRoutes "github.com/winston-ci/prole/routes"

	WinstonRoutes "github.com/winston-ci/winston/api/routes"
	"github.com/winston-ci/winston/builds"
	"github.com/winston-ci/winston/config"
	"github.com/winston-ci/winston/db"
)

var ErrBadResponse = errors.New("bad response from prole")

type Builder interface {
	Build(config.Job, config.Resources) (builds.Build, error)
}

type builder struct {
	db db.DB

	prole   *router.RequestGenerator
	winston *router.RequestGenerator
}

func NewBuilder(
	db db.DB,
	prole *router.RequestGenerator,
	winston *router.RequestGenerator,
) Builder {
	return &builder{
		db: db,

		prole:   prole,
		winston: winston,
	}
}

func (builder *builder) Build(job config.Job, resources config.Resources) (builds.Build, error) {
	log.Println("creating build")

	build, err := builder.db.CreateBuild(job.Name)
	if err != nil {
		return builds.Build{}, err
	}

	proleInputs := []ProleBuilds.Input{}
	for name, _ := range job.Inputs {
		resource, found := resources.Lookup(name)
		if !found {
			return builds.Build{}, fmt.Errorf("unknown input: %s", name)
		}

		proleInput := ProleBuilds.Input{
			Type:   resource.Type,
			Source: ProleBuilds.Source(resource.Source),

			DestinationPath: resource.Name,
		}

		if filepath.HasPrefix(job.BuildConfigPath, resource.Name) {
			proleInput.ConfigPath = job.BuildConfigPath[len(resource.Name)+1:]
		}

		proleInputs = append(proleInputs, proleInput)
	}

	complete, err := builder.winston.RequestForHandler(
		WinstonRoutes.UpdateBuild,
		router.Params{
			"job":   job.Name,
			"build": fmt.Sprintf("%d", build.ID),
		},
		nil,
	)
	if err != nil {
		panic(err)
	}

	log.Println("completion callback:", complete.URL)

	logs, err := builder.winston.RequestForHandler(
		WinstonRoutes.LogInput,
		router.Params{
			"job":   job.Name,
			"build": fmt.Sprintf("%d", build.ID),
		},
		nil,
	)
	if err != nil {
		panic(err)
	}

	log.Println("logs callback:", logs.URL)

	logs.URL.Scheme = "ws"

	proleBuild := ProleBuilds.Build{
		Privileged: job.Privileged,

		Inputs: proleInputs,

		Callback: complete.URL.String(),
		LogsURL:  logs.URL.String(),
	}

	log.Printf("creating build: %#v\n", proleBuild)

	req := new(bytes.Buffer)

	err = json.NewEncoder(req).Encode(proleBuild)
	if err != nil {
		return builds.Build{}, err
	}

	execute, err := builder.prole.RequestForHandler(
		ProleRoutes.ExecuteBuild,
		nil,
		req,
	)
	if err != nil {
		return builds.Build{}, err
	}

	execute.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(execute)
	if err != nil {
		log.Println("prole request failed:", err)
		return builds.Build{}, err
	}

	// TODO test bad response code
	if resp.StatusCode != http.StatusCreated {
		log.Println("bad prole response:", resp)
		return builds.Build{}, ErrBadResponse
	}

	resp.Body.Close()

	return build, nil
}
