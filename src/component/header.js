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
                <a className="nav-link" aria-current="page" href="#">
                  ซื้อรถยนต์
                </a>
              </li>
              <li className="nav-item">
                <a className="nav-link" href="#">
                  โปรโมชั่น
                </a>
              </li>
              <li className="nav-item">
                <a className="nav-link" href="#">
                  ติดต่อเรา
                </a>
              </li>
              <li className="nav-item">
                <a className="nav-link" href="#">
                  เกี่ยวกับ TIMELAPSE
                </a>
              </li>
              <li className="nav-item">
                <a className="nav-link" href="#">
                  รีวิว
                </a>
              </li>
            </ul>
          </div>
          <div>
            <a className="nav-link" href="#">
              เข้าสู่ระบบ/สมัครสมาชิก
            </a>
          </div>
        </div>
      </nav>
    </div>
  );
};

export default Header;
