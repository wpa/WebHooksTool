// GitHubWebHook
package main

type GitHubWebHook struct {
	Action       string  `json:"action"`
	Number       float64 `json:"number"`
	Organization struct {
		AvatarURL        string      `json:"avatar_url"`
		Description      interface{} `json:"description"`
		EventsURL        string      `json:"events_url"`
		ID               float64     `json:"id"`
		Login            string      `json:"login"`
		MembersURL       string      `json:"members_url"`
		PublicMembersURL string      `json:"public_members_url"`
		ReposURL         string      `json:"repos_url"`
		URL              string      `json:"url"`
	} `json:"organization"`
	PullRequest struct {
		Links struct {
			Comments struct {
				Href string `json:"href"`
			} `json:"comments"`
			Commits struct {
				Href string `json:"href"`
			} `json:"commits"`
			Html struct {
				Href string `json:"href"`
			} `json:"html"`
			Issue struct {
				Href string `json:"href"`
			} `json:"issue"`
			ReviewComment struct {
				Href string `json:"href"`
			} `json:"review_comment"`
			ReviewComments struct {
				Href string `json:"href"`
			} `json:"review_comments"`
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Statuses struct {
				Href string `json:"href"`
			} `json:"statuses"`
		} `json:"_links"`
		Additions float64     `json:"additions"`
		Assignee  interface{} `json:"assignee"`
		Base      struct {
			Label string `json:"label"`
			Ref   string `json:"ref"`
			Repo  struct {
				ArchiveURL       string      `json:"archive_url"`
				AssigneesURL     string      `json:"assignees_url"`
				BlobsURL         string      `json:"blobs_url"`
				BranchesURL      string      `json:"branches_url"`
				CloneURL         string      `json:"clone_url"`
				CollaboratorsURL string      `json:"collaborators_url"`
				CommentsURL      string      `json:"comments_url"`
				CommitsURL       string      `json:"commits_url"`
				CompareURL       string      `json:"compare_url"`
				ContentsURL      string      `json:"contents_url"`
				ContributorsURL  string      `json:"contributors_url"`
				CreatedAt        string      `json:"created_at"`
				DefaultBranch    string      `json:"default_branch"`
				Description      string      `json:"description"`
				DownloadsURL     string      `json:"downloads_url"`
				EventsURL        string      `json:"events_url"`
				Fork             bool        `json:"fork"`
				Forks            float64     `json:"forks"`
				ForksCount       float64     `json:"forks_count"`
				ForksURL         string      `json:"forks_url"`
				FullName         string      `json:"full_name"`
				GitCommitsURL    string      `json:"git_commits_url"`
				GitRefsURL       string      `json:"git_refs_url"`
				GitTagsURL       string      `json:"git_tags_url"`
				GitURL           string      `json:"git_url"`
				HasDownloads     bool        `json:"has_downloads"`
				HasIssues        bool        `json:"has_issues"`
				HasPages         bool        `json:"has_pages"`
				HasWiki          bool        `json:"has_wiki"`
				Homepage         interface{} `json:"homepage"`
				HooksURL         string      `json:"hooks_url"`
				HtmlURL          string      `json:"html_url"`
				ID               float64     `json:"id"`
				IssueCommentURL  string      `json:"issue_comment_url"`
				IssueEventsURL   string      `json:"issue_events_url"`
				IssuesURL        string      `json:"issues_url"`
				KeysURL          string      `json:"keys_url"`
				LabelsURL        string      `json:"labels_url"`
				Language         string      `json:"language"`
				LanguagesURL     string      `json:"languages_url"`
				MergesURL        string      `json:"merges_url"`
				MilestonesURL    string      `json:"milestones_url"`
				MirrorURL        interface{} `json:"mirror_url"`
				Name             string      `json:"name"`
				NotificationsURL string      `json:"notifications_url"`
				OpenIssues       float64     `json:"open_issues"`
				OpenIssuesCount  float64     `json:"open_issues_count"`
				Owner            struct {
					AvatarURL         string  `json:"avatar_url"`
					EventsURL         string  `json:"events_url"`
					FollowersURL      string  `json:"followers_url"`
					FollowingURL      string  `json:"following_url"`
					GistsURL          string  `json:"gists_url"`
					GravatarID        string  `json:"gravatar_id"`
					HtmlURL           string  `json:"html_url"`
					ID                float64 `json:"id"`
					Login             string  `json:"login"`
					OrganizationsURL  string  `json:"organizations_url"`
					ReceivedEventsURL string  `json:"received_events_url"`
					ReposURL          string  `json:"repos_url"`
					SiteAdmin         bool    `json:"site_admin"`
					StarredURL        string  `json:"starred_url"`
					SubscriptionsURL  string  `json:"subscriptions_url"`
					Type              string  `json:"type"`
					URL               string  `json:"url"`
				} `json:"owner"`
				Private         bool    `json:"private"`
				PullsURL        string  `json:"pulls_url"`
				PushedAt        string  `json:"pushed_at"`
				ReleasesURL     string  `json:"releases_url"`
				Size            float64 `json:"size"`
				SshURL          string  `json:"ssh_url"`
				StargazersCount float64 `json:"stargazers_count"`
				StargazersURL   string  `json:"stargazers_url"`
				StatusesURL     string  `json:"statuses_url"`
				SubscribersURL  string  `json:"subscribers_url"`
				SubscriptionURL string  `json:"subscription_url"`
				SvnURL          string  `json:"svn_url"`
				TagsURL         string  `json:"tags_url"`
				TeamsURL        string  `json:"teams_url"`
				TreesURL        string  `json:"trees_url"`
				UpdatedAt       string  `json:"updated_at"`
				URL             string  `json:"url"`
				Watchers        float64 `json:"watchers"`
				WatchersCount   float64 `json:"watchers_count"`
			} `json:"repo"`
			Sha  string `json:"sha"`
			User struct {
				AvatarURL         string  `json:"avatar_url"`
				EventsURL         string  `json:"events_url"`
				FollowersURL      string  `json:"followers_url"`
				FollowingURL      string  `json:"following_url"`
				GistsURL          string  `json:"gists_url"`
				GravatarID        string  `json:"gravatar_id"`
				HtmlURL           string  `json:"html_url"`
				ID                float64 `json:"id"`
				Login             string  `json:"login"`
				OrganizationsURL  string  `json:"organizations_url"`
				ReceivedEventsURL string  `json:"received_events_url"`
				ReposURL          string  `json:"repos_url"`
				SiteAdmin         bool    `json:"site_admin"`
				StarredURL        string  `json:"starred_url"`
				SubscriptionsURL  string  `json:"subscriptions_url"`
				Type              string  `json:"type"`
				URL               string  `json:"url"`
			} `json:"user"`
		} `json:"base"`
		Body         string      `json:"body"`
		ChangedFiles float64     `json:"changed_files"`
		ClosedAt     interface{} `json:"closed_at"`
		Comments     float64     `json:"comments"`
		CommentsURL  string      `json:"comments_url"`
		Commits      float64     `json:"commits"`
		CommitsURL   string      `json:"commits_url"`
		CreatedAt    string      `json:"created_at"`
		Deletions    float64     `json:"deletions"`
		DiffURL      string      `json:"diff_url"`
		Head         struct {
			Label string `json:"label"`
			Ref   string `json:"ref"`
			Repo  struct {
				ArchiveURL       string      `json:"archive_url"`
				AssigneesURL     string      `json:"assignees_url"`
				BlobsURL         string      `json:"blobs_url"`
				BranchesURL      string      `json:"branches_url"`
				CloneURL         string      `json:"clone_url"`
				CollaboratorsURL string      `json:"collaborators_url"`
				CommentsURL      string      `json:"comments_url"`
				CommitsURL       string      `json:"commits_url"`
				CompareURL       string      `json:"compare_url"`
				ContentsURL      string      `json:"contents_url"`
				ContributorsURL  string      `json:"contributors_url"`
				CreatedAt        string      `json:"created_at"`
				DefaultBranch    string      `json:"default_branch"`
				Description      string      `json:"description"`
				DownloadsURL     string      `json:"downloads_url"`
				EventsURL        string      `json:"events_url"`
				Fork             bool        `json:"fork"`
				Forks            float64     `json:"forks"`
				ForksCount       float64     `json:"forks_count"`
				ForksURL         string      `json:"forks_url"`
				FullName         string      `json:"full_name"`
				GitCommitsURL    string      `json:"git_commits_url"`
				GitRefsURL       string      `json:"git_refs_url"`
				GitTagsURL       string      `json:"git_tags_url"`
				GitURL           string      `json:"git_url"`
				HasDownloads     bool        `json:"has_downloads"`
				HasIssues        bool        `json:"has_issues"`
				HasPages         bool        `json:"has_pages"`
				HasWiki          bool        `json:"has_wiki"`
				Homepage         interface{} `json:"homepage"`
				HooksURL         string      `json:"hooks_url"`
				HtmlURL          string      `json:"html_url"`
				ID               float64     `json:"id"`
				IssueCommentURL  string      `json:"issue_comment_url"`
				IssueEventsURL   string      `json:"issue_events_url"`
				IssuesURL        string      `json:"issues_url"`
				KeysURL          string      `json:"keys_url"`
				LabelsURL        string      `json:"labels_url"`
				Language         string      `json:"language"`
				LanguagesURL     string      `json:"languages_url"`
				MergesURL        string      `json:"merges_url"`
				MilestonesURL    string      `json:"milestones_url"`
				MirrorURL        interface{} `json:"mirror_url"`
				Name             string      `json:"name"`
				NotificationsURL string      `json:"notifications_url"`
				OpenIssues       float64     `json:"open_issues"`
				OpenIssuesCount  float64     `json:"open_issues_count"`
				Owner            struct {
					AvatarURL         string  `json:"avatar_url"`
					EventsURL         string  `json:"events_url"`
					FollowersURL      string  `json:"followers_url"`
					FollowingURL      string  `json:"following_url"`
					GistsURL          string  `json:"gists_url"`
					GravatarID        string  `json:"gravatar_id"`
					HtmlURL           string  `json:"html_url"`
					ID                float64 `json:"id"`
					Login             string  `json:"login"`
					OrganizationsURL  string  `json:"organizations_url"`
					ReceivedEventsURL string  `json:"received_events_url"`
					ReposURL          string  `json:"repos_url"`
					SiteAdmin         bool    `json:"site_admin"`
					StarredURL        string  `json:"starred_url"`
					SubscriptionsURL  string  `json:"subscriptions_url"`
					Type              string  `json:"type"`
					URL               string  `json:"url"`
				} `json:"owner"`
				Private         bool    `json:"private"`
				PullsURL        string  `json:"pulls_url"`
				PushedAt        string  `json:"pushed_at"`
				ReleasesURL     string  `json:"releases_url"`
				Size            float64 `json:"size"`
				SshURL          string  `json:"ssh_url"`
				StargazersCount float64 `json:"stargazers_count"`
				StargazersURL   string  `json:"stargazers_url"`
				StatusesURL     string  `json:"statuses_url"`
				SubscribersURL  string  `json:"subscribers_url"`
				SubscriptionURL string  `json:"subscription_url"`
				SvnURL          string  `json:"svn_url"`
				TagsURL         string  `json:"tags_url"`
				TeamsURL        string  `json:"teams_url"`
				TreesURL        string  `json:"trees_url"`
				UpdatedAt       string  `json:"updated_at"`
				URL             string  `json:"url"`
				Watchers        float64 `json:"watchers"`
				WatchersCount   float64 `json:"watchers_count"`
			} `json:"repo"`
			Sha  string `json:"sha"`
			User struct {
				AvatarURL         string  `json:"avatar_url"`
				EventsURL         string  `json:"events_url"`
				FollowersURL      string  `json:"followers_url"`
				FollowingURL      string  `json:"following_url"`
				GistsURL          string  `json:"gists_url"`
				GravatarID        string  `json:"gravatar_id"`
				HtmlURL           string  `json:"html_url"`
				ID                float64 `json:"id"`
				Login             string  `json:"login"`
				OrganizationsURL  string  `json:"organizations_url"`
				ReceivedEventsURL string  `json:"received_events_url"`
				ReposURL          string  `json:"repos_url"`
				SiteAdmin         bool    `json:"site_admin"`
				StarredURL        string  `json:"starred_url"`
				SubscriptionsURL  string  `json:"subscriptions_url"`
				Type              string  `json:"type"`
				URL               string  `json:"url"`
			} `json:"user"`
		} `json:"head"`
		HtmlURL           string      `json:"html_url"`
		ID                float64     `json:"id"`
		IssueURL          string      `json:"issue_url"`
		Locked            bool        `json:"locked"`
		MergeCommitSha    string      `json:"merge_commit_sha"`
		Mergeable         interface{} `json:"mergeable"`
		MergeableState    string      `json:"mergeable_state"`
		Merged            bool        `json:"merged"`
		MergedAt          interface{} `json:"merged_at"`
		MergedBy          interface{} `json:"merged_by"`
		Milestone         interface{} `json:"milestone"`
		Number            float64     `json:"number"`
		PatchURL          string      `json:"patch_url"`
		ReviewCommentURL  string      `json:"review_comment_url"`
		ReviewComments    float64     `json:"review_comments"`
		ReviewCommentsURL string      `json:"review_comments_url"`
		State             string      `json:"state"`
		StatusesURL       string      `json:"statuses_url"`
		Title             string      `json:"title"`
		UpdatedAt         string      `json:"updated_at"`
		URL               string      `json:"url"`
		User              struct {
			AvatarURL         string  `json:"avatar_url"`
			EventsURL         string  `json:"events_url"`
			FollowersURL      string  `json:"followers_url"`
			FollowingURL      string  `json:"following_url"`
			GistsURL          string  `json:"gists_url"`
			GravatarID        string  `json:"gravatar_id"`
			HtmlURL           string  `json:"html_url"`
			ID                float64 `json:"id"`
			Login             string  `json:"login"`
			OrganizationsURL  string  `json:"organizations_url"`
			ReceivedEventsURL string  `json:"received_events_url"`
			ReposURL          string  `json:"repos_url"`
			SiteAdmin         bool    `json:"site_admin"`
			StarredURL        string  `json:"starred_url"`
			SubscriptionsURL  string  `json:"subscriptions_url"`
			Type              string  `json:"type"`
			URL               string  `json:"url"`
		} `json:"user"`
	} `json:"pull_request"`
	Label struct {
		Color string `json:"color"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"label"`
	Repository struct {
		ArchiveURL       string      `json:"archive_url"`
		AssigneesURL     string      `json:"assignees_url"`
		BlobsURL         string      `json:"blobs_url"`
		BranchesURL      string      `json:"branches_url"`
		CloneURL         string      `json:"clone_url"`
		CollaboratorsURL string      `json:"collaborators_url"`
		CommentsURL      string      `json:"comments_url"`
		CommitsURL       string      `json:"commits_url"`
		CompareURL       string      `json:"compare_url"`
		ContentsURL      string      `json:"contents_url"`
		ContributorsURL  string      `json:"contributors_url"`
		CreatedAt        string      `json:"created_at"`
		DefaultBranch    string      `json:"default_branch"`
		Description      string      `json:"description"`
		DownloadsURL     string      `json:"downloads_url"`
		EventsURL        string      `json:"events_url"`
		Fork             bool        `json:"fork"`
		Forks            float64     `json:"forks"`
		ForksCount       float64     `json:"forks_count"`
		ForksURL         string      `json:"forks_url"`
		FullName         string      `json:"full_name"`
		GitCommitsURL    string      `json:"git_commits_url"`
		GitRefsURL       string      `json:"git_refs_url"`
		GitTagsURL       string      `json:"git_tags_url"`
		GitURL           string      `json:"git_url"`
		HasDownloads     bool        `json:"has_downloads"`
		HasIssues        bool        `json:"has_issues"`
		HasPages         bool        `json:"has_pages"`
		HasWiki          bool        `json:"has_wiki"`
		Homepage         interface{} `json:"homepage"`
		HooksURL         string      `json:"hooks_url"`
		HtmlURL          string      `json:"html_url"`
		ID               float64     `json:"id"`
		IssueCommentURL  string      `json:"issue_comment_url"`
		IssueEventsURL   string      `json:"issue_events_url"`
		IssuesURL        string      `json:"issues_url"`
		KeysURL          string      `json:"keys_url"`
		LabelsURL        string      `json:"labels_url"`
		Language         string      `json:"language"`
		LanguagesURL     string      `json:"languages_url"`
		MergesURL        string      `json:"merges_url"`
		MilestonesURL    string      `json:"milestones_url"`
		MirrorURL        interface{} `json:"mirror_url"`
		Name             string      `json:"name"`
		NotificationsURL string      `json:"notifications_url"`
		OpenIssues       float64     `json:"open_issues"`
		OpenIssuesCount  float64     `json:"open_issues_count"`
		Owner            struct {
			AvatarURL         string  `json:"avatar_url"`
			EventsURL         string  `json:"events_url"`
			FollowersURL      string  `json:"followers_url"`
			FollowingURL      string  `json:"following_url"`
			GistsURL          string  `json:"gists_url"`
			GravatarID        string  `json:"gravatar_id"`
			HtmlURL           string  `json:"html_url"`
			ID                float64 `json:"id"`
			Login             string  `json:"login"`
			OrganizationsURL  string  `json:"organizations_url"`
			ReceivedEventsURL string  `json:"received_events_url"`
			ReposURL          string  `json:"repos_url"`
			SiteAdmin         bool    `json:"site_admin"`
			StarredURL        string  `json:"starred_url"`
			SubscriptionsURL  string  `json:"subscriptions_url"`
			Type              string  `json:"type"`
			URL               string  `json:"url"`
		} `json:"owner"`
		Private         bool    `json:"private"`
		PullsURL        string  `json:"pulls_url"`
		PushedAt        string  `json:"pushed_at"`
		ReleasesURL     string  `json:"releases_url"`
		Size            float64 `json:"size"`
		SshURL          string  `json:"ssh_url"`
		StargazersCount float64 `json:"stargazers_count"`
		StargazersURL   string  `json:"stargazers_url"`
		StatusesURL     string  `json:"statuses_url"`
		SubscribersURL  string  `json:"subscribers_url"`
		SubscriptionURL string  `json:"subscription_url"`
		SvnURL          string  `json:"svn_url"`
		TagsURL         string  `json:"tags_url"`
		TeamsURL        string  `json:"teams_url"`
		TreesURL        string  `json:"trees_url"`
		UpdatedAt       string  `json:"updated_at"`
		URL             string  `json:"url"`
		Watchers        float64 `json:"watchers"`
		WatchersCount   float64 `json:"watchers_count"`
	} `json:"repository"`
	Sender struct {
		AvatarURL         string  `json:"avatar_url"`
		EventsURL         string  `json:"events_url"`
		FollowersURL      string  `json:"followers_url"`
		FollowingURL      string  `json:"following_url"`
		GistsURL          string  `json:"gists_url"`
		GravatarID        string  `json:"gravatar_id"`
		HtmlURL           string  `json:"html_url"`
		ID                float64 `json:"id"`
		Login             string  `json:"login"`
		OrganizationsURL  string  `json:"organizations_url"`
		ReceivedEventsURL string  `json:"received_events_url"`
		ReposURL          string  `json:"repos_url"`
		SiteAdmin         bool    `json:"site_admin"`
		StarredURL        string  `json:"starred_url"`
		SubscriptionsURL  string  `json:"subscriptions_url"`
		Type              string  `json:"type"`
		URL               string  `json:"url"`
	} `json:"sender"`
}
