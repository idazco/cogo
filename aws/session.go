package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/idazco/cogo/log"
)

var Session *session.Session

func StartSession(key, secret, region string) bool {
	if len(region) < 5 {
		log.AppError("region param cannot be an empty string when calling cogo/aws/session/StartSession")
		return false
	}

	var err error
	// if static credential are provided then use those
	if len(key) > 0 && len(secret) > 0 {
		staticCreds := credentials.NewStaticCredentials(key, secret, "")
		if _, err := staticCreds.Get(); err != nil {
			log.Error("cogo/aws/StartSession(..) failed", err)
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

func StartSessionFromProfile(profile string) error {
	if profile == "" {
		profile = "default"
	}
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: profile,
		// Force enable Shared Config support
		SharedConfigState: session.SharedConfigEnable,
	})
	if err == nil {
		Session = sess
	}
	return nil
}
