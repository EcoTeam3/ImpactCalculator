-- Insert initial data into carbon_footprint_logs table
INSERT INTO carbon_footprint_logs (user_id, category, amount, unit)
VALUES
(gen_random_uuid(), 'transport', 15.50, 'kg CO2'),
(gen_random_uuid(), 'energy', 125.75, 'kWh'),
(gen_random_uuid(), 'food', 5.25, 'kg CO2'),
(gen_random_uuid(), 'consumption', 20.00, 'kg CO2');

-- Insert initial data into donations table
INSERT INTO donations (user_id, cause, amount)
VALUES
(gen_random_uuid(), 'Save the Earth', 100.00),
(gen_random_uuid(), 'Animal Rescue', 50.00),
(gen_random_uuid(), 'Tree Plantation', 75.00),
(gen_random_uuid(), 'Ocean Cleanup', 120.00);
