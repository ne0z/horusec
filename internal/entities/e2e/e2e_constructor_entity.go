package e2e

import "github.com/onsi/gomega/gexec"

type Constructor struct {
	Session           *gexec.Session
	Err               error
	Flags             map[string]string
	RepoAuthorization string
}
