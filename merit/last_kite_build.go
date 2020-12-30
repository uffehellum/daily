package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const bearerToken="cd4f8cdb69fdac86ab4cdb7efedb564a0cbc26c8"
// const kiteOrganization = "sigma"

type BuildKiteBuild struct {
	Id string //":"86d3b35e-d0f8-4093-9bf4-d23d41ffda37",
	// "graphql_id":"QnVpbGQtLS04NmQzYjM1ZS1kMGY4LTQwOTMtOWJmNC1kMjNkNDFmZmRhMzc=",
	//"url":"https://api.buildkite.com/v2/organizations/sigma/pipelines/pantsbuild-basics/builds/24",
	//"web_url":"https://buildkite.com/sigma/pantsbuild-basics/builds/24",
	//"number":24,
	State string // ":"canceled",
	//"blocked":false,"blocked_state":"",
	// "message":"Update README.md",
	Commit string // ":"11b21df853ff2b9ea2e231d4e6ffeb9332bd681e","
	Branch string // ":"master",
	// "tag":null,
	//"env":{},
	//"source":"webhook",
	//"author":{"username":"LiuVII","name":"Mikhail Filipchuk","email":"michael@sig.ma"},
	//"creator":{"id":"074ca0b3-7627-4299-9629-9b7940c275cc",
	//"graphql_id":"VXNlci0tLTA3NGNhMGIzLTc2MjctNDI5OS05NjI5LTliNzk0MGMyNzVjYw==",
	//"name":"Mikhail Filipchuk",
	//"email":"michael@gomerits.com",
	//"avatar_url":"https://www.gravatar.com/avatar/6a50642ac6f8cf822bc5d3f8596bbb04",
	//"created_at":"2017-10-27T17:55:16.767Z"},
	//"created_at":"2019-05-15T23:27:38.443Z",
	//"scheduled_at":"2019-05-15T23:27:38.404Z",
	//"started_at":null,
	FinishedAt string  "json:finished_at" //  ":"2019-05-15T23:28:13.000Z",
	//"meta_data":{},"pull_request":null,"rebuilt_from":null,
	//"pipeline":{"id":"8d0d4684-a32e-4c4a-a3ab-b7657e00686a","graphql_id":"UGlwZWxpbmUtLS04ZDBkNDY4NC1hMzJlLTRjNGEtYTNhYi1iNzY1N2UwMDY4NmE=","url":"https://api.buildkite.com/v2/organizations/sigma/pipelines/pantsbuild-basics","web_url":"https://buildkite.com/sigma/pantsbuild-basics",
	//"name":"pantsbuild-basics","description":"Pipeline for basics pantsbuild branches","slug":"pantsbuild-basics","repository":"git@github.com:LiuVII/pantsbuild_test","cluster_id":null,"branch_configuration":null,"default_branch":"master","skip_queued_branch_builds":true,"skip_queued_branch_builds_filter":"","cancel_running_branch_builds":true,"cancel_running_branch_builds_filter":"","allow_rebuilds":true,"provider":{"id":"github","settings":{"trigger_mode":"code","build_pull_requests":true,"pull_request_branch_filter_enabled":false,"skip_pull_request_builds_for_existing_commits":true,"build_pull_request_forks":false,"prefix_pull_request_fork_branch_names":true,"build_tags":false,"publish_commit_status":true,"publish_commit_status_per_step":true,"separate_pull_request_statuses":true,"publish_blocked_as_pending":false,"filter_enabled":false,"repository":"LiuVII/pantsbuild_test","pull_request_branch_filter_configuration":""},"webhook_url":"https://webhook.buildkite.com/deliver/a236f17639c7c17969b38999b9c619c70ae9b6b9ad617a5acf"},
}
func FetchSigmaBuilds() string {
	url := "https://api.buildkite.com/v2/organizations/sigma/pipelines/pantsbuild-basics/builds"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Bearer " + bearerToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	resp.Header.Add("Bearer", bearerToken)
	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	a := []BuildKiteBuild{}
	err = json.Unmarshal(body, &a)
	if err != nil {
		panic(err)
	}
	for _, build := range a {
		if build.Branch == "master" && build.State =="passed" {
			return build.Commit
		}
	}
	fmt.Println(a)
	panic("no good builds")

}

func main() {
	fmt.Println(FetchSigmaBuilds())
}