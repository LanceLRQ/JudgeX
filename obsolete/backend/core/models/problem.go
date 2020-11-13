package models


type ProblemSample struct {
	Input string 		`json:"input"`		// Input
	Output string 		`json:"output"`		// Output
}

type ProblemSampleList []ProblemSample
