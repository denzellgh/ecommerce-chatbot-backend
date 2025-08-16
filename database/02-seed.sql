INSERT INTO categories (id, name, description, created_at) VALUES
('cat_laptops', 'Laptops', 'Portable computers for work and gaming', NOW()),
('cat_desktops', 'Desktop Computers', 'High-performance desktop systems', NOW()),
('cat_phones', 'Smartphones', 'Latest mobile devices and accessories', NOW()),
('cat_tablets', 'Tablets', 'Portable tablets for productivity and entertainment', NOW()),
('cat_accessories', 'Accessories', 'Computer and mobile accessories', NOW()),
('cat_gaming', 'Gaming', 'Gaming peripherals and equipment', NOW()),
('cat_audio', 'Audio', 'Headphones, speakers, and audio equipment', NOW());

INSERT INTO products (id, name, brand, category_id, price, stock_quantity, description, specs, created_at) VALUES
-- Laptops
('prod_laptop_hp_001', 'HP Pavilion 15', 'HP', 'cat_laptops', 899.99, 8, 'Versatile laptop for work and entertainment', 
 '{"processor": "Intel i5-12500H", "ram": "8GB DDR4", "storage": "512GB SSD", "display": "15.6 FHD", "graphics": "Intel Iris Xe"}', NOW()),

('prod_laptop_dell_001', 'Dell XPS 13', 'Dell', 'cat_laptops', 1299.99, 5, 'Premium ultrabook with stunning display', 
 '{"processor": "Intel i7-1355U", "ram": "16GB LPDDR5", "storage": "1TB SSD", "display": "13.4 OLED", "graphics": "Intel Iris Xe"}', NOW()),

('prod_laptop_apple_001', 'MacBook Air M3', 'Apple', 'cat_laptops', 1499.99, 3, 'Powerful and efficient laptop with M3 chip', 
 '{"processor": "Apple M3", "ram": "16GB Unified", "storage": "512GB SSD", "display": "13.6 Retina", "graphics": "Apple M3 GPU"}', NOW()),

('prod_laptop_asus_001', 'ASUS ROG Strix G15', 'ASUS', 'cat_laptops', 1699.99, 4, 'High-performance gaming laptop', 
 '{"processor": "AMD Ryzen 7 6800H", "ram": "16GB DDR5", "storage": "1TB SSD", "display": "15.6 FHD 144Hz", "graphics": "RTX 4060"}', NOW()),

-- Desktop Computers
('prod_desktop_hp_001', 'HP Pavilion Desktop', 'HP', 'cat_desktops', 799.99, 6, 'Reliable desktop for everyday computing', 
 '{"processor": "Intel i5-13400", "ram": "16GB DDR4", "storage": "512GB SSD", "graphics": "Intel UHD 730", "ports": "USB 3.2, HDMI, DisplayPort"}', NOW()),

('prod_desktop_dell_001', 'Dell OptiPlex 7010', 'Dell', 'cat_desktops', 999.99, 7, 'Business-grade desktop computer', 
 '{"processor": "Intel i7-13700", "ram": "32GB DDR5", "storage": "1TB SSD", "graphics": "Intel UHD 770", "ports": "Multiple USB, DisplayPort"}', NOW()),

-- Smartphones
('prod_phone_iphone_001', 'iPhone 15 Pro', 'Apple', 'cat_phones', 1199.99, 12, 'Latest iPhone with titanium design', 
 '{"storage": "256GB", "display": "6.1 Super Retina XDR", "camera": "48MP Triple", "battery": "All-day battery", "connectivity": "5G"}', NOW()),

('prod_phone_samsung_001', 'Samsung Galaxy S24+', 'Samsung', 'cat_phones', 999.99, 15, 'Premium Android smartphone', 
 '{"storage": "256GB", "display": "6.7 Dynamic AMOLED 2X", "camera": "50MP Triple", "battery": "4900mAh", "connectivity": "5G"}', NOW()),

