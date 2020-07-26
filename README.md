# acee

Acee, meaning "ace the leetcode", is a program to record your daily problem solving, and generate recommended problem sets by algorithm, 
to help you reinforce the frequently incorrect problems.

Tech Stack:
Golang + embedded db(sqlite3) + cobra(CLI framework)

Plan:
  Phase 1: Design the data structure and database. 
  Phase 2: Design the CLI and implement the func from front-end to back-end
  Phase 3: Add test cases and collect feedbacks

Data: crawled leetcode problems + user records
1. help user map input problem_id -> Problem{name, link, level, freq, topic}
2. could update leetcode problems to avoid stale data


Design CLI:
acee 

1. amend: could update/delete last 10 records(default)

2. add: add a record

3. find: find history records based on date/problem id

4. get: get recommended problem set to do today, based on users' history record.

5. -h
  --help: help messages
  
6. update (update leetcode problem set)
