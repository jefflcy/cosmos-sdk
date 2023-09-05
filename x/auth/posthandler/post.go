package posthandler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// HandlerOptions are the options required for constructing a default SDK PostHandler.
type HandlerOptions struct{}

<<<<<<< HEAD
// NewPostHandler returns an empty PostHandler chain.
func NewPostHandler(_ HandlerOptions) (sdk.PostHandler, error) {
	postDecorators := []sdk.PostDecorator{}

	return sdk.ChainPostDecorators(postDecorators...), nil
=======
// NewPostHandler returns an empty posthandler chain.
func NewPostHandler(options HandlerOptions) (sdk.AnteHandler, error) {
	postDecorators := []sdk.AnteDecorator{}

	return sdk.ChainAnteDecorators(postDecorators...), nil
>>>>>>> v0.46.13-patch
}
