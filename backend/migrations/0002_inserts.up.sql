-- Склады
INSERT INTO warehouses (id, name, address, coordinates) VALUES
  (uuid_generate_v4(), 'Склад №1', 'Москва, ул. Примерная, 1', ST_GeogFromText('SRID=4326;POINT(37.618423 55.751244)')),
  (uuid_generate_v4(), 'Склад №2', 'Санкт-Петербург, Невский пр., 10', ST_GeogFromText('SRID=4326;POINT(30.314130 59.938630)'));

-- Категории
INSERT INTO categories (id, name, image_url) VALUES
  (uuid_generate_v4(), 'Фрукты', '/static/categories/fruits.png'),
  (uuid_generate_v4(), 'Овощи', '/static/categories/vegetables.png'),
  (uuid_generate_v4(), 'Молочные продукты', '/static/categories/dairy.png');

-- Привязка категорий к складам
INSERT INTO warehouse_categories (id, warehouse_id, category_id)
SELECT uuid_generate_v4(), w.id, c.id
FROM warehouses w, categories c
WHERE (w.name = 'Склад №1' AND c.name IN ('Фрукты', 'Овощи'))
   OR (w.name = 'Склад №2' AND c.name IN ('Фрукты', 'Молочные продукты'));

-- Продукты
INSERT INTO products (id, title, description, weight_grams, available) VALUES
  (uuid_generate_v4(), 'Яблоко', 'Сочное красное яблоко', 150, true),
  (uuid_generate_v4(), 'Банан', 'Спелый банан', 120, true),
  (uuid_generate_v4(), 'Молоко 1л', 'Молоко пастеризованное', 1000, true),
  (uuid_generate_v4(), 'Огурец', 'Свежий огурец', 100, true);

-- Привязка товаров к категориям
INSERT INTO product_categories (id, product_id, category_id)
SELECT uuid_generate_v4(), p.id, c.id
FROM products p, categories c
WHERE (p.title = 'Яблоко' AND c.name = 'Фрукты')
   OR (p.title = 'Банан' AND c.name = 'Фрукты')
   OR (p.title = 'Молоко 1л' AND c.name = 'Молочные продукты')
   OR (p.title = 'Огурец' AND c.name = 'Овощи');

-- Картинки товаров
INSERT INTO product_images (id, product_id, url, type)
SELECT uuid_generate_v4(), p.id, '/static/products/' || lower(replace(p.title, ' ', '_')) || '.jpg', 'main'
FROM products p;

-- Остатки на складе
INSERT INTO warehouse_stock (id, warehouse_id, product_id, quantity, price)
SELECT uuid_generate_v4(), w.id, p.id, 100, 
  CASE 
    WHEN p.title IN ('Яблоко', 'Банан') THEN 59.90
    WHEN p.title = 'Молоко 1л' THEN 79.90
    WHEN p.title = 'Огурец' THEN 49.90
  END
FROM warehouses w, products p;

-- Промо
INSERT INTO promotions (id, title, subtitle, image_url, deeplink, sort_order, is_active, starts_at, ends_at)
VALUES 
  (uuid_generate_v4(), 'Скидка на фрукты', 'Только сегодня', '/static/promos/fruits.png', '/category/fruits', 1, true, now(), now() + interval '1 day');

-- Истории
INSERT INTO stories (id, title, image_url, deeplink, sort_order, is_active)
VALUES
  (uuid_generate_v4(), 'Новые поступления', '/static/stories/new.png', '/new', 1, true);

-- Коллекция "Выгодная полка"
INSERT INTO product_collections (id, name, description, sort_order, is_active)
VALUES 
  (uuid_generate_v4(), 'Выгодная полка', 'Лучшие цены на этой неделе', 1, true);

-- Привязка товаров к коллекции
INSERT INTO product_collection_items (id, collection_id, product_id, position)
SELECT uuid_generate_v4(), pc.id, p.id, row_number() OVER ()
FROM product_collections pc, products p
WHERE pc.name = 'Выгодная полка' AND p.title IN ('Яблоко', 'Молоко 1л');
