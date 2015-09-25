package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/yosida95/golang-jenkins"
	"code.google.com/p/gcfg"
)

type Config struct {
	Main struct {
		Port string
	}
	CServer struct {
		Url     string
		JobName string
	}
	CServerAuth struct {
		Username  string
		AuthToken string
	}
	Job struct {
		TargetBranch string
		Label        string
		ParamName    string
	}
}

var cfg Config

func init() {
	err := gcfg.ReadFileInto(&cfg, "./webhooks.cfg")
	if err != nil {
		panic(err)
	}
}

func main() {
	serverCI := http.NewServeMux()
	serverCI.HandleFunc("/", postReceive)

	http.ListenAndServe(":"+cfg.Main.Port, serverCI)
}

func postReceive(w http.ResponseWriter, r *http.Request) {
	x, _ := ioutil.ReadAll(r.Body)

	var webHook GitHubWebHook
	err := json.Unmarshal(x, &webHook)
	if err != nil {
		fmt.Printf("error %v\n", err)
	}

	var action = webHook.Action
	var pullRequestRef = webHook.PullRequest.Head.Ref
	var target = webHook.PullRequest.Base.Repo.DefaultBranch
	var label = webHook.Label.Name

	if strings.EqualFold(label, cfg.Job.Label) && strings.EqualFold(target, cfg.Job.TargetBranch) && strings.EqualFold(action, "labeled") {
		fmt.Printf("Pull request action %v\n", action)
		fmt.Printf("Pull request ref %v\n", pullRequestRef)
		fmt.Printf("Pull request label %v\n", label)
		fmt.Printf("Pull request target branch %v\n", target)
		fmt.Println("Starting jenkins task")
		runJenkinsJob(pullRequestRef)
	} else {
		fmt.Printf("Pull request ref action %v\n", action)
		fmt.Printf("Pull request ref %v\n", pullRequestRef)
		fmt.Printf("Pull request label %v\n", label)
		fmt.Println("No action should be done")
	}

}

func runJenkinsJob(ref string) {
	auth := gojenkins.Auth{cfg.CServerAuth.Username, cfg.CServerAuth.AuthToken}
	jenkinsCI := gojenkins.NewJenkins(&auth, cfg.CServer.Url)
	job := gojenkins.Job{Name: cfg.CServer.JobName}
	params := url.Values{cfg.Job.ParamName: {ref}}
	err := jenkinsCI.Build(job, params)
	if err != nil {
		fmt.Printf("error %v\n", err)
	}
}
