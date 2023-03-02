package rest

import (
	"context"
	"mime/multipart"

	"github.com/cortezaproject/corteza/server/compose/rest/request"
	"github.com/cortezaproject/corteza/server/compose/service"
	"github.com/cortezaproject/corteza/server/compose/types"
	"github.com/cortezaproject/corteza/server/pkg/filter"
)

type (
	PageIcon struct {
		locale     service.ResourceTranslationsManagerService
		attachment service.AttachmentService
		ac         pageIconAccessController
	}

	pageIconAccessController interface {
		CanGrant(context.Context) bool
	}
)

func (PageIcon) New() *PageIcon {
	return &PageIcon{
		locale:     service.DefaultResourceTranslation,
		attachment: service.DefaultAttachment,
		ac:         service.DefaultAccessControl,
	}
}

func (ctrl *PageIcon) List(ctx context.Context, r *request.PageIconList) (interface{}, error) {
	var (
		err error
		f   = types.AttachmentFilter{
			Kind: types.PageIconAttachment,
		}
	)

	if f.Paging, err = filter.NewPaging(r.Limit, r.PageCursor); err != nil {
		return nil, err
	}

	if f.Sorting, err = filter.NewSorting(r.Sort); err != nil {
		return nil, err
	}

	set, f, err := ctrl.attachment.Find(ctx, f)
	return ctrl.makeIconFilterPayload(ctx, set, f, err)
}

func (ctrl *PageIcon) Upload(ctx context.Context, r *request.PageIconUpload) (interface{}, error) {
	file, err := r.Icon.Open()
	if err != nil {
		return nil, err
	}

	defer func(file multipart.File) {
		err = file.Close()
		if err != nil {
			return
		}
	}(file)

	a, err := ctrl.attachment.CreatePageIconAttachment(
		ctx,
		r.Icon.Filename,
		r.Icon.Size,
		file,
	)

	return makeAttachmentPayload(ctx, a, err)
}

func (ctrl *PageIcon) makeIconFilterPayload(ctx context.Context, nn types.AttachmentSet, f types.AttachmentFilter, err error) (*attachmentSetPayload, error) {
	if err != nil {
		return nil, err
	}

	res := &attachmentSetPayload{Filter: f, Set: make([]*attachmentPayload, len(nn))}

	for i := range nn {
		res.Set[i], _ = makeAttachmentPayload(ctx, nn[i], nil)
	}

	return res, nil
}
