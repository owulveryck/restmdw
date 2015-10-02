package api

type OfferingStructureFields struct {
	Description string                      `yaml:"description,omitempty" json:"-"`
	Value       interface{}                 `yaml:"value" json:"value"`
	Name        string                      `yaml:"name,omitempty" json:"name"`
	Type        string                      `yaml:"type" json:"-"`
	ValidValues map[interface{}]interface{} `yaml:"valid_values,omitempty" json:"-"`
	//Default     interface{}                 `yaml:"default,omitempty" json:"-"`
}

// OfferingStructure contains the offering of the API
type OfferingStructure struct {
	Description string `yaml:"description" json:"-"` // The description of the Action
	State       string `yaml:"state" json:"state"`   // the state as described in TOSCA http://docs.oasis-open.org/tosca/TOSCA-Simple-Profile-YAML/v1.0/csd03/TOSCA-Simple-Profile-YAML-v1.0-csd03.html#_Toc419746142
	Request     struct {
		Name         string                             `yaml:"request_name" json:"request_name"`
		RequestID    string                             `yaml:"id,omitempty" json:"id"`
		RequestState string                             `yaml:"state,omitempty" json:"state"` // SUBMITTED|PENDING_APPROVAL|REJECTED|APPROVED|IN_PROGRESS|COMPLETED|CANCELLED
		Fields       map[string]OfferingStructureFields `yaml:"input" json:"input"`           // The index of the map if the filed_id

	} `yaml:"request" json:"request"`
	Subscription struct {
		SubscriptionID     string `yaml:"id,omitempty" json:"id"`
		SubscriptionStatus string `yaml:"status,omitempty" json:"status"` // PENDING|ACTIVE|EXPIRED|CANCELLED|TERMINATED

	} `yaml:"subscription,omitempty" json:"subscription"`
	Instance struct {
		InstanceID    string `yaml:"id,omitempty" json:"id"`
		InstanceState string `yaml:"state,omitempty" json:"state"` // IN_PROGRESS|RESERVED|DEPLOYING| ACTIVE|FAILED|CANCELLED| MODIFYING|MODIFY_FAILED|CANCELLING| CANCEL_FAILED|EXPIRING|EXPIRE_FAILED| PUBLIC_ACTION_FAILED
	} `yaml:"instance,omitempty" json:"instance"`
	Offering struct {
		Version      string `yaml:"version" json:"version"`
		CatalogId    string `yaml:"catalog_id" json:"catalogId"`
		OfferingID   string `yaml:"offering_id" json:"offering_id"`
		CategoryName string `yaml:"category_name" json:"category_name"`
	} `yaml:"offering" json:"-"`
}
