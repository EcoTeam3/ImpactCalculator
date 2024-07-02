-- Delete initial data from carbon_footprint_logs table
DELETE FROM carbon_footprint_logs WHERE amount IN 
(15.50, 125.75, 5.25, 20.00);

-- Delete initial data from donations table
DELETE FROM donations WHERE cause IN 
('Save the Earth', 'Animal Rescue', 'Tree Plantation', 'Ocean Cleanup');
