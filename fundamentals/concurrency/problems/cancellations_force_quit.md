1. **Timeout-Based Producer-Consumer Pattern:**
   Implement a producer-consumer pattern where the producer generates random numbers and sends them to a consumer via a channel. The consumer processes and prints these numbers. Introduce a timeout mechanism using a `time.After` function to stop the producer and consumer after a specified duration. Ensure that all channels are closed properly and that the timeout is handled gracefully.

2. **Dynamic Stop Condition:**
   Create a producer-consumer pattern where the producer generates integers and sends them to a consumer. The consumer should process and print the numbers. Introduce a stop condition where the producer stops producing and signals the consumer to stop if the sum of generated numbers exceeds a specified threshold. Use a stop channel to communicate this condition and ensure proper closure of all channels.

3. **Error Handling and Recovery:**
   Design a producer-consumer system where the producer generates numbers and sends them to the consumer. The consumer processes these numbers but occasionally encounters errors (e.g., invalid data). Introduce an error channel where the consumer can send errors. Implement a mechanism to handle errors gracefully and stop both the producer and consumer when a critical error is encountered.

4. **Graceful Shutdown with Context:**
   Implement a producer-consumer pattern where the producer generates data and sends it to a consumer. Use Go’s `context` package to manage the shutdown process. Ensure that the producer and consumer can be signaled to stop processing when the context is canceled. Verify that all channels are closed properly and resources are cleaned up.

5. **Rate-Limited Producer-Consumer Pattern:**
   Develop a producer-consumer pattern where the producer generates data at a fixed rate and sends it to the consumer. Introduce a rate limit for the producer using a ticker. The consumer should process and print the data. Implement a mechanism to stop the producer and consumer when a specified count of items has been processed or a timeout occurs. Ensure that all channels are closed properly when stopping.