// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package otterizegraphql

import (
	"context"
	"encoding/json"

	"github.com/Khan/genqlient/graphql"
)

type CertificateCustomization struct {
	DnsNames []*string `json:"dnsNames"`
	Ttl      *int      `json:"ttl"`
}

// GetDnsNames returns CertificateCustomization.DnsNames, and is useful for accessing the field via an interface.
func (v *CertificateCustomization) GetDnsNames() []*string { return v.DnsNames }

// GetTtl returns CertificateCustomization.Ttl, and is useful for accessing the field via an interface.
func (v *CertificateCustomization) GetTtl() *int { return v.Ttl }

type ComponentType string

const (
	ComponentTypeIntentsOperator     ComponentType = "INTENTS_OPERATOR"
	ComponentTypeCredentialsOperator ComponentType = "CREDENTIALS_OPERATOR"
	ComponentTypeNetworkMapper       ComponentType = "NETWORK_MAPPER"
)

// GetTLSKeyPairResponse is returned by GetTLSKeyPair on success.
type GetTLSKeyPairResponse struct {
	// Get service
	Service *GetTLSKeyPairService `json:"service"`
}

// GetService returns GetTLSKeyPairResponse.Service, and is useful for accessing the field via an interface.
func (v *GetTLSKeyPairResponse) GetService() *GetTLSKeyPairService { return v.Service }

// GetTLSKeyPairService includes the requested fields of the GraphQL type Service.
type GetTLSKeyPairService struct {
	TlsKeyPair *GetTLSKeyPairServiceTlsKeyPair `json:"tlsKeyPair"`
}

// GetTlsKeyPair returns GetTLSKeyPairService.TlsKeyPair, and is useful for accessing the field via an interface.
func (v *GetTLSKeyPairService) GetTlsKeyPair() *GetTLSKeyPairServiceTlsKeyPair { return v.TlsKeyPair }

// GetTLSKeyPairServiceTlsKeyPair includes the requested fields of the GraphQL type KeyPair.
type GetTLSKeyPairServiceTlsKeyPair struct {
	TLSKeyPair `json:"-"`
}

// GetKeyPEM returns GetTLSKeyPairServiceTlsKeyPair.KeyPEM, and is useful for accessing the field via an interface.
func (v *GetTLSKeyPairServiceTlsKeyPair) GetKeyPEM() string { return v.TLSKeyPair.KeyPEM }

// GetCertPEM returns GetTLSKeyPairServiceTlsKeyPair.CertPEM, and is useful for accessing the field via an interface.
func (v *GetTLSKeyPairServiceTlsKeyPair) GetCertPEM() string { return v.TLSKeyPair.CertPEM }

// GetCaPEM returns GetTLSKeyPairServiceTlsKeyPair.CaPEM, and is useful for accessing the field via an interface.
func (v *GetTLSKeyPairServiceTlsKeyPair) GetCaPEM() string { return v.TLSKeyPair.CaPEM }

// GetRootCAPEM returns GetTLSKeyPairServiceTlsKeyPair.RootCAPEM, and is useful for accessing the field via an interface.
func (v *GetTLSKeyPairServiceTlsKeyPair) GetRootCAPEM() string { return v.TLSKeyPair.RootCAPEM }

// GetExpiresAt returns GetTLSKeyPairServiceTlsKeyPair.ExpiresAt, and is useful for accessing the field via an interface.
func (v *GetTLSKeyPairServiceTlsKeyPair) GetExpiresAt() int { return v.TLSKeyPair.ExpiresAt }

func (v *GetTLSKeyPairServiceTlsKeyPair) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*GetTLSKeyPairServiceTlsKeyPair
		graphql.NoUnmarshalJSON
	}
	firstPass.GetTLSKeyPairServiceTlsKeyPair = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	err = json.Unmarshal(
		b, &v.TLSKeyPair)
	if err != nil {
		return err
	}
	return nil
}

