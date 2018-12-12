package helpers

import (
	"encoding/json"
	"fmt"
	"strings"

	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
)

func Curl(obj interface{}, url string, props ...interface{}) {
	session := NonExperimentalCurl(fmt.Sprintf(url, props...))
	Eventually(session).Should(Exit(0))
	rawJSON := strings.TrimSpace(string(session.Out.Contents()))

	err := json.Unmarshal([]byte(rawJSON), &obj)
	Expect(err).NotTo(HaveOccurred())
}

func NonExperimentalCurl(args ...string) *Session {
	curlArgs := []string{"curl"}
	curlArgs = append(curlArgs, args...)
	return CFWithEnv(map[string]string{"CF_CLI_EXPERIMENTAL": "false"}, curlArgs...)
}
