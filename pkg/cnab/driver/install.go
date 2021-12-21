package driver

import (
	"fmt"
	"os"

	"github.com/deislabs/cnab-go/action"
	"github.com/deislabs/cnab-go/bundle"
	"github.com/deislabs/cnab-go/claim"
)

// Install installs a CNAB bundle
func Install(b *bundle.Bundle, claimstore claim.Store) error {
	c, err := claim.New("installation-foo")
	if err != nil {
		return fmt.Errorf("cannot create claim: %v", err)
	}

	// TODO - @radu-matei
	//
	// load parameters from the  CRD
	vals := make(map[string]interface{})
	params, err := bundle.ValuesOrDefaults(vals, b)
	if err != nil {
		return fmt.Errorf("cannot get parameter values: %v", err)
	}

	c.Bundle = b
	c.Parameters = params

	i := action.Install{
		Driver: &Kubernetes{},
	}

	// TODO - @radu-matei
	//
	// load credentials from secret
	var creds = make(map[string]string)
	err = i.Run(c, creds, os.Stdout)

	// Even if the action fails, we want to store a claim. This is because
	// we cannot know, based on a failure, whether or not any resources were
	// created. So we want to suggest that the user take investigative action.
	err2 := claimstore.Store(*c)
	if err != nil {
		return fmt.Errorf("Install step failed: %v", err)
	}
	return err2
}
