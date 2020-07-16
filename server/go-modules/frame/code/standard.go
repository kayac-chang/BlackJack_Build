package code

import "google.golang.org/grpc/codes"

const (
	OK                 = Code(codes.OK)
	Canceled           = Code(codes.Canceled)
	Unknown            = Code(codes.Unknown)
	InvalidArgument    = Code(codes.InvalidArgument)
	DeadlineExceeded   = Code(codes.DeadlineExceeded)
	NotFound           = Code(codes.NotFound)
	AlreadyExists      = Code(codes.AlreadyExists)
	PermissionDenied   = Code(codes.PermissionDenied)
	ResourceExhausted  = Code(codes.ResourceExhausted)
	FailedPrecondition = Code(codes.FailedPrecondition)
	Aborted            = Code(codes.Aborted)
	OutOfRange         = Code(codes.OutOfRange)
	Unimplemented      = Code(codes.Unimplemented)
	Internal           = Code(codes.Internal)
	Unavailable        = Code(codes.Unavailable)
	DataLoss           = Code(codes.DataLoss)
	Unauthenticated    = Code(codes.Unauthenticated)

)
