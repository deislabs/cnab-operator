package claimstore

import (
	"testing"

	"github.com/deislabs/duffle/pkg/claim"
	"github.com/deislabs/duffle/pkg/utils/crud"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

// make sure configMapStore implements the Duffle CRUD store
var _ crud.Store = &configMapStore{}
var _ claim.Store = claim.NewClaimStore(&configMapStore{})

func newFakeStore() claim.Store {
	return claim.NewClaimStore(NewConfigMapStore("fake-namespace", fake.NewFakeClient()))
}

func newPopulatedFakeStore() claim.Store {
	s := newFakeStore()

	c1, _ := claim.New("foo")
	c2, _ := claim.New("bar")
	s.Store(*c1)
	s.Store(*c2)

	return s
}

func TestCreateClaim(t *testing.T) {
	c, err := claim.New("foo")
	if err != nil {
		t.Fatalf("cannot create claim foo: %v", err)
	}

	s := newFakeStore()

	err = s.Store(*c)
	if err != nil {
		t.Fatalf("cannot store claim foo: %v", err)
	}

	claims, err := s.ReadAll()
	if err != nil {
		t.Fatalf("cannot read all claims: %v", err)
	}

	if len(claims) != 1 {
		t.Fatalf("expected length of claims list to be %v, got %v", 1, len(claims))
	}
}

func TestReadClaim(t *testing.T) {
	s := newPopulatedFakeStore()

	c, err := s.Read("foo")
	if err != nil {
		t.Fatalf("cannot read back claim foo from store: %v", err)
	}

	if c.Name != "foo" {
		t.Fatalf("expected claim name to be %s, got %s", "foo", c.Name)
	}
}

func TestUpdateClaim(t *testing.T) {
	s := newPopulatedFakeStore()

	c, err := s.Read("foo")
	if err != nil {
		t.Fatalf("cannot read back claim foo from store: %v", err)
	}

	c.Revision = "fake-revision"
	err = s.Store(c)
	if err != nil {
		t.Fatalf("cannot update claim: %v", err)
	}

	c2, err := s.Read("foo")
	if err != nil {
		t.Fatalf("cannot read back updated claim foo from store: %v", err)
	}

	if c2.Name != "foo" {
		t.Fatalf("expected claim name to be %s, got %s", "foo", c2.Name)
	}

	if c2.Revision != "fake-revision" {
		t.Fatalf("expected claim name to be %s, got %s", "fake-revision", c2.Revision)
	}
}

func TestDeleteClaim(t *testing.T) {
	s := newPopulatedFakeStore()
	err := s.Delete("foo")

	claims, err := s.ReadAll()
	if err != nil {
		t.Fatalf("cannot get all claims: %v", err)
	}

	if len(claims) != 1 {
		t.Fatalf("expected length of claims list to be %v, got %v", 1, len(claims))
	}
}

func TestClaimStoreLength(t *testing.T) {
	s := newPopulatedFakeStore()
	claims, err := s.ReadAll()
	if err != nil {
		t.Fatalf("cannot get all claims: %v", err)
	}

	if len(claims) != 2 {
		t.Fatalf("expected length of claims list to be %v, got %v", 2, len(claims))
	}
}
