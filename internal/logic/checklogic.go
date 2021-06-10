package logic

import (
	"context"

	"github.com/lokichoggio/check/check"
	"github.com/lokichoggio/check/internal/common/errorx"
	"github.com/lokichoggio/check/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
	"google.golang.org/grpc/codes"
)

type CheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckLogic) Check(in *check.CheckReq) (*check.CheckResp, error) {
	// 手动代码开始
	resp, err := l.svcCtx.Model.FindOne(in.Book)
	if err != nil {
		return nil, errorx.CodeMsgErrorWithStack(codes.Internal, "mysql error", err)
	}

	return &check.CheckResp{
		Found: true,
		Price: resp.Price,
	}, nil
	// 手动代码结束
}
