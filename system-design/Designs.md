# Design Rate Limiter
Limit the number of requests for particular resource.

##### Functional Reuqirements
1. Accuracy in claculating request per time for resources !! Important
2. Limit number of requests from particular IP per time basis.
3. Limit number of requests from particular user per time basis.
4. Log requests per time basis for billing and other analytics.
5. Inform users about throttled requests.

##### Non-functional Requiremnts
1. Low latency / Scalable :- Our rate limiter should not make original APIs requests looks delayed and should scale with number of requests.
2. Highly Available :- Should be highly available else our APIs becomes unavaibalible.
3. Fault tolerant: If ratelimiter went down, BE subsystem should not fail. 

##### BOE -- propose HLD and gets buy-in

##### HLD
Q: Where should be the ratelmiter sits ? In API gateway OR server side ?
A: Completely depends upon the existing tech stack used and capability of the engineers in the team. 
   - Evaluate your current technology stack, such as programming language, cache service,
        etc. Make sure your current programming language is efficient to implement rate limiting
        on the server-side.
   - Identify the rate limiting algorithm that fits your business needs. When you implement
        everything on the server-side, you have full control of the algorithm. However, your    
        choice might be limited if you use a third-party gateway.
   - If you have already used microservice architecture and included an API gateway in the
        design to perform authentication, IP whitelisting, etc., you may add a rate limiter to the
        API gateway.
   - Building your own rate limiting service takes time. If you do not have enough
        engineering resources to implement a rate limiter, a commercial API gateway is a better
        option.

###### *Algorithms for ratelimiting* 
____________________
**Token bucket** : Used by Amazon and stripe for thier APIs. In this algo,  a bucket with fixed size (Requests allowed per time) is used and refilled at preset rates (Refill Rate)

**API**

func InitTokenBucketlimiter (buketSize , refillRate int)

**Pros**:
- The algorithm is easy to implement.
- Memory efficient.
- Token bucket allows a burst of traffic for short periods. A request can go through as long
as there are tokens left.

**Cons**:
 - Two parameters in the algorithm are bucket size and token refill rate. However, it might be challenging to tune them properly.

__________
 **Leaking bucket** : It is just like token bucket algo but it uses queue.
 1. New request comes in, it checks if the queue has empty space, if yes request is being added to the queue.
 2. Else the request is dropped.
 3. Request from the queue is processed at particular constant rate so no burst trafffic problem.

**API**

func InitleakingBucketlimiter (queueSize , outFlowrate int)

**Pros**:
-  Memory efficient given the limited queue size.
- Requests are processed at a fixed rate therefore it is suitable for use cases that a stable
outflow rate is needed.

**Cons**:
-  A burst of traffic fills up the queue with old requests, and if they are not processed in
time, recent requests will be rate limited.
- There are two parameters in the algorithm. It might not be easy to tune them properly.

__________
**Fixed window counter**

- The algorithm divides the timeline into fix-sized time windows and assign a counter for
each window.
- Each request increments the counter by one.
- Once the counter reaches the pre-defined threshold, new requests are dropped until a new time window starts

**Pros**
- Memory efficient.
- Easy to understand.
- Resetting available quota at the end of a unit time window fits certain use cases.

**Cons:**
- Spike in traffic at the edges of a window could cause more requests than the allowed
quota to go through.
________
**Sliding window log (Best and Recommended by Kong API Gateway)**

- The algorithm keeps track of request timestamps. Timestamp data is usually kept in
cache, such as sorted sets of Redis
- When a new request comes in, remove all the outdated timestamps. Outdated timestamps
are defined as those older than the start of the current time window.
- Add timestamp of the new request to the log.
- If the log size is the same or lower than the allowed count, a request is accepted.
Otherwise, it is rejected.

**Pros**:

- Rate limiting implemented by this algorithm is very accurate. In any rolling window,
requests will not exceed the rate limit.
**Cons**:

- The algorithm consumes a lot of memory because even if a request is rejected, its
timestamp might still be stored in memory.

_____
- Sliding window counter
_____



