BEGIN;
CREATE TYPE rc4laundry.machine_type AS ENUM ('washer', 'dryer');
COMMENT ON TYPE rc4laundry.machine_type IS 'The machine is either a washer or a dryer.';
CREATE TABLE IF NOT EXISTS rc4laundry.machine_in_use (
	floor SMALLINT NOT NULL,
	position SMALLINT NOT NULL,
	type rc4laundry.machine_type NOT NULL,
	started_at TIMESTAMP DEFAULT NOW(),
	UNIQUE (floor, position)
);
COMMENT ON TABLE rc4laundry.machine_in_use IS 'A machine - either a washer or a dryer.';
COMMENT ON COLUMN rc4laundry.machine_in_use.floor IS 'The floor where the machine is.';
COMMENT ON COLUMN rc4laundry.machine_in_use.position IS 'The position of the machine, from left to right, starting from zero.';
COMMENT ON COLUMN rc4laundry.machine_in_use.type IS 'Type of the machine - either a washer or a dryer.';
COMMENT ON COLUMN rc4laundry.machine_in_use.started_at IS 'The time at which the machine was last started.';
COMMIT;