package cmd

import (
	"context"
	"encoding/json"
	"gopkg.in/workanator/go-floc.v2"
	"log"
)

type SequenceType int

const (
	Parallel SequenceType = 0
	Sync     SequenceType = 1
)

func (t SequenceType) String() string {
	names := [...]string{
		"Parallel",
		"Sync"}
	if t < Parallel || t > Sync {
		return "Unknown"
	}
	return names[t]
}

type Var struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	JPath string `json:"jPath"`
}

type Job struct {
	Id     string      `json:"id"`
	Url    string      `json:"url"`
	Method string      `json:"method"`
	Body   interface{} `json:"body"`
	Header interface{} `json:"header"`
	Var    []Var       `json:"var"`
}

func (j *Job) GetUrl(ctx floc.Context) string {
	return j.Url
}

func (j *Job) GetBody(ctx context.Context) string {
	bInput, _ := json.Marshal(j.Body)
	return string(bInput)
}

func (j *Job) GetHeaders(ctx floc.Context) map[string]string {
	var resp = make(map[string]string)
	var result map[string]interface{}
	bInput, _ := json.Marshal(j.Header)
	err := json.Unmarshal(bInput, &result)
	if err != nil {
		log.Print(err)
	}
	for key, item := range result {
		v := item.(string)
		resp[key] = v

		if v == "$token" {
			if v, ok := ctx.Value("token").(string); ok {
				resp[key] = v
			}
		}

	}
	return resp
}

type Sequence struct {
	Type SequenceType `json:"type"`
	Jobs []Job        `json:"jobs"`
}

type Request struct {
	Sequence []Sequence `json:"sequence"`
}
