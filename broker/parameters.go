package broker

type ProvisionParameters struct {
}

type BindParameters struct {
	AdditionalInstances []string `json:"additional_instances"`
}

type UpdateParameters struct {
	ApplyImmediately bool `json:"apply_immediately"`
}
