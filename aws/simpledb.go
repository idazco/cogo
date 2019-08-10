package aws

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/simpledb"
)

type SimpleDB struct {
	Session *session.Session
	svc     *simpledb.SimpleDB
}

// Return a new SimpleDB instance.
// Call StartSession or StartSessionFromProfile before this func.
func NewSimpleDB() (*SimpleDB, error) {
	if Session == nil {
		return nil, errors.New("session not set; a session must be started first")
	}
	return &SimpleDB{Session: Session, svc: simpledb.New(Session)}, nil
}

// Returns a list of domain, up to 100 domains in the list
func (sdb *SimpleDB) ListDomains() (*simpledb.ListDomainsOutput, error) {
	r, err := sdb.svc.ListDomains(&simpledb.ListDomainsInput{})
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Create a domain (idempotent function)
func (sdb *SimpleDB) CreateDomain(name string) (*simpledb.CreateDomainOutput, error) {
	r, err := sdb.svc.CreateDomain(&simpledb.CreateDomainInput{DomainName: aws.String(name)})
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Delete a domain
func (sdb *SimpleDB) DeleteDomain(name string) (*simpledb.DeleteDomainOutput, error) {
	r, err := sdb.svc.DeleteDomain(&simpledb.DeleteDomainInput{DomainName: aws.String(name)})
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Put attributes into AWS simple db
func (sdb *SimpleDB) PutAttributes(domain string, itemName string, attributes []*simpledb.ReplaceableAttribute) (*simpledb.PutAttributesOutput ,error) {
	r, err := sdb.svc.PutAttributes(&simpledb.PutAttributesInput{
		Attributes: attributes,
		DomainName: aws.String(domain),   // Required
		ItemName:   aws.String(itemName), // Required
		//Expected: &simpledb.UpdateCondition{
		//	Exists: aws.Bool(true),
		//	Name:   aws.String("String"),
		//	Value:  aws.String("String"),
		//},
	})
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Returns all the attributes for an item
func (sdb *SimpleDB) GetAttributes(domain string, item string) (map[string]string, error) {
	r, err := sdb.svc.GetAttributes(&simpledb.GetAttributesInput{
		DomainName:     aws.String(domain), // Required
		ItemName:       aws.String(item),   // Required
		ConsistentRead: aws.Bool(true),
	})
	if err != nil {
		return nil, err
	}

	if len(r.Attributes) > 0 {
		retval := map[string]string{}
		for n := range r.Attributes {
			name := string(*r.Attributes[n].Name)
			val := string(*r.Attributes[n].Value)
			retval[name] = val
		}
		return retval, nil
	}
	return nil, nil
}

// Delete an item and all its attributes
func (sdb *SimpleDB) DeleteItem(domain string, item string) error {
	_, err := sdb.svc.DeleteAttributes(&simpledb.DeleteAttributesInput{
		DomainName:     aws.String(domain), // Required
		ItemName:       aws.String(item),   // Required
	})
	return err
}

// Make a select query
func (sdb *SimpleDB) Select(SelectExpression string) (*simpledb.SelectOutput, error) {
	r, err := sdb.svc.Select(&simpledb.SelectInput{SelectExpression: aws.String(SelectExpression)})
	if err != nil {
		return nil, err
	}
	return r, nil
}
