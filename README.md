Benchmark testing
We’ve said that caching significantly improves the performance of our application. To support that claim, let’s perform some benchmarks.
The benchmarking tool of choice here is go-wrk, which you can install with this command:
go install github.com/tsliwowicz/go-wrk@latest
go-wrk -c 80 -d 5  localhost:8000/api/user/1
