package i18n

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// Catalog holds all localized UI strings, labels, table headers, and formats
// for the Go TUI Dashboard. It provides a static, zero-dependency translation architecture.
type Catalog struct {
	// Screen banners & general
	AppTitle       string
	OffersSummary  string
	NoOffersMatch  string
	LoadingPreview string

	// Tabs & filters
	TabAll       string
	TabEvaluated string
	TabApplied   string
	TabInterview string
	TabResponded string
	TabTop       string
	TabSkip      string
	TabRejected  string
	TabDiscarded string

	// Table column headers
	ColFit      string
	ColApplied  string
	ColCompany  string
	ColRole     string
	ColStatus   string
	ColLocation string
	ColPay      string
	ColLast     string

	// Preview labels
	LabelLoc     string
	LabelPay     string
	LabelLast    string
	LabelRemote  string
	LabelOutcome string

	// Work modes
	ModeRemote     string
	ModeRemoteFlex string
	ModeHybrid     string
	ModeFull       string

	// Progress screen
	ProgressTitle   string
	ProgressSummary string
	FunnelTitle     string
	ScoresTitle     string
	RatesTitle      string
	WeeklyTitle     string
	ActiveInfo      string

	// Relative dates
	TimeToday     string
	TimeYesterday string
	TimeDaysAgo   string

	// Status display names
	StatusEvaluated string
	StatusApplied   string
	StatusResponded string
	StatusInterview string
	StatusOffer     string
	StatusRejected  string
	StatusDiscarded string
	StatusSkip      string
	StatusHired     string

	// Additional UI strings
	NoData        string
	EmptyFile     string
	RateResponse  string
	RateInterview string
	RateOffer     string

	// Footer descriptions & hints
	HelpNav        string
	HelpTabs       string
	HelpSearch     string
	HelpSort       string
	HelpRefresh    string
	HelpReport     string
	HelpOpenURL    string
	HelpOpenPDF    string
	HelpRegenPDF   string
	HelpChange     string
	HelpColumns    string
	HelpView       string
	HelpProgress   string
	HelpQuit       string
	HelpScroll     string
	HelpPage       string
	HelpTopEnd     string
	HelpManifesto  string
	HelpBack       string
	HelpNavigate   string
	HelpToggle     string
	HelpClose      string
	HelpConfirm    string
	HelpCancel     string
	HelpFilterLive string
	HelpKeep       string
	HelpClear      string

	// Picker overlay titles & bar hints
	PickerChangeStatus string
	PickerColumnsTitle string
	SearchHintInput    string
	SearchHintNormal   string
	SearchMatching     string
	SortLabel          string
	ViewLabel          string
	ShownCount         string
	ColReport          string
	ColPDF             string

	// Sort & view modes
	SortScore    string
	SortDate     string
	SortCompany  string
	SortStatus   string
	SortLocation string
	SortPay      string
	SortLast     string
	ViewGrouped  string
	ViewFlat     string
}

// SortModeLabel returns the localized display label for a sort mode ("score", "date", etc.).
func (c *Catalog) SortModeLabel(mode string) string {
	switch strings.ToLower(strings.TrimSpace(mode)) {
	case "score":
		return c.SortScore
	case "date":
		return c.SortDate
	case "company":
		return c.SortCompany
	case "status":
		return c.SortStatus
	case "location":
		return c.SortLocation
	case "pay":
		return c.SortPay
	case "last":
		return c.SortLast
	default:
		return mode
	}
}

// ViewModeLabel returns the localized display label for a view mode ("grouped" or "flat").
func (c *Catalog) ViewModeLabel(mode string) string {
	switch strings.ToLower(strings.TrimSpace(mode)) {
	case "grouped":
		return c.ViewGrouped
	case "flat":
		return c.ViewFlat
	default:
		return mode
	}
}

// StatusLabel returns the localized display label for a canonical status ID
// (interview, offer, hired, responded, applied, evaluated, skip, rejected,
// discarded).
func (c *Catalog) StatusLabel(norm string) string {
	switch strings.ToLower(strings.TrimSpace(norm)) {
	case "interview":
		return c.StatusInterview
	case "offer":
		return c.StatusOffer
	case "hired":
		return c.StatusHired
	case "responded":
		return c.StatusResponded
	case "applied":
		return c.StatusApplied
	case "evaluated":
		return c.StatusEvaluated
	case "skip":
		return c.StatusSkip
	case "rejected":
		return c.StatusRejected
	case "discarded":
		return c.StatusDiscarded
	default:
		return norm
	}
}

// NowFunc allows injecting a mock clock for testing.
var NowFunc = time.Now

// FormatTimeAgo renders an ISO date as a relative duration in calendar days using localized strings:
// "today", "yesterday", or "Nd ago" (or Turkish equivalents).
func (c *Catalog) FormatTimeAgo(dateStr string) string {
	t, err := time.ParseInLocation("2006-01-02", dateStr, time.Local)
	if err != nil {
		return dateStr
	}
	now := NowFunc()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	contactDay := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	days := int(math.Round(today.Sub(contactDay).Hours() / 24))
	switch {
	case days <= 0:
		return c.TimeToday
	case days == 1:
		return c.TimeYesterday
	default:
		return fmt.Sprintf(c.TimeDaysAgo, days)
	}
}

