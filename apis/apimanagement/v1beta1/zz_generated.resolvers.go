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
	v1beta11 "github.com/upbound/provider-azure/apis/keyvault/v1beta1"
	v1beta12 "github.com/upbound/provider-azure/apis/network/v1beta1"
	rconfig "github.com/upbound/provider-azure/apis/rconfig"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this API.
func (mg *API) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIManagementName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.APIManagementNameRef,
		Selector:     mg.Spec.ForProvider.APIManagementNameSelector,
		To: reference.To{
			List:    &ManagementList{},
			Managed: &Management{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIManagementName")
	}
	mg.Spec.ForProvider.APIManagementName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIManagementNameRef = rsp.ResolvedReference

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

// ResolveReferences of this APIOperation.
func (mg *APIOperation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.APINameRef,
		Selector:     mg.Spec.ForProvider.APINameSelector,
		To: reference.To{
			List:    &APIList{},
			Managed: &API{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIName")
	}
	mg.Spec.ForProvider.APIName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APINameRef = rsp.ResolvedReference

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

// ResolveReferences of this APIOperationPolicy.
func (mg *APIOperationPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIManagementName),
		Extract:      resource.ExtractParamPath("api_management_name", false),
		Reference:    mg.Spec.ForProvider.APIManagementNameRef,
		Selector:     mg.Spec.ForProvider.APIManagementNameSelector,
		To: reference.To{
			List:    &APIOperationList{},
			Managed: &APIOperation{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIManagementName")
	}
	mg.Spec.ForProvider.APIManagementName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIManagementNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIName),
		Extract:      resource.ExtractParamPath("api_name", false),
		Reference:    mg.Spec.ForProvider.APINameRef,
		Selector:     mg.Spec.ForProvider.APINameSelector,
		To: reference.To{
			List:    &APIOperationList{},
			Managed: &APIOperation{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIName")
	}
	mg.Spec.ForProvider.APIName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APINameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OperationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.OperationIDRef,
		Selector:     mg.Spec.ForProvider.OperationIDSelector,
		To: reference.To{
			List:    &APIOperationList{},
			Managed: &APIOperation{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.OperationID")
	}
	mg.Spec.ForProvider.OperationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OperationIDRef = rsp.ResolvedReference

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

// ResolveReferences of this APIOperationTag.
func (mg *APIOperationTag) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIOperationID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.APIOperationIDRef,
		Selector:     mg.Spec.ForProvider.APIOperationIDSelector,
		To: reference.To{
			List:    &APIOperationList{},
			Managed: &APIOperation{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIOperationID")
	}
	mg.Spec.ForProvider.APIOperationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIOperationIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this APIPolicy.
func (mg *APIPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.APINameRef,
		Selector:     mg.Spec.ForProvider.APINameSelector,
		To: reference.To{
			List:    &APIList{},
			Managed: &API{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIName")
	}
	mg.Spec.ForProvider.APIName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APINameRef = rsp.ResolvedReference

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

// ResolveReferences of this APIRelease.
func (mg *APIRelease) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.APIIDRef,
		Selector:     mg.Spec.ForProvider.APIIDSelector,
		To: reference.To{
			List:    &APIList{},
			Managed: &API{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIID")
	}
	mg.Spec.ForProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this APISchema.
func (mg *APISchema) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.APINameRef,
		Selector:     mg.Spec.ForProvider.APINameSelector,
		To: reference.To{
			List:    &APIList{},
			Managed: &API{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIName")
	}
	mg.Spec.ForProvider.APIName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APINameRef = rsp.ResolvedReference

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

// ResolveReferences of this APIVersionSet.
func (mg *APIVersionSet) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIManagementName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.APIManagementNameRef,
		Selector:     mg.Spec.ForProvider.APIManagementNameSelector,
		To: reference.To{
			List:    &ManagementList{},
			Managed: &Management{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIManagementName")
	}
	mg.Spec.ForProvider.APIManagementName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIManagementNameRef = rsp.ResolvedReference

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

// ResolveReferences of this AuthorizationServer.
func (mg *AuthorizationServer) ResolveReferences(ctx context.Context, c client.Reader) error {
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

// ResolveReferences of this Backend.
func (mg *Backend) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIManagementName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.APIManagementNameRef,
		Selector:     mg.Spec.ForProvider.APIManagementNameSelector,
		To: reference.To{
			List:    &ManagementList{},
			Managed: &Management{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIManagementName")
	}
	mg.Spec.ForProvider.APIManagementName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIManagementNameRef = rsp.ResolvedReference

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

// ResolveReferences of this Certificate.
func (mg *Certificate) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIManagementName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.APIManagementNameRef,
		Selector:     mg.Spec.ForProvider.APIManagementNameSelector,
		To: reference.To{
			List:    &ManagementList{},
			Managed: &Management{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIManagementName")
	}
	mg.Spec.ForProvider.APIManagementName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIManagementNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KeyVaultSecretID),
		Extract:      resource.ExtractParamPath("secret_id", true),
		Reference:    mg.Spec.ForProvider.KeyVaultSecretIDRef,
		Selector:     mg.Spec.ForProvider.KeyVaultSecretIDSelector,
		To: reference.To{
			List:    &v1beta11.CertificateList{},
			Managed: &v1beta11.Certificate{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KeyVaultSecretID")
	}
	mg.Spec.ForProvider.KeyVaultSecretID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KeyVaultSecretIDRef = rsp.ResolvedReference

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

// ResolveReferences of this EmailTemplate.
func (mg *EmailTemplate) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIManagementName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.APIManagementNameRef,
		Selector:     mg.Spec.ForProvider.APIManagementNameSelector,
		To: reference.To{
			List:    &ManagementList{},
			Managed: &Management{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIManagementName")
	}
	mg.Spec.ForProvider.APIManagementName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIManagementNameRef = rsp.ResolvedReference

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

// ResolveReferences of this Gateway.
func (mg *Gateway) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIManagementID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.APIManagementIDRef,
		Selector:     mg.Spec.ForProvider.APIManagementIDSelector,
		To: reference.To{
			List:    &ManagementList{},
			Managed: &Management{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIManagementID")
	}
	mg.Spec.ForProvider.APIManagementID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIManagementIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Management.
func (mg *Management) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.AdditionalLocation); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.AdditionalLocation[i3].VirtualNetworkConfiguration); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AdditionalLocation[i3].VirtualNetworkConfiguration[i4].SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.AdditionalLocation[i3].VirtualNetworkConfiguration[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.AdditionalLocation[i3].VirtualNetworkConfiguration[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1beta12.SubnetList{},
					Managed: &v1beta12.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.AdditionalLocation[i3].VirtualNetworkConfiguration[i4].SubnetID")
			}
			mg.Spec.ForProvider.AdditionalLocation[i3].VirtualNetworkConfiguration[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.AdditionalLocation[i3].VirtualNetworkConfiguration[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}
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

	for i3 := 0; i3 < len(mg.Spec.ForProvider.VirtualNetworkConfiguration); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualNetworkConfiguration[i3].SubnetID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.VirtualNetworkConfiguration[i3].SubnetIDRef,
			Selector:     mg.Spec.ForProvider.VirtualNetworkConfiguration[i3].SubnetIDSelector,
			To: reference.To{
				List:    &v1beta12.SubnetList{},
				Managed: &v1beta12.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.VirtualNetworkConfiguration[i3].SubnetID")
		}
		mg.Spec.ForProvider.VirtualNetworkConfiguration[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.VirtualNetworkConfiguration[i3].SubnetIDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this NamedValue.
func (mg *NamedValue) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIManagementName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.APIManagementNameRef,
		Selector:     mg.Spec.ForProvider.APIManagementNameSelector,
		To: reference.To{
			List:    &ManagementList{},
			Managed: &Management{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIManagementName")
	}
	mg.Spec.ForProvider.APIManagementName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIManagementNameRef = rsp.ResolvedReference

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

// ResolveReferences of this NotificationRecipientEmail.
func (mg *NotificationRecipientEmail) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIManagementID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.APIManagementIDRef,
		Selector:     mg.Spec.ForProvider.APIManagementIDSelector,
		To: reference.To{
			List:    &ManagementList{},
			Managed: &Management{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIManagementID")
	}
	mg.Spec.ForProvider.APIManagementID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIManagementIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Policy.
func (mg *Policy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIManagementID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.APIManagementIDRef,
		Selector:     mg.Spec.ForProvider.APIManagementIDSelector,
		To: reference.To{
			List:    &ManagementList{},
			Managed: &Management{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIManagementID")
	}
	mg.Spec.ForProvider.APIManagementID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIManagementIDRef = rsp.ResolvedReference

	return nil
}
