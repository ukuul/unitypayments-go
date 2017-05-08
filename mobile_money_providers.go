package unitypayments

// MobileMoneyProviders is an enum of mobile money providers.
type MobileMoneyProviders int

// Mobile money providers
const (
	MTN MobileMoneyProviders = iota
	Airtel
	Vodafone
	Tigo
)

func (p MobileMoneyProviders) String() string {
	switch p {
	case MTN:
		return "mtn-gh"
	case Airtel:
		return "airtel-gh"
	case Vodafone:
		return "vodafone-gh"
	case Tigo:
		return "tigo-gh"
	default:
		return ""
	}
}

func (p MobileMoneyProviders) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}
