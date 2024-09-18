demo中用到的runtime.ReadMemStats方法，可以读取内存状态，而读取内存状态时需要执行stop the world。
因此demo中启动了一个独立的goroutine周期性触发stop the world，并通过压测来查看stop the world对http服务性能的影响。

有stop the world
```
❯ wrk -t20 -c800 -d30s --latency http://127.0.0.1:8080/test
Running 30s test @ http://127.0.0.1:8080/test
  20 threads and 800 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    11.02ms    8.39ms 267.43ms   97.64%
    Req/Sec     3.46k     1.31k    5.08k    87.09%
  Latency Distribution
     50%   10.14ms
     75%   10.74ms
     90%   11.56ms
     99%   40.64ms
  2009418 requests in 30.07s, 228.04MB read
  Socket errors: connect 0, read 14073, write 0, timeout 0
Requests/sec:  66833.38
Transfer/sec:      7.58MB


```

没有stop the world
```
❯ wrk -t20 -c800 -d30s --latency http://127.0.0.1:8080/test
Running 30s test @ http://127.0.0.1:8080/test
  20 threads and 800 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     9.91ms    3.99ms 216.60ms   98.68%
    Req/Sec     3.97k     0.89k    7.52k    93.81%
  Latency Distribution
     50%    9.78ms
     75%   10.22ms
     90%   10.62ms
     99%   14.18ms
  2328831 requests in 30.06s, 264.29MB read
  Socket errors: connect 0, read 7302, write 0, timeout 0
Requests/sec:  77472.12
Transfer/sec:      8.79MB
```

根据上面的对比可以看出，stw导致了长尾效应。90%的情况下，双方请求没有区别，但是p99时，stw会导致响应时间成倍增加。

