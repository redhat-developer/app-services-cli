package generate

import (
	"context"
	"io/ioutil"

	"github.com/MakeNowJust/heredoc"

	"github.com/redhat-developer/app-services-cli/pkg/localize"

	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	"github.com/spf13/cobra"
)

type options struct {
	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer
	Context    context.Context

	fileFormat       string
	overwrite        bool
	shortDescription string
	filename         string

	interactive bool
}

func NewGenerateCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		IO:         f.IOStreams,
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate configuration for current profile and active services",
		Long: heredoc.Doc(`
		Command lets you to configure access to all services in one place, 
		by generating or reusing service account and embedding configuration 
		for various services. Configuration is designed to he automatically reusable and
		discoverable by various frameworks like Quarkus, Node.js and Python.

		This command will
		1. Reuse or create service account to be used for service authentication
		2. Provide configuration details for each service 
		3. Let you to drop configuration into your project to autoconfigure it.
		For examples please check https://github.com/rhoas-examples organization

		Supported configuration formats:

		- Local Dev (Environment variables)
		- Kubernetes (configmap and secret)
		- Thrid Party Integrations (JSON format) 
		- Helm (with ArgoCD)
		- RHOAS Operator Config CR

`),
		Example: "",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) (err error) {
			return runCmd(opts)
		},
	}

	cmd.Flags().StringVar(&opts.fileFormat, "file-format", "json", "type of configuration to generate [env, kube, json, helm, rhoas]")

	flagutil.EnableStaticFlagCompletion(cmd, "file-format", []string{"env", "kube", "json", "helm", "rhoas"})

	return cmd
}

func runCmd(opts *options) error {
	if opts.fileFormat == "json" {
		fileData := heredoc.Doc(`
			{
				"kafka": {
					"id": "c694a89bbdnrh1nhmhj0",
					"name": "target",
					"bootstrapHostUrl": "target-1isy6rq3jki8q0otmjqfd3ocfrg.apps.mk-bttg0jn170hp.x5u8.s1.devshift.org",
					"clientId": "339f4248d706",
					"clientSecret": "xxxxxxxxxx"
				},
				"service-registry": {
					"id": "c4b2efb1-7360-4ef1-bc15-b5a5c13c93f7",
					"name": "test",
					"url": "https://registry.apps.example.com/t/5213600b-afc9-487e-8cc3-339f4248d706",
					"clientId": "339f4248d706",
					"clientSecret": "xxxxxxxxxx"
				}
			}
		`)

		ioutil.WriteFile("config.json", []byte(fileData), 0o600)
		opts.Logger.Info("Successfully generated configuration for json format into config.json file")
		return nil
	}

	if opts.fileFormat == "env" {
		fileData := heredoc.Doc(`
			## RHOAS KAFKA
			KAFKA_ID=c694a89bbdnrh1nhmhj0
			KAFKA_HOSTS=target-1isy6rq3jki8q0otmjqfd3ocfrg.apps.mk-bttg0jn170hp.x5u8.s1.devshift.org
			KAFKA_CLIENT_ID=339f4248d706
			KAFKA_CLIENT_SECRET=xxxxxxxxxx

			## RHOAS SERVICE REGISTRY
			SERVICE_REGISTRY_ID=c4b2efb1-7360-4ef1-bc15-b5a5c13c93f7
			SERVICE_REGISTRY_URL=https://registry.apps.example.com/t/5213600b-afc9-487e-8cc3-339f4248d706
			SERVICE_REGISTRY_CLIENT_ID=339f4248d706
			SERVICE_REGISTRY_CLIENT_SECRET=xxxxxxxxxx
		`)

		ioutil.WriteFile("services.env", []byte(fileData), 0o600)
		opts.Logger.Info("Successfully generated configuration env format into services.env file")
		return nil
	}

	if opts.fileFormat == "kube" {
		fileData := heredoc.Doc(`
		apiVersion: v1
		kind: ConfigMap
		metadata:
		  name: rhoas-service-config
		data:
			kafka-id: "c694a89bbdnrh1nhmhj0"
			kafka-bootstrapHostUrl: "target-1isy6rq3jki8q0otmjqfd3ocfrg.apps.mk-bttg0jn170hp.x5u8.s1.devshift.org"
			kafka-auth-secret: "rhoasall-services"
			registry-id: "c4b2efb1-7360-4ef1-bc15-b5a5c13c93f7"
			registry-url: "https://registry.apps.example.com/t/5213600b-afc9-487e-8cc3-339f4248d706"
			registry-auth-secret: "rhoasall-services"
		---
		apiVersion: v1
		kind: Secret
		metadata:
			name: rhoasall-services
		dataString:
			client-id: 339f4248d706
			client-secret: xxxxxxxxxx
		`)

		ioutil.WriteFile("resources.yaml", []byte(fileData), 0o600)
		opts.Logger.Info("Successfully generated configuration rhoas operator kube format into resources.yaml file")
		return nil
	}

	if opts.fileFormat == "rhoas" {
		fileData := heredoc.Doc(`
		apiVersion: rhoas.redhat.com/v1alpha1
		kind: ServiceConfig
		metadata:
		  name: rhoas-service-config
		  labels:
			app.kubernetes.io/component: external-service
			app.kubernetes.io/managed-by: rhoas
		spec:
		   services:
		     - name: "kafka"
			   id: "339f4248d706"
			   serviceAccountSecretName: "rhoasall-services"
			- name: "service-registry"
			   id: "c4b2efb1-7360-4ef1-bc15-b5a5c13c93f7"
			   serviceAccountSecretName: "rhoasall-services"
		---
		apiVersion: v1
		kind: Secret
		metadata:
			name: rhoasall-services
		dataString:
			client-id: 339f4248d706
			client-secret: xxxxxxxxxx
		`)

		ioutil.WriteFile("rhoas-config-cr.yaml", []byte(fileData), 0o600)
		opts.Logger.Info("Successfully generated configuration rhoas operator CR format into rhoas-config-cr.yaml file")
	}

	if opts.fileFormat == "helm" {
		fileData := heredoc.Doc(`
		kafka:
			bootstrapHostUrl: "target-1isy6rq3jki8q0otmjqfd3ocfrg.apps.mk-bttg0jn170hp.x5u8.s1.devshift.org",
			clientId: "339f4248d706",
			clientSecret: "xxxxxxxxxx"

		service-registry: 
			url: "https://registry.apps.example.com/t/5213600b-afc9-487e-8cc3-339f4248d706"
			clientId: "339f4248d706",
			clientSecret: "xxxxxxxxxx"
		`)

		ioutil.WriteFile("rhoas-values.yaml", []byte(fileData), 0o600)
		opts.Logger.Info("Successfully generated configuration for helm into rhoas-values.yaml file")
	}

	return nil
}
