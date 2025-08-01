// Add to existing config structure
type ConfigFileMatching struct {
    AutoMatchThreshold float64 `json:"autoMatchThreshold"` // Minimum score required for auto-matching
    AutoMatchScaleFactor float64 `json:"autoMatchScaleFactor"` // How much higher the top score must be compared to second best
    EnableAutoMatch bool `json:"enableAutoMatch"` // Whether to enable automatic matching
}
