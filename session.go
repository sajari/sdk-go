package sajari

import (
	"fmt"
	"math/rand"

	pipelinepb "code.sajari.com/protogen-go/sajari/pipeline/v2"
)

// WebSearchSession handles search sessions for website-style searching.
//
// It tracks the search text (pass to the query pipeline via queryLabel),
// and determines if the session should be reset (i.e. if the query text has
// changed significantly).
//
// A common session handler would be:
//   WebSearchSeassion("q", NewSession())
func WebSearchSession(queryLabel string, s Session) Session {
	return &webSearchSession{
		Session:    s,
		queryLabel: queryLabel,
	}
}

type webSearchSession struct {
	Session

	queryLabel string

	lastQuery string
}

func (w *webSearchSession) Reset() {
	w.Session.Reset()
}

func (w *webSearchSession) next(values map[string]string) (*pipelinepb.Tracking, error) {
	text, ok := values[w.queryLabel]
	if !ok {
		w.Reset()
		return w.Session.next(values)
	}

	if text != w.lastQuery {
		w.Reset()
	}
	return w.Session.next(values)
}

// TrackingType defines different modes of tracking which can be applied to query requests.
type TrackingType string

// TrackingType constants.
const (
	TrackingNone   TrackingType = ""        // No tracking is enabled.
	TrackingClick  TrackingType = "CLICK"   // Click tracking is enabled, Click tokens will be returned with results.
	TrackingPosNeg TrackingType = "POS_NEG" // Positive/negative interaction tokens should be returned with results.
)

func (t TrackingType) proto() (pipelinepb.Tracking_Type, error) {
	switch t {
	case TrackingNone:
		return pipelinepb.Tracking_NONE, nil

	case TrackingClick:
		return pipelinepb.Tracking_CLICK, nil

	case TrackingPosNeg:
		return pipelinepb.Tracking_POS_NEG, nil
	}
	return pipelinepb.Tracking_NONE, fmt.Errorf("unknown TrackingType: %v", t)
}

// Session is an interface which defines session handling for search.
type Session interface {
	// Reset the session.
	Reset()

	// next returns the next tracking data to be used in the query.
	next(values map[string]string) (*pipelinepb.Tracking, error)
}

// NewSession creates a Session which generates tracking information for
// performing queries within a search.
func NewSession(ty TrackingType, field string, data map[string]string) Session {
	return &session{
		trackingType: ty,
		field:        field,
		data:         data,
	}
}

// session creates tracking information for performing queries.
type session struct {
	queryID  string
	sequence int

	trackingType TrackingType
	field        string
	data         map[string]string
}

var queryIDChars = []byte("abcdefghijklmnopqrstuvwxyz0123456789")

// randString constructs a random string of 16 characters.
// This implementation is designed to mirror the one used in the
// JS SDK.
func randString() string {
	buf := make([]byte, 16)
	for i := 0; i < len(buf); i++ {
		buf[i] = queryIDChars[rand.Intn(len(queryIDChars))]
	}
	return string(buf)
}

// Reset clears the query session.
func (t *session) Reset() {
	t.queryID = ""
}

// Next implements Session.
func (t *session) next(_ map[string]string) (*pipelinepb.Tracking, error) {
	if t.queryID == "" {
		t.queryID = randString() // match JS
		t.sequence = 0
	} else {
		t.sequence++
	}

	pbType, err := t.trackingType.proto()
	if err != nil {
		return nil, err
	}

	return &pipelinepb.Tracking{
		Type:     pbType,
		QueryId:  t.queryID,
		Sequence: int32(t.sequence),
		Field:    t.field,
		Data:     t.data,
	}, nil
}
