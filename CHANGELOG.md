# 0.0.1

<p>

First Commit. Create Empty files + directory

</p>

[Add]
- ./rdb
  - dao_conn.go
  - dao_struct.go
  - re_json.go
- ./rest
  - constant.go
  - src_serve.go
- other necessary files such as .gitignore, CHANGELOG etc.

[Change]


[Fix]


[Remove]


# 0.0.2

<p>

Add JSON database structure. TODO: redis
</p>

[Add]
- ./rdb
  - dao_conn.go
    - struct DataBase
      - method Connect 
  - dao_struct.go
    - struct Employee
    - struct EmployeePlural
- ./main.go
  - func init - load environment
  
[Change]
- ./Dockerfile

