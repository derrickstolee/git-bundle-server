package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/github/git-bundle-server/internal/bundles"
	"github.com/github/git-bundle-server/internal/core"
)

type Update struct{}

func (Update) subcommand() string {
	return "update"
}

func (Update) run(args []string) error {
	if len(args) != 1 {
		return errors.New("usage: git-bundle-server update (--all|<route>)")
	}

	if args[0] == "--all" {
		return updateAll()
	}

	route := args[0]

	return updateRoute(route)
}

func updateAll() error {
	repos, err := core.GetRepositories()
	if err != nil {
		return err
	}

	for route := range repos {
		fmt.Printf("Updating route %s\n", route)
		routeErr := updateRoute(route)
		if routeErr != nil {
			fmt.Fprintf(os.Stderr, "Failed to update %s: %s", route, routeErr.Error())
			if err == nil {
				err = routeErr
			}
		}
	}

	return err
}

func updateRoute(route string) error {
	repo, err := core.CreateRepository(route)
	if err != nil {
		return err
	}

	list, err := bundles.GetBundleList(repo)
	if err != nil {
		return fmt.Errorf("failed to load bundle list: %w", err)
	}

	fmt.Printf("Creating new incremental bundle\n")
	bundle, err := bundles.CreateIncrementalBundle(repo, list)
	if err != nil {
		return err
	}

	// Nothing new!
	if bundle == nil {
		return nil
	}

	list.Bundles[bundle.CreationToken] = *bundle

	fmt.Printf("Collapsing bundle list\n")
	err = bundles.CollapseList(repo, list)
	if err != nil {
		return err
	}

	fmt.Printf("Writing updated bundle list\n")
	listErr := bundles.WriteBundleList(list, repo)
	if listErr != nil {
		return fmt.Errorf("failed to write bundle list: %w", listErr)
	}

	return nil
}
