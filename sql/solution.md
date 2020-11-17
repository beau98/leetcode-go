- 180 Consecutive Numbers 至少三个连续递增数
```sql
select distinct l1.Num AS ConsecutiveNums
from  Logs l1,Logs l2,Logs l3
where l1.Id = l2.Id-1 and l2.Id = l3.Id -1 and l1.Num = l2.Num and l2.Num = l3.Num;
```

- 181 Employees Earning More Than Their Managers 比主管薪水高的员工
```sql
select e1.Name as Employee from Employee e1,Employee e2 where e1.ManagerId = e2.Id  and e1.Salary > e2.Salary;
```

- 182 Duplicate Emails
```sql
select Email from (select Email,Count(Email) as num from Person group by Email) as statistic 
where num > 1
```
