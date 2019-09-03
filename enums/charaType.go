package enums

// TODO: Apparently the IDs are different in game code versus what is sent. How is this achieved...?

const (
	CharaTypeUnknown = iota - 1
	CharaTypeSonic   = iota + 300000 - 1 // used to match up with networked characters
	CharaTypeTails
	CharaTypeKnuckles
	CharaTypeAmy
	CharaTypeShadow
	CharaTypeBlaze
	CharaTypeRouge
	CharaTypeOmega
	CharaTypeBig
	CharaTypeCream
	CharaTypeEspio
	CharaTypeCharmy
	CharaTypeVector
	CharaTypeSilver
	CharaTypeMetalSonic
	CharaTypeClassicSonic
	CharaTypeWerehog
	CharaTypeSticks
	CharaTypeTikal
	CharaTypeMephiles
	CharaTypePSISilver
	CharaTypeAmitieAmy
	CharaTypeGothicAmy
	CharaTypeHalloweenShadow
	CharaTypeHalloweenRouge
	CharaTypeHalloweenOmega
	CharaTypeXMasSonic
	CharaTypeXMasTails
	CharaTypeXMasKnuckles
)

const (
	CTStrSonic        = "300000"
	CTStrTails        = "300001"
	CTStrKnuckles     = "300002"
	CTStrAmy          = "300003"
	CTStrShadow       = "300004"
	CTStrBlaze        = "300005"
	CTStrRouge        = "300006"
	CTStrOmega        = "300007"
	CTStrBig          = "300008"
	CTStrCream        = "300009"
	CTStrEspio        = "300010"
	CTStrCharmy       = "300011"
	CTStrVector       = "300012"
	CTStrSilver       = "300013"
	CTStrMetalSonic   = "300014"
	CTStrClassicSonic = "300015"
	CTStrWerehog      = "300016"
	CTStrSticks       = "300017"
	CTStrTikal        = "300018"
	CTStrMephiles     = "300019"
	CTStrPSISilver    = "300020"
)
