package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/woelkjulian/calc-commits/model"
)

// Configurable as flag??
var perPage = 100

func main() {
	urlPtr := flag.String("url", "", "gitlab url with the projects (e.g. https://gitlab.example.com)")
	projNamePtr := flag.String("projName", "", "namespace/projectname of the gitlab project")
	projIDPtr := flag.String("projId", "", "instead of namespace/projectname, id of the gitlab project (e.g. 999)")
	tokenPtr := flag.String("t", "", "private token of gitlab user (e.g. abcdefg123456)")
	branchPtr := flag.String("b", "", "optional name of git branch (e.g. master)")
	logPtr := flag.Bool("v", false, "optional additional log info")

	flag.Parse()

	if *urlPtr == "" || *tokenPtr == "" || (*projIDPtr == "" && *projNamePtr == "") {
		msg := "Arguments missing: \n"

		if *urlPtr == "" {
			msg += "- URL (-url https://gitlab.example.com)\n"
		}
		if *tokenPtr == "" {
			msg += "- token (-t abcdef123456)\n"
		}
		if *projIDPtr == "" && *projNamePtr == "" {
			msg += "- project id (-projId 999) or project name (-projName namespace/projectname)\n"
		}

		fmt.Printf("\n\n%v\n\n", msg)
		return
	}

	if err := calcMergeCommitsQuotient(urlPtr, tokenPtr, branchPtr, projNamePtr, projIDPtr, logPtr); err != nil {
		fmt.Printf("Error while calculating: %v\n", err)
	}
}

func calcMergeCommitsQuotient(pURL, pToken, pBranch, pProjName, pProjID *string, pLog *bool) error {

	var noOfMergeCommits int
	var pProj *string

	if *pProjName != "" {
		pProj = pProjName
		*pProj = strings.Replace(*pProj, "/", "%2F", -1)
	} else if *pProjID != "" {
		pProj = pProjID
	}

	commits, err := getAllCommits(pURL, pToken, pBranch, pProj, 0, pLog)

	if err != nil {
		return err
	}

	fmt.Printf("\n\nTotal number of Commits = %v\n\n", len(commits))

	mergeRequests, err := getMergeRequests(pURL, pToken, pBranch, pProj, 0, pLog)

	if err != nil {
		return err
	}

	fmt.Printf("\n\nTotal number of Merge Requests = %v\n\n", len(mergeRequests))

	for _, req := range mergeRequests {
		mergeRequestCommits, err := getMergeRequestCommits(pURL, pToken, pBranch, pProj, req.Id, 0, pLog)

		if err != nil {
			return err
		}

		noOfMergeCommits += len(mergeRequestCommits)
	}

	fmt.Printf("\n\nTotal number of Merge Request Commits = %v\n\n", noOfMergeCommits)

	fmt.Printf("\n\nPercentage of Merge Request Commits: %.2f%v\n\n",
		float64(noOfMergeCommits)/float64(len(commits))*float64(100),
		"%")

	return nil
}

func getAllCommits(pURL, pToken, pBranch, pProj *string, page int, pLog *bool) ([]model.Commit, error) {

	var url = fmt.Sprint(*pURL, "/api/v3/projects/", *pProj, "/repository/commits")

	if *pBranch != "" {
		branch := "%2F:" + *pBranch
		url += branch
	}

	// per_page max is 100
	url += fmt.Sprint("?per_page=", perPage, "&page=", page)

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

	if len(commits) == perPage {
		if *pLog == true {
			fmt.Printf("\ngetCommits(): found more than %v commits on page %v, searching next page...\n", perPage, page)
		}
		newCommits, err := getAllCommits(pURL, pToken, pBranch, pProj, page+1, pLog)

		if err != nil {
			return nil, err
		}

		for _, commit := range newCommits {
			commits = append(commits, commit)
		}
	}

	return commits, nil
}

func getMergeRequests(pURL, pToken, pBranch, pProj *string, page int, pLog *bool) ([]model.MergeRequest, error) {

	var url = fmt.Sprint(*pURL, "/api/v3/projects/", *pProj, "/merge_requests")

	if *pBranch != "" {
		branch := "%2F:" + *pBranch
		url += branch
	}

	// per_page max is 100
	url += fmt.Sprint("?per_page=", perPage, "&page=", page)

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

	if len(mergeRequests) == perPage {
		if *pLog == true {
			fmt.Printf("\ngetMergeRequests(): found more than %v mergeRequests on page %v, searching next page...\n", perPage, page)
		}
		newMergeRequests, err := getMergeRequests(pURL, pToken, pBranch, pProj, page+1, pLog)

		if err != nil {
			return nil, err
		}

		for _, request := range newMergeRequests {
			mergeRequests = append(mergeRequests, request)
		}
	}

	return mergeRequests, nil
}

func getMergeRequestCommits(pURL, pToken, pBranch, pProj *string, mergeReqID int, page int, pLog *bool) ([]model.Commit, error) {

	var url = fmt.Sprint(*pURL, "/api/v3/projects/", *pProj, "/merge_requests/", mergeReqID, "/commits")

	if *pBranch != "" {
		branch := "%2F:" + *pBranch
		url += branch
	}

	// per_page max is 100
	url += fmt.Sprint("?per_page=", perPage, "&page=", page)

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

	if len(mergeRequestCommits) == perPage {
		if *pLog == true {
			fmt.Printf("\ngetMergeRequestCommits(): found more than %v merge request commits on page %v, searching next page...\n", perPage, page)
		}
		newMergeRequestCommits, err := getMergeRequestCommits(pURL, pToken, pBranch, pProj, mergeReqID, page+1, pLog)

		if err != nil {
			return nil, err
		}

		for _, commit := range newMergeRequestCommits {
			mergeRequestCommits = append(mergeRequestCommits, commit)
		}
	}
	return mergeRequestCommits, nil
}
