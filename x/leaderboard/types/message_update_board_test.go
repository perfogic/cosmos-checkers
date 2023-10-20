package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/perfogic/cosmos-checkers/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgUpdateBoard_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateBoard
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateBoard{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateBoard{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
