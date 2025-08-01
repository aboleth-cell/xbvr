type AppConfig struct {
    // ... existing fields ...
    
    SceneMatching struct {
        AutoMatchThreshold float64 `json:"autoMatchThreshold"` // Minimum score required for auto-matching
        ScoreRatio        float64 `json:"scoreRatio"`         // Required ratio between top and second-best scores
        EnableAutoMatch   bool    `json:"enableAutoMatch"`    // Toggle for auto-matching feature
    } `json:"sceneMatching"`
}
