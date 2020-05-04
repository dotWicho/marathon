package marathon

import (
	"time"
)

type Queues struct {
	Queue []Queue `json:"queue"`
}

type Queue struct {
	App                    AppQueue               `json:"app"`
	Count                  int                    `json:"count"`
	Delay                  Delay                  `json:"delay"`
	Since                  interface{}            `json:"since"`
	ProcessedOffersSummary ProcessedOffersSummary `json:"processedOffersSummary"`
	LastUnusedOffers       []LastUnusedOffer      `json:"lastUnusedOffers,omitempty"`
}

type AppQueue struct {
	ID                    string                      `json:"id"`
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

type Delay struct {
	TimeLeftSeconds int  `json:"timeLeftSeconds"`
	Overdue         bool `json:"overdue"`
}

type ProcessedOffersSummary struct {
	ProcessedOffersCount int `json:"processedOffersCount"`
	UnusedOffersCount    int `json:"unusedOffersCount"`
}

type LastUnusedOffer struct {
	Offer     Offer       `json:"offer"`
	Timestamp interface{} `json:"timestamp"`
	Reason    []string    `json:"reason"`
}

type Offer struct {
	ID         string      `json:"id"`
	AgentID    string      `json:"agentId"`
	Hostname   string      `json:"hostname"`
	Resources  []Resource  `json:"resources"`
	Attributes []Attribute `json:"attributes"`
}

type Resource struct {
	Name   string   `json:"name"`
	Scalar int      `json:"scalar"`
	Ranges []Range  `json:"ranges"`
	Set    []string `json:"set"`
	Role   string   `json:"role"`
}

type Attribute struct {
	Name   string  `json:"name"`
	Scalar int     `json:"scalar"`
	Ranges []Range `json:"ranges"`
	Set []string   `json:"set"`
}

type Range struct {
	Begin int `json:"begin"`
	End   int `json:"end"`
}