package event

type V3_LOCALE string

const (
	V3_UNKNOWN_LOCALE V3_LOCALE = ""
	V3_LOCALE_KO      V3_LOCALE = "KO"
	V3_LOCALE_EN      V3_LOCALE = "EN"
	V3_LOCALE_TL      V3_LOCALE = "TL"
	V3_LOCALE_VI      V3_LOCALE = "VI"
	V3_LOCALE_ID      V3_LOCALE = "ID"
	V3_LOCALE_ZH      V3_LOCALE = "ZH"
	V3_LOCALE_TH      V3_LOCALE = "TH"
	V3_LOCALE_KM      V3_LOCALE = "KM"
	V3_LOCALE_NE      V3_LOCALE = "NE"
	V3_LOCALE_UR      V3_LOCALE = "UR"
	V3_LOCALE_BN      V3_LOCALE = "BN"
	V3_LOCALE_RU      V3_LOCALE = "RU"
	V3_LOCALE_UZ      V3_LOCALE = "UZ"
	V3_LOCALE_SI      V3_LOCALE = "SI"
	V3_LOCALE_HI      V3_LOCALE = "HI"
	V3_LOCALE_MS      V3_LOCALE = "MS"
	V3_LOCALE_MY      V3_LOCALE = "MY"
)

type EventModel struct {
	ID           int32
	EventDetail1 map[V3_LOCALE]string
}
