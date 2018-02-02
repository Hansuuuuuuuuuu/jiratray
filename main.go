package main
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/getlantern/systray"
	flag "github.com/ogier/pflag"
)

var (
	host string
	port int
	user string
	openissues = 0
)

func main() {
	flag.Parse()
	systray.Run(onReady, onExit)
	fmt.Println("user has value ", user)
}

func init() {
	flag.StringVarP(&user, "user", "u", "", "Jira Username")
	flag.StringVarP(&host, "host", "h", "", "Jira Host Address (without http://)")
	flag.IntVarP(&port, "port", "p", 443, "Jira Host Port")
}

func onReady() {
	systray.SetIcon(getIcon("assets/jira.png"))
	s := fmt.Sprintf("[%d] Jira Open Issues Tracker", openissues)
	systray.SetTitle(s)
	systray.SetTooltip("ur mom")
	getOpenIssuesCount()
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

func getOpenIssuesCount() {
	fmt.Println("sending req")
	url := fmt.Sprintf("https://%s:%d/rest/api/2/search?=resolution%%20=Unresolved%%20AND%%20assignee%%20=%%20%s", host, port, user)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	fmt.Println(resp.Body);
	defer resp.Body.Close()
}

func updateCount() {

}
// https://jira.corp.paymaya.com/rest/api/2/search?jql=resolution%20=%20Unresolved%20AND%20assignee%20=%20currentUser()