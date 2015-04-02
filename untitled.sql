SELECT *,
  (SELECT *,
     (SELECT *
      FROM items
      WHERE items.order_id = orders.order_id) AS order_items
   FROM orders
   WHERE orders.customer_id = customers.customer_id) AS cust_orders
FROM customers;