type __premarshalGetTLSKeyPairServiceTlsKeyPair struct {
	KeyPEM string `json:"keyPEM"`

	CertPEM string `json:"certPEM"`

	CaPEM string `json:"caPEM"`

	RootCAPEM string `json:"rootCAPEM"`

	ExpiresAt int `json:"expiresAt"`
}

func (v *GetTLSKeyPairServiceTlsKeyPair) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *GetTLSKeyPairServiceTlsKeyPair) __premarshalJSON() (*__premarshalGetTLSKeyPairServiceTlsKeyPair, error) {
	var retval __premarshalGetTLSKeyPairServiceTlsKeyPair

	retval.KeyPEM = v.TLSKeyPair.KeyPEM
	retval.CertPEM = v.TLSKeyPair.CertPEM
	retval.CaPEM = v.TLSKeyPair.CaPEM
	retval.RootCAPEM = v.TLSKeyPair.RootCAPEM
	retval.ExpiresAt = v.TLSKeyPair.ExpiresAt
	return &retval, nil
}

// ReportComponentStatusResponse is returned by ReportComponentStatus on success.
type ReportComponentStatusResponse struct {
	ReportComponentStatus string `json:"reportComponentStatus"`
}

// GetReportComponentStatus returns ReportComponentStatusResponse.ReportComponentStatus, and is useful for accessing the field via an interface.
func (v *ReportComponentStatusResponse) GetReportComponentStatus() string {
	return v.ReportComponentStatus
}

// ReportKubernetesWorkloadReportKubernetesWorkloadService includes the requested fields of the GraphQL type Service.
type ReportKubernetesWorkloadReportKubernetesWorkloadService struct {
	Id string `json:"id"`
}

// GetId returns ReportKubernetesWorkloadReportKubernetesWorkloadService.Id, and is useful for accessing the field via an interface.
func (v *ReportKubernetesWorkloadReportKubernetesWorkloadService) GetId() string { return v.Id }

// ReportKubernetesWorkloadResponse is returned by ReportKubernetesWorkload on success.
type ReportKubernetesWorkloadResponse struct {
	ReportKubernetesWorkload ReportKubernetesWorkloadReportKubernetesWorkloadService `json:"reportKubernetesWorkload"`
}

// GetReportKubernetesWorkload returns ReportKubernetesWorkloadResponse.ReportKubernetesWorkload, and is useful for accessing the field via an interface.
func (v *ReportKubernetesWorkloadResponse) GetReportKubernetesWorkload() ReportKubernetesWorkloadReportKubernetesWorkloadService {
	return v.ReportKubernetesWorkload
}

// TLSKeyPair includes the GraphQL fields of KeyPair requested by the fragment TLSKeyPair.
type TLSKeyPair struct {
	KeyPEM    string `json:"keyPEM"`
	CertPEM   string `json:"certPEM"`
	CaPEM     string `json:"caPEM"`
	RootCAPEM string `json:"rootCAPEM"`
	ExpiresAt int    `json:"expiresAt"`
}

// GetKeyPEM returns TLSKeyPair.KeyPEM, and is useful for accessing the field via an interface.
func (v *TLSKeyPair) GetKeyPEM() string { return v.KeyPEM }

// GetCertPEM returns TLSKeyPair.CertPEM, and is useful for accessing the field via an interface.
func (v *TLSKeyPair) GetCertPEM() string { return v.CertPEM }

// GetCaPEM returns TLSKeyPair.CaPEM, and is useful for accessing the field via an interface.
func (v *TLSKeyPair) GetCaPEM() string { return v.CaPEM }

// GetRootCAPEM returns TLSKeyPair.RootCAPEM, and is useful for accessing the field via an interface.
func (v *TLSKeyPair) GetRootCAPEM() string { return v.RootCAPEM }

// GetExpiresAt returns TLSKeyPair.ExpiresAt, and is useful for accessing the field via an interface.
func (v *TLSKeyPair) GetExpiresAt() int { return v.ExpiresAt }

