package main

import (
	"fmt"
	"io/ioutil"
	"github.com/getlantern/systray"
)


func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(getIcon("assets/jira.png"))
	systray.SetTitle("Jira Open Issues Tracker")
	systray.SetTooltip("ur mom")
}

func onExit() {

}

func getIcon(s string) [] byte {
	b, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Print(err)
	}
	return b
}
// https://jira.corp.paymaya.com/rest/api/2/search?jql=resolution%20=%20Unresolved%20AND%20assignee%20=%20currentUser()