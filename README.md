# couponService
- This back-end service delivers coupon codes to our users.
- This service can be used to create, update, retrieve & list coupons via the API

### Create Coupon

```curl -d '{"name": "Save £35 at Aldi", "brand": "Aldi", "value": 35, "createdAt": "2018-03-01 10:15:53", "expiry": "2019-03-01 10:15:53"}' -H "Content-Type:application/json" --request POST 'http://localhost:8001/coupon'```

### Get Coupon

```curl -H "Content-Type:application/json" --request GET 'http://localhost:8001/coupon/5e85b669a1f3c6873189e363'```

### List Coupons

```curl -H "Content-Type:application/json" --request GET 'http://localhost:8001/coupons'```

### Update Coupon

```curl -d '{"name": "Save £65 at Asda", "brand": "Asda", "value": 65, "createdAt": "2019-02-14 17:29:53", "expiry": "2020-07-21 17:20:45"}' -H "Content-Type:application/json" --request PUT 'http://localhost:8001/coupon/5e85b669a1f3c6873189e363'```

### Delete Coupon

```curl -H "Content-Type:application/json" --request DELETE 'http://localhost:8001/coupon/5e85b669a1f3c6873189e363'```