// Copyright 2020 The Merlin Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package imagebuilder

import (
	"fmt"

	"k8s.io/client-go/kubernetes"

	"github.com/caraml-dev/merlin/mlp"
	"github.com/caraml-dev/merlin/models"
)

// NewPredictionJobImageBuilder create ImageBuilder for building docker image of prediction job (batch)
func NewPredictionJobImageBuilder(kubeClient kubernetes.Interface, config Config) ImageBuilder {
	return newImageBuilder(kubeClient, config, &predictionJobNameGenerator{dockerRegistry: config.DockerRegistry})
}

// predictionJobNameGenerator is name generator that will be used for building docker image of prediction job
type predictionJobNameGenerator struct {
	dockerRegistry string
}

// generateBuilderJobName generate pod name that will be used to build docker image of the prediction job
func (n *predictionJobNameGenerator) generateBuilderJobName(project mlp.Project, model *models.Model, version *models.Version) string {
	return fmt.Sprintf("batch-%s-%s-%s", project.Name, model.Name, version.ID)
}

// generateDockerImageName generate the name of docker image of prediction job that will be created from given model
func (n *predictionJobNameGenerator) generateDockerImageName(project mlp.Project, model *models.Model) string {
	return fmt.Sprintf("%s/%s-%s-job", n.dockerRegistry, project.Name, model.Name)
}
