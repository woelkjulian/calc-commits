package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/woelkjulian/commitmerge-tool/model"
)

func main() {

	urlPtr := flag.String("url", "", "gitlab url with the projects")
	tokenPtr := flag.String("t", "", "private token of gitlab user")
	branchPtr := flag.String("b", "", "name of git branch")
	projPtr := flag.Int("proj", -1, "id of the gitlab project")
	logPtr := flag.Bool("v", false, "additional log info")

	flag.Parse()

	if *urlPtr == "" || *tokenPtr == "" || *projPtr == -1 {
		msg := "Arguments missing: \n"

		if *urlPtr == "" {
			msg += "- URL (-url)\n"
		}
		if *tokenPtr == "" {
			msg += "- token (-t)\n"
		}
		if *projPtr == -1 {
			msg += "- project number (-proj)\n"
		}

		fmt.Printf("\n\n%v\n\n", msg)
		return
	}

	if err := calcMergeCommitsQuotient(urlPtr, tokenPtr, branchPtr, projPtr, logPtr); err != nil {
		fmt.Printf("Error while calculating: %v\n", err)
	}
}

func calcMergeCommitsQuotient(pURL, pToken, pBranch *string, pProj *int, pLog *bool) error {

	var noOfMergeCommits int

	commits, err := getCommits(pURL, pToken, pBranch, pProj, pLog)

	if err != nil {
		return err
	}

	mergeRequests, err := getMergeRequests(pURL, pToken, pBranch, pProj, pLog)

	if err != nil {
		return err
	}

	for _, req := range mergeRequests {
		mergeRequestCommits, err := getMergeRequestCommits(pURL, pToken, pBranch, pProj, req.Id, pLog)

		if err != nil {
			return err
		}

		noOfMergeCommits += len(mergeRequestCommits)
	}

	noOfAllCommits := len(commits) + noOfMergeCommits

	fmt.Printf("\n\nPercentage of Merge Request Commits: %.2f%v\n\n",
		float64(noOfMergeCommits)/float64(noOfAllCommits)*float64(100),
		"%")

	return nil
}

func getCommits(pURL, pToken, pBranch *string, pProj *int, pLog *bool) ([]model.Commit, error) {

	var url = fmt.Sprint(*pURL, "/api/v3/projects/", *pProj, "/repository/commits")

	if *pBranch != "" {
		url += fmt.Sprint("/:", *pBranch)
	}

	if *pLog == true {
		fmt.Printf("\ngetCommits(): from %v %v", url, "...")
	}

	tr := &http.Transport{DisableKeepAlives: false}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("PRIVATE-TOKEN", *pToken)
	req.Close = false

	res, err := tr.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	body, _ := ioutil.ReadAll(res.Body)

	commits := make([]model.Commit, 0)

	if err := json.Unmarshal(body, &commits); err != nil {
		return nil, err
	}

	if *pLog == true {
		fmt.Printf("\ngetCommits(): found %v commits\n", len(commits))
	}

	return commits, nil
}

func getMergeRequests(pURL, pToken, pBranch *string, pProj *int, pLog *bool) ([]model.MergeRequest, error) {

	var url = fmt.Sprint(*pURL, "/api/v3/projects/", *pProj, "/merge_requests")

	if *pLog == true {
		fmt.Printf("\ngetMergeRequests(): from %v %v", url, "...")
	}

	tr := &http.Transport{DisableKeepAlives: false}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("PRIVATE-TOKEN", *pToken)
	req.Close = false

	res, err := tr.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	body, _ := ioutil.ReadAll(res.Body)

	mergeRequests := make([]model.MergeRequest, 0)

	if err := json.Unmarshal(body, &mergeRequests); err != nil {
		return nil, err
	}
	if *pLog == true {
		fmt.Printf("\ngetMergeRequests(): found %v merge requests\n", len(mergeRequests))
	}

	return mergeRequests, nil
}

func getMergeRequestCommits(pURL, pToken, pBranch *string, pProj *int, mergeReqID int, pLog *bool) ([]model.Commit, error) {

	var url = fmt.Sprint(*pURL, "/api/v3/projects/", *pProj, "/merge_requests/", mergeReqID, "/commits")

	if *pBranch != "" {
		url += fmt.Sprint("/:", *pBranch)
	}

	if *pLog == true {
		fmt.Printf("\ngetMergeRequestCommits(): from %v %v", url, "...")
	}

	tr := &http.Transport{DisableKeepAlives: false}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("PRIVATE-TOKEN", *pToken)
	req.Close = false

	res, err := tr.RoundTrip(req)

	if err != nil {
		return nil, err
	}

	body, _ := ioutil.ReadAll(res.Body)

	mergeRequestCommits := make([]model.Commit, 0)

	if err := json.Unmarshal(body, &mergeRequestCommits); err != nil {
		return nil, err
	}

	if *pLog == true {
		fmt.Printf("\ngetMergeRequestCommits(): found %v commits from merge requestId %v", len(mergeRequestCommits), mergeReqID)
	}

	return mergeRequestCommits, nil
}
