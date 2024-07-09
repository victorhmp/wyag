package repository

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
)

type Repository struct {
	WorkTree string
	GitDir   string
	Config   *ini.File
}

// TODO: Implement actual error handling
func NewRepository(path string, force bool) *Repository {
	repo := Repository{
		WorkTree: path,
		GitDir:   filepath.Join(path, ".git"),
	}

	_, err := os.ReadDir(repo.GitDir)

	if !force || errors.Is(err, fs.ErrNotExist) {
		// This is very bad error handling... but I'll roll with it for now
		panic("Not a git repository")
	}

	configPath, configErr := repo.RepoFilePath(false, "config")

	if configErr != nil && !force {
		panic("Configuration file missing")
	}

	parsedConfig, parsedConfigErr := ini.Load(configPath)

	if parsedConfigErr != nil {
		panic("Failed to parse configuration file")
	}

	repo.Config = parsedConfig

	if !force {
		version, err := repo.Config.Section("core").GetKey("repositoryformatversion")

		if err != nil {
			panic("Invalid configuration file. Missing repositoryformatversion")
		}

		if version.Value() != "0" {
			panic("Unsupported repositoryformatversion")
		}
	}

	return &repo
}

// Compute path under repo's gitdir
func (r *Repository) RepoPath(path ...string) string {
	args := append([]string{r.GitDir}, path...)
	return filepath.Join(args...)
}

// Same as RepoPath, but create dirname(path...) if absent.  For
// example, repo_file(r, "refs", "remotes", "origin", "HEAD") will create
// .git/refs/remotes/origin.
func (r *Repository) RepoFilePath(mkdir bool, path ...string) (string, error) {
	_, err := r.RepoDir(mkdir, path[:len(path)-1]...)

	if err != nil {
		return "", err
	}

	return r.RepoPath(path...), nil
}

// Same as RepoPath, but actually creates a file at path if absent and mkdir is true.
func (r *Repository) RepoDir(mkdir bool, path ...string) (string, error) {
	repoPath := r.RepoPath(path...)

	dir, err := os.ReadDir(repoPath)

	if dir != nil {
		return repoPath, nil
	}

	if errors.Is(err, fs.ErrNotExist) {
		if mkdir {
			err := os.MkdirAll(repoPath, 0755)
			if err != nil {
				return "", err
			}

			return repoPath, nil
		}

		panic("Not a directory")
	}

	return r.RepoPath("config"), nil
}
