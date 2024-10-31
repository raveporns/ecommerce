import React, { useState } from "react";
import Header from "../component/header";
import "../css/car-list.css";
const cars = [
  { id: 1, name: "Toyota Camry", type: "Sedan", price: "1,500,000 บาท", imageUrl: "images/car1.jpeg" },
  { id: 2, name: "Honda Civic", type: "Sedan", price: "1,200,000 บาท", imageUrl: "images/car1.jpeg" },
  { id: 3, name: "Mazda 3", type: "SUV", price: "1,100,000 บาท", imageUrl: "images/car1.jpeg" },
  { id: 4, name: "Nissan Almera", type: "Truck", price: "900,000 บาท", imageUrl: "images/car1.jpeg" },
  { id: 5, name: "Ford Ranger", type: "Truck", price: "1,300,000 บาท", imageUrl: "images/car1.jpeg" },
  { id: 6, name: "Chevrolet Colorado", type: "Truck", price: "1,400,000 บาท", imageUrl: "images/car1.jpeg" },
  { id: 7, name: "Toyota Camry", type: "Sedan", price: "1,500,000 บาท", imageUrl: "images/car1.jpeg" },
  { id: 8, name: "Honda Civic", type: "Sedan", price: "1,200,000 บาท", imageUrl: "images/car1.jpeg" },
  { id: 9, name: "Mazda 3", type: "SUV", price: "1,100,000 บาท", imageUrl: "images/car1.jpeg" },
  { id: 10, name: "Nissan Almera", type: "Truck", price: "900,000 บาท", imageUrl: "images/car1.jpeg" },
  { id: 11, name: "Ford Ranger", type: "Truck", price: "1,300,000 บาท", imageUrl: "images/car1.jpeg" },
  { id: 12, name: "Chevrolet Colorado", type: "Truck", price: "1,400,000 บาท", imageUrl: "images/car1.jpeg" },
];


const Cars = () => {
  const [filterType, setFilterType] = useState("All");
  const [filterDoors, setFilterDoors] = useState("All");

  // ฟิลเตอร์รายการรถยนต์ตามประเภทและจำนวนประตู
  const filteredCars = cars.filter(car => {
    const typeMatch = filterType === "All" || car.type === filterType;
    const doorsMatch = filterDoors === "All" || car.doors === parseInt(filterDoors);
    return typeMatch && doorsMatch;
  });

  return (
    <div>
      <Header />
      <h1>รายการรถยนต์</h1>
      <div className="container">
        {/* ฟิลเตอร์ด้านซ้าย */}
        <div className="filter">
          <h2>ประเภท</h2>
          <ul>
            <li onClick={() => setFilterType("All")}>ทั้งหมด</li>
            <li onClick={() => setFilterType("Sedan")}>Sedan</li>
            <li onClick={() => setFilterType("Convertible")}>Convertible</li>
            <li onClick={() => setFilterType("Truck")}>Truck</li>
            <li onClick={() => setFilterType("Coupe")}>Coupe</li>
          </ul>

          <h2>จำนวนประตู</h2>
          <div>
            <label>
              <input
                type="radio"
                name="doors"
                value="2"
                checked={filterDoors === "2"}
                onChange={() => setFilterDoors("2")}
              />
              2 ประตู
            </label>
            <label>
              <input
                type="radio"
                name="doors"
                value="4"
                checked={filterDoors === "4"}
                onChange={() => setFilterDoors("4")}
              />
              4 ประตู
            </label>
          </div>
        </div>

        {/* รายการรถยนต์ด้านขวา */}
        <div className="car-list">
          {filteredCars.map((car) => (
            <div key={car.id} className="car-item">
              <img src={car.imageUrl} alt={car.name} />
              <h2>{car.name}</h2>
              <p>{car.price}</p>
              <a href={`/cars/${car.id}`}>ดูรายละเอียด</a>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};
export default Cars;
