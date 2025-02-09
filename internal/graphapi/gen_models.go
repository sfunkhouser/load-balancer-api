// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphapi

import (
	"go.infratographer.com/load-balancer-api/internal/ent/generated"
	"go.infratographer.com/x/gidx"
)

// Return response from loadBalancerCreate
type LoadBalancerCreatePayload struct {
	// The created load balancer.
	LoadBalancer *generated.LoadBalancer `json:"loadBalancer"`
}

// Return response from loadBalancerDelete
type LoadBalancerDeletePayload struct {
	// The ID of the deleted load balancer.
	DeletedID gidx.PrefixedID `json:"deletedID"`
}

// Return response from loadBalancerOriginCreate
type LoadBalancerOriginCreatePayload struct {
	// The created pool origin.
	LoadBalancerOrigin *generated.Origin `json:"loadBalancerOrigin"`
}

// Return response from loadBalancerOriginDelete
type LoadBalancerOriginDeletePayload struct {
	// The deleted pool origin.
	DeletedID gidx.PrefixedID `json:"deletedID"`
}

// Return response from loadBalancerOriginUpdate
type LoadBalancerOriginUpdatePayload struct {
	// The updated pool origin.
	LoadBalancerOrigin *generated.Origin `json:"loadBalancerOrigin"`
}

// Return response from LoadBalancerPoolCreate
type LoadBalancerPoolCreatePayload struct {
	// The created pool.
	LoadBalancerPool *generated.Pool `json:"loadBalancerPool"`
}

// Return response from LoadBalancerPoolDelete
type LoadBalancerPoolDeletePayload struct {
	// The ID of the deleted pool.
	DeletedID *gidx.PrefixedID `json:"deletedID,omitempty"`
}

// Return response from LoadBalancerPoolUpdate
type LoadBalancerPoolUpdatePayload struct {
	// The updated pool.
	LoadBalancerPool *generated.Pool `json:"loadBalancerPool"`
}

// Return response from loadBalancerPortCreate
type LoadBalancerPortCreatePayload struct {
	// The created load balancer port.
	LoadBalancerPort *generated.Port `json:"loadBalancerPort"`
}

// Return response from loadBalancerPortDelete
type LoadBalancerPortDeletePayload struct {
	// The ID of the deleted load balancer port.
	DeletedID gidx.PrefixedID `json:"deletedID"`
}

// Return response from loadBalancerPortUpdate
type LoadBalancerPortUpdatePayload struct {
	// The updated load balancer port.
	LoadBalancerPort *generated.Port `json:"loadBalancerPort"`
}

// Return response from loadBalancerProviderCreate
type LoadBalancerProviderCreatePayload struct {
	// The created load balancer provider.
	LoadBalancerProvider *generated.Provider `json:"loadBalancerProvider"`
}

// Return response from loadBalancerProviderDelete
type LoadBalancerProviderDeletePayload struct {
	// The ID of the deleted load balancer provider.
	DeletedID gidx.PrefixedID `json:"deletedID"`
}

// Return response from loadBalancerProviderUpdate
type LoadBalancerProviderUpdatePayload struct {
	// The updated load balancer provider.
	LoadBalancerProvider *generated.Provider `json:"loadBalancerProvider"`
}

// Return response from loadBalancerUpdate
type LoadBalancerUpdatePayload struct {
	// The updated load balancer.
	LoadBalancer *generated.LoadBalancer `json:"loadBalancer"`
}
