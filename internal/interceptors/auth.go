package interceptors

import (
	"context"

	ch "github.com/0loff/gophkeeper_server/internal/context_helpers"
	"github.com/0loff/gophkeeper_server/internal/logger"
	"github.com/0loff/gophkeeper_server/pkg/jwt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var SkipMethodsList = map[string]struct{}{
	"/gophkeeper.Gophkeeper/UserAuth":  {},
	"/gophkeeper.Gophkeeper/UserLogin": {},
}

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if _, ok := SkipMethodsList[info.FullMethod]; ok {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("token")) == 0 {
		return "", status.Errorf(codes.PermissionDenied, "Unrecognized user")
	}

	token := md.Get("token")[0]
	userID, err := jwt.ParseToken(token)
	if err != nil {
		logger.Log.Error("Invalid user id.", zap.Error(err))
		return "", status.Errorf(codes.PermissionDenied, "Unrecognized user")
	}

	ctx = context.WithValue(ctx, ch.ContextKeyUserID, userID)

	return handler(ctx, req)
}
