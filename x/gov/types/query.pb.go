// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package types

import (
	context "context"
	types "github.com/cosmos/cosmos-sdk/types"
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Proposal queries proposal details based on ProposalID.
	Proposal(ctx context.Context, in *QueryProposalRequest, opts ...grpc.CallOption) (*QueryProposalResponse, error)
	// Proposals queries all proposals based on given status.
	Proposals(ctx context.Context, in *QueryProposalsRequest, opts ...grpc.CallOption) (*QueryProposalsResponse, error)
	// Vote queries voted information based on proposalID, voterAddr.
	Vote(ctx context.Context, in *QueryVoteRequest, opts ...grpc.CallOption) (*QueryVoteResponse, error)
	// Votes queries votes of a given proposal.
	Votes(ctx context.Context, in *QueryVotesRequest, opts ...grpc.CallOption) (*QueryVotesResponse, error)
	// Params queries all parameters of the gov module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Deposit queries single deposit information based proposalID, depositAddr.
	Deposit(ctx context.Context, in *QueryDepositRequest, opts ...grpc.CallOption) (*QueryDepositResponse, error)
	// Deposits queries all deposits of a single proposal.
	Deposits(ctx context.Context, in *QueryDepositsRequest, opts ...grpc.CallOption) (*QueryDepositsResponse, error)
	// TallyResult queries the tally of a proposal vote.
	TallyResult(ctx context.Context, in *QueryTallyResultRequest, opts ...grpc.CallOption) (*QueryTallyResultResponse, error)
}

type queryClient struct {
	cc           grpc.ClientConnInterface
	_Proposal    types.Invoker
	_Proposals   types.Invoker
	_Vote        types.Invoker
	_Votes       types.Invoker
	_Params      types.Invoker
	_Deposit     types.Invoker
	_Deposits    types.Invoker
	_TallyResult types.Invoker
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc: cc}
}

