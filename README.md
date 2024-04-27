# test-room-meet
Fitur  :
- Show list room (Status user General dan Admin)
- Update room (admin)
- Add new room (admin)
- Delete room (admin)
- Update status room (admin)
- Order Room (Status user General)
- Login and Register user
- Akses masing2 fitur (belum sempat fungsi check JWT)
- Request booking room meeting (General) (nantinya akan ada check request berapa orang dalam room dan capacity dari roomnya)
  
- Fitur Belum Sempat Di Develop
- CRUD Inventory (Belum sempat Di Develop)

Role :
-"general"
-"admin"

Database :
- Users : id, name, password, role

  CREATE TABLE "users" (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(255) NOT NULL,
  password TEXT NOT NULL,
  role VARCHAR(50) NOT NULL DEFAULT 'general',
  email VARCHAR(255) NOT NULL UNIQUE,
  created_at TIMESTAMP DEFAULT current_timestamp
  );
- Room : id, name, location, location, status, capacity,updated_at,created_at,deleted_at

CREATE TABLE Rooms (
id SERIAL PRIMARY KEY,
name VARCHAR(255) NOT NULL,
location VARCHAR(255) NOT NULL,
status VARCHAR(50) NOT NULL,
capacity INT NOT NULL,
created_at TIMESTAMP DEFAULT current_timestamp,
deleted_at TIMESTAMP NULL,
updated_at TIMESTAMP NULL

);
Table Belum Sempat Di Develop
- Inventory : id, name, status, qty 