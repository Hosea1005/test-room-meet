# room-meet

Room Meet merupakan aplikasi untuk management room meeting, pada aplikasi ini terdiri dari 2 role antara lain "admin" dan "general". Admin dapat melakukan CRUD pada Room Meeting dan admin dapat rejected atau approved bookingan room dari general. Admin juga dapat Melakukan CRUD Inventory, yang dimana inventory merupakan item-item apa saja yang ada pada room meet. General dapat melihat list room meet yang available, dan juga dapat melakukan booking pada room tersebut, namun berdasarkan capacity dari room tersebut.

Fitur  :
- Show list room (Status user General dan Admin)
- http://localhost:8080/room"
- Update room (admin)
- http://localhost:8080/room/update-room/{id}"
- Add new room (admin)
- http://localhost:8080/room/add-room
- Delete room (admin)
- http://localhost:8080/room/delete-room/{id}
- Update status room (admin))
- http://localhost:8080/room/update-status/{id}  [status only:"rejected" or "approved"]
- Login and Register user :
- http://localhost:8080/login
- http://localhost:8080/register
- Akses masing2 fitur (belum sempat fungsi check JWT)
- Request booking room meeting (General) (nantinya akan ada check request berapa orang dalam room dan capacity dari roomnya)
- http://localhost:8080/book/request-room/{id}

  Fitur Belum Sempat Di Develop
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