import React from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import "../css/header.css";
import { Link } from "react-router-dom";

const Header = () => {
  return (
    <div className="square-header">
      <nav className="navbar navbar-expand">
        <div className="container-fluid">
          <Link className="navbar-brand" to="/">
            TIMELAPSE
          </Link>
          <div className="navbar-collapse">
            <ul className="navbar-nav">
              <li className="nav-item">
                <Link className="nav-link" to="/cars">
                  ซื้อรถยนต์
                </Link>
              </li>
              <li className="nav-item">
                <Link className="nav-link" to="/promotion">
                  โปรโมชั่น
                </Link>
              </li>
              <li className="nav-item">
                <Link className="nav-link" to="/contact">
                  ติดต่อเรา
                </Link>
              </li>
              <li className="nav-item">
                <Link className="nav-link" to="/about">
                  เกี่ยวกับ TIMELAPSE
                </Link>
              </li>
              <li className="nav-item">
                <Link className="nav-link" to="/review">
                  รีวิว
                </Link>
              </li>
            </ul>
          </div>
          <div>
            <Link className="nav-link" to="/login">
              เข้าสู่ระบบ/สมัครสมาชิก
            </Link>
          </div>
        </div>
      </nav>
    </div>
  );
};

export default Header;
