/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/project-flotta/flotta-operator/api/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePlaybookExecutions implements PlaybookExecutionInterface
type FakePlaybookExecutions struct {
	Fake *FakeManagementV1alpha1
	ns   string
}

var playbookexecutionsResource = schema.GroupVersionResource{Group: "management.project-flotta.io", Version: "v1alpha1", Resource: "playbookexecutions"}

var playbookexecutionsKind = schema.GroupVersionKind{Group: "management.project-flotta.io", Version: "v1alpha1", Kind: "PlaybookExecution"}

// Get takes name of the playbookExecution, and returns the corresponding playbookExecution object, and an error if there is any.
func (c *FakePlaybookExecutions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PlaybookExecution, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(playbookexecutionsResource, c.ns, name), &v1alpha1.PlaybookExecution{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PlaybookExecution), err
}

// List takes label and field selectors, and returns the list of PlaybookExecutions that match those selectors.
func (c *FakePlaybookExecutions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PlaybookExecutionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(playbookexecutionsResource, playbookexecutionsKind, c.ns, opts), &v1alpha1.PlaybookExecutionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PlaybookExecutionList{ListMeta: obj.(*v1alpha1.PlaybookExecutionList).ListMeta}
	for _, item := range obj.(*v1alpha1.PlaybookExecutionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested playbookExecutions.
func (c *FakePlaybookExecutions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(playbookexecutionsResource, c.ns, opts))

}

// Create takes the representation of a playbookExecution and creates it.  Returns the server's representation of the playbookExecution, and an error, if there is any.
func (c *FakePlaybookExecutions) Create(ctx context.Context, playbookExecution *v1alpha1.PlaybookExecution, opts v1.CreateOptions) (result *v1alpha1.PlaybookExecution, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(playbookexecutionsResource, c.ns, playbookExecution), &v1alpha1.PlaybookExecution{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PlaybookExecution), err
}

// Update takes the representation of a playbookExecution and updates it. Returns the server's representation of the playbookExecution, and an error, if there is any.
func (c *FakePlaybookExecutions) Update(ctx context.Context, playbookExecution *v1alpha1.PlaybookExecution, opts v1.UpdateOptions) (result *v1alpha1.PlaybookExecution, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(playbookexecutionsResource, c.ns, playbookExecution), &v1alpha1.PlaybookExecution{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PlaybookExecution), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePlaybookExecutions) UpdateStatus(ctx context.Context, playbookExecution *v1alpha1.PlaybookExecution, opts v1.UpdateOptions) (*v1alpha1.PlaybookExecution, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(playbookexecutionsResource, "status", c.ns, playbookExecution), &v1alpha1.PlaybookExecution{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PlaybookExecution), err
}

// Delete takes name of the playbookExecution and deletes it. Returns an error if one occurs.
func (c *FakePlaybookExecutions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(playbookexecutionsResource, c.ns, name, opts), &v1alpha1.PlaybookExecution{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePlaybookExecutions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(playbookexecutionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PlaybookExecutionList{})
	return err
}

// Patch applies the patch and returns the patched playbookExecution.
func (c *FakePlaybookExecutions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PlaybookExecution, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(playbookexecutionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PlaybookExecution{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PlaybookExecution), err
}
