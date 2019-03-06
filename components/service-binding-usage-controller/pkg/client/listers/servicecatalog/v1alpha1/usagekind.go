// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kyma-project/kyma/components/service-binding-usage-controller/pkg/apis/servicecatalog/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// UsageKindLister helps list UsageKinds.
type UsageKindLister interface {
	// List lists all UsageKinds in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.UsageKind, err error)
	// Get retrieves the UsageKind from the index for a given name.
	Get(name string) (*v1alpha1.UsageKind, error)
	UsageKindListerExpansion
}

// usageKindLister implements the UsageKindLister interface.
type usageKindLister struct {
	indexer cache.Indexer
}

// NewUsageKindLister returns a new UsageKindLister.
func NewUsageKindLister(indexer cache.Indexer) UsageKindLister {
	return &usageKindLister{indexer: indexer}
}

// List lists all UsageKinds in the indexer.
func (s *usageKindLister) List(selector labels.Selector) (ret []*v1alpha1.UsageKind, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.UsageKind))
	})
	return ret, err
}

// Get retrieves the UsageKind from the index for a given name.
func (s *usageKindLister) Get(name string) (*v1alpha1.UsageKind, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("usagekind"), name)
	}
	return obj.(*v1alpha1.UsageKind), nil
}