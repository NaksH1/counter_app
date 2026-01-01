-- ============================================
-- SAMPLE DATA FOR ASHRAM CONNECT DATABASE
-- ============================================

-- Clear existing data (optional - use with caution)
-- TRUNCATE TABLE feedbacks, schedules, visits, lockers, profiles CASCADE;

-- ============================================
-- 1. PROFILES (10 volunteers)
-- ============================================
INSERT INTO profiles (id, name, email, phone_number, gender, category, is_blocked, remarks, created_at, updated_at) VALUES
('a1b2c3d4-e5f6-7890-abcd-ef1234567890', 'Rajesh Kumar', 'rajesh.kumar@example.com', '+91-9876543210', 'Male', 'Short Term Volunteer', false, 'Experienced in kitchen seva', NOW() - INTERVAL '30 days', NOW()),
('b2c3d4e5-f6a7-8901-bcde-f12345678901', 'Priya Sharma', 'priya.sharma@example.com', '+91-9876543211', 'Female', 'Long Term Volunteer', false, 'First time volunteer', NOW() - INTERVAL '25 days', NOW()),
('c3d4e5f6-a7b8-9012-cdef-123456789012', 'Amit Patel', 'amit.patel@example.com', '+91-9876543212', 'Male', 'Short Term Volunteer', false, NULL, NOW() - INTERVAL '20 days', NOW()),
('d4e5f6a7-b8c9-0123-def1-234567890123', 'Sneha Reddy', 'sneha.reddy@example.com', '+91-9876543213', 'Female', 'Overseas Volunteer', false, 'Prefers morning shifts', NOW() - INTERVAL '15 days', NOW()),
('e5f6a7b8-c9d0-1234-ef12-345678901234', 'Vikram Singh', 'vikram.singh@example.com', '+91-9876543214', 'Male', 'Long Term Volunteer', false, 'Good with coordination', NOW() - INTERVAL '10 days', NOW()),
('f6a7b8c9-d0e1-2345-f123-456789012345', 'Ananya Iyer', 'ananya.iyer@example.com', '+91-9876543215', 'Female', 'Short Term Volunteer', false, NULL, NOW() - INTERVAL '8 days', NOW()),
('a7b8c9d0-e1f2-3456-1234-567890123456', 'Karthik Menon', 'karthik.menon@example.com', '+91-9876543216', 'Male', 'Short Term Volunteer', true, 'Blocked due to misconduct', NOW() - INTERVAL '5 days', NOW()),
('b8c9d0e1-f2a3-4567-2345-678901234567', 'Divya Nair', 'divya.nair@example.com', '+91-9876543217', 'Female', 'Long Term Volunteer', false, 'Excellent at dining hall seva', NOW() - INTERVAL '3 days', NOW()),
('c9d0e1f2-a3b4-5678-3456-789012345678', 'Arjun Desai', 'arjun.desai@example.com', '+91-9876543218', 'Male', 'Overseas Volunteer', false, NULL, NOW() - INTERVAL '2 days', NOW()),
('d0e1f2a3-b4c5-6789-4567-890123456789', 'Meera Joshi', 'meera.joshi@example.com', '+91-9876543219', 'Female', 'Short Term Volunteer', false, 'Prefers counter duty', NOW() - INTERVAL '1 day', NOW());

