# SIMPLE DUMMY POS RESTful API USING GO

## FINAL PROJECT BOOTCAMP GO-LANG SANBERCODE

### USAGE
To use this app, simply open your favorite code editor that supports go environments such as vscode, goland, etc. Make sure you have install go in your computer if not please follow installation instructions here [Go](https://go.dev/dl/).
Then continue by opening the terminal and run 

"""
go run main.go
"""

If the app is run successfully, continue to the next step.

> [!NOTE]
> You can also open this link instead if you don't want to run it locally.
> https://sb-pos-production.up.railway.app/
> And make a request using Postman or other API tester app.

There are 3 roles of user in this app that is admin, customer and cashier, here are list of dummy accounts you can use
1. ADMIN
email: admin@mail.com
password: admin

2. CUSTOMER
email: customer@mail.com
password: customer

3. CASHIER
email: cashier@mail.com
password: cashier

```
You can use these credentials to get jwt token and have full access of the app, simply use `/login` url and provide request body in JSON
like shown below:

{
    "email": email,
    "password": password
}
```
```
If the credential is recognized, the API will give response like below:

{
    "token": token
}
```

You must put this token to your Authorization header to gain access of this app. In postman you can open `Authorization` tab and select `Bearer Token` authorization and put the token returned from login in the token field.

### API REFERENCES
1. REGISTER
    This endpoint requires request body in JSON Format like provided example below.

    ```
   **`/register/admin`** --> method POST
    This endpoint is used to create new user with admin role.
    ```

    ```
   **`/register/customer`** --> method POST
    Used to create new user with customer role.
    ```

    ```
   **`/register/cashier`** --> method POST
    Create new user with cashier role.
    ```

    ```
    {
        "username": "admin",
        "email": "admin@mail.com",
        "password": "admin"
    }
    ```
    
3. LOGIN
    This endpoint requires request body in JSON Format like provided example below.

    ```
   **`/login`** --> method POST
    Used to gain full access of this app, returns token after successfull api calls.
    ```

   ```
    {
        "email": "admin@mail.com",
        "password": "admin"
    }
    ```

    ```
   Sample response
    {
        "token": "token"
    }
    ```

5. USERS
    This endpoint requires token to access, the token itself is retrieved after successful login.

    ```
   **`/users`** --> method GET
    Used to get users registered in this app database.
    ```

    ```
   Sample response
    {
        "result": [
            {
                "id": 1,
                "username": "admin",
                "email": "email",
                "role_name": "rolename",
                "role_id": "roleid",
                "created_at": "createdat",
                "updated_at": "updatedat"
            }
        ]
    }
    ```

    ```
   **`/users/:id`** --> method GET
    Get user registered in this app using the registered id.
    ```

    ```
   Sample response
    {
        "result": {
                "id": 1,
                "username": "admin",
                "email": "email",
                "role_name": "rolename",
                "role_id": "roleid",
                "created_at": "createdat",
                "updated_at": "updatedat"
            }
    }
    ```

    ```
   **`/users/:id`** --> method PUT
    Update user registered in the database by referencing it's registered id. Requires request body as shown below, the field is not required though.
    ```

    ```
    {
        "username": "user",
        "email": "mail@mail.com"
    }
    ```

    ```
   Sample response
    {
        "message": "berhasil update user"
    }
    ```

    ```
   **`/users/update_role/:id`** --> method PUT
    Update user role registered in the database by referencing it's registered id. Requires request body as shown below, the field is not required though. Only admins can access this resource.
    ```

    ```
    {
        "role_id": 2
    }
    ```

    ```
   Sample response
    {
        "message": "berhasil update role user"
    }
    ```

    ```
   **`/users/:id`** --> method DELETE
    Deletes user from the database, requires id as it's parameter. Can only be done by admins.
    ```

    ```
   Sample response
    {
        "message": "berhasil delete user"
    }
    ```

7. ROLES
    This is a roles masterdata, only admins can have access for this API.

    ```
   **`/roles`** --> method GET
    Used to get roles stored in this app database.
    ```

    ```
   Sample response
    {
        "result": [
            {
                "id": 1,
                "name": "admin",
                "created_at": "createdat",
                "updated_at": "updatedat"
            }
        ]
    }
    ```

    ```
   **`/roles/:id`** --> method GET
    Used to get role by id stored in this app database.
    ```

    ```
   Sample response
    {
        "result": {
                "id": 1,
                "name": "admin",
                "created_at": "createdat",
                "updated_at": "updatedat"
            }
    }
    ```

    ```
   **`/roles`** --> method POST
    Used to create role for this app. Requires request body as shown below.
    ```

    ```
    {
        "name": "role"
    }
    ```

    ```
   Sample response
    {
        "result": "berhasil tambah role"
    }
    ```

    ```
   **`/roles/:id`** --> method PUT
    Used to update role by id stored in this app database. Sample request body is shown below.
    ```

    ```
    {
        "name": "newname"
    }
    ```

    ```
   Sample response
    {
        "message": "berhasil update role"
    }
    ```

    ```
   **`/roles/:id`** --> method DELETE
    Used to delete role by id stored in this app database. Requires id params.

    Sample response
    {
        "message": "berhasil delete role"
    }
    ```

9. PRODUCTS CATEGORY
    This is a product category masterdata, only admins can have manipulate this API. Other roles can get it's data.

    ```
   **`/products/categories`** --> method GET
    Used to get product categories stored in this app database.

    Sample response
    {
        "result": [
            {
                "id": 1,
                "name": "category",
                "created_at": "createdat",
                "updated_at": "updatedat"
            }
        ]
    }
    ```

    ```
   **`/products/categories/:id`** --> method GET
    Used to get product category by id stored in this app database.

    Sample response
    {
        "result": {
                "id": 1,
                "name": "category",
                "created_at": "createdat",
                "updated_at": "updatedat"
            }
    }
    ```

    ```
   **`/products/categories`** --> method POST
    Used to create product category for this app. Requires request body as shown below.

    {
        "name": "sembako"
    }

    Sample response
    {
        "result": "berhasil menambah product category"
    }
    ```

    ```
   **`/products/categories/:id`** --> method PUT
    Used to update product category by id stored in this app database. Sample request body is shown below.

    {
        "name": "newname"
    }

    Sample response
    {
        "message": "berhasil update product category"
    }
    ```

    ```
   **`/products/categories/:id`** --> method DELETE
    Used to delete product category by id stored in this app database. Requires id params.

    Sample response
    {
        "message": "berhasil delete product category"
    }
    ```

11. PRODUCTS
    This is a products masterdata, only admins can have manipulate this API. Other roles can get it's data.

    ```
    **`/products`** --> method GET
    Used to get products stored in this app database.

    Sample response
    {
        "result": [
            {
                "id": 1,
                "name": "category",
                "price": 20000,
                "stock_quantity": 10,
                "category_id": 1,
                "category_name": "name",
                "created_at": "createdat",
                "updated_at": "updatedat"
            }
        ]
    }
    ```

    ```
    **`/products/:id`** --> method GET
    Used to get product by id stored in this app database.

    Sample response
    {
        "result": {
                "id": 1,
                "name": "category",
                "price": 20000,
                "stock_quantity": 10,
                "category_id": 1,
                "category_name": "name",
                "created_at": "createdat",
                "updated_at": "updatedat"
            }
    }
    ```

    ```
    **`/products`** --> method POST
    Used to create product for this app. Requires request body as shown below.

    {
        "name": "bakso",
        "price": 200000,
        "stock_quantity": 20,
        "category_id": 1
    }

    Sample response
    {
        "result": "berhasil menambah product"
    }
    ```

    ```
    **`/products/:id`** --> method PUT
    Used to update product by id stored in this app database. Sample request body is shown below.

    {
        "name": "bakso",
        "price": 200000,
        "stock_quantity": 20,
        "category_id": 1
    }

    Sample response
    {
        "message": "berhasil update product"
    }
    ```

    ```
    **`/products/:id`** --> method DELETE
    Used to delete product by id stored in this app database. Requires id params.

    Sample response
    {
        "message": "berhasil delete product"
    }
    ```
    
13. ORDER STATUS
    This is a order status masterdata, only admins can have manipulate this API. Other roles can get it's data.

    ```
    **`/order_status`** --> method GET
    Used to get order status stored in this app database.

    Sample response
    {
        "result": [
            {
                "id": 1,
                "name": "paid",
                "created_at": "createdat",
                "updated_at": "updatedat"
            }
        ]
    }
    ```

    ```
    **`/order_status/:id`** --> method GET
    Used to get order status by id stored in this app database.

    Sample response
    {
        "result": {
                "id": 1,
                "name": "paid",
                "created_at": "createdat",
                "updated_at": "updatedat"
            }
    }
    ```

    ```
    **`/order_status`** --> method POST
    Used to create order status for this app. Requires request body as shown below.

    {
        "name": "name"
    }

    Sample response
    {
        "result": "berhasil menambah order status"
    }
    ```

    ```
    **`/order_status/:id`** --> method PUT
    Used to update order status by id stored in this app database. Sample request body is shown below.

    {
        "name": "newname"
    }

    Sample response
    {
        "message": "berhasil update order status"
    }
    ```

    ```
    **`/order_status/:id`** --> method DELETE
    Used to delete order status by id stored in this app database. Requires id params.

    Sample response
    {
        "message": "berhasil delete order status"
    }
    ```

15. TRANSACTION TYPE
    This is a transaction type masterdata, only admins can have manipulate this API. Other roles can get it's data.

    ```
    **`/transaction_type`** --> method GET
    Used to get transaction type stored in this app database.

    Sample response
    {
        "result": [
            {
                "id": 1,
                "name": "sale",
                "created_at": "createdat",
                "updated_at": "updatedat"
            }
        ]
    }
    ```

    ```
    **`/transaction_type/:id`** --> method GET
    Used to get transaction type by id stored in this app database.

    Sample response
    {
        "result": {
                "id": 1,
                "name": "sale",
                "created_at": "createdat",
                "updated_at": "updatedat"
            }
    }
    ```

    ```
    **`/transaction_type`** --> method POST
    Used to create transaction type for this app. Requires request body as shown below.

    {
        "name": "name"
    }

    Sample response
    {
        "result": "berhasil menambah transaction type"
    }
    ```

    ```
    **`/transaction_type/:id`** --> method PUT
    Used to update transaction type by id stored in this app database. Sample request body is shown below.

    {
        "name": "newname"
    }

    Sample response
    {
        "message": "berhasil update transaction type"
    }
    ```

    ```
    **`/transaction_type/:id`** --> method DELETE
    Used to delete transaction type by id stored in this app database. Requires id params.

    Sample response
    {
        "message": "berhasil delete transaction type"
    }
    ```

17. CUSTOMERS
    Endpoints for user with customer role to view their profile and balance and also topup their balance. Only admins and customers can access this resource.

    ```
    **`/customers/profile`** --> method GET
    Used to view customer's username, email and balance.

    Sample response
    {
        "id": 1,
        "balance": 2000000,
        "username": "customer",
        "email": "customer@mail.com",
    }
    ```

    ```
    **`/customers/topup`** --> method POST
    Used to topup customer's balance. Requires request body as shown below.

    {
        "balance": 2000000,
    }

    Sample response
    {
        "message": "topup berhasil"
    }
    ```

19. ORDERS
    Order data, main bussiness flow for this app, in this resource, each role have different access for this table.

    ```
    **`/orders`** --> method GET
    Used to get orders in this app, any authorized user can access this data.

    Sample response
    {
        "result": [
        {
            "id": 1,
            "name": "order_2024-4-25",
            "status_id": 1,
            "status_name": "paid",
            "Products": [
                {
                    "id": 1,
                    "product_id": 1,
                    "product_name": "Susu SGM Explor",
                    "product_quantity": 20,
                    "product_price": 1000000,
                    "created_at": "2024-04-25T15:20:31.70502Z",
                    "updated_at": "2024-04-25T15:20:31.705022Z"
                },
                {
                    "id": 2,
                    "product_id": 3,
                    "product_name": "Susu S26",
                    "product_quantity": 20,
                    "product_price": 200000,
                    "created_at": "2024-04-25T15:20:31.707023Z",
                    "updated_at": "2024-04-25T15:20:31.707026Z"
                }
            ],
            "created_at": "2024-04-25T15:20:31.696055Z",
            "updated_at": "2024-04-25T15:30:47.829814Z"
        }
    ]
    }
    ```

    ```
    **`/orders`** --> method POST
    Create new order based on selected products by the customer. Required request body as below.

    {
        "products": [
        {
            "quantity": 20,
            "product_id": 1
        },
        {
            "quantity": 20,
            "product_id": 3
        }
        ]
    }

    Sample response
    {
        "message": "berhasil membuat order"
    }
    ```

    ```
    **`/orders/:id/pay`** --> method POST
    Used to pay the order up by the customer, if his balance is sufficient. Requires only order id as it's url parameter.

    Sample response
    {
        "result": "pembayaran berhasil"
    }
    ```

    ```
    **`/orders/:id/complete`** --> method POST
    Cashiers will close this order after successful payments is received. Admins can also do the same though. Requires only order id as it's url parameter.

    Sample response
    {
        "message": "order berhasil di tutup"
    }
    ```

21. TRANSACTION SALES
    This API is used to get sales report in this app.

    ```
    **`/transaction/sales`** --> method GET
    Used to get sales report including total earnings for the whole orders.

    Sample response
    {
        "result": [
        {
            "id": 1,
            "order_id": 1,
            "order_name": "order_2024-4-25",
            "customer_id": 1,
            "customer_name": "customer",
            "products_sold": [
                {
                    "id": 1,
                    "product_id": 1,
                    "product_name": "Susu SGM Explor",
                    "product_quantity": 20,
                    "product_price": 1000000,
                    "created_at": "2024-04-25T15:20:31.70502Z",
                    "updated_at": "2024-04-25T15:20:31.705022Z"
                },
                {
                    "id": 2,
                    "product_id": 3,
                    "product_name": "Susu S26",
                    "product_quantity": 20,
                    "product_price": 200000,
                    "created_at": "2024-04-25T15:20:31.707023Z",
                    "updated_at": "2024-04-25T15:20:31.707026Z"
                }
            ],
            "total_earnings": 24000000,
            "created_at": "2024-04-25T15:34:15.678006Z",
            "updated_at": "2024-04-25T15:34:15.678009Z"
        }
    ]
    }
    ```

    ```
    **`/transaction/sales/:id`** --> method GET
    Used to get sales report including total earnings by transaction id.

    Sample response
    {
        "result": {
            "id": 1,
            "order_id": 1,
            "order_name": "order_2024-4-25",
            "customer_id": 1,
            "customer_name": "customer",
            "products_sold": [
                {
                    "id": 1,
                    "product_id": 1,
                    "product_name": "Susu SGM Explor",
                    "product_quantity": 20,
                    "product_price": 1000000,
                    "created_at": "2024-04-25T15:20:31.70502Z",
                    "updated_at": "2024-04-25T15:20:31.705022Z"
                },
                {
                    "id": 2,
                    "product_id": 3,
                    "product_name": "Susu S26",
                    "product_quantity": 20,
                    "product_price": 200000,
                    "created_at": "2024-04-25T15:20:31.707023Z",
                    "updated_at": "2024-04-25T15:20:31.707026Z"
                }
            ],
            "total_earnings": 24000000,
            "created_at": "2024-04-25T15:34:15.678006Z",
            "updated_at": "2024-04-25T15:34:15.678009Z"
        }
    }
    ```
