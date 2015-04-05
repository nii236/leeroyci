package web

import (
	"leeroy/config"
	"leeroy/logging"
	"net/http"
	"net/http/httptest"
	//"strings"
	"testing"
)

func TestStatus(t *testing.T) {
	rw := httptest.NewRecorder()
	req, _ := http.NewRequest("Get", "127.0.0.1", nil)
	c := &config.Config{}
	blog := &logging.Buildlog{
		Jobs: []*logging.Job{
			&logging.Job{},
		},
	}

	Status(rw, req, c, blog)

	if rw.Code != 200 {
		t.Error("Wrong status code: ", 200)
	}
}

func TestRepo(t *testing.T) {
	rw := httptest.NewRecorder()
	req, _ := http.NewRequest("Get", "127.0.0.1/status/repo/75726c", nil)
	c := &config.Config{}
	blog := &logging.Buildlog{
		Jobs: []*logging.Job{
			&logging.Job{
				URL:    "url",
				Branch: "foo",
			},
		},
	}

	Repo(rw, req, c, blog)

	if rw.Code != 200 {
		t.Error("Wrong status code: ", 200)
	}
}

func TestBranch(t *testing.T) {
	rw := httptest.NewRecorder()
	req, _ := http.NewRequest("Get", "127.0.0.1/status/branch/75726c/foo", nil)
	c := &config.Config{}
	blog := &logging.Buildlog{
		Jobs: []*logging.Job{
			&logging.Job{
				URL:    "url",
				Branch: "foo",
			},
		},
	}

	Branch(rw, req, c, blog)

	if rw.Code != 200 {
		t.Error("Wrong status code: ", 200)
	}
}

func TestCommit(t *testing.T) {
	rw := httptest.NewRecorder()
	req, _ := http.NewRequest("Get", "127.0.0.1/status/branch/75726c/foo", nil)
	c := &config.Config{}
	blog := &logging.Buildlog{
		Jobs: []*logging.Job{
			&logging.Job{
				URL:    "url",
				Commit: "foo",
			},
		},
	}

	Commit(rw, req, c, blog)

	if rw.Code != 200 {
		t.Error("Wrong status code: ", 200)
	}
}

func TestBadge(t *testing.T) {
	rw := httptest.NewRecorder()
	req, _ := http.NewRequest("Get", "127.0.0.1/status/badge/75726c/foo", nil)
	c := &config.Config{}
	blog := &logging.Buildlog{
		Jobs: []*logging.Job{
			&logging.Job{},
		},
	}

	Badge(rw, req, c, blog)

	if rw.Code != 200 {
		t.Error("Wrong status code: ", 200)
	}

	if ct, ok := rw.Header()["Content-Type"]; ok == true {
		if ct[0] != "image/svg+xml" {
			t.Error("Wrong content type: ", ct[0])
		}
	} else {
		t.Error("Content Type not found")
	}
}