-- ============================================
-- 2. LOCKERS (20 lockers across 2 sections)
-- ============================================
INSERT INTO lockers (id, locker_number, section, is_occupied, created_at) VALUES
('10000000-0000-0000-0000-000000000001', 'A-001', 'Section A', true, NOW() - INTERVAL '60 days'),
('10000000-0000-0000-0000-000000000002', 'A-002', 'Section A', true, NOW() - INTERVAL '60 days'),
('10000000-0000-0000-0000-000000000003', 'A-003', 'Section A', false, NOW() - INTERVAL '60 days'),
('10000000-0000-0000-0000-000000000004', 'A-004', 'Section A', true, NOW() - INTERVAL '60 days'),
('10000000-0000-0000-0000-000000000005', 'A-005', 'Section A', false, NOW() - INTERVAL '60 days'),
('10000000-0000-0000-0000-000000000006', 'A-006', 'Section A', false, NOW() - INTERVAL '60 days'),
('10000000-0000-0000-0000-000000000007', 'A-007', 'Section A', true, NOW() - INTERVAL '60 days'),
('10000000-0000-0000-0000-000000000008', 'A-008', 'Section A', false, NOW() - INTERVAL '60 days'),
('10000000-0000-0000-0000-000000000009', 'A-009', 'Section A', false, NOW() - INTERVAL '60 days'),
('10000000-0000-0000-0000-000000000010', 'A-010', 'Section A', false, NOW() - INTERVAL '60 days'),
('20000000-0000-0000-0000-000000000001', 'B-001', 'Section B', true, NOW() - INTERVAL '60 days'),
('20000000-0000-0000-0000-000000000002', 'B-002', 'Section B', false, NOW() - INTERVAL '60 days'),
('20000000-0000-0000-0000-000000000003', 'B-003', 'Section B', true, NOW() - INTERVAL '60 days'),
('20000000-0000-0000-0000-000000000004', 'B-004', 'Section B', false, NOW() - INTERVAL '60 days'),
('20000000-0000-0000-0000-000000000005', 'B-005', 'Section B', false, NOW() - INTERVAL '60 days'),
('20000000-0000-0000-0000-000000000006', 'B-006', 'Section B', true, NOW() - INTERVAL '60 days'),
('20000000-0000-0000-0000-000000000007', 'B-007', 'Section B', false, NOW() - INTERVAL '60 days'),
('20000000-0000-0000-0000-000000000008', 'B-008', 'Section B', false, NOW() - INTERVAL '60 days'),
('20000000-0000-0000-0000-000000000009', 'B-009', 'Section B', true, NOW() - INTERVAL '60 days'),
('20000000-0000-0000-0000-000000000010', 'B-010', 'Section B', false, NOW() - INTERVAL '60 days');

-- ============================================
-- 3. VISITS (Mix of checked-in, checked-out, and pending)
-- ============================================
-- Rajesh Kumar - Currently checked in
INSERT INTO visits (id, profile_id, arrival_date, departure_date, stay_area, status, locker_id, remarks, created_at) VALUES
('01000000-0000-0000-0000-000000000001', 'a1b2c3d4-e5f6-7890-abcd-ef1234567890', NOW() - INTERVAL '5 days', NULL, 'Dormitory A', 'checked-in', '10000000-0000-0000-0000-000000000001', 'Extended stay', NOW() - INTERVAL '5 days');

-- Priya Sharma - Currently checked in
INSERT INTO visits (id, profile_id, arrival_date, departure_date, stay_area, status, locker_id, remarks, created_at) VALUES
('02000000-0000-0000-0000-000000000001', 'b2c3d4e5-f6a7-8901-bcde-f12345678901', NOW() - INTERVAL '3 days', NULL, 'Dormitory B', 'checked-in', '10000000-0000-0000-0000-000000000002', NULL, NOW() - INTERVAL '3 days');

-- Amit Patel - Checked out
INSERT INTO visits (id, profile_id, arrival_date, departure_date, stay_area, status, locker_id, remarks, created_at) VALUES
('03000000-0000-0000-0000-000000000001', 'c3d4e5f6-a7b8-9012-cdef-123456789012', NOW() - INTERVAL '15 days', NOW() - INTERVAL '8 days', 'Dormitory A', 'checked-out', NULL, 'Completed seva period', NOW() - INTERVAL '15 days');

-- Sneha Reddy - Currently checked in
INSERT INTO visits (id, profile_id, arrival_date, departure_date, stay_area, status, locker_id, remarks, created_at) VALUES
('04000000-0000-0000-0000-000000000001', 'd4e5f6a7-b8c9-0123-def1-234567890123', NOW() - INTERVAL '7 days', NULL, 'Guest House', 'checked-in', '10000000-0000-0000-0000-000000000004', 'Overseas volunteer', NOW() - INTERVAL '7 days');

