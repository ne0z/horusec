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
	"github.com/ZupIT/horusec/e2e/commands/start/flags"
	"github.com/ZupIT/horusec/internal/entities/e2e"
	"github.com/ZupIT/horusec/internal/utils/testutil"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega/gexec"
)


var _ = Describe("running binary Horusec with start parameter", func() {
	constructor := &e2e.Constructor{}
	JustBeforeEach(func() {
		cmd := testutil.GinkgoGetHorusecCmdWithFlags(testutil.CmdStart, constructor.Flags)
		constructor.Session, constructor.Err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	})

	When("--project-path is passed", flags.FlagProjectPath(constructor))

	When("--analysis-timeout is passed", flags.FlagAnalysisTimeOut(constructor))

	When("--authorization is passed", flags.FlagAuthorization(constructor))
})
