package symembed

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

type SymLink *string

type SymManifest struct {
	Links map[string]string
}

type SymWalkFunc func(path string, symlink SymLink, info fs.FileInfo, err error) error

func GenManifest(root string) (SymManifest, error) {

	manifest := SymManifest{Links: make(map[string]string)}

	SymWalk(root, func(path string, symlink SymLink, info fs.FileInfo, err error) error {

		if symlink != nil {
			relSymlink, relSymlinkErr := filepath.Rel(root, *symlink)
			relPath, relPathErr := filepath.Rel(root, path)

			if relSymlinkErr != nil {
				log.Println("Failed to get relative path for symlink:", *symlink, "Error:", relSymlinkErr)
				return relSymlinkErr
			}

			if relPathErr != nil {
				log.Println("Failed to get relative path for path:", path, "Error:", relPathErr)
				return relPathErr
			}

			log.Println("link: ", relPath, " -> ", relSymlink)
			manifest.Links[relPath] = relSymlink
		}

		if err != nil {
			return err
		}

		return nil
	})

	return manifest, nil
}

func (manifest *SymManifest) ApplyManifest(root string) error {

	for path, symlink := range manifest.Links {
		linkPath := filepath.Join(root, path)
		linkTarget := filepath.Join(root, symlink)

		linkDir := filepath.Dir(linkPath)

		relTarget, relTargetErr := filepath.Rel(linkDir, linkTarget)

		if relTargetErr != nil {
			log.Println("Failed to get relative path for link target:", linkTarget, "Error:", relTargetErr)
		}

		log.Println("link: ", linkPath, " -> ", linkTarget)

		if err := os.Symlink(relTarget, linkPath); err != nil {
			log.Println("Failed to create symlink:", linkPath, "Error:", err)
		}

	}

	return nil
}

// Walk is similar to filepath.Walk (https://golang.org/pkg/path/filepath/#Walk) except it follows
// the symbolic links it finds. The walk function keeps a list of all visited directories to avoid
// endless loop resulted from symbolic loops
func SymWalk(root string, fn SymWalkFunc) error {
	rr, err := filepath.EvalSymlinks(root) // Find real base if there is any symlinks in the path
	if err != nil {
		return err
	}

	visitedDirs := make(map[string]struct{})
	return filepath.Walk(rr, getWalkFn(visitedDirs, fn))
}

func getWalkFn(visitedDirs map[string]struct{}, fn SymWalkFunc) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fn(path, nil, info, err)
		}

		if info.IsDir() {
			if _, ok := visitedDirs[path]; ok {
				return filepath.SkipDir
			}
			visitedDirs[path] = struct{}{}
		}

		if info.Mode()&os.ModeSymlink == 0 {
			if fnerr := fn(path, nil, info, err); fnerr != nil {
				return fnerr
			}

		} else {
			// path is a symlink
			rp, err := filepath.EvalSymlinks(path)
			if err != nil {
				return err
			}

			ri, err := os.Stat(rp)
			if err != nil {
				return err
			}

			if fnerr := fn(path, &rp, info, err); fnerr != nil {
				return fnerr
			}

			if ri.IsDir() {
				return filepath.Walk(rp, getWalkFn(visitedDirs, fn))
			}
		}

		return nil
	}
}