-- Vikram Singh - Currently checked in
INSERT INTO visits (id, profile_id, arrival_date, departure_date, stay_area, status, locker_id, remarks, created_at) VALUES
('05000000-0000-0000-0000-000000000001', 'e5f6a7b8-c9d0-1234-ef12-345678901234', NOW() - INTERVAL '2 days', NULL, 'Dormitory C', 'checked-in', '10000000-0000-0000-0000-000000000007', NULL, NOW() - INTERVAL '2 days');

-- Ananya Iyer - Pending arrival
INSERT INTO visits (id, profile_id, arrival_date, departure_date, stay_area, status, locker_id, remarks, created_at) VALUES
('06000000-0000-0000-0000-000000000001', 'f6a7b8c9-d0e1-2345-f123-456789012345', NOW() + INTERVAL '2 days', NULL, NULL, 'pending', NULL, 'Arriving soon', NOW());

-- Divya Nair - Currently checked in
INSERT INTO visits (id, profile_id, arrival_date, departure_date, stay_area, status, locker_id, remarks, created_at) VALUES
('08000000-0000-0000-0000-000000000001', 'b8c9d0e1-f2a3-4567-2345-678901234567', NOW() - INTERVAL '10 days', NULL, 'Dormitory B', 'checked-in', '20000000-0000-0000-0000-000000000001', NULL, NOW() - INTERVAL '10 days');

-- Arjun Desai - Currently checked in
INSERT INTO visits (id, profile_id, arrival_date, departure_date, stay_area, status, locker_id, remarks, created_at) VALUES
('09000000-0000-0000-0000-000000000001', 'c9d0e1f2-a3b4-5678-3456-789012345678', NOW() - INTERVAL '4 days', NULL, 'Guest House', 'checked-in', '20000000-0000-0000-0000-000000000003', 'International volunteer', NOW() - INTERVAL '4 days');

-- Meera Joshi - Checked out (past visit)
INSERT INTO visits (id, profile_id, arrival_date, departure_date, stay_area, status, locker_id, remarks, created_at) VALUES
('0a000000-0000-0000-0000-000000000001', 'd0e1f2a3-b4c5-6789-4567-890123456789', NOW() - INTERVAL '20 days', NOW() - INTERVAL '15 days', 'Dormitory A', 'checked-out', NULL, NULL, NOW() - INTERVAL '20 days');

-- Meera Joshi - Currently checked in (new visit)
INSERT INTO visits (id, profile_id, arrival_date, departure_date, stay_area, status, locker_id, remarks, created_at) VALUES
('0a000000-0000-0000-0000-000000000002', 'd0e1f2a3-b4c5-6789-4567-890123456789', NOW() - INTERVAL '1 day', NULL, 'Dormitory C', 'checked-in', '20000000-0000-0000-0000-000000000006', 'Second visit', NOW() - INTERVAL '1 day');

-- ============================================
-- 4. SCHEDULES (Seva assignments for checked-in volunteers)
-- ============================================
-- Rajesh Kumar schedules
INSERT INTO schedules (id, profile_id, visit_id, date, seva_type, location, notes, created_at) VALUES
('b1000000-0000-0000-0000-000000000001', 'a1b2c3d4-e5f6-7890-abcd-ef1234567890', '01000000-0000-0000-0000-000000000001', NOW() - INTERVAL '2 days', 'Kitchen Support', 'Main Kitchen', NULL, NOW() - INTERVAL '3 days'),
('b1000000-0000-0000-0000-000000000002', 'a1b2c3d4-e5f6-7890-abcd-ef1234567890', '01000000-0000-0000-0000-000000000001', NOW() - INTERVAL '1 day', 'Kitchen Support', 'Main Kitchen', NULL, NOW() - INTERVAL '2 days'),
('b1000000-0000-0000-0000-000000000003', 'a1b2c3d4-e5f6-7890-abcd-ef1234567890', '01000000-0000-0000-0000-000000000001', NOW(), 'Dining', 'Dining Hall', NULL, NOW() - INTERVAL '1 day'),
('b1000000-0000-0000-0000-000000000004', 'a1b2c3d4-e5f6-7890-abcd-ef1234567890', '01000000-0000-0000-0000-000000000001', NOW() + INTERVAL '1 day', 'Kitchen Support', 'Main Kitchen', NULL, NOW());

