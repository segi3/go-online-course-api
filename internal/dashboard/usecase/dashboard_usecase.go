package dashboard

import (
	adminUseCase "online-course/internal/admin/usecase"
	dto "online-course/internal/dashboard/dto"
	orderUseCase "online-course/internal/order/usecase"
	productUseCase "online-course/internal/product/usecase"
	userUseCase "online-course/internal/user/usecase"
)

type DashboardUseCase interface {
	GetDataDashboard() dto.DashboardResponseBody
}

type DashboardUseCaseImpl struct {
	userUseCase    userUseCase.UserUseCase
	adminUseCase   adminUseCase.AdminUseCase
	orderUseCase   orderUseCase.OrderUseCase
	productUseCase productUseCase.ProductUseCase
}

// GetDataDashboard implements DashboardUseCase
func (usecase *DashboardUseCaseImpl) GetDataDashboard() dto.DashboardResponseBody {
	totalUser := usecase.userUseCase.Count()
	totalAdmin := usecase.adminUseCase.Count()
	totalOrder := usecase.orderUseCase.Count()
	totalProduct := usecase.productUseCase.Count()

	return dto.DashboardResponseBody{
		TotalUser:    int64(totalUser),
		TotalProduct: int64(totalAdmin),
		TotalOrder:   int64(totalOrder),
		TotalAdmin:   int64(totalProduct),
	}
}

func NewDashboardUseCase(
	userUseCase userUseCase.UserUseCase,
	adminUseCase adminUseCase.AdminUseCase,
	orderUseCase orderUseCase.OrderUseCase,
	productUseCase productUseCase.ProductUseCase,
) DashboardUseCase {
	return &DashboardUseCaseImpl{userUseCase, adminUseCase, orderUseCase, productUseCase}
}
