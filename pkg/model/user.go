package model

import "time"

type ArtifactRefs struct {
	OneToMany

	//Values []Artifact
}

type ProjectRef struct {
	OneToOne

	//Project Project
}

type ProjectRefs struct {
	OneToMany

	//Projects []Project
}

type UserPermissionRefs struct {
	OneToMany

	//Projects []Project
}

type ProfileImageRef struct {
	OneToOne

	//Project Project
}

type UserProfileRef struct {
	OneToOne

	//Project Project
}

type RevisionHistoryRef struct {
	OneToOne

	//Project Project
}

type User struct {
	AccountLockedUntil             bool
	ArtifactsCreated               ArtifactRefs
	ArtifactsOwned                 ArtifactRefs
	CostCenter                     string
	DateFormat                     string
	DateTimeFormat                 string
	DefaultDetailPageToViewingMode bool
	DefaultProject                 ProjectRef
	Deleted                        bool
	Department                     string
	Disabled                       bool
	DisplayName                    string
	EmailAddress                   string
	EmailNotificationEnabled       bool
	FirstName                      string
	InvestmentAdmin                bool
	LandingPage                    string
	Language                       string
	LastActiveDate                 time.Time
	LastLoginDate                  time.Time
	LastName                       string
	LastPasswordUpdateDate         time.Time
	LastSystemTimeZoneName         string
	LdapUuid                       string
	Locale                         string
	MiddleName                     string
	NetworkID                      string
	OfficeLocation                 string
	OnpremLdapUsername             string
	PasswordExpires                int64
	Phone                          string
	Planner                        bool
	ProfileImage                   ProfileImageRef
	ProjectScopeDown               bool
	ProjectScopeUp                 bool
	RevisionHistory                RevisionHistoryRef
	Role                           string
	sessionTimeout                 int64
	SessionTimeoutWarning          bool
	ShortDisplayName               string
	SubscriptionAdmin              bool
	SubscriptionID                 int64
	SubscriptionPermission         string
	TeamMemberships                ProjectRefs
	UserName                       string
	UserPermissions                UserPermissionRefs
	UserProfile                    UserProfileRef
	WorkspacePermission            string
	ZuulID                         string
}