-- Priya Sharma schedules
INSERT INTO schedules (id, profile_id, visit_id, date, seva_type, location, notes, created_at) VALUES
('b2000000-0000-0000-0000-000000000001', 'b2c3d4e5-f6a7-8901-bcde-f12345678901', '02000000-0000-0000-0000-000000000001', NOW(), 'Counter', 'Reception Counter', NULL, NOW() - INTERVAL '1 day'),
('b2000000-0000-0000-0000-000000000002', 'b2c3d4e5-f6a7-8901-bcde-f12345678901', '02000000-0000-0000-0000-000000000001', NOW() + INTERVAL '1 day', 'Counter', 'Reception Counter', NULL, NOW());

-- Sneha Reddy schedules
INSERT INTO schedules (id, profile_id, visit_id, date, seva_type, location, notes, created_at) VALUES
('b4000000-0000-0000-0000-000000000001', 'd4e5f6a7-b8c9-0123-def1-234567890123', '04000000-0000-0000-0000-000000000001', NOW() - INTERVAL '1 day', 'Dining', 'Dining Hall', 'Morning shift', NOW() - INTERVAL '2 days'),
('b4000000-0000-0000-0000-000000000002', 'd4e5f6a7-b8c9-0123-def1-234567890123', '04000000-0000-0000-0000-000000000001', NOW(), 'Dining', 'Dining Hall', 'Morning shift', NOW() - INTERVAL '1 day'),
('b4000000-0000-0000-0000-000000000003', 'd4e5f6a7-b8c9-0123-def1-234567890123', '04000000-0000-0000-0000-000000000001', NOW() + INTERVAL '1 day', 'Dining', 'Dining Hall', NULL, NOW());

-- Vikram Singh schedules
INSERT INTO schedules (id, profile_id, visit_id, date, seva_type, location, notes, created_at) VALUES
('b5000000-0000-0000-0000-000000000001', 'e5f6a7b8-c9d0-1234-ef12-345678901234', '05000000-0000-0000-0000-000000000001', NOW(), 'Counter', 'Information Desk', NULL, NOW() - INTERVAL '1 day'),
('b5000000-0000-0000-0000-000000000002', 'e5f6a7b8-c9d0-1234-ef12-345678901234', '05000000-0000-0000-0000-000000000001', NOW() + INTERVAL '2 days', 'Counter', 'Information Desk', NULL, NOW());

-- Divya Nair schedules
INSERT INTO schedules (id, profile_id, visit_id, date, seva_type, location, notes, created_at) VALUES
('b8000000-0000-0000-0000-000000000001', 'b8c9d0e1-f2a3-4567-2345-678901234567', '08000000-0000-0000-0000-000000000001', NOW() - INTERVAL '3 days', 'Dining', 'Dining Hall', NULL, NOW() - INTERVAL '4 days'),
('b8000000-0000-0000-0000-000000000002', 'b8c9d0e1-f2a3-4567-2345-678901234567', '08000000-0000-0000-0000-000000000001', NOW(), 'Dining', 'Dining Hall', NULL, NOW() - INTERVAL '1 day'),
('b8000000-0000-0000-0000-000000000003', 'b8c9d0e1-f2a3-4567-2345-678901234567', '08000000-0000-0000-0000-000000000001', NOW() + INTERVAL '1 day', 'Dining', 'Dining Hall', NULL, NOW());

-- Arjun Desai schedules
INSERT INTO schedules (id, profile_id, visit_id, date, seva_type, location, notes, created_at) VALUES
('b9000000-0000-0000-0000-000000000001', 'c9d0e1f2-a3b4-5678-3456-789012345678', '09000000-0000-0000-0000-000000000001', NOW(), 'Kitchen Support', 'Main Kitchen', NULL, NOW() - INTERVAL '1 day'),
('b9000000-0000-0000-0000-000000000002', 'c9d0e1f2-a3b4-5678-3456-789012345678', '09000000-0000-0000-0000-000000000001', NOW() + INTERVAL '1 day', 'Kitchen Support', 'Main Kitchen', NULL, NOW());

