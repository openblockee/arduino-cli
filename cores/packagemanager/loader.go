/*
 * This file is part of arduino-cli.
 *
 * arduino-cli is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA  02110-1301  USA
 *
 * As a special exception, you may use this file as part of a free software
 * library without restriction.  Specifically, if other files instantiate
 * templates or use macros or inline functions from this file, or you compile
 * this file and link it with other files to produce an executable, this
 * file does not by itself cause the resulting executable to be covered by
 * the GNU General Public License.  This exception does not however
 * invalidate any other reasons why the executable file might be covered by
 * the GNU General Public License.
 *
 * Copyright 2017-2018 ARDUINO AG (http://www.arduino.cc/)
 */

package packagemanager

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	properties "github.com/arduino/go-properties-map"
	"github.com/bcmi-labs/arduino-cli/configs"
	"github.com/bcmi-labs/arduino-cli/cores"
)

// LoadHardwareFromDirectories read all plaforms from the configured paths
func (pm *packageManager) LoadHardware() error {
	dirs, err := configs.HardwareDirectories()
	if err != nil {
		return fmt.Errorf("getting hardware folder: %s", err)
	}
	return pm.LoadHardwareFromDirectories(dirs)
}

// LoadHardwareFromDirectories read all plaforms from a set of directories
func (pm *packageManager) LoadHardwareFromDirectories(hardwarePaths []string) error {
	for _, path := range hardwarePaths {
		path, err := filepath.Abs(path)
		if err != nil {
			return fmt.Errorf("find abs path: %s", err)
		}
		// TODO: IS THIS CHECK NEEDED? can we ignore and let it fail at next ReadDir?
		if stat, err := os.Stat(path); err != nil {
			return fmt.Errorf("reading %s stat info: %s", path, err)
		} else if !stat.IsDir() {
			return fmt.Errorf("%s is not a folder", path)
		}

		// TODO: IS THIS REALLY NEEDED? this is used only to get ctags properties AFAIK
		platformTxtPath := filepath.Join(path, "platform.txt")
		if props, err := properties.SafeLoad(platformTxtPath); err == nil {
			pm.packages.Properties.Merge(props)
		} else {
			return fmt.Errorf("reading %s: %s", platformTxtPath, err)
		}

		// Scan subfolders.
		files, err := ioutil.ReadDir(path)
		if err != nil {
			return fmt.Errorf("reading %s directory: %s", path, err)
		}
		for _, packagerPathInfo := range files {
			// First exclude all "tools" folders
			packager := packagerPathInfo.Name()
			if packager == "tools" {
				continue
			}

			// Follow symlinks
			packagerPath := filepath.Join(path, packager)
			packagerPath, err := filepath.EvalSymlinks(packagerPath)
			if err != nil {
				return fmt.Errorf("following possible symlink %s: %s", path, err)
			}

			// There are two possible package folder structures:
			// - PACKAGER/ARCHITECTURE/...
			// - PACKAGER/hardware/ARCHITECTURE/VERSION/...
			// if we found the latter we just move into "hardware" folder and continue
			hardwareSubdirPath := filepath.Join(packagerPath, "hardware")
			if info, err := os.Stat(hardwareSubdirPath); err == nil && info.IsDir() {
				// move down into the "hardware" folder
				packagerPath = hardwareSubdirPath
			} else if info, err := os.Stat(packagerPath); err == nil && info.IsDir() {
				// we are already at the correct level
			} else {
				// error: do nothing
				continue
			}

			targetPackage := pm.packages.GetOrCreatePackage(packager)
			if err := pm.LoadPlatforms(targetPackage, packagerPath); err != nil {
				return fmt.Errorf("loading package %s: %s", packager, err)
			}
		}
	}

	return nil
}

