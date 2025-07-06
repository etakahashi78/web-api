package main

//
//import (
//	"context"
//	"errors"
//	"net/http"
//	"testing"
//	"time"
//
//	"github.com/stretchr/testify/assert"
//
//	"web-api/config"
//	"web-api/infra"
//	"web-api/router"
//	"web-api/usecase/interactor"
//)
//
//// mockInfraRepositories は infra.Repositories のモック
//type mockInfraRepositories struct {
//	ProductRepository interactor.ProductRepository
//	ConnectError      error
//}
//
//func (m *mockInfraRepositories) NewMySQLProductRepository(dsn string) (interactor.ProductRepository, error) {
//	return m.ProductRepository, m.ConnectError
//}
//
//func (m *mockInfraRepositories) Close() {}
//
//// mockProductInteractor は interactor.ProductInteractor のモック
//type mockProductInteractor struct{}
//
//func (m *mockProductInteractor) RegisterProduct(name string, price float64) (*interactor.Product, error) {
//	return nil, nil
//}
//func (m *mockProductInteractor) ListProducts() ([]*interactor.Product, error) {
//	return nil, nil
//}
//
//// mockRouter は router.SetupRoutes のモック
//type mockRouter struct {
//	Handler http.Handler
//}
//
//func (m *mockRouter) SetupRoutes(productController interface{}) http.Handler {
//	return m.Handler
//}
//
//func TestRun(t *testing.T) {
//	testCases := []struct {
//		name        string
//		cfg         *config.Config
//		mockRepos   *mockInfraRepositories
//		mockRouter  *mockRouter
//		wantErr     bool
//		wantErrorIs error
//	}{
//		{
//			name: "Successful Server Start",
//			cfg: &config.Config{
//				DBDSNPrimary: "test_dsn",
//				DBDSNReplica: "test_dsn",
//			},
//			mockRepos: &mockInfraRepositories{
//				ConnectError: nil,
//			},
//			mockRouter: &mockRouter{
//				Handler: http.NewServeMux(), // 空の ServeMux で成功とみなす
//			},
//			wantErr: false,
//		},
//		{
//			name: "Failed Database Connection",
//			cfg: &config.Config{
//				DBDSNPrimary: "invalid_dsn",
//				DBDSNReplica: "invalid_dsn",
//			},
//			mockRepos: &mockInfraRepositories{
//				ConnectError: errors.New("database connection failed"),
//			},
//			mockRouter: &mockRouter{
//				Handler: http.NewServeMux(),
//			},
//			wantErr:     true,
//			wantErrorIs: errors.New("database connection failed"), // 期待するエラー
//		},
//		{
//			name: "Failed Server ListenAndServe",
//			cfg: &config.Config{
//				DBDSNPrimary: "test_dsn",
//				DBDSNReplica: "test_dsn",
//			},
//			mockRepos: &mockInfraRepositories{
//				ConnectError: nil,
//			},
//			mockRouter: &mockRouter{
//				Handler: &errorServer{}, // ListenAndServe が常にエラーを返すサーバーモック
//			},
//			wantErr:     true,
//			wantErrorIs: errors.New("failed to start server"), // run 関数内でラップされるため、部分一致で検証
//		},
//	}
//
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//			// infra.NewRepositories をモックで置き換える
//			originalNewRepositories := infra.NewRepositories
//			infra.NewRepositories = func(cfg *config.Config) *infra.Repositories {
//				return &infra.Repositories{
//					ProductRepository: tc.mockRepos.ProductRepository,
//					ConnectError:      tc.mockRepos.ConnectError,
//				}
//			}
//			defer func() {
//				infra.NewRepositories = originalNewRepositories // テスト後に元の関数に戻す
//			}()
//
//			// router.SetupRoutes をモックで置き換える
//			originalSetupRoutes := router.SetupRoutes
//			router.SetupRoutes = func(productController interface{}) http.Handler {
//				return tc.mockRouter.Handler
//			}
//			defer func() {
//				router.SetupRoutes = originalSetupRoutes // テスト後に元の関数に戻す
//			}()
//
//			err := run(tc.cfg) // 直接 run 関数を呼び出す
//
//			if tc.wantErr {
//				assert.Error(t, err, "TestRun(%s) expected error", tc.name)
//				if tc.wantErrorIs != nil {
//					assert.ErrorIs(t, err, tc.wantErrorIs, "TestRun(%s) error mismatch", tc.name)
//				}
//			} else {
//				assert.NoError(t, err, "TestRun(%s) unexpected error", tc.name)
//			}
//
//			// サーバーが起動した場合は、念のため少し待ってからシャットダウンを試みる (テストが無限に続くのを防ぐため)
//			if !tc.wantErr && tc.mockRouter.Handler != nil {
//				ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
//				defer cancel()
//				if srv, ok := tc.mockRouter.Handler.(*http.Server); ok {
//					srv.Shutdown(ctx)
//				}
//			}
//		})
//	}
//}
//
//// ListenAndServe が常にエラーを返す http.Handler のモック
//type errorServer struct{}
//
//func (e *errorServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	// Do nothing
//}
//
//func (e *errorServer) ListenAndServe(addr string, handler http.Handler) error {
//	return errors.New("failed to start server")
//}
//
//func (e *errorServer) Shutdown(ctx context.Context) error {
//	return nil
//}
