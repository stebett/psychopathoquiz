# psychopathoquiz
A simple terminal game written in Go to learn psychopathology, with definitions from DSM 5 

This toy project has been done to learn Psychopathology and Golang at the same time, easing the study of both. 

The games give you random questions to answer, with a probability changing based on your performance on every quastion. 
The program calculates the Levenshtein distance between your answer and the correct one, and gives you a score based on that.

The Levenshtein algorithm is implemented with dynamic programming.

![pic](https://github.com/stebett/psychopathoquiz/blob/master/screenshot.png)
