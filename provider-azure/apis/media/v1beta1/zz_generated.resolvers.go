/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1"
	rconfig "github.com/upbound/official-providers/provider-azure/apis/rconfig"
	v1beta11 "github.com/upbound/official-providers/provider-azure/apis/storage/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Asset.
func (mg *Asset) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MediaServicesAccountName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.MediaServicesAccountNameRef,
		Selector:     mg.Spec.ForProvider.MediaServicesAccountNameSelector,
		To: reference.To{
			List:    &ServicesAccountList{},
			Managed: &ServicesAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MediaServicesAccountName")
	}
	mg.Spec.ForProvider.MediaServicesAccountName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MediaServicesAccountNameRef = rsp.ResolvedReference

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

// ResolveReferences of this LiveEvent.
func (mg *LiveEvent) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MediaServicesAccountName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.MediaServicesAccountNameRef,
		Selector:     mg.Spec.ForProvider.MediaServicesAccountNameSelector,
		To: reference.To{
			List:    &ServicesAccountList{},
			Managed: &ServicesAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MediaServicesAccountName")
	}
	mg.Spec.ForProvider.MediaServicesAccountName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MediaServicesAccountNameRef = rsp.ResolvedReference

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

// ResolveReferences of this LiveEventOutput.
func (mg *LiveEventOutput) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AssetName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.AssetNameRef,
		Selector:     mg.Spec.ForProvider.AssetNameSelector,
		To: reference.To{
			List:    &AssetList{},
			Managed: &Asset{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AssetName")
	}
	mg.Spec.ForProvider.AssetName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AssetNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LiveEventID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.LiveEventIDRef,
		Selector:     mg.Spec.ForProvider.LiveEventIDSelector,
		To: reference.To{
			List:    &LiveEventList{},
			Managed: &LiveEvent{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LiveEventID")
	}
	mg.Spec.ForProvider.LiveEventID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LiveEventIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ServicesAccount.
func (mg *ServicesAccount) ResolveReferences(ctx context.Context, c client.Reader) error {
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

	for i3 := 0; i3 < len(mg.Spec.ForProvider.StorageAccount); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageAccount[i3].ID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.StorageAccount[i3].IDRef,
			Selector:     mg.Spec.ForProvider.StorageAccount[i3].IDSelector,
			To: reference.To{
				List:    &v1beta11.AccountList{},
				Managed: &v1beta11.Account{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.StorageAccount[i3].ID")
		}
		mg.Spec.ForProvider.StorageAccount[i3].ID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.StorageAccount[i3].IDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this StreamingEndpoint.
func (mg *StreamingEndpoint) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MediaServicesAccountName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.MediaServicesAccountNameRef,
		Selector:     mg.Spec.ForProvider.MediaServicesAccountNameSelector,
		To: reference.To{
			List:    &ServicesAccountList{},
			Managed: &ServicesAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MediaServicesAccountName")
	}
	mg.Spec.ForProvider.MediaServicesAccountName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MediaServicesAccountNameRef = rsp.ResolvedReference

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

// ResolveReferences of this StreamingLocator.
func (mg *StreamingLocator) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AssetName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.AssetNameRef,
		Selector:     mg.Spec.ForProvider.AssetNameSelector,
		To: reference.To{
			List:    &AssetList{},
			Managed: &Asset{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AssetName")
	}
	mg.Spec.ForProvider.AssetName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AssetNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MediaServicesAccountName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.MediaServicesAccountNameRef,
		Selector:     mg.Spec.ForProvider.MediaServicesAccountNameSelector,
		To: reference.To{
			List:    &ServicesAccountList{},
			Managed: &ServicesAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MediaServicesAccountName")
	}
	mg.Spec.ForProvider.MediaServicesAccountName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MediaServicesAccountNameRef = rsp.ResolvedReference

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

// ResolveReferences of this StreamingPolicy.
func (mg *StreamingPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MediaServicesAccountName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.MediaServicesAccountNameRef,
		Selector:     mg.Spec.ForProvider.MediaServicesAccountNameSelector,
		To: reference.To{
			List:    &ServicesAccountList{},
			Managed: &ServicesAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MediaServicesAccountName")
	}
	mg.Spec.ForProvider.MediaServicesAccountName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MediaServicesAccountNameRef = rsp.ResolvedReference

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

// ResolveReferences of this Transform.
func (mg *Transform) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MediaServicesAccountName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.MediaServicesAccountNameRef,
		Selector:     mg.Spec.ForProvider.MediaServicesAccountNameSelector,
		To: reference.To{
			List:    &ServicesAccountList{},
			Managed: &ServicesAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MediaServicesAccountName")
	}
	mg.Spec.ForProvider.MediaServicesAccountName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MediaServicesAccountNameRef = rsp.ResolvedReference

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
