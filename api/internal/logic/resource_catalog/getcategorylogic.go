package resource_catalog

import (
	"context"

	"idrm/api/internal/svc"
	"idrm/api/internal/types"
	"idrm/pkg/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryLogic {
	return &GetCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryLogic) GetCategory(req *types.CategoryReq) (resp *types.CategoryResp, err error) {
	// 调用model层（无论是sqlx还是gorm，接口一致）
	category, err := l.svcCtx.CategoryModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, errorx.NewWithMsg(404, "类别不存在")
	}

	return &types.CategoryResp{
		Id:          category.Id,
		Name:        category.Name,
		Code:        category.Code,
		ParentId:    category.ParentId,
		Level:       category.Level,
		Sort:        category.Sort,
		Description: category.Description,
		Status:      category.Status,
	}, nil
}
