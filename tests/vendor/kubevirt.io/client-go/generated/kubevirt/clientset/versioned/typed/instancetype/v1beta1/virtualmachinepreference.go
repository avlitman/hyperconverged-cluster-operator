/*
Copyright The KubeVirt Authors.

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

package v1beta1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1beta1 "kubevirt.io/api/instancetype/v1beta1"
	scheme "kubevirt.io/client-go/generated/kubevirt/clientset/versioned/scheme"
)

// VirtualMachinePreferencesGetter has a method to return a VirtualMachinePreferenceInterface.
// A group's client should implement this interface.
type VirtualMachinePreferencesGetter interface {
	VirtualMachinePreferences(namespace string) VirtualMachinePreferenceInterface
}

// VirtualMachinePreferenceInterface has methods to work with VirtualMachinePreference resources.
type VirtualMachinePreferenceInterface interface {
	Create(ctx context.Context, virtualMachinePreference *v1beta1.VirtualMachinePreference, opts v1.CreateOptions) (*v1beta1.VirtualMachinePreference, error)
	Update(ctx context.Context, virtualMachinePreference *v1beta1.VirtualMachinePreference, opts v1.UpdateOptions) (*v1beta1.VirtualMachinePreference, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.VirtualMachinePreference, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.VirtualMachinePreferenceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VirtualMachinePreference, err error)
	VirtualMachinePreferenceExpansion
}

// virtualMachinePreferences implements VirtualMachinePreferenceInterface
type virtualMachinePreferences struct {
	client rest.Interface
	ns     string
}

// newVirtualMachinePreferences returns a VirtualMachinePreferences
func newVirtualMachinePreferences(c *InstancetypeV1beta1Client, namespace string) *virtualMachinePreferences {
	return &virtualMachinePreferences{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the virtualMachinePreference, and returns the corresponding virtualMachinePreference object, and an error if there is any.
func (c *virtualMachinePreferences) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VirtualMachinePreference, err error) {
	result = &v1beta1.VirtualMachinePreference{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinepreferences").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VirtualMachinePreferences that match those selectors.
func (c *virtualMachinePreferences) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VirtualMachinePreferenceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.VirtualMachinePreferenceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinepreferences").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested virtualMachinePreferences.
func (c *virtualMachinePreferences) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinepreferences").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a virtualMachinePreference and creates it.  Returns the server's representation of the virtualMachinePreference, and an error, if there is any.
func (c *virtualMachinePreferences) Create(ctx context.Context, virtualMachinePreference *v1beta1.VirtualMachinePreference, opts v1.CreateOptions) (result *v1beta1.VirtualMachinePreference, err error) {
	result = &v1beta1.VirtualMachinePreference{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualmachinepreferences").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachinePreference).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a virtualMachinePreference and updates it. Returns the server's representation of the virtualMachinePreference, and an error, if there is any.
func (c *virtualMachinePreferences) Update(ctx context.Context, virtualMachinePreference *v1beta1.VirtualMachinePreference, opts v1.UpdateOptions) (result *v1beta1.VirtualMachinePreference, err error) {
	result = &v1beta1.VirtualMachinePreference{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachinepreferences").
		Name(virtualMachinePreference.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachinePreference).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the virtualMachinePreference and deletes it. Returns an error if one occurs.
func (c *virtualMachinePreferences) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachinepreferences").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *virtualMachinePreferences) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachinepreferences").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched virtualMachinePreference.
func (c *virtualMachinePreferences) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VirtualMachinePreference, err error) {
	result = &v1beta1.VirtualMachinePreference{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualmachinepreferences").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