-- Meera Joshi schedules (current visit)
INSERT INTO schedules (id, profile_id, visit_id, date, seva_type, location, notes, created_at) VALUES
('ba000000-0000-0000-0000-000000000001', 'd0e1f2a3-b4c5-6789-4567-890123456789', '0a000000-0000-0000-0000-000000000002', NOW(), 'Counter', 'Reception Counter', NULL, NOW() - INTERVAL '1 day'),
('ba000000-0000-0000-0000-000000000002', 'd0e1f2a3-b4c5-6789-4567-890123456789', '0a000000-0000-0000-0000-000000000002', NOW() + INTERVAL '2 days', 'Counter', 'Reception Counter', NULL, NOW());

-- ============================================
-- 5. FEEDBACK (Mix of positive, negative, neutral)
-- ============================================
INSERT INTO feedbacks (id, profile_id, visit_id, content, type, created_by, created_at) VALUES
('c1000000-0000-0000-0000-000000000001', 'a1b2c3d4-e5f6-7890-abcd-ef1234567890', '01000000-0000-0000-0000-000000000001', 'Very dedicated and punctual. Great team player.', 'Positive', 'Kitchen Coordinator', NOW() - INTERVAL '2 days'),
('c2000000-0000-0000-0000-000000000001', 'b2c3d4e5-f6a7-8901-bcde-f12345678901', '02000000-0000-0000-0000-000000000001', 'Needs more training on counter procedures.', 'Neutral', 'Counter Manager', NOW() - INTERVAL '1 day'),
('c3000000-0000-0000-0000-000000000001', 'c3d4e5f6-a7b8-9012-cdef-123456789012', '03000000-0000-0000-0000-000000000001', 'Excellent work throughout the stay. Highly recommended.', 'Positive', 'Seva Coordinator', NOW() - INTERVAL '8 days'),
('c4000000-0000-0000-0000-000000000001', 'd4e5f6a7-b8c9-0123-def1-234567890123', '04000000-0000-0000-0000-000000000001', 'Always arrives early for shifts. Very enthusiastic.', 'Positive', 'Dining Manager', NOW() - INTERVAL '1 day'),
('c5000000-0000-0000-0000-000000000001', 'e5f6a7b8-c9d0-1234-ef12-345678901234', '05000000-0000-0000-0000-000000000001', 'Good coordination skills. Handles queries well.', 'Positive', 'Reception Head', NOW()),
('c7000000-0000-0000-0000-000000000001', 'a7b8c9d0-e1f2-3456-1234-567890123456', NULL, 'Inappropriate behavior reported. Blocked from future visits.', 'Negative', 'Admin', NOW() - INTERVAL '5 days'),
('c8000000-0000-0000-0000-000000000001', 'b8c9d0e1-f2a3-4567-2345-678901234567', '08000000-0000-0000-0000-000000000001', 'Outstanding performance in dining hall. Very helpful to guests.', 'Positive', 'Dining Manager', NOW() - INTERVAL '3 days'),
('c9000000-0000-0000-0000-000000000001', 'c9d0e1f2-a3b4-5678-3456-789012345678', '09000000-0000-0000-0000-000000000001', 'Adapting well to the environment. Positive attitude.', 'Positive', 'Kitchen Coordinator', NOW() - INTERVAL '1 day'),
('ca000000-0000-0000-0000-000000000001', 'd0e1f2a3-b4c5-6789-4567-890123456789', '0a000000-0000-0000-0000-000000000002', 'Prefers counter duty and excels at it.', 'Positive', 'Counter Manager', NOW());

-- ============================================
-- VERIFICATION QUERIES
-- ============================================
-- Run these to verify the data was inserted correctly:

SELECT COUNT(*) as profile_count FROM profiles;
SELECT COUNT(*) as locker_count FROM lockers;
SELECT COUNT(*) as visit_count FROM visits;
SELECT COUNT(*) as schedule_count FROM schedules;
SELECT COUNT(*) as feedback_count FROM feedbacks;

View currently checked-in volunteers:
SELECT p.name, v.arrival_date, v.stay_area, l.locker_number 
FROM profiles p
JOIN visits v ON p.id = v.profile_id
LEFT JOIN lockers l ON v.locker_id = l.id
WHERE v.status = 'checked-in'
ORDER BY v.arrival_date;