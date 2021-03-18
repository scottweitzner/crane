package internal

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/scottweitzner/crane/common"
)

func CloneAndSwitchVersion(url string, version string) error {

	r, err := git.PlainClone(common.GitClonePath, false, &git.CloneOptions{
		URL: url,
	})
	if err != nil {
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		return err
	}

	if err = w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName(version),
	}); err != nil {
		return err
	}

	return nil
}
