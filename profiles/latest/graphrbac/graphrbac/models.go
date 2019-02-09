// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package graphrbac

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ObjectType = original.ObjectType

const (
	ObjectTypeApplication      ObjectType = original.ObjectTypeApplication
	ObjectTypeDirectoryObject  ObjectType = original.ObjectTypeDirectoryObject
	ObjectTypeGroup            ObjectType = original.ObjectTypeGroup
	ObjectTypeServicePrincipal ObjectType = original.ObjectTypeServicePrincipal
	ObjectTypeUser             ObjectType = original.ObjectTypeUser
)

type UserType = original.UserType

const (
	Guest  UserType = original.Guest
	Member UserType = original.Member
)

type ADGroup = original.ADGroup
type AddOwnerParameters = original.AddOwnerParameters
type AppRole = original.AppRole
type Application = original.Application
type ApplicationCreateParameters = original.ApplicationCreateParameters
type ApplicationListResult = original.ApplicationListResult
type ApplicationListResultIterator = original.ApplicationListResultIterator
type ApplicationListResultPage = original.ApplicationListResultPage
type ApplicationUpdateParameters = original.ApplicationUpdateParameters
type ApplicationsClient = original.ApplicationsClient
type BaseClient = original.BaseClient
type BasicDirectoryObject = original.BasicDirectoryObject
type CheckGroupMembershipParameters = original.CheckGroupMembershipParameters
type CheckGroupMembershipResult = original.CheckGroupMembershipResult
type DeletedApplicationsClient = original.DeletedApplicationsClient
type DirectoryObject = original.DirectoryObject
type DirectoryObjectListResult = original.DirectoryObjectListResult
type DirectoryObjectListResultIterator = original.DirectoryObjectListResultIterator
type DirectoryObjectListResultPage = original.DirectoryObjectListResultPage
type Domain = original.Domain
type DomainListResult = original.DomainListResult
type DomainsClient = original.DomainsClient
type ErrorMessage = original.ErrorMessage
type GetObjectsParameters = original.GetObjectsParameters
type GraphError = original.GraphError
type GroupAddMemberParameters = original.GroupAddMemberParameters
type GroupCreateParameters = original.GroupCreateParameters
type GroupGetMemberGroupsParameters = original.GroupGetMemberGroupsParameters
type GroupGetMemberGroupsResult = original.GroupGetMemberGroupsResult
type GroupListResult = original.GroupListResult
type GroupListResultIterator = original.GroupListResultIterator
type GroupListResultPage = original.GroupListResultPage
type GroupsClient = original.GroupsClient
type KeyCredential = original.KeyCredential
type KeyCredentialListResult = original.KeyCredentialListResult
type KeyCredentialsUpdateParameters = original.KeyCredentialsUpdateParameters
type OAuth2Client = original.OAuth2Client
type ObjectsClient = original.ObjectsClient
type OdataError = original.OdataError
type PasswordCredential = original.PasswordCredential
type PasswordCredentialListResult = original.PasswordCredentialListResult
type PasswordCredentialsUpdateParameters = original.PasswordCredentialsUpdateParameters
type PasswordProfile = original.PasswordProfile
type Permissions = original.Permissions
type PermissionsListResult = original.PermissionsListResult
type RequiredResourceAccess = original.RequiredResourceAccess
type ResourceAccess = original.ResourceAccess
type ServicePrincipal = original.ServicePrincipal
type ServicePrincipalCreateParameters = original.ServicePrincipalCreateParameters
type ServicePrincipalListResult = original.ServicePrincipalListResult
type ServicePrincipalListResultIterator = original.ServicePrincipalListResultIterator
type ServicePrincipalListResultPage = original.ServicePrincipalListResultPage
type ServicePrincipalUpdateParameters = original.ServicePrincipalUpdateParameters
type ServicePrincipalsClient = original.ServicePrincipalsClient
type SignInName = original.SignInName
type SignedInUserClient = original.SignedInUserClient
type User = original.User
type UserBase = original.UserBase
type UserCreateParameters = original.UserCreateParameters
type UserGetMemberGroupsParameters = original.UserGetMemberGroupsParameters
type UserGetMemberGroupsResult = original.UserGetMemberGroupsResult
type UserListResult = original.UserListResult
type UserListResultIterator = original.UserListResultIterator
type UserListResultPage = original.UserListResultPage
type UserUpdateParameters = original.UserUpdateParameters
type UsersClient = original.UsersClient

