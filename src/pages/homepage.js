import React from "react";
import Header from "../component/header";
import "../css/homepage.css";
import "bootstrap/dist/js/bootstrap.bundle.min.js";

const Homepage = () => {
  return (
    <div>
      <Header />

      <div
        id="carouselExampleControls"
        className="carousel slide"
        data-bs-ride="carousel"
      >
        <div className="carousel-inner">
          <div className="carousel-item active">
            <img
              src="images/car1.jpeg"
              className="d-block w-100"
              alt="..."
            ></img>
          </div>
          <div className="carousel-item">
            <img
              src="images/car2.jpeg"
              className="d-block w-100"
              alt="..."
            ></img>
          </div>
          <div className="carousel-item">
            <img
              src="images/car3.jpeg"
              className="d-block w-100"
              alt="..."
            ></img>
          </div>
        </div>
        <button
          className="carousel-control-prev"
          type="button"
          data-bs-target="#carouselExampleControls"
          data-bs-slide="prev"
        >
          <span
            className="carousel-control-prev-icon"
            aria-hidden="true"
          ></span>
          <span className="visually-hidden">Previous</span>
        </button>
        <button
          className="carousel-control-next"
          type="button"
          data-bs-target="#carouselExampleControls"
          data-bs-slide="next"
        >
          <span
            className="carousel-control-next-icon"
            aria-hidden="true"
          ></span>
          <span className="visually-hidden">Next</span>
        </button>
      </div>

      <div className="card-content">
        <div className="row row-cols-1 row-cols-md-3 g-4">
          <div className="col">
            <div className="card h-100">
              <img src="images/car1.jpeg" className="card-img-top" alt="..."></img>
              <div className="card-body">
                <h5 className="card-title">Card title</h5>
                <p className="card-text">test</p>
              </div>
            </div>
          </div>
          <div className="col">
            <div className="card h-100">
              <img src="images/car1.jpeg" className="card-img-top" alt="..."></img>
              <div className="card-body">
                <h5 className="card-title">Card title</h5>
                <p className="card-text">test</p>
              </div>
            </div>
          </div>
          <div className="col">
            <div className="card h-100">
              <img src="images/car1.jpeg" className="card-img-top" alt="..."></img>
              <div className="card-body">
                <h5 className="card-title">Card title</h5>
                <p className="card-text">test</p>
              </div>
            </div>
          </div>

          <div className="col">
            <div className="card h-100">
              <img src="images/car1.jpeg" className="card-img-top" alt="..."></img>
              <div className="card-body">
                <h5 className="card-title">Card title</h5>
                <p className="card-text">test</p>
              </div>
            </div>
          </div>
          <div className="col">
            <div className="card h-100">
              <img src="images/car1.jpeg" className="card-img-top" alt="..."></img>
              <div className="card-body">
                <h5 className="card-title">Card title</h5>
                <p className="card-text">test</p>
              </div>
            </div>
          </div>
          <div className="col">
            <div className="card h-100">
              <img src="images/car1.jpeg" className="card-img-top" alt="..."></img>
              <div className="card-body">
                <h5 className="card-title">Card title</h5>
                <p className="card-text">test</p>
              </div>
            </div>
          </div>
          <div className="col">
            <div className="card h-100">
              <img src="images/car1.jpeg" className="card-img-top" alt="..."></img>
              <div className="card-body">
                <h5 className="card-title">Card title</h5>
                <p className="card-text">test</p>
              </div>
            </div>
          </div>
          <div className="col">
            <div className="card h-100">
              <img src="images/car1.jpeg" className="card-img-top" alt="..."></img>
              <div className="card-body">
                <h5 className="card-title">Card title</h5>
                <p className="card-text">test</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};
export default Homepage;
