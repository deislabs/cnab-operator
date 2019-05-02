// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/deislabs/cnab-operator/pkg/apis/cnab/v1alpha1.Bundle":       schema_pkg_apis_cnab_v1alpha1_Bundle(ref),
		"github.com/deislabs/cnab-operator/pkg/apis/cnab/v1alpha1.BundleSpec":   schema_pkg_apis_cnab_v1alpha1_BundleSpec(ref),
		"github.com/deislabs/cnab-operator/pkg/apis/cnab/v1alpha1.BundleStatus": schema_pkg_apis_cnab_v1alpha1_BundleStatus(ref),
	}
}

func schema_pkg_apis_cnab_v1alpha1_Bundle(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Bundle is the Schema for the bundles API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/deislabs/cnab-operator/pkg/apis/cnab/v1alpha1.BundleSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/deislabs/cnab-operator/pkg/apis/cnab/v1alpha1.BundleStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/deislabs/cnab-operator/pkg/apis/cnab/v1alpha1.BundleSpec", "github.com/deislabs/cnab-operator/pkg/apis/cnab/v1alpha1.BundleStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_cnab_v1alpha1_BundleSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BundleSpec defines the desired state of Bundle",
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name of the bundle",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"url": {
						SchemaProps: spec.SchemaProps{
							Description: "URL of the bundle in a remote OCI registry",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_cnab_v1alpha1_BundleStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BundleStatus defines the observed state of Bundle",
				Properties: map[string]spec.Schema{
					"pulled": {
						SchemaProps: spec.SchemaProps{
							Description: "Pulled indicates the controller successfully pulled the bundle from the registry",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"processed": {
						SchemaProps: spec.SchemaProps{
							Description: "Processed indicates the controller successfully processed the pulled bundle",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}
