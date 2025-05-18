package util

import (
	"os"
	"path/filepath"
)

func ExtractAssets(outputDir string) error {
	for _, name := range AssetNames() {
		data, err := Asset(name)
		if err != nil {
			return err
		}

		destPath := filepath.Join(outputDir, name)
		err = os.MkdirAll(filepath.Dir(destPath), os.ModePerm)
		if err != nil {
			return err
		}

		err = os.WriteFile(destPath, data, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