func New(tenantID string) BaseClient {
	return original.New(tenantID)
}
func NewApplicationListResultIterator(page ApplicationListResultPage) ApplicationListResultIterator {
	return original.NewApplicationListResultIterator(page)
}
func NewApplicationListResultPage(getNextPage func(context.Context, ApplicationListResult) (ApplicationListResult, error)) ApplicationListResultPage {
	return original.NewApplicationListResultPage(getNextPage)
}
func NewApplicationsClient(tenantID string) ApplicationsClient {
	return original.NewApplicationsClient(tenantID)
}
func NewApplicationsClientWithBaseURI(baseURI string, tenantID string) ApplicationsClient {
	return original.NewApplicationsClientWithBaseURI(baseURI, tenantID)
}
func NewDeletedApplicationsClient(tenantID string) DeletedApplicationsClient {
	return original.NewDeletedApplicationsClient(tenantID)
}
func NewDeletedApplicationsClientWithBaseURI(baseURI string, tenantID string) DeletedApplicationsClient {
	return original.NewDeletedApplicationsClientWithBaseURI(baseURI, tenantID)
}
func NewDirectoryObjectListResultIterator(page DirectoryObjectListResultPage) DirectoryObjectListResultIterator {
	return original.NewDirectoryObjectListResultIterator(page)
}
func NewDirectoryObjectListResultPage(getNextPage func(context.Context, DirectoryObjectListResult) (DirectoryObjectListResult, error)) DirectoryObjectListResultPage {
	return original.NewDirectoryObjectListResultPage(getNextPage)
}
func NewDomainsClient(tenantID string) DomainsClient {
	return original.NewDomainsClient(tenantID)
}
func NewDomainsClientWithBaseURI(baseURI string, tenantID string) DomainsClient {
	return original.NewDomainsClientWithBaseURI(baseURI, tenantID)
}
func NewGroupListResultIterator(page GroupListResultPage) GroupListResultIterator {
	return original.NewGroupListResultIterator(page)
}
func NewGroupListResultPage(getNextPage func(context.Context, GroupListResult) (GroupListResult, error)) GroupListResultPage {
	return original.NewGroupListResultPage(getNextPage)
}
func NewGroupsClient(tenantID string) GroupsClient {
	return original.NewGroupsClient(tenantID)
}
func NewGroupsClientWithBaseURI(baseURI string, tenantID string) GroupsClient {
	return original.NewGroupsClientWithBaseURI(baseURI, tenantID)
}
func NewOAuth2Client(tenantID string) OAuth2Client {
	return original.NewOAuth2Client(tenantID)
}
func NewOAuth2ClientWithBaseURI(baseURI string, tenantID string) OAuth2Client {
	return original.NewOAuth2ClientWithBaseURI(baseURI, tenantID)
}
func NewObjectsClient(tenantID string) ObjectsClient {
	return original.NewObjectsClient(tenantID)
}
func NewObjectsClientWithBaseURI(baseURI string, tenantID string) ObjectsClient {
	return original.NewObjectsClientWithBaseURI(baseURI, tenantID)
}
func NewServicePrincipalListResultIterator(page ServicePrincipalListResultPage) ServicePrincipalListResultIterator {
	return original.NewServicePrincipalListResultIterator(page)
}
func NewServicePrincipalListResultPage(getNextPage func(context.Context, ServicePrincipalListResult) (ServicePrincipalListResult, error)) ServicePrincipalListResultPage {
	return original.NewServicePrincipalListResultPage(getNextPage)
}
func NewServicePrincipalsClient(tenantID string) ServicePrincipalsClient {
	return original.NewServicePrincipalsClient(tenantID)
}
func NewServicePrincipalsClientWithBaseURI(baseURI string, tenantID string) ServicePrincipalsClient {
	return original.NewServicePrincipalsClientWithBaseURI(baseURI, tenantID)
}
func NewSignedInUserClient(tenantID string) SignedInUserClient {
	return original.NewSignedInUserClient(tenantID)
}
func NewSignedInUserClientWithBaseURI(baseURI string, tenantID string) SignedInUserClient {
	return original.NewSignedInUserClientWithBaseURI(baseURI, tenantID)
}
func NewUserListResultIterator(page UserListResultPage) UserListResultIterator {
	return original.NewUserListResultIterator(page)
}
func NewUserListResultPage(getNextPage func(context.Context, UserListResult) (UserListResult, error)) UserListResultPage {
	return original.NewUserListResultPage(getNextPage)
}
func NewUsersClient(tenantID string) UsersClient {
	return original.NewUsersClient(tenantID)
}
func NewUsersClientWithBaseURI(baseURI string, tenantID string) UsersClient {
	return original.NewUsersClientWithBaseURI(baseURI, tenantID)
}
func NewWithBaseURI(baseURI string, tenantID string) BaseClient {
	return original.NewWithBaseURI(baseURI, tenantID)
}
func PossibleObjectTypeValues() []ObjectType {
	return original.PossibleObjectTypeValues()
}
func PossibleUserTypeValues() []UserType {
	return original.PossibleUserTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
