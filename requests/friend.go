package requests

type FacebookIncentiveRequest struct {
    Base
    Type             int64 `json:"type,string"`
    AchievementCount int64 `json:"achievementCount,string"`
}