('prod_phone_pixel_001', 'Google Pixel 8 Pro', 'Google', 'cat_phones', 899.99, 10, 'AI-powered photography smartphone', 
 '{"storage": "128GB", "display": "6.7 LTPO OLED", "camera": "50MP Triple with AI", "battery": "5050mAh", "connectivity": "5G"}', NOW()),

-- Tablets
('prod_tablet_ipad_001', 'iPad Pro 12.9"', 'Apple', 'cat_tablets', 1299.99, 6, 'Professional tablet with M4 chip', 
 '{"processor": "Apple M4", "storage": "512GB", "display": "12.9 Liquid Retina XDR", "connectivity": "Wi-Fi 6E", "accessories": "Apple Pencil Pro compatible"}', NOW()),

('prod_tablet_samsung_001', 'Galaxy Tab S9+', 'Samsung', 'cat_tablets', 899.99, 8, 'Premium Android tablet with S Pen', 
 '{"processor": "Snapdragon 8 Gen 2", "storage": "256GB", "display": "12.4 Dynamic AMOLED 2X", "connectivity": "Wi-Fi 6E", "accessories": "S Pen included"}', NOW()),

-- Gaming Accessories
('prod_gaming_logitech_001', 'Logitech G Pro X Superlight', 'Logitech', 'cat_gaming', 149.99, 20, 'Ultra-lightweight wireless gaming mouse', 
 '{"dpi": "25600 DPI", "weight": "63g", "battery": "70 hours", "connectivity": "LIGHTSPEED Wireless", "switches": "LIGHTFORCE hybrid"}', NOW()),

('prod_gaming_razer_001', 'Razer BlackWidow V4 Pro', 'Razer', 'cat_gaming', 229.99, 12, 'Premium mechanical gaming keyboard', 
 '{"switches": "Razer Green Mechanical", "lighting": "Razer Chroma RGB", "connectivity": "Wired/Wireless", "features": "Command Dial, Media Keys"}', NOW()),

-- Audio Equipment
('prod_audio_sony_001', 'Sony WH-1000XM5', 'Sony', 'cat_audio', 399.99, 18, 'Industry-leading noise canceling headphones', 
 '{"driver": "30mm", "noise_canceling": "V1 Processor", "battery": "30 hours", "connectivity": "Bluetooth 5.2, NFC", "features": "Quick Attention Mode"}', NOW()),

('prod_audio_bose_001', 'Bose QuietComfort Earbuds', 'Bose', 'cat_audio', 299.99, 14, 'Premium wireless earbuds with ANC', 
 '{"driver": "Custom 6.4mm", "noise_canceling": "Active", "battery": "24 hours total", "connectivity": "Bluetooth 5.1", "features": "IPX4 rated"}', NOW()),

-- Accessories
('prod_acc_anker_001', 'Anker PowerCore 26800', 'Anker', 'cat_accessories', 59.99, 25, 'High-capacity portable power bank', 
 '{"capacity": "26800mAh", "ports": "3 USB outputs", "input": "Micro USB, USB-C", "charging_time": "6-7 hours", "features": "MultiProtect Safety"}', NOW()),

('prod_acc_belkin_001', 'Belkin 3-in-1 Wireless Charger', 'Belkin', 'cat_accessories', 149.99, 16, 'Fast wireless charging station', 
 '{"compatibility": "iPhone, Apple Watch, AirPods", "power": "15W fast charging", "design": "Magnetic alignment", "features": "LED indicator, foreign object detection"}', NOW());

INSERT INTO user_preferences (id, user_session, preferred_categories, price_range_min, price_range_max, preferred_brands, created_at) VALUES
('pref_session_001', 'demo_session_123', '{"cat_laptops", "cat_gaming"}', 500.00, 2000.00, '{"Apple", "ASUS", "Dell"}', NOW()),
('pref_session_002', 'demo_session_456', '{"cat_phones", "cat_audio"}', 200.00, 1500.00, '{"Apple", "Samsung", "Sony"}', NOW()),
('pref_session_003', 'demo_session_789', '{"cat_accessories", "cat_audio"}', 50.00, 500.00, '{"Anker", "Bose", "Logitech"}', NOW());