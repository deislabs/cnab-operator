package claimstore

import (
	"context"
	"fmt"

	"github.com/deislabs/cnab-go/utils/crud"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ClaimLabelSelector represents the label applied to all claims
const ClaimLabelSelector = "claims.cnab.io/v1alpha1"

// ClaimKey represents the key under which the claim is stored in the ConfigMap
const ClaimKey = "claim"

// ConfigMapStore represents a CNAB claim store
// backed by Kubernetes claims, stored in the operator namespace
//
// TODO - @radu-matei
// decide if storing claims is better suited for a CRD.
// See https://github.com/deislabs/cnab-operator/issues/8
type configMapStore struct {
	namespace string
	client    client.Client
}

// NewConfigMapStore returns a new claim store
// backed by Kubernetes ConfigMaps
func NewConfigMapStore(namespace string, client client.Client) crud.Store {
	return &configMapStore{namespace: namespace, client: client}
}

// List returns all claims in the operator namespace
func (s *configMapStore) List() ([]string, error) {
	otps := &client.ListOptions{}
	otps.InNamespace(s.namespace)
	otps.SetLabelSelector(ClaimLabelSelector)

	cml := &corev1.ConfigMapList{}

	err := s.client.List(context.Background(), otps, cml)
	if err != nil {
		return nil, fmt.Errorf("cannot list claims from the ConfigMap store: %v", err)
	}

	items := []string{}
	for _, cm := range cml.Items {
		name := cm.Name
		items = append(items, name)
	}

	return items, nil
}

// Store saves a new claim as a Kubernetes ConfigMap
func (s *configMapStore) Store(name string, data []byte) error {
	cm := &corev1.ConfigMap{}
	key := client.ObjectKey{Name: name, Namespace: s.namespace}
	err := s.client.Get(context.Background(), key, cm)
	if err != nil {
		if errors.IsNotFound(err) {
			return s.createConfigMap(name, s.namespace, data)
		}
		return err
	}
	return s.updateConfigMap(cm, data)
}

// Read returns the value of a claim, or error if it doesn't exist
func (s *configMapStore) Read(name string) ([]byte, error) {
	cm := &corev1.ConfigMap{}
	key := client.ObjectKey{Name: name, Namespace: s.namespace}
	err := s.client.Get(context.Background(), key, cm)
	if err != nil {
		return nil, fmt.Errorf("cannot get config map: %v", err)
	}

	return cm.BinaryData[ClaimKey], nil
}

// Delete removes a claim from the Kubernetes ConfigMaps
func (s *configMapStore) Delete(name string) error {
	cm := &corev1.ConfigMap{}
	key := client.ObjectKey{Name: name, Namespace: s.namespace}
	err := s.client.Get(context.Background(), key, cm)
	if err != nil {
		return fmt.Errorf("cannot get config map: %v", err)
	}

	return s.client.Delete(context.Background(), cm, client.GracePeriodSeconds(5))
}

func (s *configMapStore) createConfigMap(name string, namespace string, data []byte) error {
	cm := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		BinaryData: map[string][]byte{
			ClaimKey: data,
		},
	}

	return s.client.Create(context.Background(), cm)
}

// TODO - @radu-matei
//
// Should consumers of this package be able to change the name of a claim
// and store it back? Under the current ConfigMap store logic, it becomes another claim
func (s *configMapStore) updateConfigMap(cm *corev1.ConfigMap, data []byte) error {
	cm.BinaryData[ClaimKey] = data
	return s.client.Update(context.Background(), cm)
}
