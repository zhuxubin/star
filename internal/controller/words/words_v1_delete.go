package words

import (
	"context"

	"star/internal/logic/users"
	"star/internal/logic/words"

	"star/api/words/v1"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	uid, err := users.GetUid(ctx)
	if err != nil {
		return nil, err
	}

	err = words.Delete(ctx, uid, req.Id)
	return
}
