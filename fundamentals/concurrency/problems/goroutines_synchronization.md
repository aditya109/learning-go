1. **Question 1:**
   - You are given a string representing a sequence of operations, "RRRRLLLRRRLLLRLRL". You have two goroutines: `right` and `left`. The `right` goroutine returns true when it encounters an "R", and the `left` goroutine returns true when it encounters an "L". The goal is to count how many times the sequence "RLL" appears in the string using these two goroutines.

2. **Question 2:**
   - A series of commands are represented as a string "ABCCBABCA". You have two goroutines, `taskA` and `taskB`, which process the string. The `taskA` goroutine returns true for every "A", and the `taskB` goroutine returns true for every "B". Whenever both goroutines return true consecutively, you can count one occurrence of the sequence "AB". The goal is to find the maximum number of "AB" sequences in the string.

3. **Question 3:**
   - You are provided with a string "MMTTMMTTMMMTT". You have two goroutines, `morning` and `evening`, which return true if they receive an "M" or "T", respectively. Each time the `morning` goroutine processes "M" and the `evening` goroutine processes "T", it counts as one valid "MT" pair. Calculate the maximum number of "MT" pairs that can be formed from the string.

4. **Question 4:**
   - Given a string "XYXXYYXYYXY", you have two goroutines, `processX` and `processY`. The `processX` goroutine returns true when it processes an "X", and the `processY` goroutine returns true when it processes a "Y". The aim is to count how many times the sequence "XYX" can be formed using these goroutines.

5. **Question 5:**
   - You receive a signal sequence "PSPSPPSSPPS". You have two goroutines, `signalP` and `signalS`, which return true when they encounter "P" and "S", respectively. When both goroutines identify the sequence "PS", it counts as one valid signal. Determine the total number of valid "PS" signals that can be identified in the sequence.