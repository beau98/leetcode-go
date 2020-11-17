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
- 183 Customers Who Never Order 
```sql
select Name as Customers from Customers left join Orders on Customers.Id = Orders.CustomerId where Orders.CustomerId is null
```

- 184 Department Highest Salary
```sql
select Department.Name as Department,Employee.Name as Employee,Salary from Employee join Department on Employee.DepartmentId = Department.Id 
where (Employee.DepartmentId , Salary) IN
(select DepartmentId,Max(Salary) from Employee  group by DepartmentId);
```

- 196 Delete Duplicate Emails
```sql
delete from Person where Id not in  (select * from (select Min(Id) from Person group by Email) as tmp);
```

- 197 Rising Temperature
```sql
select w1.id from Weather as w1 join Weather as w2 on DATEDIFF(w1.recordDate, w2.recordDate) = 1 and w1.Temperature > w2.Temperature;
```

- 595 Big Countries
```sql
select name,population,area from World where area > 3000000 or population > 25000000;
```

- 596 Classes More Than 5 Students
```sql
select class from courses group by class HAVING COUNT(DISTINCT student) >= 5;
```

- 620 Not Boring Movies
```sql
select * from cinema where description != 'boring' and id % 2=1 order by rating desc;
```

- 626 Exchange Seats
```sql
SELECT
	CASE
		WHEN seat.id % 2 <> 0 AND seat.id = (SELECT COUNT(*) FROM seat) THEN seat.id
		WHEN seat.id % 2 = 0 THEN seat.id - 1
		ELSE
			seat.id + 1
	END as id,
	student 
FROM seat
ORDER BY id;
```

- 627 Swap Salary
```sql
update salary set sex  = 
    CASE sex
        when 'm' then 'f'
        else 'm'
    end;
```