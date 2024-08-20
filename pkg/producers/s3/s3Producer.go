// Copyright © 2024 JR team
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package s3

import (
	"bytes"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
)

type Config struct {
	AWSRegion string `json:"aws_region"`
	Bucket    string `json:"bucket"`
}

type S3Producer struct {
	client s3.S3
	bucket string
}

func (p *S3Producer) Initialize(configFile string) {
	var config Config
	file, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to ReadFile")
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse configuration parameters")
	}

	sess, err := session.NewSession(&aws.Config{Region: &config.AWSRegion})

	if err != nil {
		log.Fatal().Err(err).Msg("Can't establish a session to S3")
		return
	}

	s3Client := s3.New(sess)

	p.client = *s3Client
	p.bucket = config.Bucket
}

func (p *S3Producer) Produce(k []byte, v []byte, o any) {

	bucket := p.bucket
	var key string

	if len(k) == 0 || strings.ToLower(string(k)) == "null" {
		// generate a UUID as index
		key = uuid.New().String()
	} else {
		key = string(k)
	}

	//object will be stored with no content type
	_, err := p.client.PutObject(&s3.PutObjectInput{
		Body:   bytes.NewReader(v),
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to write data in s3")
	}
}

func (p *S3Producer) Close() error {
	log.Warn().Msg("S3 Client doesn't provide a close method!")
	return nil
}
