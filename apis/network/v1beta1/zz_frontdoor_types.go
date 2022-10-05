/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BackendObservation struct {
}

type BackendParameters struct {

	// Location of the backend (IP address or FQDN)
	// +kubebuilder:validation:Required
	Address *string `json:"address" tf:"address,omitempty"`

	// Enable or Disable use of this Backend Routing Rule. Permitted values are true or false. Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The HTTP TCP port number. Possible values are between 1 - 65535.
	// +kubebuilder:validation:Required
	HTTPPort *float64 `json:"httpPort" tf:"http_port,omitempty"`

	// The HTTPS TCP port number. Possible values are between 1 - 65535.
	// +kubebuilder:validation:Required
	HTTPSPort *float64 `json:"httpsPort" tf:"https_port,omitempty"`

	// The value to use as the host header sent to the backend.
	// +kubebuilder:validation:Required
	HostHeader *string `json:"hostHeader" tf:"host_header,omitempty"`

	// Priority to use for load balancing. Higher priorities will not be used for load balancing if any lower priority backend is healthy. Defaults to 1.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Weight of this endpoint for load balancing purposes. Defaults to 50.
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type BackendPoolHealthProbeObservation struct {

	// The ID of the Azure Front Door Backend.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BackendPoolHealthProbeParameters struct {

	// Is this health probe enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The number of seconds between each Health Probe. Defaults to 120.
	// +kubebuilder:validation:Optional
	IntervalInSeconds *float64 `json:"intervalInSeconds,omitempty" tf:"interval_in_seconds,omitempty"`

	// Specifies the name of the Health Probe.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The path to use for the Health Probe. Default is /.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Specifies HTTP method the health probe uses when querying the backend pool instances. Possible values include: Get and Head. Defaults to Get.
	// +kubebuilder:validation:Optional
	ProbeMethod *string `json:"probeMethod,omitempty" tf:"probe_method,omitempty"`

	// Protocol scheme to use for the Health Probe. Defaults to Http.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type BackendPoolLoadBalancingObservation struct {

	// The ID of the Azure Front Door Backend.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BackendPoolLoadBalancingParameters struct {

	// The additional latency in milliseconds for probes to fall into the lowest latency bucket. Defaults to 0.
	// +kubebuilder:validation:Optional
	AdditionalLatencyMilliseconds *float64 `json:"additionalLatencyMilliseconds,omitempty" tf:"additional_latency_milliseconds,omitempty"`

	// Specifies the name of the Load Balancer.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The number of samples to consider for load balancing decisions. Defaults to 4.
	// +kubebuilder:validation:Optional
	SampleSize *float64 `json:"sampleSize,omitempty" tf:"sample_size,omitempty"`

	// The number of samples within the sample period that must succeed. Defaults to 2.
	// +kubebuilder:validation:Optional
	SuccessfulSamplesRequired *float64 `json:"successfulSamplesRequired,omitempty" tf:"successful_samples_required,omitempty"`
}

type BackendPoolObservation struct {

	// The ID of the Azure Front Door Backend.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BackendPoolParameters struct {

	// A backend block as defined below.
	// +kubebuilder:validation:Required
	Backend []BackendParameters `json:"backend" tf:"backend,omitempty"`

	// Specifies the name of the backend_pool_health_probe block within this resource to use for this Backend Pool.
	// +kubebuilder:validation:Required
	HealthProbeName *string `json:"healthProbeName" tf:"health_probe_name,omitempty"`

	// Specifies the name of the backend_pool_load_balancing block within this resource to use for this Backend Pool.
	// +kubebuilder:validation:Required
	LoadBalancingName *string `json:"loadBalancingName" tf:"load_balancing_name,omitempty"`

	// Specifies the name of the Backend Pool.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type BackendPoolSettingsObservation struct {
}

type BackendPoolSettingsParameters struct {

	// Specifies the send and receive timeout on forwarding request to the backend. When the timeout is reached, the request fails and returns. Possible values are between 0 - 240. Defaults to 60.
	// +kubebuilder:validation:Optional
	BackendPoolsSendReceiveTimeoutSeconds *float64 `json:"backendPoolsSendReceiveTimeoutSeconds,omitempty" tf:"backend_pools_send_receive_timeout_seconds,omitempty"`

	// Enforce certificate name check on HTTPS requests to all backend pools, this setting will have no effect on HTTP requests. Permitted values are true or false.
	// +kubebuilder:validation:Required
	EnforceBackendPoolsCertificateNameCheck *bool `json:"enforceBackendPoolsCertificateNameCheck" tf:"enforce_backend_pools_certificate_name_check,omitempty"`
}

type ExplicitResourceOrderObservation struct {
	BackendPoolHealthProbeIds []*string `json:"backendPoolHealthProbeIds,omitempty" tf:"backend_pool_health_probe_ids,omitempty"`

	BackendPoolIds []*string `json:"backendPoolIds,omitempty" tf:"backend_pool_ids,omitempty"`

	BackendPoolLoadBalancingIds []*string `json:"backendPoolLoadBalancingIds,omitempty" tf:"backend_pool_load_balancing_ids,omitempty"`

	FrontendEndpointIds []*string `json:"frontendEndpointIds,omitempty" tf:"frontend_endpoint_ids,omitempty"`

	RoutingRuleIds []*string `json:"routingRuleIds,omitempty" tf:"routing_rule_ids,omitempty"`
}

type ExplicitResourceOrderParameters struct {
}

type ForwardingConfigurationObservation struct {
}

type ForwardingConfigurationParameters struct {

	// Specifies the name of the Backend Pool to forward the incoming traffic to.
	// +kubebuilder:validation:Required
	BackendPoolName *string `json:"backendPoolName" tf:"backend_pool_name,omitempty"`

	// Specify the caching duration (in ISO8601 notation e.g. P1DT2H for 1 day and 2 hours). Needs to be greater than 0 and smaller than 365 days. cache_duration works only in combination with cache_enabled set to true.
	// +kubebuilder:validation:Optional
	CacheDuration *string `json:"cacheDuration,omitempty" tf:"cache_duration,omitempty"`

	// Specifies whether to Enable caching or not. Valid options are true or false. Defaults to false.
	// +kubebuilder:validation:Optional
	CacheEnabled *bool `json:"cacheEnabled,omitempty" tf:"cache_enabled,omitempty"`

	// Defines cache behaviour in relation to query string parameters. Valid options are StripAll, StripAllExcept, StripOnly or StripNone. Defaults to StripAll.
	// +kubebuilder:validation:Optional
	CacheQueryParameterStripDirective *string `json:"cacheQueryParameterStripDirective,omitempty" tf:"cache_query_parameter_strip_directive,omitempty"`

	// Specify query parameters (array). Works only in combination with cache_query_parameter_strip_directive set to StripAllExcept or StripOnly.
	// +kubebuilder:validation:Optional
	CacheQueryParameters []*string `json:"cacheQueryParameters,omitempty" tf:"cache_query_parameters,omitempty"`

	// Whether to use dynamic compression when caching. Valid options are true or false. Defaults to false.
	// +kubebuilder:validation:Optional
	CacheUseDynamicCompression *bool `json:"cacheUseDynamicCompression,omitempty" tf:"cache_use_dynamic_compression,omitempty"`

	// Path to use when constructing the request to forward to the backend. This functions as a URL Rewrite. Default behaviour preserves the URL path.
	// +kubebuilder:validation:Optional
	CustomForwardingPath *string `json:"customForwardingPath,omitempty" tf:"custom_forwarding_path,omitempty"`

	// Protocol to use when redirecting. Valid options are HttpOnly, HttpsOnly, or MatchRequest. Defaults to HttpsOnly.
	// +kubebuilder:validation:Optional
	ForwardingProtocol *string `json:"forwardingProtocol,omitempty" tf:"forwarding_protocol,omitempty"`
}

type FrontDoorObservation struct {

	// A backend_pool block as defined below.
	// +kubebuilder:validation:Required
	BackendPool []BackendPoolObservation `json:"backendPool,omitempty" tf:"backend_pool,omitempty"`

	// A backend_pool_health_probe block as defined below.
	// +kubebuilder:validation:Required
	BackendPoolHealthProbe []BackendPoolHealthProbeObservation `json:"backendPoolHealthProbe,omitempty" tf:"backend_pool_health_probe,omitempty"`

	// A map/dictionary of Backend Pool Health Probe Names (key) to the Backend Pool Health Probe ID (value)
	BackendPoolHealthProbes map[string]*string `json:"backendPoolHealthProbes,omitempty" tf:"backend_pool_health_probes,omitempty"`

	// A backend_pool_load_balancing block as defined below.
	// +kubebuilder:validation:Required
	BackendPoolLoadBalancing []BackendPoolLoadBalancingObservation `json:"backendPoolLoadBalancing,omitempty" tf:"backend_pool_load_balancing,omitempty"`

	// A map/dictionary of Backend Pool Load Balancing Setting Names (key) to the Backend Pool Load Balancing Setting ID (value)
	BackendPoolLoadBalancingSettings map[string]*string `json:"backendPoolLoadBalancingSettings,omitempty" tf:"backend_pool_load_balancing_settings,omitempty"`

	// A map/dictionary of Backend Pool Names (key) to the Backend Pool ID (value)
	BackendPools map[string]*string `json:"backendPools,omitempty" tf:"backend_pools,omitempty"`

	// The host that each frontendEndpoint must CNAME to.
	CNAME *string `json:"cname,omitempty" tf:"cname,omitempty"`

	ExplicitResourceOrder []ExplicitResourceOrderObservation `json:"explicitResourceOrder,omitempty" tf:"explicit_resource_order,omitempty"`

	// A frontend_endpoint block as defined below.
	// +kubebuilder:validation:Required
	FrontendEndpoint []FrontendEndpointObservation `json:"frontendEndpoint,omitempty" tf:"frontend_endpoint,omitempty"`

	// A map/dictionary of Frontend Endpoint Names (key) to the Frontend Endpoint ID (value)
	FrontendEndpoints map[string]*string `json:"frontendEndpoints,omitempty" tf:"frontend_endpoints,omitempty"`

	// The unique ID of the Front Door which is embedded into the incoming headers X-Azure-FDID attribute and maybe used to filter traffic sent by the Front Door to your backend.
	HeaderFrontdoorID *string `json:"headerFrontdoorId,omitempty" tf:"header_frontdoor_id,omitempty"`

	// The ID of the Azure Front Door Backend.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A routing_rule block as defined below.
	// +kubebuilder:validation:Required
	RoutingRule []RoutingRuleObservation `json:"routingRule,omitempty" tf:"routing_rule,omitempty"`

	// A map/dictionary of Routing Rule Names (key) to the Routing Rule ID (value)
	RoutingRules map[string]*string `json:"routingRules,omitempty" tf:"routing_rules,omitempty"`
}

type FrontDoorParameters struct {

	// A backend_pool block as defined below.
	// +kubebuilder:validation:Required
	BackendPool []BackendPoolParameters `json:"backendPool" tf:"backend_pool,omitempty"`

	// A backend_pool_health_probe block as defined below.
	// +kubebuilder:validation:Required
	BackendPoolHealthProbe []BackendPoolHealthProbeParameters `json:"backendPoolHealthProbe" tf:"backend_pool_health_probe,omitempty"`

	// A backend_pool_load_balancing block as defined below.
	// +kubebuilder:validation:Required
	BackendPoolLoadBalancing []BackendPoolLoadBalancingParameters `json:"backendPoolLoadBalancing" tf:"backend_pool_load_balancing,omitempty"`

	// +kubebuilder:validation:Optional
	BackendPoolSettings []BackendPoolSettingsParameters `json:"backendPoolSettings,omitempty" tf:"backend_pool_settings,omitempty"`

	// A friendly name for the Front Door service.
	// +kubebuilder:validation:Optional
	FriendlyName *string `json:"friendlyName,omitempty" tf:"friendly_name,omitempty"`

	// A frontend_endpoint block as defined below.
	// +kubebuilder:validation:Required
	FrontendEndpoint []FrontendEndpointParameters `json:"frontendEndpoint" tf:"frontend_endpoint,omitempty"`

	// Should the Front Door Load Balancer be Enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	LoadBalancerEnabled *bool `json:"loadBalancerEnabled,omitempty" tf:"load_balancer_enabled,omitempty"`

	// Specifies the name of the Resource Group in which the Front Door service should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A routing_rule block as defined below.
	// +kubebuilder:validation:Required
	RoutingRule []RoutingRuleParameters `json:"routingRule" tf:"routing_rule,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type FrontendEndpointObservation struct {

	// The ID of the Azure Front Door Backend.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FrontendEndpointParameters struct {

	// Specifies the host name of the frontend_endpoint. Must be a domain name. In order to use a name.azurefd.net domain, the name value must match the Front Door name.
	// +kubebuilder:validation:Required
	HostName *string `json:"hostName" tf:"host_name,omitempty"`

	// Specifies the name of the frontend_endpoint.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Whether to allow session affinity on this host. Valid options are true or false Defaults to false.
	// +kubebuilder:validation:Optional
	SessionAffinityEnabled *bool `json:"sessionAffinityEnabled,omitempty" tf:"session_affinity_enabled,omitempty"`

	// The TTL to use in seconds for session affinity, if applicable. Defaults to 0.
	// +kubebuilder:validation:Optional
	SessionAffinityTTLSeconds *float64 `json:"sessionAffinityTtlSeconds,omitempty" tf:"session_affinity_ttl_seconds,omitempty"`

	// Defines the Web Application Firewall policy ID for each host.
	// +kubebuilder:validation:Optional
	WebApplicationFirewallPolicyLinkID *string `json:"webApplicationFirewallPolicyLinkId,omitempty" tf:"web_application_firewall_policy_link_id,omitempty"`
}

type RoutingRuleObservation struct {

	// The ID of the Azure Front Door Backend.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RoutingRuleParameters struct {

	// Protocol schemes to match for the Backend Routing Rule. Defaults to Http.
	// +kubebuilder:validation:Required
	AcceptedProtocols []*string `json:"acceptedProtocols" tf:"accepted_protocols,omitempty"`

	// Enable or Disable use of this Backend Routing Rule. Permitted values are true or false. Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A forwarding_configuration block as defined below.
	// +kubebuilder:validation:Optional
	ForwardingConfiguration []ForwardingConfigurationParameters `json:"forwardingConfiguration,omitempty" tf:"forwarding_configuration,omitempty"`

	// The names of the frontend_endpoint blocks within this resource to associate with this routing_rule.
	// +kubebuilder:validation:Required
	FrontendEndpoints []*string `json:"frontendEndpoints" tf:"frontend_endpoints,omitempty"`

	// Specifies the name of the Routing Rule.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The route patterns for the Backend Routing Rule. Defaults to /*.
	// +kubebuilder:validation:Required
	PatternsToMatch []*string `json:"patternsToMatch" tf:"patterns_to_match,omitempty"`

	// A redirect_configuration block as defined below.
	// +kubebuilder:validation:Optional
	RedirectConfiguration []RoutingRuleRedirectConfigurationParameters `json:"redirectConfiguration,omitempty" tf:"redirect_configuration,omitempty"`
}

type RoutingRuleRedirectConfigurationObservation struct {
}

type RoutingRuleRedirectConfigurationParameters struct {

	// The destination fragment in the portion of URL after '#'. Set this to add a fragment to the redirect URL.
	// +kubebuilder:validation:Optional
	CustomFragment *string `json:"customFragment,omitempty" tf:"custom_fragment,omitempty"`

	// Set this to change the URL for the redirection.
	// +kubebuilder:validation:Optional
	CustomHost *string `json:"customHost,omitempty" tf:"custom_host,omitempty"`

	// The path to retain as per the incoming request, or update in the URL for the redirection.
	// +kubebuilder:validation:Optional
	CustomPath *string `json:"customPath,omitempty" tf:"custom_path,omitempty"`

	// Replace any existing query string from the incoming request URL.
	// +kubebuilder:validation:Optional
	CustomQueryString *string `json:"customQueryString,omitempty" tf:"custom_query_string,omitempty"`

	// Protocol to use when redirecting. Valid options are HttpOnly, HttpsOnly, or MatchRequest. Defaults to MatchRequest
	// +kubebuilder:validation:Required
	RedirectProtocol *string `json:"redirectProtocol" tf:"redirect_protocol,omitempty"`

	// Status code for the redirect. Valida options are Moved, Found, TemporaryRedirect, PermanentRedirect.
	// +kubebuilder:validation:Required
	RedirectType *string `json:"redirectType" tf:"redirect_type,omitempty"`
}

// FrontDoorSpec defines the desired state of FrontDoor
type FrontDoorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FrontDoorParameters `json:"forProvider"`
}

// FrontDoorStatus defines the observed state of FrontDoor.
type FrontDoorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FrontDoorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FrontDoor is the Schema for the FrontDoors API. Manages an Azure Front Door instance.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FrontDoor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FrontDoorSpec   `json:"spec"`
	Status            FrontDoorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FrontDoorList contains a list of FrontDoors
type FrontDoorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FrontDoor `json:"items"`
}

// Repository type metadata.
var (
	FrontDoor_Kind             = "FrontDoor"
	FrontDoor_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FrontDoor_Kind}.String()
	FrontDoor_KindAPIVersion   = FrontDoor_Kind + "." + CRDGroupVersion.String()
	FrontDoor_GroupVersionKind = CRDGroupVersion.WithKind(FrontDoor_Kind)
)

func init() {
	SchemeBuilder.Register(&FrontDoor{}, &FrontDoorList{})
}