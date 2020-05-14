# go-crawler

Take a look at the whole picture. How we designed the whole project.

![project](https://github.com/yngyuan/go-crawler/blob/master/images/project.png?raw=true)

## A single-task crawler
design of a single-task crawler, so far we don't have to save anything.
![single](https://github.com/yngyuan/go-crawler/blob/master/images/single.png?raw=true)

## A concurrent crawler using goroutine
Go from the single-task crawler, we made this into a concurrent crawler with the help of goroutines.
![concurrent](https://github.com/yngyuan/go-crawler/blob/master/images/concurrent.png?raw=true)

The core part is how we implemented a task scheduler.
![scheduler](https://github.com/yngyuan/go-crawler/blob/master/images/scheduler.png?raw=true)

## A distributed crawler.
From concurrent to distributed, the main change is how the workers communicate: from goroutines to the use of rpc.

If you don't know rpc, [here](https://github.com/yngyuan/min-rpc/) is a minimal rpc project from me.

![distributed](https://github.com/yngyuan/go-crawler/blob/master/images/distributed.png?raw=true)

## Next Steps
The purpose of this project is more about learning the go programming language. However, there are more steps to make this a better crawler. 
1. Crawler more websites. Use css selector/xpath to parse. We used regex in this project.
2. More anti anti-crawler techs.
3. Support dynamic websites. Simulate login, and search.
4. Remove duplicate in distributed crawler.
