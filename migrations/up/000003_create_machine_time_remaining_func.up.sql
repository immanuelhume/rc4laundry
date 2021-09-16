CREATE OR REPLACE FUNCTION rc4laundry.machine_in_use_time_remaining(machine rc4laundry.machine_in_use) RETURNS BIGINT AS $$
DECLARE time_passed BIGINT;
BEGIN
SELECT EXTRACT(
		EPOCH
		FROM (NOW() - machine.started_at)
	) INTO time_passed;
RETURN CASE
	WHEN machine.type = 'washer' then 1800 - time_passed
	ELSE 2400 - time_passed
END;
END;
$$ LANGUAGE plpgsql STABLE;
COMMENT ON FUNCTION rc4laundry.machine_in_use_time_remaining IS 'The estimated time left, in seconds, for this machine to finish its work.';