// LoadPlatforms load plaftorms from the specified directory assuming that they belongs
// to the targetPackage object passed as parameter.
func (pm *packageManager) LoadPlatforms(targetPackage *cores.Package, packageFolder string) error {
	// packagePlatformTxt, err := properties.SafeLoad(filepath.Join(folder, constants.FILE_PLATFORM_TXT))
	// if err != nil {
	// 	return err
	// }
	// targetPackage.Properties.Merge(packagePlatformTxt)

	files, err := ioutil.ReadDir(packageFolder)
	if err != nil {
		return fmt.Errorf("reading directory %s: %s", packageFolder, err)
	}

	for _, file := range files {
		architecure := file.Name()
		platformPath := filepath.Join(packageFolder, architecure)
		if architecure == "tools" ||
			architecure == "platform.txt" { // TODO: Check if this "platform.txt" condition should be here....
			continue
		}

		// There are two possible platform folder structures:
		// - ARCHITECTURE/boards.txt
		// - ARCHITECTURE/VERSION/boards.txt
		// We identify them by checking where is the bords.txt file
		possibleBoardTxtPath := filepath.Join(platformPath, "boards.txt")
		if _, err := os.Stat(possibleBoardTxtPath); err == nil {
			// case: ARCHITECTURE/boards.txt
			// this is an unversioned Platform

			platform := targetPackage.GetOrCreatePlatform(architecure)
			release := platform.GetOrCreateRelease("")
			if err := pm.LoadPlatformRelease(release, platformPath); err != nil {
				return fmt.Errorf("loading platform release: %s", err)
			}

		} else if os.IsNotExist(err) {
			// case: ARCHITECTURE/VERSION/boards.txt
			// let's dive into VERSION folders

			platform := targetPackage.GetOrCreatePlatform(architecure)
			versionDirs, err := ioutil.ReadDir(platformPath)
			if err != nil {
				return fmt.Errorf("reading dir %s: %s", platformPath, err)
			}
			for _, versionDir := range versionDirs {
				if !versionDir.IsDir() {
					continue
				}
				version := versionDir.Name()
				release := platform.GetOrCreateRelease(version)
				platformWithVersionPath := filepath.Join(platformPath, version)

				if err := pm.LoadPlatformRelease(release, platformWithVersionPath); err != nil {
					return fmt.Errorf("loading platform release %s: %s", version, err)
				}
			}
		} else {
			return fmt.Errorf("looking for boards.txt in %s: %s", possibleBoardTxtPath, err)
		}
	}

	return nil
}

func (pm *packageManager) LoadPlatformRelease(platform *cores.PlatformRelease, folder string) error {
	if _, err := os.Stat(filepath.Join(folder, "boards.txt")); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("opening boards.txt: %s", err)
	} else if os.IsNotExist(err) {
		return fmt.Errorf("invalid platform directory %s: boards.txt not found", folder)
	}
	platform.Folder = folder

	// Some useful paths
	platformTxtPath := filepath.Join(folder, "platform.txt")
	platformTxtLocalPath := filepath.Join(folder, "platform.local.txt")
	programmersTxtPath := filepath.Join(folder, "programmers.txt")

	// Create platform properties
	platform.Properties = platform.Properties.Clone() // TODO: why CLONE?
	if p, err := properties.SafeLoad(platformTxtPath); err == nil {
		platform.Properties.Merge(p)
	} else {
		return fmt.Errorf("loading %s: %s", platformTxtPath, err)
	}
	if p, err := properties.SafeLoad(platformTxtLocalPath); err == nil {
		platform.Properties.Merge(p)
	} else {
		return fmt.Errorf("loading %s: %s", platformTxtLocalPath, err)
	}

	// Create programmers properties
	if programmersProperties, err := properties.SafeLoad(programmersTxtPath); err == nil {
		platform.Programmers = properties.MergeMapsOfProperties(
			map[string]properties.Map{},
			platform.Programmers, // TODO: Very weird, why not an empty one?
			programmersProperties.FirstLevelOf())
	} else {
		return err
	}

	if err := platform.LoadBoards(); err != nil {
		return err
	}

	return nil
}
