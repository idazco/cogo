package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/idazco/cogo/log"
	"github.com/aws/aws-sdk-go/aws"
)

var Session *session.Session

func StartSession(key, secret, region string) bool {
	if len(region) < 5 {
		log.AppError("region param cannot be an emptry string when calling cogo/aws/session/StartSession")
		return false
	}

	var err error
	// if static credential are provided then use those
	if len(key) > 0 && len(secret) > 0 {
		staticCreds := credentials.NewStaticCredentials(key, secret, "")
		if _, err := staticCreds.Get(); err != nil {
			log.Error("wpp-golang-lib/awssession.New(..) failed", err)
			return false
		}
		config := &aws.Config{Credentials: staticCreds, Region: aws.String(region)}
		Session, err = session.NewSession(config)
		if err != nil {
			log.Error("creating AWS session with credentials failed", err)
			return false
		}

	} else {
		Session, err = session.NewSession(&aws.Config{Region: aws.String(region)})
		if err != nil {
			log.Error("creating AWS session without credentials failed", err)
			return false
		}
	}
	return true
}