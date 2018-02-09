package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/idazco/cogo/log"
)

// ListObjects lists objects in an S3 AWS bucket
func S3_ListObjects(bucket, prefix, delimeter string) (*s3.ListObjectsOutput, bool) {
	if Session == nil {
		log.AppError("tried calling S3 function without starting an AWS session first")
		return nil, false
	}

	svc := s3.New(Session)

	params := &s3.ListObjectsInput{
		Bucket:    aws.String(bucket), // Required
		MaxKeys:   aws.Int64(10000),
		Delimiter: aws.String(delimeter),
		//		EncodingType: aws.String("EncodingType"),
		//		Marker:       aws.String("Marker"),
	}
	if prefix != "" {
		params.Prefix = aws.String(prefix)
	}
	resp, err := svc.ListObjects(params)

	if err != nil {
		if aErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			log.AppError(fmt.Sprintf("AWS S3 standard error %v: %s (origin: %v)", aErr.Code(), aErr.Message(), aErr.OrigErr()))
			return nil, false
		} else if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			log.AppError(fmt.Sprintf("AWS S3 service error %v: %s (status: %v)", reqErr.Code(), reqErr.Message(), reqErr.StatusCode()))
			return nil, false
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			log.Error("S3 ListObjects failed", err)
			return nil, false
		}
	}
	return resp, true
}