// __GetTLSKeyPairInput is used internally by genqlient
type __GetTLSKeyPairInput struct {
	Id                        *string                   `json:"id"`
	CertificateCustomizations *CertificateCustomization `json:"certificateCustomizations"`
}

// GetId returns __GetTLSKeyPairInput.Id, and is useful for accessing the field via an interface.
func (v *__GetTLSKeyPairInput) GetId() *string { return v.Id }

// GetCertificateCustomizations returns __GetTLSKeyPairInput.CertificateCustomizations, and is useful for accessing the field via an interface.
func (v *__GetTLSKeyPairInput) GetCertificateCustomizations() *CertificateCustomization {
	return v.CertificateCustomizations
}

// __ReportComponentStatusInput is used internally by genqlient
type __ReportComponentStatusInput struct {
	Component ComponentType `json:"component"`
}

// GetComponent returns __ReportComponentStatusInput.Component, and is useful for accessing the field via an interface.
func (v *__ReportComponentStatusInput) GetComponent() ComponentType { return v.Component }

// __ReportKubernetesWorkloadInput is used internally by genqlient
type __ReportKubernetesWorkloadInput struct {
	Namespace    string `json:"namespace"`
	PodOwnerName string `json:"podOwnerName"`
}

// GetNamespace returns __ReportKubernetesWorkloadInput.Namespace, and is useful for accessing the field via an interface.
func (v *__ReportKubernetesWorkloadInput) GetNamespace() string { return v.Namespace }

// GetPodOwnerName returns __ReportKubernetesWorkloadInput.PodOwnerName, and is useful for accessing the field via an interface.
func (v *__ReportKubernetesWorkloadInput) GetPodOwnerName() string { return v.PodOwnerName }

func GetTLSKeyPair(
	ctx context.Context,
	client graphql.Client,
	id *string,
	certificateCustomizations *CertificateCustomization,
) (*GetTLSKeyPairResponse, error) {
	req := &graphql.Request{
		OpName: "GetTLSKeyPair",
		Query: `
query GetTLSKeyPair ($id: ID!, $certificateCustomizations: CertificateCustomization) {
	service(id: $id) {
		tlsKeyPair(certificateCustomization: $certificateCustomizations) {
			... TLSKeyPair
		}
	}
}
fragment TLSKeyPair on KeyPair {
	keyPEM
	certPEM
	caPEM
	rootCAPEM
	expiresAt
}
`,
		Variables: &__GetTLSKeyPairInput{
			Id:                        id,
			CertificateCustomizations: certificateCustomizations,
		},
	}
	var err error

	var data GetTLSKeyPairResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func ReportComponentStatus(
	ctx context.Context,
	client graphql.Client,
	component ComponentType,
) (*ReportComponentStatusResponse, error) {
	req := &graphql.Request{
		OpName: "ReportComponentStatus",
		Query: `
mutation ReportComponentStatus ($component: ComponentType!) {
	reportComponentStatus(component: $component)
}
`,
		Variables: &__ReportComponentStatusInput{
			Component: component,
		},
	}
	var err error

	var data ReportComponentStatusResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func ReportKubernetesWorkload(
	ctx context.Context,
	client graphql.Client,
	namespace string,
	podOwnerName string,
) (*ReportKubernetesWorkloadResponse, error) {
	req := &graphql.Request{
		OpName: "ReportKubernetesWorkload",
		Query: `
mutation ReportKubernetesWorkload ($namespace: String!, $podOwnerName: String!) {
	reportKubernetesWorkload(namespace: $namespace, podOwnerName: $podOwnerName) {
		id
	}
}
`,
		Variables: &__ReportKubernetesWorkloadInput{
			Namespace:    namespace,
			PodOwnerName: podOwnerName,
		},
	}
	var err error

	var data ReportKubernetesWorkloadResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
