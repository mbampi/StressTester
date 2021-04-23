# Stress Tester

Golang script to test multiple requests to same endpoint separated in levels

example result:

```
2021/04/22 23:49:31 Stress Tester: RequestCentralizer
2021/04/22 23:51:11 Level: 0 | Fires: 100 | Req/Second: 1 | Total: 100 | Errors: 0 (0.00%) 
2021/04/22 23:52:52 Level: 1 | Fires: 100 | Req/Second: 5 | Total: 500 | Errors: 0 (0.00%) 
2021/04/22 23:54:32 Level: 2 | Fires: 100 | Req/Second: 10 | Total: 1000 | Errors: 0 (0.00%) 
2021/04/22 23:56:13 Level: 3 | Fires: 100 | Req/Second: 15 | Total: 1500 | Errors: 0 (0.00%) 
2021/04/22 23:57:54 Level: 4 | Fires: 100 | Req/Second: 20 | Total: 2000 | Errors: 7 (0.35%) 
2021/04/22 23:59:35 Level: 5 | Fires: 100 | Req/Second: 25 | Total: 2500 | Errors: 14 (0.56%) 
2021/04/23 00:01:17 Level: 6 | Fires: 100 | Req/Second: 30 | Total: 3000 | Errors: 22 (0.73%) 
2021/04/23 00:01:17 Time  11m45.966114678s
```