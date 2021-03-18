package internal

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/scottweitzner/crane/common"
)

// CloneVersion clones a specific version of a repository
func CloneVersion(url string, version string) error {

	if _, err := git.PlainClone(common.GitClonePath, false, &git.CloneOptions{
		URL:           url,
		ReferenceName: plumbing.ReferenceName(version),
	}); err != nil {
		return err
	}

	return nil
}
