CREATE TYPE footprint_category AS ENUM ('transport', 'energy', 'food', 'consumption');

CREATE TABLE carbon_footprint_logs (
    foot_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID,
    category footprint_category,
    amount DECIMAL(10, 2),
    unit VARCHAR(20),
    logged_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Donations table
CREATE TABLE donations (
    donation_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID,
    cause VARCHAR(100),
    amount DECIMAL(10, 2),
    donated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);