package zoom

import "fmt"

// UpdateMeetingOptions are the options to update a meeting with
type UpdateMeetingOptions struct {
	MeetingID      int
	DataParameters UpdateMeetingDataParameters
	URLParameters  UpdateMeetingURLParameters
}

type UpdateMeetingDataParameters struct {
	Agenda         string          `json:"agenda,omitempty"`
	Duration       int             `json:"duration,omitempty"`
	Password       string          `json:"password,omitempty"` // Max 10 characters. [a-z A-Z 0-9 @ - _ *]
	Settings       MeetingSettings `json:"settings,omitempty"`
	StartTime      *Time           `json:"start_time,omitempty"`
	Timezone       string          `json:"timezone,omitempty"`
	Topic          string          `json:"topic,omitempty"`
	TrackingFields []TrackingField `json:"tracking_fields,omitempty"`
	Type           MeetingType     `json:"type,omitempty"`
}

type UpdateMeetingURLParameters struct {
	OccurrenceID string `url:"occurrence_id,omitempty"`
}

// UpdateMeetingPath - v2 update a meeting
const UpdateMeetingPath = "/meetings/%d"

// UpdateMeeting calls PATCH /meetings/{meetingId}
func UpdateMeeting(opts UpdateMeetingOptions) (Meeting, error) {
	return defaultClient.UpdateMeeting(opts)
}

// UpdateMeeting calls PATCH /meetings/{meetingId}
// https://marketplace.zoom.us/docs/api-reference/zoom-api/meetings/meetingupdate
func (c *Client) UpdateMeeting(opts UpdateMeetingOptions) (Meeting, error) {
	var ret = Meeting{}
	return ret, c.requestV2(requestV2Opts{
		Method:         Patch,
		Path:           fmt.Sprintf(UpdateMeetingPath, opts.MeetingID),
		DataParameters: &opts.DataParameters,
		URLParameters:  &opts.URLParameters,
		Ret:            &ret,
	})
}
