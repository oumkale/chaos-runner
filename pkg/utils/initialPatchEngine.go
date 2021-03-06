package utils

import (
	"github.com/litmuschaos/chaos-operator/pkg/apis/litmuschaos/v1alpha1"
	"github.com/pkg/errors"
)

// ExperimentStatus is wrapper for v1alpha1.ExperimentStatuses
type ExperimentStatus v1alpha1.ExperimentStatuses

// InitialPatchEngine patches the chaosEngine with the initial ExperimentStatuses
func InitialPatchEngine(engineDetails EngineDetails, clients ClientSets, experimentList []ExperimentDetails) error {

	// Get chaosengine Object
	expEngine, err := engineDetails.GetChaosEngine(clients)
	if err != nil {
		return errors.Errorf("Unable to get ChaosEngine, error: %v", err)
	}

	// patch the experiment status in chaosengine
	for _, v := range experimentList {
		var expStatus ExperimentStatus
		expStatus.InitialExperimentStatus(v.Name, engineDetails.Name)
		expEngine.Status.Experiments = append(expEngine.Status.Experiments, v1alpha1.ExperimentStatuses(expStatus))
	}
	_, updateErr := clients.LitmusClient.LitmuschaosV1alpha1().ChaosEngines(engineDetails.EngineNamespace).Update(expEngine)
	if updateErr != nil {
		return errors.Errorf("Unable to update ChaosEngine in namespace: %v, error: %v", engineDetails.EngineNamespace, err)
	}
	return nil
}
