insert into student (name,age,grade ) values ('张三',18,'三年级');

select * from student where age > 18;

update student set grade = '四年级' where name = '张三';
delete from student where age  < 15;