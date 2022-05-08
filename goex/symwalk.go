package goex

import (
	"io/fs"
	"os"
	"path/filepath"
)

func Walk(fsys fs.FS, root string, fn fs.WalkDirFunc) error {
	rr, err := filepath.EvalSymlinks(root) // Find real base if there is any symlinks in the path
	if err != nil {
		return err
	}

	visitedDirs := make(map[string]struct{})
	return fs.WalkDir(fsys, rr, getWalkFn(fsys, visitedDirs, fn))
}

func getWalkFn(fsys fs.FS, visitedDirs map[string]struct{}, fn fs.WalkDirFunc) fs.WalkDirFunc {
	return func(path string, dirinfo os.DirEntry, err error) error {
		if err != nil {
			return fn(path, dirinfo, err)
		}

		if dirinfo.IsDir() {
			if _, ok := visitedDirs[path]; ok {
				return filepath.SkipDir
			}
			visitedDirs[path] = struct{}{}
		}

		if err := fn(path, dirinfo, err); err != nil {
			return err
		}

		fileInfo, fileInfoErr := dirinfo.Info()

		if fileInfoErr != nil {
			return fileInfoErr
		}

		if fileInfo.Mode()&os.ModeSymlink == 0 {
			return nil
		}

		// path is a symlink
		rp, err := filepath.EvalSymlinks(path)
		if err != nil {
			return err
		}

		ri, err := os.Stat(rp)
		if err != nil {
			return err
		}

		if ri.IsDir() {
			return fs.WalkDir(fsys, rp, getWalkFn(fsys, visitedDirs, fn))
		}

		return nil
	}
}
