package panels

// for main
const (
	AppTitle         = "Bar Code Generator"
	MainWindowWidth  = 800
	MainWindowHeight = 600
)

// DataPanel
const (
	DataPaneTitlelString = "Enter content"
)

const (
	DTText = iota
	DTURL
	DTSMS
)

// NavigatorPanel
const (
	// AdjustColorsString     = "Adjust colors"
	// AdjustDesignString     = "Adjust design"
	CustomizeQRCodeString  = "Customize QR code"
	DTContainerTitleString = "Data type"
	DTURLString            = "URL"
	DTTextString           = "Text"
	DTSMSString            = "SMS"
)

// PreviewPanel
const (
	QrIconString  = "qr-code-250.png"
	MinQrIconSize = 250
)

// TextPanel
const (
	TextLabelString = "Text:"
)

// URLPanel
const (
	UrlLabelString = "URL:"
)

// SMSPanel
const (
	SmsNrLabelString   = "Phone:"
	SmsTextLabelString = "Text:"
)
