import React from "react";
import { useState, useEffect } from "react";

import Container from "react-bootstrap/Container";
import Navbar from "react-bootstrap/Navbar";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";

import "./app.scss";
import { Tile } from "./tile/tile";
import { Device, DeviceType } from "./interfaces/device.interface";

export default function App() {
  const [devices, setDevices] = useState([]);

  useEffect(() => {
    fetch("/devices")
      .then((res: Response) => res.json())
      .then((devices: Device[]) =>
        devices.filter((device: Device) => device.type !== DeviceType.BUTTON)
      )
      .then((devices: Device[]) => setDevices(devices))
      .catch((err) => {
        console.error(err.message);
      });
  }, []);

  return (
    <>
      <Navbar sticky="top" className="bg-body-tertiary">
        <Container>
          <Navbar.Brand href="#home">Smarthome</Navbar.Brand>
        </Container>
      </Navbar>
      <ul className="devices-container">
        {devices.map((device: Device) => (
          <li key={device.id}>
            <Tile device={device} />
          </li>
        ))}
      </ul>
    </>
  );
}
