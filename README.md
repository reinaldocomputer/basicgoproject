# Benchmark

## Insert

| Name           | Duration     | Quantity of data | Method                 | Workers | Processor                    | Memory              | Database |
| -------------- | ------------ | ---------------- | ---------------------- | ------- | ---------------------------- | ------------------- | -------- |
| CLock 6workers | 3.043115ms   | 1000             | Channel and mutex lock | 6       | 2,6 GHz 6-Core Intel Core i7 | 16 GB 2667 MHz DDR4 | map      |
| CLock 6workers | 27.125452ms  | 10000            | Channel and mutex lock | 6       | 2,6 GHz 6-Core Intel Core i7 | 16 GB 2667 MHz DDR4 | map      |
| CLock 6workers | 157.20645ms  | 100000           | Channel and mutex lock | 6       | 2,6 GHz 6-Core Intel Core i7 | 16 GB 2667 MHz DDR4 | map      |
| CLock 6workers | 1.243096578s | 1000000          | Channel and mutex lock | 6       | 2,6 GHz 6-Core Intel Core i7 | 16 GB 2667 MHz DDR4 | map      |
| CLock 6workers | 3.713871219s | 3000000          | Channel and mutex lock | 6       | 2,6 GHz 6-Core Intel Core i7 | 16 GB 2667 MHz DDR4 | map      |