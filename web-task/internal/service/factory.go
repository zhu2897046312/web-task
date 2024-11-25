package service

type ServiceFactory struct {
    base *Service
}

func NewServiceFactory(base *Service) *ServiceFactory {
    return &ServiceFactory{base: base}
}

func (f *ServiceFactory) GetUserService() *UserService {
    return NewUserService(f.base)
}

func (f *ServiceFactory) GetProductService() *ProductService {
    return NewProductService(f.base)
}

func (f *ServiceFactory) GetOrderService() *OrderService {
    return NewOrderService(f.base)
}

func (f *ServiceFactory) GetCartService() *CartService {
    return NewCartService(f.base)
}

func (f *ServiceFactory) GetAdvertisementService() *AdvertisementService {
    return NewAdvertisementService(f.base)
}

func (f *ServiceFactory) GetReviewService() *ReviewService {
    return NewReviewService(f.base)
} 