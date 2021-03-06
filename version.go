// This module ensures that you embed the current git revision into the
// application binary. You can do so with the following command:
//
//     go build -ldflags "-X main.gitRevisionId $(git rev-parse HEAD)"

package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var gitRevisionId string

var printVersion = flag.Bool("version", false, "Print the git revision id and exit")

func printVersionIfAsked() {
	if !*printVersion {
		return
	}

	fmt.Println(gitRevisionId)

	if gitRevisionId == "" {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func versionServer(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, fmt.Sprintf("https://github.com/sburnett/encore/commit/%[1]s", gitRevisionId), http.StatusFound)
	fmt.Fprint(w, gitRevisionId)
}
