/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-azure/apis/azure/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Monitor.
func (mg *Monitor) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SubAccount.
func (mg *SubAccount) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LogzMonitorID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.LogzMonitorIDRef,
		Selector:     mg.Spec.ForProvider.LogzMonitorIDSelector,
		To: reference.To{
			List:    &MonitorList{},
			Managed: &Monitor{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LogzMonitorID")
	}
	mg.Spec.ForProvider.LogzMonitorID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LogzMonitorIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SubAccountTagRule.
func (mg *SubAccountTagRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LogzSubAccountID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.LogzSubAccountIDRef,
		Selector:     mg.Spec.ForProvider.LogzSubAccountIDSelector,
		To: reference.To{
			List:    &SubAccountList{},
			Managed: &SubAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LogzSubAccountID")
	}
	mg.Spec.ForProvider.LogzSubAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LogzSubAccountIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TagRule.
func (mg *TagRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LogzMonitorID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.LogzMonitorIDRef,
		Selector:     mg.Spec.ForProvider.LogzMonitorIDSelector,
		To: reference.To{
			List:    &MonitorList{},
			Managed: &Monitor{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LogzMonitorID")
	}
	mg.Spec.ForProvider.LogzMonitorID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LogzMonitorIDRef = rsp.ResolvedReference

	return nil
}