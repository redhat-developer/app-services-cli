// Cluster package provides way to work with RHOAS and Service binding operators
// Package will ofer integration with various RHOAS services
//
// Structure:
//
// v1alpha - end user api
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
// kubeClients, err := kubeclient.NewKubernetesClusterClients(cliProperties, opts.kubeconfigLocation)
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
package cluster
