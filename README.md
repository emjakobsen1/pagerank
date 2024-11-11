## Linear Algebra PageRank algorithm

This is a demonstration/implementation of the PageRank algorithm, which contributed to Google's success by ranking pages on the web for their search engine. Made as part of the Linear Algebra course fall 2024 at IT-University of Copenhagen. 

### Prerequisties 

You will need the Go programming language to run this program:

- [Go Programming Language](https://go.dev/doc/install)
- You will need to have whichever dataset.txt to test against in the same folder as main.go for the program to handle it.

### How to run 

The program takes a command-line argument with datasets on graphs adhering to the format provided in PageRankExampleData.

1. From the repository root, run `main.go` and the file name. Examples 

    ```
     go run main.go
     go run main.go bigRandom.txt
    ```

If no argument is given the default data set will be p2p-Gnutella08-mod.txt.

### Runtimes 

Expect some long runtimes on the huge graphs, especially bigRandom.txt. Execution time with 1 million steps on Windows with AMD Ryzen 5 5600U/8GB RAM took 30 seconds.