func (c *queryClient) Proposal(ctx context.Context, in *QueryProposalRequest, opts ...grpc.CallOption) (*QueryProposalResponse, error) {
	if invoker := c._Proposal; invoker != nil {
		var out QueryProposalResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._Proposal, err = invokerConn.Invoker("/cosmos.gov.v1beta1.QueryProposal")
		if err != nil {
			var out QueryProposalResponse
			err = c._Proposal(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryProposalResponse)
	err := c.cc.Invoke(ctx, "/cosmos.gov.v1beta1.QueryProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Proposals(ctx context.Context, in *QueryProposalsRequest, opts ...grpc.CallOption) (*QueryProposalsResponse, error) {
	if invoker := c._Proposals; invoker != nil {
		var out QueryProposalsResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._Proposals, err = invokerConn.Invoker("/cosmos.gov.v1beta1.QueryProposals")
		if err != nil {
			var out QueryProposalsResponse
			err = c._Proposals(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryProposalsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.gov.v1beta1.QueryProposals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Vote(ctx context.Context, in *QueryVoteRequest, opts ...grpc.CallOption) (*QueryVoteResponse, error) {
	if invoker := c._Vote; invoker != nil {
		var out QueryVoteResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._Vote, err = invokerConn.Invoker("/cosmos.gov.v1beta1.QueryVote")
		if err != nil {
			var out QueryVoteResponse
			err = c._Vote(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryVoteResponse)
	err := c.cc.Invoke(ctx, "/cosmos.gov.v1beta1.QueryVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Votes(ctx context.Context, in *QueryVotesRequest, opts ...grpc.CallOption) (*QueryVotesResponse, error) {
	if invoker := c._Votes; invoker != nil {
		var out QueryVotesResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._Votes, err = invokerConn.Invoker("/cosmos.gov.v1beta1.QueryVotes")
		if err != nil {
			var out QueryVotesResponse
			err = c._Votes(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryVotesResponse)
	err := c.cc.Invoke(ctx, "/cosmos.gov.v1beta1.QueryVotes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	if invoker := c._Params; invoker != nil {
		var out QueryParamsResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._Params, err = invokerConn.Invoker("/cosmos.gov.v1beta1.QueryParams")
		if err != nil {
			var out QueryParamsResponse
			err = c._Params(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.gov.v1beta1.QueryParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Deposit(ctx context.Context, in *QueryDepositRequest, opts ...grpc.CallOption) (*QueryDepositResponse, error) {
	if invoker := c._Deposit; invoker != nil {
		var out QueryDepositResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._Deposit, err = invokerConn.Invoker("/cosmos.gov.v1beta1.QueryDeposit")
		if err != nil {
			var out QueryDepositResponse
			err = c._Deposit(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryDepositResponse)
	err := c.cc.Invoke(ctx, "/cosmos.gov.v1beta1.QueryDeposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Deposits(ctx context.Context, in *QueryDepositsRequest, opts ...grpc.CallOption) (*QueryDepositsResponse, error) {
	if invoker := c._Deposits; invoker != nil {
		var out QueryDepositsResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._Deposits, err = invokerConn.Invoker("/cosmos.gov.v1beta1.QueryDeposits")
		if err != nil {
			var out QueryDepositsResponse
			err = c._Deposits(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryDepositsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.gov.v1beta1.QueryDeposits", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TallyResult(ctx context.Context, in *QueryTallyResultRequest, opts ...grpc.CallOption) (*QueryTallyResultResponse, error) {
	if invoker := c._TallyResult; invoker != nil {
		var out QueryTallyResultResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._TallyResult, err = invokerConn.Invoker("/cosmos.gov.v1beta1.QueryTallyResult")
		if err != nil {
			var out QueryTallyResultResponse
			err = c._TallyResult(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryTallyResultResponse)
	err := c.cc.Invoke(ctx, "/cosmos.gov.v1beta1.QueryTallyResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Proposal queries proposal details based on ProposalID.
	Proposal(types.Context, *QueryProposalRequest) (*QueryProposalResponse, error)
	// Proposals queries all proposals based on given status.
	Proposals(types.Context, *QueryProposalsRequest) (*QueryProposalsResponse, error)
	// Vote queries voted information based on proposalID, voterAddr.
	Vote(types.Context, *QueryVoteRequest) (*QueryVoteResponse, error)
	// Votes queries votes of a given proposal.
	Votes(types.Context, *QueryVotesRequest) (*QueryVotesResponse, error)
	// Params queries all parameters of the gov module.
	Params(types.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Deposit queries single deposit information based proposalID, depositAddr.
	Deposit(types.Context, *QueryDepositRequest) (*QueryDepositResponse, error)
	// Deposits queries all deposits of a single proposal.
	Deposits(types.Context, *QueryDepositsRequest) (*QueryDepositsResponse, error)
	// TallyResult queries the tally of a proposal vote.
	TallyResult(types.Context, *QueryTallyResultRequest) (*QueryTallyResultResponse, error)
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Proposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProposalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Proposal(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.gov.v1beta1.QueryProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Proposal(types.UnwrapSDKContext(ctx), req.(*QueryProposalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Proposals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProposalsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Proposals(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.gov.v1beta1.QueryProposals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Proposals(types.UnwrapSDKContext(ctx), req.(*QueryProposalsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Vote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Vote(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.gov.v1beta1.QueryVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Vote(types.UnwrapSDKContext(ctx), req.(*QueryVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Votes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryVotesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Votes(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.gov.v1beta1.QueryVotes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Votes(types.UnwrapSDKContext(ctx), req.(*QueryVotesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.gov.v1beta1.QueryParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(types.UnwrapSDKContext(ctx), req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Deposit(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.gov.v1beta1.QueryDeposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Deposit(types.UnwrapSDKContext(ctx), req.(*QueryDepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Deposits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDepositsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Deposits(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.gov.v1beta1.QueryDeposits",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Deposits(types.UnwrapSDKContext(ctx), req.(*QueryDepositsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TallyResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTallyResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TallyResult(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.gov.v1beta1.QueryTallyResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TallyResult(types.UnwrapSDKContext(ctx), req.(*QueryTallyResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.gov.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Proposal",
			Handler:    _Query_Proposal_Handler,
		},
		{
			MethodName: "Proposals",
			Handler:    _Query_Proposals_Handler,
		},
		{
			MethodName: "Vote",
			Handler:    _Query_Vote_Handler,
		},
		{
			MethodName: "Votes",
			Handler:    _Query_Votes_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Deposit",
			Handler:    _Query_Deposit_Handler,
		},
		{
			MethodName: "Deposits",
			Handler:    _Query_Deposits_Handler,
		},
		{
			MethodName: "TallyResult",
			Handler:    _Query_TallyResult_Handler,
		},
	},
	Metadata: "cosmos/gov/v1beta1/query.proto",
}

const (
	QueryProposalMethod    = "/cosmos.gov.v1beta1.QueryProposal"
	QueryProposalsMethod   = "/cosmos.gov.v1beta1.QueryProposals"
	QueryVoteMethod        = "/cosmos.gov.v1beta1.QueryVote"
	QueryVotesMethod       = "/cosmos.gov.v1beta1.QueryVotes"
	QueryParamsMethod      = "/cosmos.gov.v1beta1.QueryParams"
	QueryDepositMethod     = "/cosmos.gov.v1beta1.QueryDeposit"
	QueryDepositsMethod    = "/cosmos.gov.v1beta1.QueryDeposits"
	QueryTallyResultMethod = "/cosmos.gov.v1beta1.QueryTallyResult"
)
