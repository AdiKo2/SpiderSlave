package Repository

import "time"

type HomeUser struct {
	DisplayName string `xml:"displayName"`
	FirstName   string `xml:"firstName"`
	LastName    string `xml:"lastName"`
	Email       string `xml:"email"`
	//IpPhone
}

type HomeUserPointer struct {
	DisplayName *string
	FirstName   *string
	LastName    *string
	Email       *string
}

type EnvStatus struct {
	IsOn            bool
	InProcess       bool
	LastTimeProcess time.Time
}

type ReceiveContacts struct {
	Identity    string
	DisplayName string
	FirstName   string
	LastName    string
	Email       string
	ReceiveTime time.Time
	Enable      bool
}

type CreateNewContacts struct {
	DisplayName string
	FirstName   string
	LastName    string
	Email       string
}

type TrustedPartners struct {
	PartnerName        string
	DisplayNamePartner string
	EmailEnding        string
}

type SystemSettings struct {
	ServiceName string
	InUse       bool
	LastTimeUse time.Time
}

type DBRepository struct {
}
