package main

import "testing"

func Test_buildRepo(t *testing.T) {
	tests := []struct {
		name     string
		registry string
		repo     string
		want     string
	}{
		{
			name: "dockerhub",
			repo: "golang",
			want: "golang",
		},
		{
			name:     "internal",
			registry: "artifactory.example.com",
			repo:     "service",
			want:     "artifactory.example.com/service",
		},
		{
			name:     "backward_compatibility",
			registry: "artifactory.example.com",
			repo:     "artifactory.example.com/service",
			want:     "artifactory.example.com/service",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildRepo(tt.registry, tt.repo, true); got != tt.want {
				t.Errorf("buildRepo(%q, %q) = %v, want %v", tt.registry, tt.repo, got, tt.want)
			}
		})
	}
}

func TestFormatRepoName(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Example test case 1",
			input: "MyRepo",
			want:  "myrepo",
		},
		{
			name:  "Example test case 2",
			input: "my-repo",
			want:  "my-repo",
		},
		{
			name:  "Example test case 3",
			input: "My.Repo",
			want:  "my-repo",
		},
		{
			name:  "Example test case 4",
			input: "my/repo",
			want:  "my-repo",
		},
		{
			name:  "Example test case 5",
			input: "saas/service.app",
			want:  "saas-service-app",
		},
		{
			name:  "Example test case 6",
			input: "jarvis2/jarvis2.api",
			want:  "jarvis2-jarvis2-api",
		},
		{
			name:  "Example test case 7",
			input: "saas/goCrawler",
			want:  "saas-gocrawler",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatRepoName(tt.input); got != tt.want {
				t.Errorf("formatRepoName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatRepoPath(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Trim repo path",
			input: "/path/to/repo",
			want:  "path/to/repo",
		},
		{
			name:  "Trim empty repo path",
			input: "",
			want:  "",
		},
		{
			name:  "Trim repo path with trailing slash",
			input: "xhj-prod-registry-vpc.cn-hangzhou.cr.aliyuncs.com/xhj-image/",
			want:  "xhj-prod-registry-vpc.cn-hangzhou.cr.aliyuncs.com/xhj-image",
		},
		{
			name:  "Trim repo path with trailing slash",
			input: "xhj-prod-registry-vpc.cn-hangzhou.cr.aliyuncs.com/xhj-image",
			want:  "xhj-prod-registry-vpc.cn-hangzhou.cr.aliyuncs.com/xhj-image",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := formatRepoPath(test.input)
			if got != test.want {
				t.Errorf("formatRepoPath(%q) = %q, want %q", test.input, got, test.want)
			}
		})
	}
}