// En is the static English translation catalog.
var En = Catalog{
	// Screen banners & general
	AppTitle:       "CAREER PIPELINE",
	OffersSummary:  "%d offers | Avg %s/5",
	NoOffersMatch:  "No offers match this filter",
	LoadingPreview: "Loading preview...",

	// Tabs & filters
	TabAll:       "ALL",
	TabEvaluated: "EVALUATED",
	TabApplied:   "APPLIED",
	TabInterview: "INTERVIEW",
	TabResponded: "RESPONDED",
	TabTop:       "TOP ≥4",
	TabSkip:      "SKIP",
	TabRejected:  "REJECTED",
	TabDiscarded: "DISCARDED",

	// Table column headers
	ColFit:      "FIT",
	ColApplied:  "APPLIED",
	ColCompany:  "COMPANY",
	ColRole:     "ROLE",
	ColStatus:   "STATUS",
	ColLocation: "LOCATION",
	ColPay:      "PAY",
	ColLast:     "LAST",

	// Preview labels
	LabelLoc:     "Loc: ",
	LabelPay:     "Pay: ",
	LabelLast:    "Last contact: ",
	LabelRemote:  "Remote: ",
	LabelOutcome: "Outcome: ",

	// Work modes
	ModeRemote:     "Remote",
	ModeRemoteFlex: "RemoteFlex",
	ModeHybrid:     "Hybrid",
	ModeFull:       "Full",

	// Progress screen
	ProgressTitle:   "SEARCH PROGRESS",
	ProgressSummary: "%d evaluated | %.1f avg score",
	FunnelTitle:     "Pipeline Funnel",
	ScoresTitle:     "Score Distribution",
	RatesTitle:      "Conversion Rates",
	WeeklyTitle:     "Weekly Activity",
	ActiveInfo:      "%d active applications | %d total offers",

	// Relative dates
	TimeToday:     "today",
	TimeYesterday: "yesterday",
	TimeDaysAgo:   "%dd ago",

	// Status display names
	StatusEvaluated: "Evaluated",
	StatusApplied:   "Applied",
	StatusResponded: "Responded",
	StatusInterview: "Interview",
	StatusOffer:     "Offer",
	StatusRejected:  "Rejected",
	StatusDiscarded: "Discarded",
	StatusSkip:      "SKIP",
	StatusHired:     "Hired",

	// Additional UI strings
	NoData:        "No data",
	EmptyFile:     "(empty file)",
	RateResponse:  "Response Rate: ",
	RateInterview: "Interview Rate: ",
	RateOffer:     "Offer Rate: ",

	// Footer descriptions & hints
	HelpNav:        " nav  ",
	HelpTabs:       " tabs  ",
	HelpSearch:     " search  ",
	HelpSort:       " sort  ",
	HelpRefresh:    " refresh  ",
	HelpReport:     " report  ",
	HelpOpenURL:    " open URL  ",
	HelpOpenPDF:    " open PDF  ",
	HelpRegenPDF:   " regen PDF  ",
	HelpChange:     " change  ",
	HelpColumns:    " columns  ",
	HelpView:       " view  ",
	HelpProgress:   " progress  ",
	HelpQuit:       " quit",
	HelpScroll:     " scroll  ",
	HelpPage:       " page  ",
	HelpTopEnd:     " top/end  ",
	HelpManifesto:  " manifesto  ",
	HelpBack:       " back",
	HelpNavigate:   " navigate  ",
	HelpToggle:     " toggle  ",
	HelpClose:      " close",
	HelpConfirm:    " confirm  ",
	HelpCancel:     " cancel",
	HelpFilterLive: " filter live  ",
	HelpKeep:       " keep  ",
	HelpClear:      " clear  ",

	// Picker overlay titles & bar hints
	PickerChangeStatus: "Change status:",
	PickerColumnsTitle: "─── Columns (SPACE toggle · ESC close) ───",
	SearchHintInput:    "   Enter: keep   Esc: cancel   Ctrl+U: clear",
	SearchHintNormal:   "   Esc: clear   /: edit",
	SearchMatching:     "  %d/%d matching",
	SortLabel:          "[Sort: %s]",
	ViewLabel:          "[View: %s]",
	ShownCount:         "%d shown",
	ColReport:          "RPT",
	ColPDF:             "PDF",

	// Sort & view modes
	SortScore:    "score",
	SortDate:     "date",
	SortCompany:  "company",
	SortStatus:   "status",
	SortLocation: "location",
	SortPay:      "pay",
	SortLast:     "last",
	ViewGrouped:  "grouped",
	ViewFlat:     "flat",
}

// Current points to the active language catalog. Defaults to English (&En).
var Current = &En

// SetLang is retained for API compatibility; the dashboard is English-only.
func SetLang(lang string) {
	Current = &En
}

// GetLang returns the active language code (always "en").
func GetLang() string {
	return "en"
}
