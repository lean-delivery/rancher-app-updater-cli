package _context

var Pipeline pipeline

func InitPipelineContext(branchName string, buildNumber int, serviceName string) *pipeline{
	Pipeline = pipeline{
		BranchName: branchName,
		BuildNumber: buildNumber,
		ServiceName: serviceName,
	}

	return &Pipeline
}

type pipeline struct {
	BranchName string
	BuildNumber int
	ServiceName string
}