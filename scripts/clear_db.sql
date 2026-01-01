-- SQL script to clear all data from the database
-- Use this if you prefer SQL approach over Go script

-- Clear data in order (respecting foreign key constraints)
TRUNCATE TABLE schedules CASCADE;
TRUNCATE TABLE feedbacks CASCADE;
TRUNCATE TABLE visits CASCADE;
TRUNCATE TABLE profiles CASCADE;
TRUNCATE TABLE lockers CASCADE;
TRUNCATE TABLE stay_areas CASCADE;
TRUNCATE TABLE seva_types CASCADE;

-- Verify tables are empty
SELECT 'profiles' as table_name, COUNT(*) as count FROM profiles
UNION ALL
SELECT 'visits', COUNT(*) FROM visits
UNION ALL
SELECT 'schedules', COUNT(*) FROM schedules
UNION ALL
SELECT 'feedbacks', COUNT(*) FROM feedbacks
UNION ALL
SELECT 'lockers', COUNT(*) FROM lockers
UNION ALL
SELECT 'stay_areas', COUNT(*) FROM stay_areas
UNION ALL
SELECT 'seva_types', COUNT(*) FROM seva_types;
