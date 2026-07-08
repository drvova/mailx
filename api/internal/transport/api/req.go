package api

type UserReq struct {
	Email    string `json:"email" validate:"required,emailx"`
	Password string `json:"password" validate:"password"`
	OTP      string `json:"otp" validate:"min=0,max=8"`
}

type AuthReq struct {
	AccessKey string `json:"access_key" validate:"required,min=85,max=85"`
}

type EmailReq struct {
	Email string `json:"email" validate:"required,emailx"`
}

type SignupUserReq struct {
	Email    string `json:"email" validate:"required,emailx"`
	Password string `json:"password" validate:"password"`
	SubID    string `json:"subid" validate:"omitempty,uuid"`
}

type SignupEmailReq struct {
	Email string `json:"email" validate:"required,emailx"`
	SubID string `json:"subid" validate:"required,uuid"`
}

type SubscriptionReq struct {
	ID    string `json:"id" validate:"required,uuid"`
	SubID string `json:"subid"`
}

type AliasReq struct {
	Description    string `json:"description"`
	Enabled        bool   `json:"enabled"`
	Recipients     string `json:"recipients"`
	FromName       string `json:"from_name"`
	Format         string `json:"format"`
	Domain         string `json:"domain" validate:"required"`
	CatchAllSuffix string `json:"catch_all_suffix" validate:"omitempty,alphanum,min=6,max=12"`
	Type           string `json:"type" validate:"omitempty,oneof=relay inbox"`
	TTLHours       int    `json:"ttl_hours" validate:"omitempty,min=1,max=720"`
}

type RecipientReq struct {
	ID         string `json:"id" validate:"required,uuid"`
	PGPKey     string `json:"pgp_key" validate:"omitempty,pgp"`
	PGPEnabled bool   `json:"pgp_enabled"`
	PGPInline  bool   `json:"pgp_inline"`
}

type DeleteRecipientReq struct {
	Recipients string `json:"recipients"`
}

type SettingsReq struct {
	ID           string `json:"id" validate:"required,uuid"`
	Domain       string `json:"domain"`
	Recipient    string `json:"recipient"`
	FromName     string `json:"from_name"`
	AliasFormat  string `json:"alias_format"`
	LogIssues    bool   `json:"log_issues"`
	RemoveHeader bool   `json:"remove_header"`
}

type DeleteUserReq struct {
	OTP string `json:"otp" validate:"required,len=8"`
}

type ChangePasswordReq struct {
	Password string `json:"password" validate:"password"`
}

type ResetPasswordReq struct {
	OTP      string `json:"otp" validate:"required,len=32"`
	Password string `json:"password" validate:"password"`
}

type ActivateReq struct {
	OTP string `json:"otp" validate:"required,len=6"`
}

type TotpReq struct {
	OTP string `json:"otp" validate:"required,min=6,max=8"`
}

type PASessionReq struct {
	ID        string `json:"id" validate:"required,uuid"`
	PreauthId string `json:"preauth_id" validate:"required,uuid"`
	Token     string `json:"token" validate:"required"`
}

type RotatePASessionReq struct {
	ID string `json:"sessionid" validate:"required,uuid"`
}

type AccessKeyReq struct {
	Name      string `json:"name" validate:"required"`
	ExpiresAt string `json:"expires_at"`
}

type DomainReq struct {
	Name string `json:"name" validate:"required,fqdn"`
}

type UpdateDomainReq struct {
	ID          string `json:"id" validate:"required,uuid"`
	Description string `json:"description"`
	Recipient   string `json:"recipient"`
	FromName    string `json:"from_name"`
	Enabled     bool   `json:"enabled"`
}

type PlanReq struct {
	Name              string `json:"name" validate:"required,min=2,max=50"`
	DisplayName       string `json:"display_name" validate:"required,min=2,max=100"`
	PriceCents        int    `json:"price_cents" validate:"min=0"`
	Currency          string `json:"currency" validate:"required,oneof=usd eur gbp"`
	Interval          string `json:"interval" validate:"required,oneof=monthly yearly one_time"`
	MaxRecipients     int    `json:"max_recipients" validate:"min=0"`
	MaxCredentials    int    `json:"max_credentials" validate:"min=0"`
	MaxDailyAliases   int    `json:"max_daily_aliases" validate:"min=0"`
	MaxDailySendReply int    `json:"max_daily_send_reply" validate:"min=0"`
	MaxSessions       int    `json:"max_sessions" validate:"min=0"`
	SortOrder         int    `json:"sort_order" validate:"min=0"`
}

type CheckoutReq struct {
	PlanID string `json:"plan_id" validate:"required,uuid"`
}
