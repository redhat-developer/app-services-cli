// Cluster package provides way to work with RHOAS and Service binding operators
// Package will ofer integration with various RHOAS services
//
// Structure:
//
// v1alpha - end user API
// kubeclient - provides kubernetes clients
// service - individual service implementations
//
//
// Usage:
//
// cliProperties := v1alpha.CommandEnvironment{
//     IO:         opts.IO,
//     Logger:     opts.Logger,
//     Localizer:  opts.localizer,
//     Config:     opts.Config,
//     Connection: conn,
// }
//
// kubeClients, err := kubeclient.NewKubernetesClusterClients(&cliProperties, opts.kubeconfigLocation)
// if err != nil {
//     return err
// }
//
// clusterAPI := cluster.KubernetesClusterAPIImpl{
//     KubernetesClients:  kubeClients,
//     CommandEnvironment: &cliProperties,
// }
//
// err = clusterAPI.ExecuteServiceBinding(&v1alpha.BindOperationOptions{
//     ServiceName:             opts.serviceName,
//     Namespace:               opts.namespace,
//     AppName:                 opts.appName,
//     ForceCreationWithoutAsk: opts.forceCreationWithoutAsk,
//     BindingName:             opts.bindingName,
//     BindAsFiles:             !opts.bindAsEnv,
//     DeploymentConfigEnabled: opts.deploymentConfigEnabled,
// })
//
// Adding new service in "4 Simple Steps" ™
//
// 1. Review services/defitinions.go and create definitions for your own service
// 2. Add new file with Structure that represent your own service CRD.
// For example see services/resources/KafkaConnection.go
// 3. Copy ./services/kafka.go and implement it for your own service.
// 4. Add reference to the latest resource in status.go to check if operator is up to date
//
// After service is created you need to review `createServiceInstance`
// method that assings specific instance of service depending of string provided by user.
package cluster
