package lambda

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/pocockn/models/sns"
)

type (
	// ImageSimilarity is the payload sent when 2 successful shout uploads happen.
	ImageSimilarity struct {
		Source  string `json:"source"`
		Target  string `json:"target"`
		ShoutID string `json:"shout_id"`
	}
)

func NewRawJsonImageSimilarity(source string, target string, shoutID string) (json.RawMessage, error) {
	imageSimilartity := ImageSimilarity{
		Source:  source,
		Target:  target,
		ShoutID: shoutID,
	}

	payloadBytes, err := json.Marshal(&imageSimilartity)
	if err != nil {
		return nil, nil
	}

	return json.RawMessage(payloadBytes), nil
}

func NewImageSimilarityFromSNSPayload(data json.RawMessage) (ImageSimilarity, error) {
	var imageSimilarity ImageSimilarity
	notification, err := sns.NewNotification(data)
	if err != nil {
		return imageSimilarity, errors.Wrapf(err, "problem creating new notification from data %+v", data)
	}

	err = json.Unmarshal(json.RawMessage(notification.Message), &imageSimilarity)
	if err != nil {
		return imageSimilarity, errors.Wrapf(
			err, "problem unmarshalling notificaiton message into image simialrity struct %+v",
			notification.Message,
		)
	}

	return imageSimilarity, nil
}
