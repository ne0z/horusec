// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package start_test

import (
	"fmt"
	"time"

	"github.com/ZupIT/horusec/internal/utils/testutil"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("running binary Horusec with start parameter", func() {
	var (
		session           *gexec.Session
		err               error
		flags             map[string]string
		repoAuthorization string
	)

	JustBeforeEach(func() {
		cmd := testutil.GinkgoGetHorusecCmdWithFlags(testutil.CmdStart, flags)
		session, err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

	})

	When("--project-path is passed", func() {
		BeforeEach(func() {
			flags = map[string]string{
				testutil.StartFlagProjectPath: testutil.GoExample1,
			}
		})

		It("Returns no error on execution", func() {
			Expect(err).NotTo(HaveOccurred())
		})

		It("Shows the project path property", func() {
			Expect(session.Wait(2 * time.Minute).Out.Contents()).To(ContainSubstring("project_path"))
		})

		It("The path is set", func() {
			Eventually(session).Should(gbytes.Say(fmt.Sprintf("%s", testutil.GoExample1)))
		})

		It("Finish with exit code 0", func() {
			Eventually(session.Wait(4 * time.Minute)).Should(gexec.Exit(0))
		})
	})

	When("--analysis-timeout is passed", func() {
		BeforeEach(func() {
			flags = map[string]string{
				testutil.StartFlagProjectPath:     testutil.GoExample1,
				testutil.StartFlagAnalysisTimeout: "500",
			}
		})
		It("Returns no error on execution", func() {
			Expect(err).NotTo(HaveOccurred())
		})

		It("Show warning message with timeout set", func() {
			Expect(session.Wait(2 * time.Minute).Out.Contents()).To(ContainSubstring("Horusec will return a timeout after 500 seconds."))
		})

		It("Finish with exit code 0", func() {
			Eventually(session.Wait(2 * time.Minute)).Should(gexec.Exit(0))
		})
	})

	When("--authorization is passed", func() {
		BeforeEach(func() {
			repoAuthorization = uuid.New().String()

			flags = map[string]string{
				testutil.StartFlagProjectPath:   testutil.GoExample1,
				testutil.StartFlagAuthorization: repoAuthorization,
			}
		})

		It("Returns no error on execution", func() {
			Expect(err).NotTo(HaveOccurred())
		})

		It("Shows the repository property", func() {
			Expect(session.Wait(2 * time.Minute).Out.Contents()).To(ContainSubstring("repository_authorization"))
		})

		It("The authorization token is set", func() {
			Eventually(session).Should(gbytes.Say(fmt.Sprintf("%s", repoAuthorization)))
		})

		It("Finish with exit code 0", func() {
			Eventually(session.Wait(2 * time.Minute)).Should(gexec.Exit(0))
		})
	})
})
