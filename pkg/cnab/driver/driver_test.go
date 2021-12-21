package driver

import (
	"context"
	"testing"

	"github.com/deislabs/cnab-go/driver"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

var _ driver.Driver = &Kubernetes{}

func newFakeKubernetesDriver() *Kubernetes {
	return &Kubernetes{
		client:    fake.NewFakeClient(),
		namespace: "fake-namespace",
	}
}

func newFakeOperation() *driver.Operation {
	return &driver.Operation{
		Installation: "test",
		Image:        "test:1.2.3",
		ImageType:    "oci",
		Revision:     "revision",
		Action:       "install",
	}
}

func TestHandles(t *testing.T) {
	d := newFakeKubernetesDriver()
	imgTypes := []string{driver.ImageTypeOCI, driver.ImageTypeDocker}

	for _, imgType := range imgTypes {
		if !d.Handles(imgType) {
			t.Fatalf("Kubernetes driver does not handle image type %v", imgType)
		}
	}
}

func TestCreateSecret(t *testing.T) {
	d := newFakeKubernetesDriver()
	op := newFakeOperation()
	err := d.Run(op)
	if err != nil {
		t.Fatalf("cannot run operation: %v", err)
	}

	s := &corev1.Secret{}
	key := client.ObjectKey{Name: genName(op.Installation, op.Revision), Namespace: d.namespace}
	err = d.client.Get(context.Background(), key, s)
	if err != nil {
		t.Fatalf("cannot get secret: %v", err)
	}

	testLabels(t, op, s.ObjectMeta.Labels)

	// TODO - @radu-matei
	//
	// also test the files and environment from the secret
}

func TestCreatePod(t *testing.T) {
	d := newFakeKubernetesDriver()
	op := newFakeOperation()
	err := d.Run(op)
	if err != nil {
		t.Fatalf("cannot run operation: %v", err)
	}

	p := &corev1.Pod{}
	key := client.ObjectKey{Name: genName(op.Installation, op.Revision), Namespace: d.namespace}
	err = d.client.Get(context.Background(), key, p)
	if err != nil {
		t.Fatalf("cannot get pod: %v", err)
	}

	testLabels(t, op, p.ObjectMeta.Labels)

	if p.Spec.RestartPolicy != corev1.RestartPolicyNever {
		t.Fatalf("expected pod restart policy to be %v, got %v", corev1.RestartPolicyNever, p.Spec.RestartPolicy)
	}

	if len(p.Spec.Containers) != 1 {
		t.Fatalf("expected pod to have %v containers, got %v", 1, len(p.Spec.Containers))
	}

	if p.Spec.Containers[0].Name != "invocationimage" {
		t.Fatalf("expected container name to be %v, got %v", "invocationimage", p.Spec.Containers[0].Name)
	}

	if p.Spec.Containers[0].Image != op.Image {
		t.Fatalf("expected container image to be %v, got %v", op.Image, p.Spec.Containers[0].Image)
	}

	if p.Spec.Containers[0].Env[0].Name != "CNAB_ACTION" {
		t.Fatalf("expected first environment variable name to be %v, got %v", "CNAB_ACTION", p.Spec.Containers[0].Env[0].Name)
	}

	if p.Spec.Containers[0].Env[0].Value != op.Action {
		t.Fatalf("expected first environment variable value to be %v, got %v", op.Action, p.Spec.Containers[0].Env[0].Value)
	}
}

func testLabels(t *testing.T, op *driver.Operation, labels map[string]string) {
	rl, ok := labels["release"]
	if !ok || rl != op.Installation {
		t.Fatalf("label for %s not correctly set in secret: expected %s, got %s", "release", op.Installation, rl)
	}

	al, ok := labels["action"]
	if !ok || al != op.Action {
		t.Fatalf("label for %s not correctly set in secret: expected %s, got %s", "action", op.Action, al)
	}

	hl, ok := labels["heritage"]
	if !ok || hl != cnabOperator {
		t.Fatalf("label for %s not correctly set in secret: expected %s, got %s", "heritage", cnabOperator, hl)
	}

	rvl, ok := labels["revision"]
	if !ok || rvl != op.Revision {
		t.Fatalf("label for %s not correctly set in secret: expected %s, got %s", "revision", op.Revision, rvl)
	}
}
