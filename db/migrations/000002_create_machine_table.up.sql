begin;
create type rc4laundry.machine_type as enum ('washer', 'dryer');
comment on type rc4laundry.machine_type is 'The machine is either a washer or a dryer.';
create table if not exists rc4laundry.machine (
	floor integer not null,
	position integer not null,
	type rc4laundry.machine_type not null,
	last_started_at timestamp default now() not null,
	approx_duration integer,
	is_in_use boolean default false not null,
	unique (floor, position)
);
comment on table rc4laundry.machine is 'A machine.';
comment on column rc4laundry.machine.floor is 'The floor where the machine is.';
comment on column rc4laundry.machine.position is 'The position of the machine, from left to right, starting from zero.';
comment on column rc4laundry.machine.type is 'Type of the machine - either a washer or a dryer.';
comment on column rc4laundry.machine.last_started_at is 'The time at which the machine was last started.';
comment on column rc4laundry.machine.is_in_use is 'Whether the machine is currently in use.';
commit;