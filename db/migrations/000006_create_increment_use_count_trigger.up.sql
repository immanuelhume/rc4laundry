begin;
create or replace function rc4laundry.increment_machine_use_count() returns trigger as $$ begin new.total_use_count := new.total_use_count + 1;
return new;
end;
$$ language plpgsql;
comment on function rc4laundry.increment_machine_use_count() is 'Records one instance of use, incrementing the total_use_count column.';
create trigger machine_use_count before
update on rc4laundry.machine for each row
	when (new.is_in_use is true) execute procedure rc4laundry.increment_machine_use_count();
commit;