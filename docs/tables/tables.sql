
CREATE TABLE seat_class (
    id INT PRIMARY KEY,
    class_name VARCHAR(5),
    min_price varchar(10),
    normal_price varchar(10),
    max_price varchar(10)
);

CREATE TABLE bookings (
    id SERIAL PRIMARY KEY,
    seat_id varchar(15) NOT NULL,
    booking_id varchar(10) NOT NULL unique,
    user_id INT NOT NULL,
    date_created timestamp,
    FOREIGN KEY (user_id) REFERENCES users (user_id),
    FOREIGN KEY (seat_id) REFERENCES seats (seat_id)

);

CREATE TABLE booked_seat (
    id SERIAL PRIMARY KEY,
    booking_id INT NOT NULL,
    seat_id varchar(15) NOT NULL,
    FOREIGN KEY (booking_id) REFERENCES bookings,
    FOREIGN KEY (seat_id) REFERENCES seats
);


  
  
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
	 user_name VARCHAR(20) NOT NULL,
    user_phone VARCHAR(20) NOT NULL,
	user_email varchar(50) NoT NULL unique,
    seat_id varchar(15) NOT NULL,
	 seat_class VARCHAR(50) NOT NULL,
    FOREIGN KEY (seat_id) REFERENCES seats (seat_id)
);


CREATE TABLE seat_pricing (
    id INT PRIMARY KEY,
    class_name VARCHAR(5) PRIMARY KEY,
    min_price FLOAT,
    max_price FLOAT,
    normal_price FLOAT
);



