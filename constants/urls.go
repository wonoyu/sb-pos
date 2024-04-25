package constants

const (
	Login            = "/login"
	register         = "/register"
	RegisterAdmin    = register + "/admin"
	RegisterCashier  = register + "/cashier"
	RegisterCustomer = register + "/customer"
)

const (
	Users          = "/users"
	UserById       = "/users/:id"
	UpdateUserRole = "/users/update_role/:id"
)

const (
	Roles    = "/roles"
	RoleById = "/roles/:id"
)

const (
	Products     = "/products"
	ProductById  = "/products/:id"
	Categories   = "/products/categories"
	CategoryById = "/products/categories/:id"
)

const (
	Coupons    = "/coupons"
	CouponById = "/coupons/:id"
)

const (
	OrderStatus     = "/order_status"
	OrderStatusById = "/order_status/:id"
	Orders          = "/orders"
	OrderById       = "/orders/:id"
	OrdersPay       = OrderById + "/pay"
	OrdersComplete  = OrderById + "/complete"
)

const (
	TransactionType      = "/transaction_type"
	TransactionTypeById  = "/transaction_type/:id"
	Transactions         = "/transactions"
	TransactionById      = "/transactions/:id"
	SalesTransactions    = Transactions + "/sales"
	SalesTransactionById = Transactions + "/sales/:id"
)

const (
	Customers       = "/customers"
	CustomerById    = "/customers/:id"
	CustomerProfile = Customers + "/profile"
	CustomerTopup   = Customers + "/topup"
)
