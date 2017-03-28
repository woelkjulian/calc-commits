package model

type MergeRequest struct {
	Id                       int       `json:"id"`
	Iid                      int       `json:"iid"`
	ProjectId                int       `json:"project_id"`
	Title                    string    `json:"title"`
	Description              string    `json:"description"`
	State                    string    `json:"state"`
	CreatedAt                string    `json:"created_at"`
	UpdatedAt                string    `json:"updated_at"`
	TargetBranch             string    `json:"target_branch"`
	SourceBranch             string    `json:"source_branch"`
	Upvotes                  int       `json:"upvotes"`
	Downvotes                int       `json:"downvotes"`
	Author                   Author    `json:"author"`
	Assignee                 Author    `json:"assignee"`
	SourceProjectId          int       `json:"source_project_id"`
	TargetProjectId          int       `json:"target_project_id"`
	Labels                   []string  `json:"labels"`
	WorkInProgress           bool      `json:"work_in_progress"`
	Milestone                Milestone `json:"milestone"`
	MergeWhenBuildSucceeds   bool      `json:"merge_when_build_succeeds"`
	MergeStatus              string    `json:"merge_status"`
	Sha                      string    `json:"sha"`
	MergeCommitSha           string    `json:"merge_commit_sha"`
	Subscribed               bool      `json:"subscribed"`
	UserNotesCount           int       `json:"user_notes_count"`
	ShouldRemoveSourceBranch bool      `json:"should_remove_source_branch"`
	ForceRemoveSourceBranch  bool      `json:"force_remove_source_branch"`
	WebUrl                   string    `json:"web_url"`
}
