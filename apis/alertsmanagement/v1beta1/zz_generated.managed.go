/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this MonitorActionRuleActionGroup.
func (mg *MonitorActionRuleActionGroup) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MonitorActionRuleActionGroup.
func (mg *MonitorActionRuleActionGroup) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this MonitorActionRuleActionGroup.
func (mg *MonitorActionRuleActionGroup) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this MonitorActionRuleActionGroup.
func (mg *MonitorActionRuleActionGroup) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MonitorActionRuleActionGroup.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MonitorActionRuleActionGroup) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this MonitorActionRuleActionGroup.
func (mg *MonitorActionRuleActionGroup) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MonitorActionRuleActionGroup.
func (mg *MonitorActionRuleActionGroup) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MonitorActionRuleActionGroup.
func (mg *MonitorActionRuleActionGroup) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MonitorActionRuleActionGroup.
func (mg *MonitorActionRuleActionGroup) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this MonitorActionRuleActionGroup.
func (mg *MonitorActionRuleActionGroup) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this MonitorActionRuleActionGroup.
func (mg *MonitorActionRuleActionGroup) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MonitorActionRuleActionGroup.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MonitorActionRuleActionGroup) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this MonitorActionRuleActionGroup.
func (mg *MonitorActionRuleActionGroup) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MonitorActionRuleActionGroup.
func (mg *MonitorActionRuleActionGroup) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MonitorActionRuleSuppression.
func (mg *MonitorActionRuleSuppression) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MonitorActionRuleSuppression.
func (mg *MonitorActionRuleSuppression) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this MonitorActionRuleSuppression.
func (mg *MonitorActionRuleSuppression) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this MonitorActionRuleSuppression.
func (mg *MonitorActionRuleSuppression) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MonitorActionRuleSuppression.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MonitorActionRuleSuppression) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this MonitorActionRuleSuppression.
func (mg *MonitorActionRuleSuppression) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MonitorActionRuleSuppression.
func (mg *MonitorActionRuleSuppression) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MonitorActionRuleSuppression.
func (mg *MonitorActionRuleSuppression) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MonitorActionRuleSuppression.
func (mg *MonitorActionRuleSuppression) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this MonitorActionRuleSuppression.
func (mg *MonitorActionRuleSuppression) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this MonitorActionRuleSuppression.
func (mg *MonitorActionRuleSuppression) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MonitorActionRuleSuppression.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MonitorActionRuleSuppression) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this MonitorActionRuleSuppression.
func (mg *MonitorActionRuleSuppression) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MonitorActionRuleSuppression.
func (mg *MonitorActionRuleSuppression) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MonitorAlertProcessingRuleActionGroup.
func (mg *MonitorAlertProcessingRuleActionGroup) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MonitorAlertProcessingRuleActionGroup.
func (mg *MonitorAlertProcessingRuleActionGroup) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this MonitorAlertProcessingRuleActionGroup.
func (mg *MonitorAlertProcessingRuleActionGroup) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this MonitorAlertProcessingRuleActionGroup.
func (mg *MonitorAlertProcessingRuleActionGroup) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MonitorAlertProcessingRuleActionGroup.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MonitorAlertProcessingRuleActionGroup) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this MonitorAlertProcessingRuleActionGroup.
func (mg *MonitorAlertProcessingRuleActionGroup) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MonitorAlertProcessingRuleActionGroup.
func (mg *MonitorAlertProcessingRuleActionGroup) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MonitorAlertProcessingRuleActionGroup.
func (mg *MonitorAlertProcessingRuleActionGroup) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MonitorAlertProcessingRuleActionGroup.
func (mg *MonitorAlertProcessingRuleActionGroup) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this MonitorAlertProcessingRuleActionGroup.
func (mg *MonitorAlertProcessingRuleActionGroup) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this MonitorAlertProcessingRuleActionGroup.
func (mg *MonitorAlertProcessingRuleActionGroup) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MonitorAlertProcessingRuleActionGroup.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MonitorAlertProcessingRuleActionGroup) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this MonitorAlertProcessingRuleActionGroup.
func (mg *MonitorAlertProcessingRuleActionGroup) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MonitorAlertProcessingRuleActionGroup.
func (mg *MonitorAlertProcessingRuleActionGroup) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MonitorAlertProcessingRuleSuppression.
func (mg *MonitorAlertProcessingRuleSuppression) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MonitorAlertProcessingRuleSuppression.
func (mg *MonitorAlertProcessingRuleSuppression) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this MonitorAlertProcessingRuleSuppression.
func (mg *MonitorAlertProcessingRuleSuppression) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this MonitorAlertProcessingRuleSuppression.
func (mg *MonitorAlertProcessingRuleSuppression) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MonitorAlertProcessingRuleSuppression.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MonitorAlertProcessingRuleSuppression) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this MonitorAlertProcessingRuleSuppression.
func (mg *MonitorAlertProcessingRuleSuppression) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MonitorAlertProcessingRuleSuppression.
func (mg *MonitorAlertProcessingRuleSuppression) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MonitorAlertProcessingRuleSuppression.
func (mg *MonitorAlertProcessingRuleSuppression) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MonitorAlertProcessingRuleSuppression.
func (mg *MonitorAlertProcessingRuleSuppression) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this MonitorAlertProcessingRuleSuppression.
func (mg *MonitorAlertProcessingRuleSuppression) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this MonitorAlertProcessingRuleSuppression.
func (mg *MonitorAlertProcessingRuleSuppression) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MonitorAlertProcessingRuleSuppression.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MonitorAlertProcessingRuleSuppression) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this MonitorAlertProcessingRuleSuppression.
func (mg *MonitorAlertProcessingRuleSuppression) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MonitorAlertProcessingRuleSuppression.
func (mg *MonitorAlertProcessingRuleSuppression) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MonitorSmartDetectorAlertRule.
func (mg *MonitorSmartDetectorAlertRule) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MonitorSmartDetectorAlertRule.
func (mg *MonitorSmartDetectorAlertRule) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this MonitorSmartDetectorAlertRule.
func (mg *MonitorSmartDetectorAlertRule) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this MonitorSmartDetectorAlertRule.
func (mg *MonitorSmartDetectorAlertRule) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MonitorSmartDetectorAlertRule.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MonitorSmartDetectorAlertRule) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this MonitorSmartDetectorAlertRule.
func (mg *MonitorSmartDetectorAlertRule) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MonitorSmartDetectorAlertRule.
func (mg *MonitorSmartDetectorAlertRule) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MonitorSmartDetectorAlertRule.
func (mg *MonitorSmartDetectorAlertRule) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MonitorSmartDetectorAlertRule.
func (mg *MonitorSmartDetectorAlertRule) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this MonitorSmartDetectorAlertRule.
func (mg *MonitorSmartDetectorAlertRule) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this MonitorSmartDetectorAlertRule.
func (mg *MonitorSmartDetectorAlertRule) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MonitorSmartDetectorAlertRule.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MonitorSmartDetectorAlertRule) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this MonitorSmartDetectorAlertRule.
func (mg *MonitorSmartDetectorAlertRule) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MonitorSmartDetectorAlertRule.
func (mg *MonitorSmartDetectorAlertRule) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
