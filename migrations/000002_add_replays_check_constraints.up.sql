ALTER TABLE replays ADD CONSTRAINT replays_runtime_check CHECK (runtime >= 0);
ALTER TABLE replays ADD CONSTRAINT replays_year_check CHECK (year BETWEEN 2011 AND date_part('year', now()));
ALTER TABLE replays ADD CONSTRAINT heroes_length_check CHECK (array_length(heroes, 1) 10);
