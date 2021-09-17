create or replace function rc4laundry.machine_time_remaining(machine rc4laundry.machine) returns integer as $$
declare time_passed integer;
begin if machine.is_in_use is false then return 0;
else
select extract(
		epoch
		from (now() - machine.last_started_at)
	) into time_passed;
return greatest(1, machine.approx_duration - time_passed);
end if;
end;
$$ language plpgsql stable;
comment on function rc4laundry.machine_time_remaining is 'The estimated time left, in seconds, for this machine to finish its work.';