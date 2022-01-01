package usecase

import (
	"github.com/afif0808/bobobox_test/models"
	"github.com/afif0808/bobobox_test/payloads"
	"github.com/afif0808/bobobox_test/pkg/structs"
)

func convertCreateStayPayloadToStayModel(ps []payloads.CreateStayPayload) ([]models.Stay, error) {
	var err error
	ss := make([]models.Stay, len(ps))
	for i := range ps {
		err = structs.Merge(&ss[i], ps[i])
		if err != nil {
			return nil, err
		}
	}
	return ss, err
}
