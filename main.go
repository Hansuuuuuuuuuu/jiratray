package main
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
	"time"
	"encoding/json"
	"github.com/getlantern/systray"
	flag "github.com/ogier/pflag"
)

type Jira struct {
	Expand     string `json:"expand"`
	StartAt    int    `json:"startAt"`
	MaxResults int    `json:"maxResults"`
	Total      int    `json:"total"`
	Issues     []struct {
		Expand string `json:"expand"`
		ID     string `json:"id"`
		Self   string `json:"self"`
		Key    string `json:"key"`
	} `json:"issues"`
}

var (
	host string
	port int
	user string
	password string
	openissues = 0
	interval int
)

func main() {
	flag.Parse()
	systray.Run(onReady, onExit)
}

func init() {
	flag.StringVarP(&user, "user", "u", "", "Jira Username")
	flag.StringVarP(&password, "password", "p", "", "Jira Password")
	flag.StringVarP(&host, "host", "h", "", "Jira Host Address (without http://)")
	flag.IntVarP(&port, "port", "x", 443, "Jira Host Port")
	flag.IntVarP(&interval, "interval", "i", 30, "Update Interval (in seconds)")
}

func onReady() {
	systray.SetIcon(getIcon("assets/jira.png"))
	s := fmt.Sprintf("[%d] Jira Open Issues Tracker", openissues)
	systray.SetTitle(s)
	tt := fmt.Sprintf("Logged in as %s", user)
	systray.SetTooltip(tt)
	for {
	    updateCount()
	    time.Sleep(time.Duration(interval) * time.Second)
	}
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

func getOpenIssuesCount() int {
	client := &http.Client{}
	url := fmt.Sprintf("https://%s:%d/rest/api/2/search?jql=resolution%%20=Unresolved%%20AND%%20assignee%%20=%%20%s", host, port, user)

	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(user, password)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	    log.Fatal(err)
	}

	var jiraJson Jira
	json.Unmarshal([]byte(respData), &jiraJson)
	//respString := string(respData)
	return jiraJson.Total
}

func updateCount() {
	tot := getOpenIssuesCount()
	s := fmt.Sprintf("[%d] Jira Open Issues Tracker", tot)
	systray.SetTitle(s)
}