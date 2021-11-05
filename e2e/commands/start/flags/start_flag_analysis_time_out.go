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

package flags

import (
	"fmt"
	"github.com/ZupIT/horusec/internal/entities/e2e"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"

	"github.com/ZupIT/horusec/internal/utils/testutil"
)

func FlagAnalysisTimeOut(constructor *e2e.Constructor) func() {
	return func() {
		BeforeEach(func() {
			constructor.Flags = map[string]string{
				testutil.StartFlagProjectPath: testutil.GoExample1,
			}
		})

		It("Returns no error on execution", func() {
			Expect(constructor.Err).NotTo(HaveOccurred())
		})

		It("Shows the project path property", func() {
			Expect(constructor.Session.Wait(2 * time.Minute).Out.Contents()).To(ContainSubstring("project_path"))
		})

		It("The path is set", func() {
			Eventually(constructor.Session).Should(gbytes.Say(fmt.Sprintf("%s", testutil.GoExample1)))
		})

		It("Finish with exit code 0", func() {
			Eventually(constructor.Session.Wait(4 * time.Minute)).Should(gexec.Exit(0))
		})
	}
}
