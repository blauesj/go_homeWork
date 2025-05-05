start transaction ;
if (select balance from accouts where id = 1) < 100 then
	rollback ;
end if ;

update accouts set balance = balance - 100 where id = 1 ;
update accouts set balance = balance + 100 where id = 2 ;

insert into transactions (from_account, to_account, amount)values (1, 2, 100);

commit ;