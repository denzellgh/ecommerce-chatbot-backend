INSERT INTO roles (id, name, description, created_at) VALUES
('role_admin', 'admin', 'Administrator with full access to all features including dashboard and metrics', NOW()),
('role_user', 'user', 'Regular user with access to chatbot and recommendations only', NOW()),
('role_manager', 'manager', 'Store manager with limited admin access for inventory oversight', NOW()),
('role_support', 'support', 'Customer support representative with read-only access to most features', NOW());

-- Password: admin123
INSERT INTO users (id, name, email, password, role_id, is_active, created_at) VALUES
('user_admin_001', 'Admin User', 'admin@makerstech.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'role_admin', TRUE, NOW()),

-- Password: manager123  
('user_manager_001', 'Store Manager', 'manager@makerstech.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'role_manager', TRUE, NOW()),

-- Password: support123
('user_support_001', 'Customer Support', 'support@makerstech.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'role_support', TRUE, NOW()),

-- Password: user123
('user_customer_001', 'John Doe', 'john.doe@email.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'role_user', TRUE, NOW()),
('user_customer_002', 'Jane Smith', 'jane.smith@email.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'role_user', TRUE, NOW()),
('user_customer_003', 'Mike Johnson', 'mike.johnson@email.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'role_user', TRUE, NOW()),

-- Password: demo123
('user_demo_001', 'Demo User', 'demo@makerstech.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'role_user', TRUE, NOW()),

-- Password: testadmin123
('user_test_admin', 'Test Administrator', 'testadmin@makerstech.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'role_admin', TRUE, NOW()),

-- Some inactive users for testing
('user_inactive_001', 'Inactive User', 'inactive@makerstech.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'role_user', FALSE, NOW()),

-- Users with different creation dates for testing
('user_old_001', 'Old User', 'olduser@makerstech.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'role_user', TRUE, NOW() - INTERVAL '30 days'),
('user_recent_001', 'Recent User', 'recent@makerstech.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'role_user', TRUE, NOW() - INTERVAL '2 days');

UPDATE users SET last_login = NOW() - INTERVAL '1 hour' WHERE id = 'user_admin_001';
UPDATE users SET last_login = NOW() - INTERVAL '3 hours' WHERE id = 'user_customer_001';
UPDATE users SET last_login = NOW() - INTERVAL '1 day' WHERE id = 'user_demo_001';
UPDATE users SET last_login = NOW() - INTERVAL '2 days' WHERE id = 'user_manager_001';

INSERT INTO user_preferences (id, user_session, preferred_categories, price_range_min, price_range_max, preferred_brands, created_at) VALUES
('pref_admin_001', 'user_admin_001', '{"cat_laptops", "cat_desktops", "cat_gaming"}', 1000.00, 5000.00, '{"Apple", "Dell", "ASUS"}', NOW()),
('pref_customer_001', 'user_customer_001', '{"cat_phones", "cat_audio", "cat_accessories"}', 100.00, 1000.00, '{"Apple", "Samsung", "Sony"}', NOW()),
('pref_customer_002', 'user_customer_002', '{"cat_laptops", "cat_tablets"}', 500.00, 1500.00, '{"Apple", "Samsung"}', NOW()),
('pref_demo_001', 'user_demo_001', '{"cat_gaming", "cat_audio"}', 50.00, 800.00, '{"Logitech", "Razer", "Bose"}', NOW());