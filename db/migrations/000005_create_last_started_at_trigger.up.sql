begin;
create or replace function rc4laundry.set_machine_last_started_at() returns trigger as $$ begin new.last_started_at := now();
return new;
end;
$$ language plpgsql;
comment on function rc4laundry.set_machine_last_started_at() is 'Sets a timestamp for when the machine was last started.';
create trigger machine_last_started_at before
update on rc4laundry.machine for each row
	when (new.is_in_use is true) execute procedure rc4laundry.set_machine_last_started_at();
commit;