package ante_test

import (
	"testing"

<<<<<<< HEAD
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
=======
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
>>>>>>> v0.46.13-patch

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/testutil"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/keeper"

	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	antetestutil "github.com/cosmos/cosmos-sdk/x/auth/ante/testutil"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtestutil "github.com/cosmos/cosmos-sdk/x/auth/testutil"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
<<<<<<< HEAD
=======
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
>>>>>>> v0.46.13-patch
)

// TestAccount represents an account used in the tests in x/auth/ante.
type TestAccount struct {
	acc  types.AccountI
	priv cryptotypes.PrivKey
}

// AnteTestSuite is a test suite to be used with ante handler tests.
type AnteTestSuite struct {
	anteHandler    sdk.AnteHandler
	ctx            sdk.Context
	clientCtx      client.Context
	txBuilder      client.TxBuilder
	accountKeeper  keeper.AccountKeeper
	bankKeeper     *authtestutil.MockBankKeeper
	feeGrantKeeper *antetestutil.MockFeegrantKeeper
	encCfg         moduletestutil.TestEncodingConfig
}

func TestAnteTestSuite(t *testing.T) {
	suite.Run(t, new(AnteTestSuite))
}

// SetupTest setups a new test, with new app, context, and anteHandler.
<<<<<<< HEAD
func SetupTestSuite(t *testing.T, isCheckTx bool) *AnteTestSuite {
	suite := &AnteTestSuite{}
	ctrl := gomock.NewController(t)
	suite.bankKeeper = authtestutil.NewMockBankKeeper(ctrl)

	suite.feeGrantKeeper = antetestutil.NewMockFeegrantKeeper(ctrl)

	key := sdk.NewKVStoreKey(types.StoreKey)
	testCtx := testutil.DefaultContextWithDB(t, key, sdk.NewTransientStoreKey("transient_test"))
	suite.ctx = testCtx.Ctx.WithIsCheckTx(isCheckTx).WithBlockHeight(1) // app.BaseApp.NewContext(isCheckTx, tmproto.Header{}).WithBlockHeight(1)
	suite.encCfg = moduletestutil.MakeTestEncodingConfig(auth.AppModuleBasic{})

	maccPerms := map[string][]string{
		"fee_collector":          nil,
		"mint":                   {"minter"},
		"bonded_tokens_pool":     {"burner", "staking"},
		"not_bonded_tokens_pool": {"burner", "staking"},
		"multiPerm":              {"burner", "minter", "staking"},
		"random":                 {"random"},
	}

	suite.accountKeeper = keeper.NewAccountKeeper(
		suite.encCfg.Codec, key, types.ProtoBaseAccount, maccPerms, sdk.Bech32MainPrefix, types.NewModuleAddress("gov").String(),
	)
	suite.accountKeeper.GetModuleAccount(suite.ctx, types.FeeCollectorName)
	err := suite.accountKeeper.SetParams(suite.ctx, types.DefaultParams())
	require.NoError(t, err)
=======
func (s *AnteTestSuite) SetupTest(isCheckTx bool) {
	s.app, s.ctx = createTestApp(s.T(), isCheckTx)
	s.ctx = s.ctx.WithBlockHeight(1)
>>>>>>> v0.46.13-patch

	// We're using TestMsg encoding in some tests, so register it here.
	suite.encCfg.Amino.RegisterConcrete(&testdata.TestMsg{}, "testdata.TestMsg", nil)
	testdata.RegisterInterfaces(suite.encCfg.InterfaceRegistry)

<<<<<<< HEAD
	suite.clientCtx = client.Context{}.
		WithTxConfig(suite.encCfg.TxConfig)

	anteHandler, err := ante.NewAnteHandler(
		ante.HandlerOptions{
			AccountKeeper:   suite.accountKeeper,
			BankKeeper:      suite.bankKeeper,
			FeegrantKeeper:  suite.feeGrantKeeper,
			SignModeHandler: suite.encCfg.TxConfig.SignModeHandler(),
=======
	s.clientCtx = client.Context{}.
		WithTxConfig(encodingConfig.TxConfig)

	anteHandler, err := ante.NewAnteHandler(
		ante.HandlerOptions{
			AccountKeeper:   s.app.AccountKeeper,
			BankKeeper:      s.app.BankKeeper,
			FeegrantKeeper:  s.app.FeeGrantKeeper,
			SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
>>>>>>> v0.46.13-patch
			SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
		},
	)

<<<<<<< HEAD
	require.NoError(t, err)
	suite.anteHandler = anteHandler

	suite.txBuilder = suite.clientCtx.TxConfig.NewTxBuilder()

	return suite
}

func (suite *AnteTestSuite) CreateTestAccounts(numAccs int) []TestAccount {
=======
	s.Require().NoError(err)
	s.anteHandler = anteHandler
}

// CreateTestAccounts creates `numAccs` accounts, and return all relevant
// information about them including their private keys.
func (s *AnteTestSuite) CreateTestAccounts(numAccs int) []TestAccount {
>>>>>>> v0.46.13-patch
	var accounts []TestAccount

	for i := 0; i < numAccs; i++ {
		priv, _, addr := testdata.KeyTestPubAddr()
<<<<<<< HEAD
		acc := suite.accountKeeper.NewAccountWithAddress(suite.ctx, addr)
		acc.SetAccountNumber(uint64(i))
		suite.accountKeeper.SetAccount(suite.ctx, acc)
=======
		acc := s.app.AccountKeeper.NewAccountWithAddress(s.ctx, addr)
		err := acc.SetAccountNumber(uint64(i))
		s.Require().NoError(err)
		s.app.AccountKeeper.SetAccount(s.ctx, acc)
		someCoins := sdk.Coins{
			sdk.NewInt64Coin("atom", 10000000),
		}
		err = s.app.BankKeeper.MintCoins(s.ctx, minttypes.ModuleName, someCoins)
		s.Require().NoError(err)

		err = s.app.BankKeeper.SendCoinsFromModuleToAccount(s.ctx, minttypes.ModuleName, addr, someCoins)
		s.Require().NoError(err)

>>>>>>> v0.46.13-patch
		accounts = append(accounts, TestAccount{acc, priv})
	}

	return accounts
}

// TestCase represents a test case used in test tables.
type TestCase struct {
	desc     string
	malleate func(*AnteTestSuite) TestCaseArgs
	simulate bool
	expPass  bool
	expErr   error
}

type TestCaseArgs struct {
	chainID   string
	accNums   []uint64
	accSeqs   []uint64
	feeAmount sdk.Coins
	gasLimit  uint64
	msgs      []sdk.Msg
	privs     []cryptotypes.PrivKey
}

// DeliverMsgs constructs a tx and runs it through the ante handler. This is used to set the context for a test case, for
// example to test for replay protection.
func (suite *AnteTestSuite) DeliverMsgs(t *testing.T, privs []cryptotypes.PrivKey, msgs []sdk.Msg, feeAmount sdk.Coins, gasLimit uint64, accNums, accSeqs []uint64, chainID string, simulate bool) (sdk.Context, error) {
	require.NoError(t, suite.txBuilder.SetMsgs(msgs...))
	suite.txBuilder.SetFeeAmount(feeAmount)
	suite.txBuilder.SetGasLimit(gasLimit)

	tx, txErr := suite.CreateTestTx(privs, accNums, accSeqs, chainID)
	require.NoError(t, txErr)
	return suite.anteHandler(suite.ctx, tx, simulate)
}

func (suite *AnteTestSuite) RunTestCase(t *testing.T, tc TestCase, args TestCaseArgs) {
	require.NoError(t, suite.txBuilder.SetMsgs(args.msgs...))
	suite.txBuilder.SetFeeAmount(args.feeAmount)
	suite.txBuilder.SetGasLimit(args.gasLimit)

	// Theoretically speaking, ante handler unit tests should only test
	// ante handlers, but here we sometimes also test the tx creation
	// process.
	tx, txErr := suite.CreateTestTx(args.privs, args.accNums, args.accSeqs, args.chainID)
	newCtx, anteErr := suite.anteHandler(suite.ctx, tx, tc.simulate)

	if tc.expPass {
		require.NoError(t, txErr)
		require.NoError(t, anteErr)
		require.NotNil(t, newCtx)

		suite.ctx = newCtx
	} else {
		switch {
		case txErr != nil:
			require.Error(t, txErr)
			require.ErrorIs(t, txErr, tc.expErr)

		case anteErr != nil:
			require.Error(t, anteErr)
			require.ErrorIs(t, anteErr, tc.expErr)

		default:
			t.Fatal("expected one of txErr, anteErr to be an error")
		}
	}
}

// CreateTestTx is a helper function to create a tx given multiple inputs.
func (s *AnteTestSuite) CreateTestTx(privs []cryptotypes.PrivKey, accNums []uint64, accSeqs []uint64, chainID string) (xauthsigning.Tx, error) {
	// First round: we gather all the signer infos. We use the "set empty
	// signature" hack to do that.
	var sigsV2 []signing.SignatureV2
	for i, priv := range privs {
		sigV2 := signing.SignatureV2{
			PubKey: priv.PubKey(),
			Data: &signing.SingleSignatureData{
				SignMode:  s.clientCtx.TxConfig.SignModeHandler().DefaultMode(),
				Signature: nil,
			},
			Sequence: accSeqs[i],
		}

		sigsV2 = append(sigsV2, sigV2)
	}
	err := s.txBuilder.SetSignatures(sigsV2...)
	if err != nil {
		return nil, err
	}

	// Second round: all signer infos are set, so each signer can sign.
	sigsV2 = []signing.SignatureV2{}
	for i, priv := range privs {
		signerData := xauthsigning.SignerData{
			ChainID:       chainID,
			AccountNumber: accNums[i],
			Sequence:      accSeqs[i],
		}
		sigV2, err := tx.SignWithPrivKey(
			s.clientCtx.TxConfig.SignModeHandler().DefaultMode(), signerData,
			s.txBuilder, priv, s.clientCtx.TxConfig, accSeqs[i])
		if err != nil {
			return nil, err
		}

		sigsV2 = append(sigsV2, sigV2)
	}
	err = s.txBuilder.SetSignatures(sigsV2...)
	if err != nil {
		return nil, err
	}

	return s.txBuilder.GetTx(), nil
}
<<<<<<< HEAD
=======

// TestCase represents a test case used in test tables.
type TestCase struct {
	desc     string
	malleate func()
	simulate bool
	expPass  bool
	expErr   error
}

// CreateTestTx is a helper function to create a tx given multiple inputs.
func (s *AnteTestSuite) RunTestCase(privs []cryptotypes.PrivKey, msgs []sdk.Msg, feeAmount sdk.Coins, gasLimit uint64, accNums, accSeqs []uint64, chainID string, tc TestCase) {
	s.Run(fmt.Sprintf("Case %s", tc.desc), func() {
		s.Require().NoError(s.txBuilder.SetMsgs(msgs...))
		s.txBuilder.SetFeeAmount(feeAmount)
		s.txBuilder.SetGasLimit(gasLimit)

		// Theoretically speaking, ante handler unit tests should only test
		// ante handlers, but here we sometimes also test the tx creation
		// process.
		tx, txErr := s.CreateTestTx(privs, accNums, accSeqs, chainID)
		newCtx, anteErr := s.anteHandler(s.ctx, tx, tc.simulate)

		if tc.expPass {
			s.Require().NoError(txErr)
			s.Require().NoError(anteErr)
			s.Require().NotNil(newCtx)

			s.ctx = newCtx
		} else {
			switch {
			case txErr != nil:
				s.Require().Error(txErr)
				s.Require().True(errors.Is(txErr, tc.expErr))

			case anteErr != nil:
				s.Require().Error(anteErr)
				s.Require().True(errors.Is(anteErr, tc.expErr))

			default:
				s.Fail("expected one of txErr,anteErr to be an error")
			}
		}
	})
}
>>>>>>> v0.46.13-patch
