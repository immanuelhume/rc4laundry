begin;
create or replace function rc4laundry.set_machine_approx_duration_default() returns trigger as $$ begin if new.type = 'washer' then new.approx_duration := 1800;
elsif new.type = 'dryer' then new.approx_duration := 2400;
else raise 'Unknown machine type: %',
new.type;
end if;
return new;
end;
$$ language plpgsql;
create trigger machine_approx_duration_default before
insert on rc4laundry.machine for each row
	when (new.approx_duration is null) execute procedure rc4laundry.set_machine_approx_duration_default();
comment on function rc4laundry.set_machine_approx_duration_default() is 'Hook to insert a default value for the approximate total duration of each machine if none is provided.';
commit;