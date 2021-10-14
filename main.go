package main

import (
	git "github.com/libgit2/git2go/v31"
	"github.com/sirupsen/logrus"
	"os"
)

type BasicAuth struct {
	Username, Password string
}

func (a *BasicAuth) credentialsCallback(url string, usernameFromURL string, allowedTypes git.CredentialType) (*git.Credential, error) {
	cert, err := git.NewCredentialUserpassPlaintext(a.Username, a.Password)
	if err != nil {
		logrus.Errorf("[Credential] NewUserpassPlaintext err is [%v]", err)
		return nil, err
	}
	return cert, err
}

func main() {
	url, directory, username, password := os.Args[1], os.Args[2], os.Args[3], os.Args[4]
	logrus.Infof("%v   %v   %v  %v ", url, directory, username, password)

	auth := BasicAuth{Username: username, Password: password}
	cloneOptions := &git.CloneOptions{
		FetchOptions: &git.FetchOptions{
			RemoteCallbacks: git.RemoteCallbacks{
				CredentialsCallback: auth.credentialsCallback,
			},
		},
	}

	repo, err := git.Clone(url, directory, cloneOptions)
	if err != nil {
		logrus.Errorf("[Clone] err is %v", err)
		return
	}
	logrus.Infof("repo is %v", repo)
	head, err := repo.Head()
	if err != nil {
		logrus.Errorf("[Head] err is %v", err)
		return
	}
	headCommit, err := repo.LookupCommit(head.Target())
	if err != nil {
		logrus.Errorf("[HeadCommit] err is %v", err)
		return
	}
	commitID, err := headCommit.AsObject().ShortId()
	if err != nil {
		logrus.Errorf("[HeadCommit] get short id err is %v", err)
		return
	}
	logrus.Infof("Author is [%v], Hash is [%v], Short hash is [%v], Commit msg is [%v]", headCommit.Author(), headCommit.AsObject().Id(), commitID, headCommit.Message())
}
