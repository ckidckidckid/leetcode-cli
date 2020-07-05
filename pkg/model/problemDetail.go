package model

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/kavimaluskam/leetcode-cli/pkg/utils"
)

// ProblemDetail is the response from leetcode GraphQL API
// concerning individual problem detail
type ProblemDetail struct {
	QuestionID            string                `json:"questionId"`
	QuestionFrontendID    string                `json:"questionFrontendId"`
	BoundTopicID          string                `json:"boundTopicId"`
	Title                 string                `json:"title"`
	TitleSlug             string                `json:"titleSlug"`
	Content               string                `json:"content"`
	TranslatedTitle       string                `json:"translatedTitle"`
	TrnslatedContent      string                `json:"translatedContent"`
	IsPaidOnly            bool                  `json:"isPaidOnly"`
	Diffculty             string                `json:"difficulty"`
	Likes                 int                   `json:"likes"`
	Dislikes              int                   `json:"dislikes"`
	IsLiked               bool                  `json:"isLiked"`
	SimilarQuestions      string                `json:"similarQuestions"`
	Contributors          []ProblemContributor  `json:"contributors"`
	LangToValidPlayground string                `json:"langToValidPlayground"`
	TopicTags             []ProblemTag          `json:"topicTags"`
	CompanyTagStats       string                `json:"companyTagStats"`
	CodeSnippets          []ProblemCodeSnippets `json:"codeSnippets"`
	Stats                 string                `json:"stats"`
	Hints                 []string              `json:"hints"`
	Solution              ProblemSolution       `json:"solution"`
	Status                string                `json:"status"`
	SampleTestCase        string                `json:"sampleTestCase"`
	MetaData              string                `json:"metaData"`
	JudgerAvailable       bool                  `json:"judgerAvailable"`
	JudgeType             string                `json:"judgeType"`
	MySQLSchemas          []string              `json:"mysqlSchemas"`
	EnableRuneCode        bool                  `json:"enableRunCode"`
	EnableTestMode        bool                  `json:"enableTestMode"`
	EnableDebugger        bool                  `json:"enableDebugger"`
	EnvInfo               string                `json:"envInfo"`
	LibraryURL            string                `json:"libraryUrl"`
	AdminURL              string                `json:"adminUrl"`
	TypeName              string                `json:"__typename"`
}

// ProblemContributor is the response from leetcode GraphQL API
// concering problem contributor
type ProblemContributor struct {
	Username   string `json:"username"`
	ProfileURL string `json:"profileUrl"`
	AvatarURL  string `json:"avatarUrl"`
	TypeName   string `json:"__typename"`
}

// ProblemTag is the response from leetcode GraphQL API
// concerning problem tagging
type ProblemTag struct {
	Name           string `json:"name"`
	Slug           string `json:"slug"`
	TranslatedName string `json:"translatedName"`
	TypeName       string `json:"__typename"`
}

// ProblemCodeSnippets is the response from leetcode GraphQL API
// concerning problem codes
type ProblemCodeSnippets struct {
	Lang     string `json:"lang"`
	LangSlug string `json:"langSlug"`
	Code     string `json:"code"`
	TypeName string `json:"__typename"`
}

// ProblemSolution is the response from leetcode GraphQL API
// concerning problem solutions
type ProblemSolution struct {
	ID           string `json:"id"`
	CanSeeDetail bool   `json:"canSeeDetail"`
	PaidOnly     bool   `json:"paidOnly"`
	TypeName     string `json:"__typename"`
}

// ProblemStats is the string response from leetcode GraphQL API
// concerning problem stats
type ProblemStats struct {
	TotalAccepted      string `json:"totalAccepted"`
	TotalSubmission    string `json:"totalSubmission"`
	TotalAcceptedRaw   int    `json:"totalAcceptedRaw"`
	TotalSubmissionRaw int    `json:"totalSubmissionRaw"`
	AcceptRate         string `json:"acRate"`
}

// GetDiffculty is a mapper function from problem diffculty level to string
func (pd ProblemDetail) GetDiffculty() string {
	switch pd.Diffculty {
	case "Easy":
		return utils.Green("Easy")
	case "Medium":
		return utils.Yellow("Medium")
	default:
		return utils.Red("Hard")
	}
}

// GetStats is a property function unmarshal json string field `stats`
func (pd ProblemDetail) GetStats() (*ProblemStats, error) {
	ps := &ProblemStats{}
	err := json.Unmarshal([]byte(pd.Stats), ps)
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (pd ProblemDetail) ExportStdoutDetail() error {
	pds, err := pd.GetStats()
	if err != nil {
		return err
	}

	p := strings.NewReader(pd.Content)
	parsedContent, err := goquery.NewDocumentFromReader(p)
	if err != nil {
		return err
	}

	fmt.Printf("[%s] %s\n", pd.QuestionID, pd.Title)
	fmt.Println()
	fmt.Println(utils.Gray(strings.Replace(utils.ProblemURL, "$slug", pd.TitleSlug, 1)))
	fmt.Println()
	fmt.Printf("* %s (%s)\n", pd.GetDiffculty(), pds.AcceptRate)
	fmt.Printf("* Total Accepted:    %d\n", pds.TotalAcceptedRaw)
	fmt.Printf("* Total Submissions: %d\n", pds.TotalSubmissionRaw)
	fmt.Println("* Testcase Example:")
	for _, line := range strings.Split(pd.SampleTestCase, "\n") {
		fmt.Printf("  %s\n", line)
	}
	fmt.Println()
	fmt.Println(parsedContent.Text())

	return nil
}