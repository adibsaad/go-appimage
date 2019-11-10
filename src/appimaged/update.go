// Instead of putting the updater into the desktop files,
// we call ourselves. This is so that the updater can be updated
// without us having to update all desktop files.
// TODO: This could be made generic
// into a way of launching the most recent version of an application
// we know of.

package main

import (
	"fmt"
	"os"

	"github.com/probonopd/appimage/internal/helpers"
)

func update() {
	if len(os.Args) < 2 {
		fmt.Println("Argument missing")
		os.Exit(1)
	}

	// I think this way of doing things is really clever because
	// this way we can even put the update action into menus if
	// no updater is on the system yet at that time, and give the user
	// instructions what to do once they want to use the upate action
	// in the menu. Also, this way the updater does not have to be at
	// a static location on the $PATH but can be put into any location
	// from which it gets integrated.

	// For now we don't implement updating functionality ourselves
	// but merely launch an updater we found among the integrated
	// AppImages. In the future we may do the updating ourselves.

	// aiu := "gh-releases-zsync|antony-jr|AppImageUpdater|continuous|AppImageUpdater*-x86_64.AppImage.zsync"
	aiu := "gh-releases-zsync|AppImage|AppImageUpdate|continuous|AppImageUpdate-*x86_64.AppImage.zsync"

	a := FindMostRecentAppImageWithMatchingUpdateInformation(aiu)
	if a == "" {
		SimpleNotify("AppImageUpdate missing", "Please download the AppImageUpdate\nAppImage and try again", 30000)
		// Tried making a hyperlink but when I click it in Xfce, nothing happens.
	} else {
		cmd := []string{a}
		cmd = append(cmd, os.Args[2:]...)

		helpers.RunCmdTransparently(cmd)
	}
}