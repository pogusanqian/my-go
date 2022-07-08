# 创建user表
CREATE TABLE IF NOT EXISTS users(
  id INT PRIMARY KEY AUTO_INCREMENT,
  username VARCHAR(32) NOT NULL UNIQUE,
  passward VARCHAR(32) NOT NULL,
  email VARCHAR(64) NOT NULL
);

# 创建图书表
CREATE TABLE IF NOT EXISTS books(
  id INT PRIMARY KEY AUTO_INCREMENT,
  title VARCHAR(64) NOT NULL,
  author VARCHAR(32) NOT NULL,
  price DOUBLE(11, 2) NOT NULL,
  sales INT NOT NULL,
  stock INT NOT NULL,
  img_path text NOT NULL
);

# 创建sessions表
CREATE TABLE IF NOT EXISTS sessions(
  session_id VARCHAR(100) PRIMARY KEY,
  username VARCHAR(100) NOT NULL,
  user_id INT NOT NULL,
  FOREIGN KEY(user_id) REFERENCES users(id)
);

# 创建购物车表
CREATE TABLE IF NOT EXISTS carts(
  id VARCHAR(100) PRIMARY KEY,
  total_count INT NOT NULL,
  total_amount DOUBLE(11, 2) NOT NULL,
  user_id INT NOT NULL,
  FOREIGN KEY(user_id) REFERENCES users(id)
);

# 创建购物项表
CREATE TABLE IF NOT EXISTS cart_itmes(
  id INT PRIMARY KEY AUTO_INCREMENT,
  COUNT INT NOT NULL,
  amount DOUBLE(11, 2) NOT NULL,
  book_id INT NOT NULL,
  cart_id VARCHAR(100) NOT NULL,
  FOREIGN KEY(book_id) REFERENCES books(id),
  FOREIGN KEY(cart_id) REFERENCES carts(id)
);

# 创建订单表
CREATE TABLE IF NOT EXISTS orders(
  id VARCHAR(100) PRIMARY KEY,
  create_time DATETIME NOT NULL,
  total_count INT NOT NULL,
  total_amount DOUBLE(11, 2) NOT NULL,
  state INT NOT NULL,
  user_id INT,
  FOREIGN KEY(user_id) REFERENCES users(id)
);

# 创建订单项表
CREATE TABLE IF NOT EXISTS order_items(
  id INT PRIMARY KEY AUTO_INCREMENT,
  COUNT INT NOT NULL,
  amount DOUBLE(11, 2) NOT NULL,
  title VARCHAR(100) NOT NULL,
  author VARCHAR(100) NOT NULL,
  price DOUBLE(11, 2) NOT NULL,
  img_path VARCHAR(100) NOT NULL,
  order_id VARCHAR(100) NOT NULL,
  FOREIGN KEY(order_id) REFERENCES orders(id)
);