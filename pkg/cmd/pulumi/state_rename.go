// Copyright 2016-2022, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"errors"
	"fmt"

	"github.com/pulumi/pulumi/pkg/v3/backend/display"
	"github.com/pulumi/pulumi/pkg/v3/resource/deploy"
	"github.com/pulumi/pulumi/pkg/v3/resource/deploy/providers"
	"github.com/pulumi/pulumi/pkg/v3/resource/edit"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/cmdutil"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"

	"github.com/spf13/cobra"
)

func updateDependencies(dependencies []resource.URN, oldUrn resource.URN, newUrn resource.URN) []resource.URN {
	var updatedDependencies []resource.URN
	for _, dependency := range dependencies {
		if dependency == oldUrn {
			// replace old URN with new URN
			updatedDependencies = append(updatedDependencies, newUrn)
		} else {
			updatedDependencies = append(updatedDependencies, dependency)
		}
	}
	return updatedDependencies
}

// stateRenameOperation renames a resource (or provider) and mutates/rewrites references to it in the snapshot.
func stateRenameOperation(
	urn resource.URN, newResourceName tokens.QName, opts display.Options, snap *deploy.Snapshot,
) error {
	// Check whether the input URN corresponds to an existing resource
	existingResources := edit.LocateResource(snap, urn)
	if len(existingResources) != 1 {
		return errors.New("The input URN does not correspond to an existing resource")
	}

	contract.Assertf(tokens.IsQName(string(newResourceName)),
		"QName must be valid")

	inputResource := existingResources[0]
	oldUrn := inputResource.URN
	// update the URN with only the name part changed
	newUrn := oldUrn.Rename(string(newResourceName))
	// Check whether the new URN _does not_ correspond to an existing resource
	candidateResources := edit.LocateResource(snap, newUrn)
	if len(candidateResources) > 0 {
		return errors.New("The chosen new name for the state corresponds to an already existing resource")
	}

	// Update the URN of the input resource
	inputResource.URN = newUrn
	// Update the dependants of the input resource
	for _, existingResource := range snap.Resources {
		// update resources other than the input resource
		if existingResource.URN != inputResource.URN {
			// Update dependencies
			existingResource.Dependencies = updateDependencies(existingResource.Dependencies, oldUrn, newUrn)
			// Update property dependencies
			for property, dependencies := range existingResource.PropertyDependencies {
				existingResource.PropertyDependencies[property] = updateDependencies(dependencies, oldUrn, newUrn)
			}

			// Update parent, if any.
			if existingResource.Parent == oldUrn {
				existingResource.Parent = newUrn
			}
		}
	}

	updateProvider := func(newRef providers.Reference) error {
		// Loop through all resources and rename references to the provider.
		for _, curResource := range snap.Resources {

			if curResource.Provider == "" {
				// Skip resources that don't use a provider.
				continue
			}
			curResourceProviderRef, err := providers.ParseReference(curResource.Provider)
			if err != nil {
				return err
			}

			// Skip resources that don't use the renamed provider.
			if curResourceProviderRef.URN() != oldUrn {
				continue
			}

			// Update the provider.
			curResource.Provider = newRef.String()
		}
		return nil
	}

	// If the renamed resource is a Provider, fix all resources referring to the old name.
	if providers.IsProviderType(inputResource.Type) {
		newRef, err := providers.NewReference(newUrn, inputResource.ID)
		if err != nil {
			return err
		}
		return updateProvider(newRef)
	}

	return nil
}

//nolint:lll
func newStateRenameCommand() *cobra.Command {
	var stack string
	var yes bool

	cmd := &cobra.Command{
		Use:   "rename [resource URN] [new name]",
		Short: "Renames a resource from a stack's state",
		Long: `Renames a resource from a stack's state

This command renames a resource from a stack's state. The resource is specified
by its Pulumi URN and the new name of the resource.

Make sure that URNs are single-quoted to avoid having characters unexpectedly interpreted by the shell.

To see the list of URNs in a stack, use ` + "`pulumi stack --show-urns`" + `.
`,
		Example: "pulumi state rename 'urn:pulumi:stage::demo::eks:index:Cluster$pulumi:providers:kubernetes::eks-provider' new-name-here",
		Args:    cmdutil.MaximumNArgs(2),
		Run: cmdutil.RunFunc(func(cmd *cobra.Command, args []string) error {
			ctx := commandContext()
			yes = yes || skipConfirmations()

			if len(args) < 2 && !cmdutil.Interactive() {
				return missingNonInteractiveArg("resource URN", "new name")
			}

			var urn resource.URN
			var newResourceName tokens.QName
			switch len(args) {
			case 0: // We got neither the URN nor the name.
				var snap *deploy.Snapshot
				err := surveyStack(
					func() (err error) {
						urn, err = getURNFromState(ctx, stack, &snap, "Select a resource to rename:")
						return
					},
					func() (err error) {
						newResourceName, err = getNewResourceName()
						return
					},
				)
				if err != nil {
					return err
				}
			case 1: // We got the urn but not the name
				urn = resource.URN(args[0])
				if !urn.IsValid() {
					return errors.New("The provided input URN is not valid")
				}
				var err error
				newResourceName, err = getNewResourceName()
				if err != nil {
					return err
				}
			case 2: // We got the URN and the name.
				urn = resource.URN(args[0])
				if !urn.IsValid() {
					return errors.New("The provided input URN is not valid")
				}
				rName := args[1]
				if !tokens.IsQName(rName) {
					reason := "resource names may only contain alphanumerics, underscores, hyphens, dots, and slashes"
					return fmt.Errorf("invalid name %q: %s", rName, reason)
				}
				newResourceName = tokens.QName(rName)
			}

			// Show the confirmation prompt if the user didn't pass the --yes parameter to skip it.
			showPrompt := !yes

			err := runTotalStateEdit(ctx, stack, showPrompt,
				func(opts display.Options, snap *deploy.Snapshot) error {
					return stateRenameOperation(urn, newResourceName, opts, snap)
				})
			if err != nil {
				// an error occurred
				// return it
				return err
			}

			fmt.Println("Resource renamed")
			return nil
		}),
	}

	cmd.PersistentFlags().StringVarP(
		&stack, "stack", "s", "",
		"The name of the stack to operate on. Defaults to the current stack")

	cmd.Flags().BoolVarP(&yes, "yes", "y", false, "Skip confirmation prompts")
	return cmd
}
