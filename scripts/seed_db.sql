-- SQL script to seed initial data
-- Run this after clearing the database

-- Seed SevaTypes
INSERT INTO seva_types (id, name, description, is_active, created_at) VALUES
(gen_random_uuid(), 'Dining', 'Food service and dining hall seva', true, NOW()),
(gen_random_uuid(), 'Kitchen Support', 'Kitchen preparation and cleaning seva', true, NOW()),
(gen_random_uuid(), 'Counter', 'Reception and counter service seva', true, NOW()),
(gen_random_uuid(), 'Cleaning', 'General cleaning and maintenance seva', true, NOW()),
(gen_random_uuid(), 'Security', 'Security and monitoring seva', true, NOW()),
(gen_random_uuid(), 'Garden', 'Gardening and landscaping seva', true, NOW()),
(gen_random_uuid(), 'Transport', 'Transportation and logistics seva', true, NOW()),
(gen_random_uuid(), 'Medical', 'Medical assistance and first aid seva', true, NOW());

-- Seed StayAreas
INSERT INTO stay_areas (id, name, capacity) VALUES
(gen_random_uuid(), 'Dormitory A - Men', 50),
(gen_random_uuid(), 'Dormitory B - Men', 50),
(gen_random_uuid(), 'Dormitory C - Women', 40),
(gen_random_uuid(), 'Dormitory D - Women', 40),
(gen_random_uuid(), 'Family Room Block 1', 20),
(gen_random_uuid(), 'Family Room Block 2', 20),
(gen_random_uuid(), 'Guest House - VIP', 10),
(gen_random_uuid(), 'Cottage Area', 15);

-- Seed Lockers (100 lockers)
INSERT INTO lockers (id, locker_number, is_occupied, created_at)
SELECT 
    gen_random_uuid(),
    'L' || LPAD(generate_series::text, 3, '0'),
    false,
    NOW()
FROM generate_series(1, 100);

-- Verify seeded data
SELECT 'seva_types' as table_name, COUNT(*) as count FROM seva_types
UNION ALL
SELECT 'stay_areas', COUNT(*) FROM stay_areas
UNION ALL
SELECT 'lockers', COUNT(*) FROM lockers;

-- Display seeded data
SELECT 'SevaTypes:' as info;
SELECT name, description FROM seva_types ORDER BY name;

SELECT 'StayAreas:' as info;
SELECT name, capacity FROM stay_areas ORDER BY name;

SELECT 'Lockers:' as info;
SELECT COUNT(*) as total_lockers, 
       SUM(CASE WHEN is_occupied THEN 1 ELSE 0 END) as occupied,
       SUM(CASE WHEN NOT is_occupied THEN 1 ELSE 0 END) as available
FROM lockers;
