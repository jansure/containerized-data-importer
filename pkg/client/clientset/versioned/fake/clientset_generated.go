/*
Copyright 2018 The CDI Authors.

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
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
	clientset "kubevirt.io/containerized-data-importer/pkg/client/clientset/versioned"
	cdiv1alpha1 "kubevirt.io/containerized-data-importer/pkg/client/clientset/versioned/typed/core/v1alpha1"
	fakecdiv1alpha1 "kubevirt.io/containerized-data-importer/pkg/client/clientset/versioned/typed/core/v1alpha1/fake"
	cdiv1beta1 "kubevirt.io/containerized-data-importer/pkg/client/clientset/versioned/typed/core/v1beta1"
	fakecdiv1beta1 "kubevirt.io/containerized-data-importer/pkg/client/clientset/versioned/typed/core/v1beta1/fake"
	uploadv1alpha1 "kubevirt.io/containerized-data-importer/pkg/client/clientset/versioned/typed/upload/v1alpha1"
	fakeuploadv1alpha1 "kubevirt.io/containerized-data-importer/pkg/client/clientset/versioned/typed/upload/v1alpha1/fake"
	uploadv1beta1 "kubevirt.io/containerized-data-importer/pkg/client/clientset/versioned/typed/upload/v1beta1"
	fakeuploadv1beta1 "kubevirt.io/containerized-data-importer/pkg/client/clientset/versioned/typed/upload/v1beta1/fake"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var _ clientset.Interface = &Clientset{}

// CdiV1alpha1 retrieves the CdiV1alpha1Client
func (c *Clientset) CdiV1alpha1() cdiv1alpha1.CdiV1alpha1Interface {
	return &fakecdiv1alpha1.FakeCdiV1alpha1{Fake: &c.Fake}
}

// CdiV1beta1 retrieves the CdiV1beta1Client
func (c *Clientset) CdiV1beta1() cdiv1beta1.CdiV1beta1Interface {
	return &fakecdiv1beta1.FakeCdiV1beta1{Fake: &c.Fake}
}

// UploadV1alpha1 retrieves the UploadV1alpha1Client
func (c *Clientset) UploadV1alpha1() uploadv1alpha1.UploadV1alpha1Interface {
	return &fakeuploadv1alpha1.FakeUploadV1alpha1{Fake: &c.Fake}
}

// UploadV1beta1 retrieves the UploadV1beta1Client
func (c *Clientset) UploadV1beta1() uploadv1beta1.UploadV1beta1Interface {
	return &fakeuploadv1beta1.FakeUploadV1beta1{Fake: &c.Fake}
}