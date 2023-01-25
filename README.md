# Activity Selector

This project aims to assist in the selection of activities to be carried out in my free time

For execution, I receive a list of activities in csv format, and the application randomly selects one of the activities from this list

### Installation

To install the external libs in the project, run the command below:
```shell
go mod download
```

### Execution

To execute, use the run command below passing as parameter the input file to be run
```shell
go run main.go activity_sample.csv
```

# Future works

For the evolution and improvement of the project, it would be interesting:

* add a list of "weights" for each activity, and take it into account when selecting tasks, so that activities with a higher weight are more likely to be selected