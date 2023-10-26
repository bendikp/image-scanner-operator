// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// ScanConfigApplyConfiguration represents an declarative configuration of the ScanConfig type for use
// with apply.
type ScanConfigApplyConfiguration struct {
	MinSeverity   *string `json:"minSeverity,omitempty"`
	IgnoreUnfixed *bool   `json:"ignoreUnfixed,omitempty"`
}

// ScanConfigApplyConfiguration constructs an declarative configuration of the ScanConfig type for use with
// apply.
func ScanConfig() *ScanConfigApplyConfiguration {
	return &ScanConfigApplyConfiguration{}
}

// WithMinSeverity sets the MinSeverity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MinSeverity field is set to the value of the last call.
func (b *ScanConfigApplyConfiguration) WithMinSeverity(value string) *ScanConfigApplyConfiguration {
	b.MinSeverity = &value
	return b
}

// WithIgnoreUnfixed sets the IgnoreUnfixed field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IgnoreUnfixed field is set to the value of the last call.
func (b *ScanConfigApplyConfiguration) WithIgnoreUnfixed(value bool) *ScanConfigApplyConfiguration {
	b.IgnoreUnfixed = &value
	return b
}