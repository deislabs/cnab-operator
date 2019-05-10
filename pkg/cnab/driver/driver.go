package driver

import (
	"context"
	"fmt"
	"strings"

	"github.com/deislabs/cnab-go/driver"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const cnabOperator = "operator.cnab.io"

// Kubernetes represents a CNAB driver for Kubernetes
type Kubernetes struct {
	client    client.Client
	namespace string
}

// Handles returns true if the image type is docker or oci.
func (d *Kubernetes) Handles(imgType string) bool {
	switch strings.ToLower(imgType) {
	case "docker", "oci":
		return true
	}
	return false
}

// Run executes the operation inside the invocation image
func (d *Kubernetes) Run(op *driver.Operation) error {
	if !d.Handles(op.ImageType) {
		return fmt.Errorf("driver for Kubernetes does not handle type %q", op.ImageType)
	}

	runName := genName(op.Installation, op.Revision)
	err := d.createSecret(runName, op)
	if err != nil {
		return fmt.Errorf("cannot create secret: %v", err)
	}

	// TODO - @radu-matei
	// after the invocation image pod has finished, this secret should be deleted
	// the reconcile function for Bundles should *not* wait for the pod to finish,
	// but rather there should be an additional controller that watches for invocation
	// pods and after they exit, it should also delete the secrets used.
	// This currently not implemented.
	//defer d.destroySecret(runName)

	return d.createPod(runName, op)
}

func genName(installation, revision string) string {
	return strings.ToLower(fmt.Sprintf("%s-%s", installation, revision))
}

// createSecret creates a secret that stores all of the envs and files
func (d *Kubernetes) createSecret(name string, op *driver.Operation) error {
	combinedSecrets := map[string]string{}
	for k, v := range op.Environment {
		combinedSecrets[k] = v
	}
	for k, v := range op.Files {
		combinedSecrets[k] = v
	}

	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: d.namespace,
			Labels: map[string]string{
				"release":  op.Installation,
				"action":   op.Action,
				"heritage": cnabOperator,
				"revision": op.Revision,
			},
		},
		Type: cnabOperator,
	}
	secret.StringData = combinedSecrets

	return d.client.Create(context.Background(), secret)
}

// destroySecret deletes the secret that stored all envs and files
// after the invocation image was executed
// func (d *Kubernetes) destroySecret(name string) error {
// 	s := &corev1.Secret{}
// 	key := client.ObjectKey{Name: name, Namespace: d.namespace}
// 	err := d.client.Get(context.Background(), key, s)
// 	if err != nil {
// 		return fmt.Errorf("cannot get secret %v: %v", name, err)
// 	}

// 	return d.client.Delete(context.Background(), s)
// }

// createPod runs the action in the invocation image as a Kubernetes pod
func (d *Kubernetes) createPod(name string, op *driver.Operation) error {
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: d.namespace,
			Labels: map[string]string{
				"release":  op.Installation,
				"action":   op.Action,
				"heritage": cnabOperator,
				"revision": op.Revision,
			},
		},
		Spec: corev1.PodSpec{
			RestartPolicy: corev1.RestartPolicyNever,
			Containers: []corev1.Container{
				{
					Name:         "invocationimage",
					Image:        op.Image,
					VolumeMounts: []corev1.VolumeMount{},
					Env:          []corev1.EnvVar{},
				},
			},
			Volumes: []corev1.Volume{},
		},
	}

	// Copy env var definitions into pod
	// Because credentials may be passed here, we don't put the values in. Instead, we
	// reference the secret.
	vars := []corev1.EnvVar{
		{
			Name:  "CNAB_ACTION",
			Value: op.Action,
		},
	}
	trueVal := true
	for k := range op.Environment {
		vars = append(vars, corev1.EnvVar{
			Name: k,
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					Key: k,
					LocalObjectReference: corev1.LocalObjectReference{
						Name: name,
					},
					Optional: &trueVal,
				},
			},
		})
	}
	pod.Spec.Containers[0].Env = vars

	// Copy volumes into pod
	// Again, we use secrets because the data inside of these may be credential info
	for k := range op.Files {
		pod.Spec.Volumes = append(pod.Spec.Volumes, corev1.Volume{
			Name: name,
			VolumeSource: corev1.VolumeSource{
				Secret: &corev1.SecretVolumeSource{
					SecretName: name,
					Optional:   &trueVal,
				},
			},
		})

		for i := range pod.Spec.Containers {
			pod.Spec.Containers[i].VolumeMounts = append(
				pod.Spec.Containers[i].VolumeMounts,
				corev1.VolumeMount{
					Name:      name,
					MountPath: k,
				},
			)
		}
	}

	return d.client.Create(context.Background(), pod)
}
