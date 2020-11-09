/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-logr/logr"
	"github.com/golang/glog"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	mas "github.com/bf2fc6cc711aee1a0c2a/cli/client/mas"
	v1 "github.com/bf2fc6cc711aee1a0c2a/cli/operator/api/v1"
)

// ManagedKafkaReconciler reconciles a ManagedKafka object
type ManagedKafkaReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=streaming.my.domain,resources=managedkafkas,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=streaming.my.domain,resources=managedkafkas/status,verbs=get;update;patch

func (r *ManagedKafkaReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	_ = r.Log.WithValues("managedkafka", req.NamespacedName)

	kafka := &v1.ManagedKafka{}
	r.Get(ctx, req.NamespacedName, kafka)

	client := BuildMasClient()

	kafkaRequest := mas.KafkaRequest{Name: *&kafka.Spec.Name, Region: *&kafka.Spec.Region, CloudProvider: *&kafka.Spec.CloudProvider}

	response, status, err := client.DefaultApi.ApiManagedServicesApiV1KafkasPost(context.Background(), false, kafkaRequest)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while requesting new Kafka cluster: %v", err)
	}
	if status.StatusCode == 200 {
		jsonResponse, _ := json.MarshalIndent(response, "", "  ")
		fmt.Print("Created API \n ", string(jsonResponse))
		return ctrl.Result{}, nil
	} else {
		fmt.Print("Creation failed", response, status)
		return ctrl.Result{}, status
	}

}

func BuildMasClient() *mas.APIClient {
	// TODO config abstraction
	testHost := "localhost:8000"
	testScheme := "http"
	// Based on https://github.com/OpenAPITools/openapi-generator/blob/master/samples/client/petstore/go/pet_api_test.go

	cfg := mas.NewConfiguration()
	// TODO read flag from config
	cfg.AddDefaultHeader("Authorization", "Bearer 9f4068b1c2cc720dd44dc2c6157569ae")
	cfg.Host = testHost
	cfg.Scheme = testScheme
	client := mas.NewAPIClient(cfg)

	return client
}

func (r *ManagedKafkaReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1.ManagedKafka{}).
		Complete(r)
}
