package marathon

import (
	"time"
)

// Queues Array of
type Queues struct {
	Queue []QueueType `json:"queue"`
}

// QueueType holds definitions of Tasks queued up or waiting to be scheduled
type QueueType struct {
	App                    AppDefinition          `json:"app"`
	Count                  int                    `json:"count"`
	Delay                  Delay                  `json:"delay"`
	Since                  interface{}            `json:"since"`
	ProcessedOffersSummary ProcessedOffersSummary `json:"processedOffersSummary"`
	LastUnusedOffers       []LastUnusedOffer      `json:"lastUnusedOffers,omitempty"`
}

// AppQueue holds definitions of Apps in Queue
type AppQueue struct {
	ID                    string          `json:"id"`
	Instances             int             `json:"instances"`
	Cpus                  int             `json:"cpus"`
	Mem                   int             `json:"mem"`
	Disk                  int             `json:"disk"`
	Constraints           [][]string      `json:"constraints"`
	Ports                 []int           `json:"ports"`
	RequirePorts          bool            `json:"requirePorts"`
	BackoffSeconds        int             `json:"backoffSeconds"`
	BackoffFactor         float64         `json:"backoffFactor"`
	MaxLaunchDelaySeconds int             `json:"maxLaunchDelaySeconds"`
	Container             Queue           `json:"container"`
	UpgradeStrategy       UpgradeStrategy `json:"upgradeStrategy"`
	Version               time.Time       `json:"version"`
	VersionInfo           VersionInfo     `json:"versionInfo"`
}

// Delay Queue representation
type Delay struct {
	TimeLeftSeconds int  `json:"timeLeftSeconds"`
	Overdue         bool `json:"overdue"`
}

// ProcessedOffersSummary is Processed Offers Summary info
type ProcessedOffersSummary struct {
	ProcessedOffersCount int `json:"processedOffersCount"`
	UnusedOffersCount    int `json:"unusedOffersCount"`
}

// LastUnusedOffer Offers not used info
type LastUnusedOffer struct {
	Offer     Offer       `json:"offer"`
	Timestamp interface{} `json:"timestamp"`
	Reason    []string    `json:"reason"`
}

// Offer holds info of Mesos Offered resources
type Offer struct {
	ID         string      `json:"id"`
	AgentID    string      `json:"agentId"`
	Hostname   string      `json:"hostname"`
	Resources  []Resource  `json:"resources"`
	Attributes []Attribute `json:"attributes"`
}

// Resource is Offer resource representation
type Resource struct {
	Name   string   `json:"name"`
	Scalar int      `json:"scalar"`
	Ranges []Range  `json:"ranges"`
	Set    []string `json:"set"`
	Role   string   `json:"role"`
}

// Attribute of an Offer from Mesos
type Attribute struct {
	Name   string   `json:"name"`
	Scalar int      `json:"scalar"`
	Ranges []Range  `json:"ranges"`
	Set    []string `json:"set"`
}

// Range scalar of an Offer from Mesos
type Range struct {
	Begin int `json:"begin"`
	End   int `json:"end"`
}
