1. **Multi-Producer Multi-Consumer Scenario:**
   - Implement a system with two producers and two consumers. Producers generate numbers and send them to a shared channel. One consumer calculates the sum of the numbers received, and the other calculates the average. Introduce a quit channel that stops all producers and consumers when the sum reaches a certain limit. Ensure all channels are properly closed and use `range` to receive values.

2. **Prime Number Filter:**
   - Create a producer that generates numbers and sends them to a consumer via a channel. The consumer should filter out prime numbers and send them to another channel for processing. Introduce a quit channel that stops the producer and consumer when a certain number of primes have been processed. Ensure that all channels are properly closed and use `range` to receive values.

3. **Dynamic Termination Condition:**
   - Develop a producer-consumer system where the producer generates numbers and sends them to multiple consumers via separate channels. Each consumer has a different task, such as calculating the sum, product, or identifying even numbers. Introduce a quit channel that signals the end of processing when any consumer reaches a certain condition (e.g., sum exceeds a limit, product exceeds a threshold, or a specific even number is found). Ensure proper closing of channels and use `range` for receiving values.

4. **Load Balancing with Multiple Consumers:**
   - Implement a producer-consumer pattern with one producer and multiple consumers. The producer generates numbers and uses a channel to distribute them to consumers in a round-robin fashion. Each consumer processes the numbers and sends results back to a results channel. Introduce a quit channel that signals when the processing should stop (e.g., after a certain number of results or a specific condition is met). Ensure that all channels are properly closed and use `range` to receive values.

5. **Timed Producer and Consumer:**
   - Create a system where a producer generates numbers and sends them to a consumer via a channel. The producer should stop generating numbers and close the channel after a fixed time period or when a certain number of numbers have been produced, whichever comes first. The consumer processes the numbers and stops when the channel is closed. Introduce a quit channel that can also signal an immediate stop to the producer and consumer if necessary. Ensure proper channel closing and use `range` to receive values.